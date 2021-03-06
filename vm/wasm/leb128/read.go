package leb128

import (
	"io"
)

func ReadVarUint32Size(r io.Reader) (res uint32, size uint, err error) {
	b := make([]byte, 1)
	var shift uint
	for {
		if _, err = io.ReadFull(r, b); err != nil {
			return
		}

		size++

		cur := uint32(b[0])
		res |= (cur & 0x7f) << (shift)
		if cur&0x80 == 0 {
			return res, size, nil
		}
		shift += 7
	}
}

func ReadVarUint32(r io.Reader) (uint32, error) {
	n, _, err := ReadVarUint32Size(r)
	return n, err
}

func ReadVarint32Size(r io.Reader) (res int32, size uint, err error) {
	res64, size, err := ReadVarint64Size(r)
	res = int32(res64)
	return
}

func ReadVarint32(r io.Reader) (int32, error) {
	n, _, err := ReadVarint32Size(r)
	return n, err
}

func ReadVarint64Size(r io.Reader) (res int64, size uint, err error) {
	var shift uint
	var sign int64 = -1
	b := make([]byte, 1)

	for {
		if _, err = io.ReadFull(r, b); err != nil {
			return
		}
		size++

		cur := int64(b[0])
		res |= (cur & 0x7f) << shift
		shift += 7
		sign <<= 7
		if cur&0x80 == 0 {
			break
		}
	}

	if ((sign >> 1) & res) != 0 {
		res |= sign
	}
	return res, size, nil
}

func ReadVarint64(r io.Reader) (int64, error) {
	n, _, err := ReadVarint64Size(r)
	return n, err
}
