package domain

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T, log *log){
		"append and read a record succeeds": testAppendRead,
		"offset out of range error":         testOutOfRangeErr,
		"init with existing segments":       testInitExisting,
		"reader":                            testReader,
		"truncate":                          testTruncate,
	} {
		t.Run(scenario, func(t *testing.T) {
			dir, err := os.MkdirTemp("", "store-test")
			require.NoError(t, err)
			defer os.RemoveAll(dir)

			c := Config{}
			c.Segment.MaxStoreBytes = 32
			log, err := NewLog(dir, c)
			require.NoError(t, err)

			fn(t, log)
		})
	}
}

func testAppendRead(t *testing.T, log *log) {
	record := &Record{Value: []byte("hello world")}
	offset, err := log.Append(record)
	require.NoError(t, err)
	require.Equal(t, uint64(0), offset)

	got, err := log.Read(offset)
	require.NoError(t, err)
	require.Equal(t, got.Value, record.Value)
}

func testOutOfRangeErr(t *testing.T, log *log) {
	got, err := log.Read(1)
	require.Nil(t, got)
	require.Equal(t, err, ErrOffsetOutOfRange)
}

func testInitExisting(t *testing.T, log *log) {
	record := &Record{Value: []byte("hello world")}
	for i := 0; i < 3; i++ {
		_, err := log.Append(record)
		require.NoError(t, err)
	}
	err := log.Close()
	require.NoError(t, err)

	offset, err := log.LowestOffset()
	require.NoError(t, err)
	require.Equal(t, uint64(0), offset)
	offset, err = log.HighestOffset()
	require.NoError(t, err)
	require.Equal(t, uint64(2), offset)

	newLog, err := NewLog(log.Dir, log.Config)
	require.NoError(t, err)

	offset, err = newLog.LowestOffset()
	require.NoError(t, err)
	require.Equal(t, uint64(0), offset)
	offset, err = newLog.HighestOffset()
	require.NoError(t, err)
	require.Equal(t, uint64(2), offset)
}

func testReader(t *testing.T, log *log) {
	record := &Record{Value: []byte("hello world")}
	offset, err := log.Append(record)
	require.NoError(t, err)
	require.Equal(t, uint64(0), offset)

	reader := log.Reader()
	readAll, err := io.ReadAll(reader)
	require.NoError(t, err)

	got := &Record{}
	err = proto.Unmarshal(readAll[LenWidth:], got)
	require.NoError(t, err)
	require.Equal(t, record.Value, got.Value)
}

func testTruncate(t *testing.T, log *log) {
	record := &Record{Value: []byte("hello world")}
	for i := 0; i < 3; i++ {
		_, err := log.Append(record)
		require.NoError(t, err)
	}

	err := log.Truncate(1)
	require.NoError(t, err)

	got, err := log.Read(0)
	require.Nil(t, got)
	require.Error(t, err)
	require.Equal(t, err, ErrOffsetOutOfRange)
}
