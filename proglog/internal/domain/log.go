package domain

import (
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/stackus/errors"
)

var ErrOffsetOutOfRange = errors.Wrap(errors.ErrBadRequest, "offset out of range")

type (
	Log interface {
		Append(record *Record) (uint64, error)
		Read(uint642 uint64) (*Record, error)
	}
	originReader struct {
		*store
		off int64
	}
	log struct {
		mu sync.RWMutex

		Dir    string
		Config Config

		activeSegment *segment
		segments      []*segment
	}
)

func NewLog(dir string, c Config) (*log, error) {
	if c.Segment.MaxStoreBytes == 0 {
		c.Segment.MaxStoreBytes = 1024
	}
	if c.Segment.MaxIndexBytes == 0 {
		c.Segment.MaxIndexBytes = 1024
	}
	l := &log{Dir: dir, Config: c}

	return l, l.setup()
}
func (l *log) setup() error {
	files, err := os.ReadDir(l.Dir)
	if err != nil {
		return err
	}
	var baseOffsets []uint64
	for _, file := range files {
		offStr := strings.TrimSuffix(file.Name(), path.Ext(file.Name()))
		off, err := strconv.ParseUint(offStr, 10, 0)
		if err != nil {
			return err
		}
		baseOffsets = append(baseOffsets, off)
	}

	sort.Slice(baseOffsets, func(i, j int) bool {
		return baseOffsets[i] < baseOffsets[j]
	})
	for i := 0; i < len(baseOffsets); i++ {
		if err = l.newSegment(baseOffsets[i]); err != nil {
			return err
		}
		// baseOffset contains dup for index and store so we skip
		// the dup
		i++
	}
	if l.segments == nil {
		if err = l.newSegment(l.Config.Segment.InitialOffset); err != nil {
			return err
		}
	}
	return nil
}

func (l *log) Append(record *Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	offset, err := l.activeSegment.Append(record)
	if err != nil {
		return 0, err
	}
	if l.activeSegment.IsMaxed() {
		err = l.newSegment(offset + 1)
	}
	return offset, err
}

func (l *log) Read(offset uint64) (*Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var seg *segment
	for _, segment := range l.segments {
		if segment.baseOffset <= offset && offset < segment.nextOffset {
			seg = segment
			break
		}
	}
	if seg == nil || seg.nextOffset <= offset {
		return nil, ErrOffsetOutOfRange
	}

	return seg.Read(offset)
}

func (l *log) newSegment(baseOffset uint64) error {
	seg, err := newSegment(l.Dir, baseOffset, l.Config)
	if err != nil {
		return err
	}
	l.segments = append(l.segments, seg)
	l.activeSegment = seg

	return nil
}

func (l *log) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, segment := range l.segments {
		if err := segment.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (l *log) Remove() error {
	if err := l.Close(); err != nil {
		return err
	}
	return os.RemoveAll(l.Dir)
}

func (l *log) Reset() error {
	if err := l.Remove(); err != nil {
		return err
	}
	return l.setup()
}

func (l *log) LowestOffset() (uint64, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.segments[0].baseOffset, nil
}

func (l *log) HighestOffset() (uint64, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	off := l.segments[len(l.segments)-1].nextOffset
	if off == 0 {
		return 0, nil
	}
	return off - 1, nil
}

func (l *log) Truncate(lowest uint64) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var segments []*segment
	for _, seg := range l.segments {
		if seg.nextOffset <= lowest+1 {
			if err := seg.Remove(); err != nil {
				return err
			}
			continue
		}
		segments = append(segments, seg)
	}
	l.segments = segments
	return nil
}

func (l *log) Reader() io.Reader {
	l.mu.RLock()
	defer l.mu.RUnlock()

	readers := make([]io.Reader, len(l.segments))
	for i, seg := range l.segments {
		readers[i] = &originReader{seg.store, 0}
	}
	return io.MultiReader(readers...)
}

func (o *originReader) Read(data []byte) (int, error) {
	n, err := o.ReadAt(data, o.off)
	o.off += int64(n)
	return n, err
}
