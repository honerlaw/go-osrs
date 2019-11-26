package io

import (
	"encoding/binary"
	"errors"
	"io"
)

var BIT_MASK = []int32{
	0, 0x1, 0x3, 0x7, 0xf, 0x1f, 0x3f, 0x7f, 0xff, 0x1ff, 0x3ff, 0x7ff, 0xfff,
	0x1fff, 0x3fff, 0x7fff, 0xffff, 0x1ffff, 0x3ffff, 0x7ffff, 0xfffff, 0x1fffff,
	0x3fffff, 0x7fffff, 0xffffff, 0x1ffffff, 0x3ffffff, 0x7ffffff, 0xfffffff,
	0x1fffffff, 0x3fffffff, 0x7fffffff, -1,
};

type Buffer struct {
	internal            []byte // the internal byte array this buffer is operating on
	length              uint32 // how many bytes of are actually in the buffer (vs capacity)
	index               uint32 // the location of the reader index
	isCompacted         bool
	isBitMode           bool   // whether we are trying to read / write bits
	bitIndex            uint32 // the current bitIndex to write to
	variableOpcodeIndex uint32 // the index in the buffer where we write the variable packet length
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		internal:            make([]byte, size),
		length:              0,
		index:               0,
		isCompacted:         false,
		isBitMode:           false,
		bitIndex:            0,
		variableOpcodeIndex: 0,
	}
}

func (b *Buffer) AsByteArray() []byte {
	return b.internal
}

func (b *Buffer) Length() uint32 {
	return b.length
}

func (b *Buffer) Capacity() int {
	return cap(b.internal)
}

func (b *Buffer) Compact() {
	b.isCompacted = true
}

func (b *Buffer) Remaining() uint32 {
	return b.length - b.index
}

func (b *Buffer) EnableBitMode() {
	b.bitIndex = b.index * 8;
	b.isBitMode = true;
}

func (b *Buffer) DisableBitMode() {
	b.index = (b.bitIndex + 7) / 8
	b.isBitMode = false
}

// reads data into the buffer
func (b *Buffer) Read(reader io.Reader) (error) {
	if b.isCompacted {
		b.isCompacted = false

		// we want to read into the buffer from where it last was,
		var slice = b.internal[b.index:]
		var length, err = reader.Read(slice)
		b.index = 0                // start the index over, so reading starts from beginning again
		b.length += uint32(length) // increment the number of bytes read
		return err
	}
	var length, err = reader.Read(b.internal)
	b.index = 0
	b.length = uint32(length)
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

func (b *Buffer) WriteBits(count byte, value int32) {
	var bytePos = b.bitIndex >> 3
	var bitOffset = byte(8 - (b.bitIndex & 7))
	b.bitIndex = b.bitIndex + uint32(count)

	for ; count > bitOffset; bitOffset = 8 {
		var tmp = int32(b.internal[bytePos])
		tmp &= ^BIT_MASK[bitOffset]
		tmp |= value >> (count - bitOffset) & BIT_MASK[bitOffset]
		b.internal[bytePos+1] = byte(tmp)
		count -= bitOffset
	}
	if count == bitOffset {
		var tmp = int32(b.internal[bytePos])
		tmp &= ^BIT_MASK[bitOffset]
		tmp |= value & BIT_MASK[bitOffset]
		b.internal[bytePos] = byte(tmp)
	} else {
		var tmp = int32(b.internal[bytePos])
		tmp &= ^(BIT_MASK[count] << (bitOffset - count))
		tmp |= (value & BIT_MASK[count]) << (bitOffset - count)
		b.internal[bytePos] = byte(tmp)
	}
}

func (b *Buffer) WriteByte(value byte) {
	b.internal[b.index] = value
	b.index += 1
}

func (b *Buffer) WriteByteA(value byte) {
	b.internal[b.index] = value + 128
	b.index += 1
}

func (b *Buffer) WriteByteC(value byte) {
	b.internal[b.index] = -value
	b.index += 1
}

func (b *Buffer) WriteByteS(value byte) {
	b.internal[b.index] = 128 - value
	b.index += 1
}

func (b *Buffer) WriteLEShort(value int16) {
	var slice = b.internal[b.index : b.index+2]
	binary.LittleEndian.PutUint16(slice, uint16(value))
	b.index += 2
}

func (b *Buffer) WriteBEShort(value int16) {
	var slice = b.internal[b.index : b.index+2]
	binary.BigEndian.PutUint16(slice, uint16(value))
	b.index += 2
}

func (b *Buffer) WriteBEShortA(value int16) {
	b.WriteByte(byte(value >> 8))
	b.WriteByteA(byte(value))
}

func (b *Buffer) WriteLEShortA(value int16) {
	b.WriteByteA(byte(value))
	b.WriteByte(byte(value >> 8))
}

func (b *Buffer) WriteLong(value int64) {
	var slice = b.internal[b.index : b.index+8]
	binary.LittleEndian.PutUint64(slice, uint64(value))
	b.index += 8
}

func (b *Buffer) WriteOpcode(opcode uint32, offset uint32) {
	b.WriteByte(byte(opcode + offset))
}

func (b *Buffer) WriteVariableShortOpcode(opcode uint32, offset uint32) {
	b.WriteOpcode(opcode, offset)
	b.variableOpcodeIndex = b.index
	b.WriteLEShort(0)
}

func (b *Buffer) WriteVariableShortLength() {
	var slice = b.internal[b.variableOpcodeIndex : b.variableOpcodeIndex+2]
	binary.LittleEndian.PutUint16(slice, uint16(b.index))
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
	b.index = b.index + uint32(endIndex) + 1

	return string(slice[:endIndex]), nil
}
