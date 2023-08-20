package domain

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"path"
)

type segment struct {
	store                  *store
	index                  *index
	baseOffset, nextOffset uint64
	config                 Config
}

func newSegment(dir string, baseOffset uint64, config Config) (*segment, error) {
	s := &segment{
		baseOffset: baseOffset,
		config:     config,
	}
	var err error
	storeFileName := path.Join(dir, fmt.Sprintf("%d:%s", baseOffset, ".store"))
	storeFileFlag := os.O_RDWR | os.O_APPEND | os.O_CREATE
	storeFile, err := os.OpenFile(storeFileName, storeFileFlag, 0644)
	if err != nil {
		return nil, err
	}

	if s.store, err = newStore(storeFile); err != nil {
		return nil, err
	}

	indexFileName := path.Join(dir, fmt.Sprintf("%d:%s", baseOffset, ".index"))
	indexFileFlag := os.O_RDWR | os.O_CREATE
	indexFile, err := os.OpenFile(indexFileName, indexFileFlag, 0644)
	if err != nil {
		return nil, err
	}

	if s.index, err = newIndex(indexFile, config); err != nil {
		return nil, err
	}

	if offset, _, err := s.index.Read(-1); err != nil {
		s.nextOffset = baseOffset
	} else {
		s.nextOffset = baseOffset + uint64(offset) + 1
	}

	return s, nil
}

func (s *segment) Append(record *Record) (uint64, error) {
	cur := s.nextOffset
	record.Offset = cur
	recordBytes, err := proto.Marshal(record)

	_, pos, err := s.store.Append(recordBytes)
	if err != nil {
		return 0, err
	}

	// index offsets are relative to base offset
	offset := s.nextOffset - s.baseOffset
	if err = s.index.Write(uint32(offset), pos); err != nil {
		return 0, err
	}
	s.nextOffset++
	return cur, nil
}

func (s *segment) Read(offset uint64) (*Record, error) {
	_, pos, err := s.index.Read(int64(offset - s.baseOffset))
	if err != nil {
		return nil, err
	}

	data, err := s.store.Read(pos)
	if err != nil {
		return nil, err
	}
	record := &Record{}
	err = proto.Unmarshal(data, record)
	return record, err
}

func (s *segment) IsMaxed() bool {
	maxStore := s.store.size >= s.config.Segment.MaxStoreBytes
	maxIndex := s.index.size >= s.config.Segment.MaxIndexBytes

	return maxIndex || maxStore
}
func (s *segment) Close() error {
	if err := s.index.Close(); err != nil {
		return err
	}
	if err := s.store.Close(); err != nil {
		return err
	}
	return nil
}

func (s *segment) Remove() error {
	if err := s.Close(); err != nil {
		return err
	}
	if err := os.Remove(s.index.Name()); err != nil {
		return err
	}
	if err := os.Remove(s.store.Name()); err != nil {
		return err
	}
	return nil
}
