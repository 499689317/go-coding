package coding

import (
	"errors"
	"encoding/binary"
)

// 读取tcp报文data域byte流
// Reader作为中间层起到取长补短的效果
type Reader struct {
	cursor int
	bytes []byte
}

func NewReader(buf []byte) *Reader {
	return &Reader{0, buf}
}

func (r *Reader) ResetCursor() {
	r.cursor = 0
}

// 未读取byte长度
func (r *Reader) UnreadLen() int {
	if len(r.bytes) <= r.cursor {
		return 0
	}
	return len(r.bytes) - r.cursor
}

func (r *Reader) Bytes() []byte {
	return r.bytes[r.cursor:]
}

func (r *Reader) Read(buf []byte, l int) (int, error) {
	var n int = 0
	if l == 0 || r.UnreadLen() <= l {
		n = copy(buf, r.bytes[r.cursor:])
	} else {
		n = copy(buf, r.bytes[r.cursor:r.cursor+l])
	}
	if n != l {
		return 0, errors.New("copy n != l")
	}
	r.cursor += l
	return n, nil
}

// l取决与string字节码长度，rune类型要特殊处理
func (r *Reader) ReadString(l int) (string, error) {
	if r.UnreadLen() < l {
		return "", errors.New("UnreadLen < l")
	}
	s := string(r.bytes[r.cursor:r.cursor+l])
	r.cursor += l
	return s, nil
}

// pc一般是小端读取字节数据
func (r *Reader) ReadUint16() (uint16, error) {
	if r.UnreadLen() < 2 {
		return 0, errors.New("UnreadLen < 2")
	}
	b := r.bytes[r.cursor:r.cursor+2]
	r.cursor += 2
	return binary.LittleEndian.Uint16(b), nil
}

func (r *Reader) ReadInt16() (int16, error) {
	c, e := r.ReadUint16()
	if e != nil {
		return 0, e
	}
	return int16(c), nil
}

func (r *Reader) ReadUint32() (uint32, error) {
	if r.UnreadLen() < 4 {
		return 0, errors.New("UnreadLen < 4")
	}
	b := r.bytes[r.cursor:r.cursor+4]
	r.cursor += 4
	return binary.LittleEndian.Uint32(b), nil
}

func (r *Reader) ReadInt32() (int32, error) {
	c, e := r.ReadUint32()
	if e != nil {
		return 0, e
	}
	return int32(c), nil
}

func (r *Reader) ReadUint64() (uint64, error) {
	if r.UnreadLen() < 8 {
		return 0, errors.New("UnreadLen < 8")
	}
	b := r.bytes[r.cursor:r.cursor+8]
	r.cursor += 8
	return binary.LittleEndian.Uint64(b), nil
}

func (r *Reader) ReadInt64() (int64, error) {
	c, e := r.ReadUint64()
	if e != nil {
		return 0, e
	}
	return int64(c), nil
}