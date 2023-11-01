package domain

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestSegment(t *testing.T) {
	dir, err := os.MkdirTemp("", "segment-test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	want := &Record{Value: []byte("hello world")}

	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = EntWidth * 3

	seq, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), seq.nextOffset)
	require.False(t, seq.IsMaxed())

	for i := uint64(0); i < 3; i++ {
		off, err := seq.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)
		got, err := seq.Read(off)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}

	_, err = seq.Append(want)
	require.Equal(t, io.EOF, err)

	// maxed index
	require.True(t, seq.IsMaxed())

	c.Segment.MaxStoreBytes = uint64(len(want.Value) * 3)
	c.Segment.MaxIndexBytes = 1024

	seq, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	// maxed store
	require.True(t, seq.IsMaxed())

	err = seq.Remove()
	require.NoError(t, err)
	seq, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.False(t, seq.IsMaxed())
}
