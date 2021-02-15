package coding

import (
	//"errors"
	"unsafe"
	"encoding/binary"
)

const MAX_CAP = 128

type Writer struct {
	cursor int
	bytes []byte
}

// 字节暂存区，编解码字节数据源
// 网卡单次最大可通过约1600个字节
func NewWriter() *Writer {
	return &Writer{0, make([]byte, MAX_CAP)}
}

func (w *Writer) ResetCursor() {
	w.cursor = 0
}
// 未使用空间长度
func (w *Writer) UnwriteLen() int {
	if len(w.bytes) <= w.cursor {
		return 0
	}
	return len(w.bytes) - w.cursor
}

// 已使用空间长度
func (w *Writer) WriteLen() int {
	return w.cursor
}

func (w *Writer) Bytes() []byte {
	return w.bytes[:w.cursor]
}

// 自动扩容
func (w *Writer) apply(l int) {
	a := (l/MAX_CAP + 1) * MAX_CAP
	bytes := make([]byte, len(w.bytes)+a)
	copy(bytes, w.bytes[:w.cursor])
	w.bytes = bytes
}

func (w *Writer) Write(buf []byte) (int, error) {
	l := len(buf)
	wl := w.UnwriteLen()
	if wl < l {
		w.apply(l-wl)
	}
	n := copy(w.bytes[w.cursor:], buf)
	w.cursor += n
	return n, nil
}

func (w *Writer) WriteString(s string) (int, error) {
	return w.Write(*(*[]byte)(unsafe.Pointer(&s)))
}

func (w *Writer) WriteUint16(n uint16) (int, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(n))
	return w.Write(buf[:2])
}

func (w *Writer) WriteUint32(n uint32) (int, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(n))
	return w.Write(buf[:4])
}

func (w *Writer) WriteUint64(n uint64) (int, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, n)
	return w.Write(buf)
}
