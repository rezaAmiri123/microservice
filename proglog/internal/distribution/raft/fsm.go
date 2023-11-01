package raft

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/raft"
	"github.com/rezaAmiri123/microservice/proglog/internal/constats"
	"github.com/rezaAmiri123/microservice/proglog/internal/domain"
	"io"
)

type (
	fsm struct {
		log domain.Log
	}
	snapshot struct {
		reader io.Reader
	}
)

var _ raft.FSM = (*fsm)(nil)

func (f *fsm) Apply(record *raft.Log) interface{} {
	buf := record.Data
	reqType := constats.RequestType(buf[0])
	switch reqType {
	case constats.AppendRequestType:
		return f.applyAppend(buf[1:])
	}
	return nil
}

func (f *fsm) applyAppend(data []byte) interface{} {
	var req domain.Record
	if err := proto.Unmarshal(data, &req); err != nil {
		return err
	}
	offset, err := f.log.Append(&req)
	if err != nil {
		return err
	}
	return &domain.Record{Offset: offset}
}

func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	reader := f.log.Reader()
	return &snapshot{reader: reader}, nil
}
func (f *fsm) Restore(reader io.ReadCloser) error {
	if err := f.log.Reset(); err != nil {
		return err
	}
	lenWidth := make([]byte, domain.LenWidth)
	var buf bytes.Buffer
	for {
		_, err := io.ReadFull(reader, lenWidth)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		size := int64(domain.Enc.Uint64(lenWidth))
		if _, err := io.CopyN(&buf, reader, size); err != nil {
			return err
		}

		record := &domain.Record{}
		if err = proto.Unmarshal(buf.Bytes(), record); err != nil {
			return err
		}
		if _, err = f.log.Append(record); err != nil {
			return err
		}
		buf.Reset()
	}
	return nil
}

var _ raft.FSMSnapshot = (*snapshot)(nil)

func (s *snapshot) Persist(sink raft.SnapshotSink) error {
	if _, err := io.Copy(sink, s.reader); err != nil {
		_ = sink.Cancel()
	}
	return sink.Close()
}

func (s *snapshot) Release() {}
