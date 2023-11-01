package domain

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"sync"
)

const (
	LenWidth = 8
)

var (
	Enc = binary.BigEndian
)

type store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}

var _ interface {
	io.ReaderAt
	io.Closer
} = (*store)(nil)

func newStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	s := &store{
		File: f,
		size: uint64(fi.Size()),
		buf:  bufio.NewWriter(f),
		//mu:   sync.Mutex{},
	}

	return s, nil
}

func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return 0, err
	}

	return s.File.ReadAt(p, off)
}

func (s *store) Append(p []byte) (n, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	pos = s.size
	if err = binary.Write(s.buf, Enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}

	byteCount, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}

	byteCount += LenWidth
	s.size += uint64(byteCount)

	return uint64(byteCount), pos, nil
}

func (s *store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Flush writes any buffered data to the underlying io.Writer.
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	posBytes := make([]byte, LenWidth)
	if _, err := s.File.ReadAt(posBytes, int64(pos)); err != nil {
		return nil, err
	}

	dataBytes := make([]byte, Enc.Uint64(posBytes))
	if _, err := s.File.ReadAt(dataBytes, int64(pos+LenWidth)); err != nil {
		return nil, err
	}

	return dataBytes, nil
}

func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return err
	}

	return s.File.Close()
}
