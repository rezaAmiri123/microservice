package raft

import (
	"github.com/hashicorp/raft"
	"github.com/rezaAmiri123/microservice/proglog/internal/domain"
)

var _ raft.LogStore = (*logStore)(nil)

type logStore struct {
	domain.Log
}

func newLogStore(dir string, config domain.Config) (*logStore, error) {
	log, err := domain.NewLog(dir, config)
	if err != nil {
		return nil, err
	}
	return &logStore{Log: log}, nil
}

func (l *logStore) FirstIndex() (uint64, error) {
	return l.LowestOffset()
}
func (l *logStore) LastIndex() (uint64, error) {
	offset, err := l.HighestOffset()
	return offset, err
}
func (l *logStore) GetLog(index uint64, outLog *raft.Log) error {
	record, err := l.Read(index)
	if err != nil {
		return err
	}

	outLog.Data = record.Value
	outLog.Index = record.Offset
	outLog.Type = raft.LogType(record.Type)
	outLog.Term = record.Term

	return nil
}

func (l *logStore) StoreLog(log *raft.Log) error {
	return l.StoreLogs([]*raft.Log{log})
}
func (l *logStore) StoreLogs(logs []*raft.Log) error {
	for _, log := range logs {
		if _, err := l.Append(&domain.Record{
			Value: log.Data,
			Term:  log.Term,
			Type:  uint32(log.Type),
		}); err != nil {
			return err
		}
	}
	return nil
}
func (l *logStore) DeleteRange(min, max uint64) error {
	return l.Truncate(max)
}
