package coding

import (
	"errors"
	"time"
	"fmt"
)

type Coding struct {
	bytes []byte
}

func NewCoding() *Coding {
	c := new(Coding)
	return c
}

func (c *Coding) toMessage(msg Messager) *Message {
	for {
		if msg == nil {
			break
		}
		m, ok := msg.(*Message)
		if !ok {
			break
		}
		return m
	}
	return nil
}

func (c *Coding) Bytes() []byte {
	return c.bytes
}

// 编码
func (c *Coding) Encode(msg Messager) ([]byte, error) {
	m := c.toMessage(msg)
	if m == nil {
		return nil, errors.New("msg toMessage failed")
	}
	w := NewWriter()
	m.TimeStamp = uint32(time.Now().Unix())
	m.Size = uint16(len(m.Body))

	w.WriteUint16(m.Size)
	w.WriteUint16(m.Version)
	w.WriteUint16(m.Server)
	w.WriteUint32(m.TimeStamp)
	w.Write(m.Body)
	return w.Bytes(), nil
}
// 解码
func (c *Coding) Redecode(bytes []byte) ([]Messager, error) {
	ms := []Messager{}
	l := len(c.bytes)
	b := len(bytes)
	bs := make([]byte, l+b)
	copy(bs, c.bytes)
	copy(bs[l:], bytes)
	c.bytes = bs
	for {
		fmt.Println(c.bytes)
		m, e := c.Decode(c.bytes)
		if e != nil {
			break
		}
		if m == nil {
			break
		}
		ms = append(ms, m)
	}
	return ms, nil
}
func (c *Coding) Decode(bytes []byte) (Messager, error) {
	l := len(bytes)
	if l < HEADER_LEN {
		return nil, nil
	}
	r := NewReader(bytes)
	// Body Size
	b, e := r.ReadUint16()
	if e != nil {
		return nil, e
	}
	// 判断是不是完整的消息
	if l < HEADER_LEN + int(b) {
		return nil, nil
	}
	// Version
	v, e := r.ReadUint16()
	if e != nil {
		return nil, e
	}
	// Server
	s, e := r.ReadUint16()
	if e != nil {
		return nil, e
	}
	// Time
	t, e := r.ReadUint32()
	if e != nil {
		return nil, e
	}
	// Body
	buf := make([]byte, b)
	_, e = r.Read(buf, int(b))
	if e != nil {
		return nil, e
	}

	m := &Message{
		Size: uint16(b),
		Version: uint16(v),
		Server: uint16(s),
		TimeStamp: uint32(t),
		Body: buf,
	}

	c.bytes = r.Bytes()

	return m, nil
}