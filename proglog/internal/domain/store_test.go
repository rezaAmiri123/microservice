package domain

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	msg      = []byte("hello world")
	widthLen = uint64(len(msg) + LenWidth)
)

func TestStoreAppendRead(t *testing.T) {
	file, err := os.CreateTemp("", "store_append_read_test")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	s, err := newStore(file)
	require.NoError(t, err)

	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)

	s, err = newStore(file)
	require.NoError(t, err)
	testRead(t, s)

}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	for i := uint64(1); i < 4; i++ {
		bytesNumber, pos, err := s.Append(msg)
		require.NoError(t, err)
		require.Equal(t, pos+bytesNumber, widthLen*i)
		require.True(t, s.buf.Size() > 0)
	}
}
func testRead(t *testing.T, s *store) {
	t.Helper()
	var pos uint64
	for i := uint64(1); i < 4; i++ {
		read, err := s.Read(pos)
		require.NoError(t, err)
		require.Equal(t, msg, read)
		pos += widthLen
	}
}

func testReadAt(t *testing.T, s *store) {
	t.Helper()
	for i, off := uint64(1), int64(0); i < 4; i++ {
		b := make([]byte, LenWidth)
		n, err := s.ReadAt(b, off)
		require.NoError(t, err)
		require.Equal(t, LenWidth, n)
		off += int64(n)

		size := Enc.Uint64(b)
		b = make([]byte, size)
		n, err = s.ReadAt(b, off)
		require.NoError(t, err)
		require.Equal(t, msg, b)
		require.Equal(t, int(size), n)
		off += int64(n)
	}
}

func TestStoreClose(t *testing.T) {
	file, err := os.CreateTemp("", "store_create_test")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	s, err := newStore(file)
	require.NoError(t, err)
	_, _, err = s.Append(msg)
	require.NoError(t, err)

	file, beforeSize, err := openFile(file.Name())
	require.NoError(t, err)

	err = s.Close()
	require.NoError(t, err)

	file, afterSize, err := openFile(file.Name())
	require.NoError(t, err)

	require.True(t, afterSize > beforeSize)
}

func openFile(name string) (file *os.File, size int64, err error) {
	fileFlag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	file, err = os.OpenFile(name, fileFlag, 0644)
	if err != nil {
		return nil, 0, err
	}
	fileInfo, err := os.Stat(name)
	if err != nil {
		return nil, 0, err
	}
	return file, fileInfo.Size(), nil
}
