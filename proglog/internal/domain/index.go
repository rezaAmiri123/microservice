package domain

import (
	"github.com/tysonmote/gommap"
	"io"
	"os"
)

var (
	OffWidth uint64 = 4
	PosWidth uint64 = 8
	EntWidth        = OffWidth + PosWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func newIndex(file *os.File, c Config) (*index, error) {
	idx := &index{file: file}
	fileInfo, err := os.Stat(file.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fileInfo.Size())
	err = os.Truncate(file.Name(), int64(c.Segment.MaxIndexBytes))
	if err != nil {
		return nil, err
	}
	protFlag := gommap.PROT_READ | gommap.PROT_WRITE
	idx.mmap, err = gommap.Map(idx.file.Fd(), protFlag, gommap.MAP_SHARED)
	if err != nil {
		return nil, err
	}

	return idx, nil
}

func (i *index) Close() error {
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}
	if err := i.file.Sync(); err != nil {
		return err
	}
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}
	return i.file.Close()
}

// we want offset and position
func (i *index) Read(input int64) (output uint32, pos uint64, err error) {
	if i.size == 0 {
		return 0, 0, io.EOF
	}
	if input == -1 {
		// we want the last record
		output = uint32((i.size / EntWidth) - 1)
	} else {
		output = uint32(input)
	}
	pos = uint64(output) * EntWidth
	if i.size < pos+EntWidth {
		return 0, 0, io.EOF
	}
	output = Enc.Uint32(i.mmap[pos : pos+OffWidth])
	pos = Enc.Uint64(i.mmap[pos+OffWidth : pos+EntWidth])

	return output, pos, nil
}

func (i *index) Write(offset uint32, pos uint64) error {
	if uint64(len(i.mmap)) < i.size+EntWidth {
		return io.EOF
	}
	Enc.PutUint32(i.mmap[i.size:i.size+OffWidth], offset)
	Enc.PutUint64(i.mmap[i.size+OffWidth:i.size+EntWidth], pos)
	i.size += EntWidth
	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}
