package io

import (
	"encoding/binary"
	"errors"
	"io"
)

type Buffer struct {
	internal    []byte // the internal byte array this buffer is operating on
	length      int    // how many bytes of are actually in the buffer (vs capacity)
	index       int    // the location of the reader index
	isCompacted bool
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		internal:    make([]byte, size),
		length:      0,
		index:       0,
		isCompacted: false,
	}
}

func (b *Buffer) AsByteArray() []byte {
	return b.internal
}

func (b *Buffer) Length() int {
	return b.length
}

func (b *Buffer) Capacity() int {
	return cap(b.internal)
}

func (b *Buffer) Compact() {
	b.isCompacted = true
}

func (b *Buffer) Remaining() int {
	return b.length - b.index
}

// reads data into the buffer
func (b *Buffer) Read(reader io.Reader) (error) {
	if b.isCompacted {
		b.isCompacted = false

		// we want to read into the buffer from where it last was,
		var slice = b.internal[b.index:]
		var length, err = reader.Read(slice)
		b.index = 0        // start the index over, so reading starts from beginning again
		b.length += length // increment the number of bytes read
		return err
	}
	var length, err = reader.Read(b.internal)
	b.index = 0
	b.length = length
	return err
}

func (b *Buffer) ReadByte() (value byte) {
	value = b.internal[b.index]
	b.index += 1
	return
}

func (b *Buffer) ReadShort() (value uint16) {
	value = binary.BigEndian.Uint16(b.internal[b.index : b.index+2])
	b.index += 2
	return
}

func (b *Buffer) ReadInt() (value uint32) {
	value = binary.LittleEndian.Uint32(b.internal[b.index : b.index+4])
	b.index += 4
	return
}

func (b *Buffer) ReadBEInt() (value uint32) {
	value = binary.BigEndian.Uint32(b.internal[b.index : b.index+4])
	b.index += 4
	return
}

func (b *Buffer) ReadLong() (value uint64) {
	value = binary.LittleEndian.Uint64(b.internal[b.index : b.index+8])
	b.index += 8
	return
}

func (b *Buffer) WriteByte(value byte) {
	b.internal[b.index] = value
	b.index += 1
}

func (b *Buffer) WriteInt(value uint32) {
	var slice = b.internal[b.index : b.index+4]
	binary.LittleEndian.PutUint32(slice, value)
	b.index += 4
}

func (b *Buffer) WriteLong(value int64) {
	var slice = b.internal[b.index : b.index+8]
	binary.LittleEndian.PutUint64(slice, uint64(value))
	b.index += 8
}

func (b *Buffer) ReadRSString() (string, error) {

	var endIndex = -1
	var slice = b.internal[b.index:]

	// find the index in the slice of the last character
	for index, val := range slice {
		if val == 10 {
			endIndex = index
			break
		}
	}

	if endIndex == -1 {
		return "", errors.New("Failed to read string")
	}

	// endIndex is relative to the slice (which starts at b.index), so add it back and then one more to skip the
	// previous string's ending 10 byte
	b.index = b.index + endIndex + 1

	return string(slice[:endIndex]), nil
}