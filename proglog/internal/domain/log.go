package domain

type (
	Record struct {
		Value  []byte
		Offset uint64
	}
	Log interface {
		Append(record Record) (uint64, error)
		Read(uint642 uint64) (Record, error)
	}
)
