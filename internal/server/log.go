package server

import (
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)

	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset > uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}

	return c.records[offset], nil
}

type Record struct {
	Offset uint64
	Value  []byte
}

type logError string

func (e logError) Error() string {
	return string(e)
}

const ErrOffsetNotFound = logError("offset not found")
