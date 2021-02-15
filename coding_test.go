package coding

import (
	"testing"
)

func TestCoding(t *testing.T) {
	c := NewCoding()

	str := "ad2fkaAaldakllkjflkafdlfjdsfhlakjsdfhlkafjka"
	m := &Message{
		Version: 1,
		Server: 1,
		Body: []byte(str),
	}

	// 编码
	buf, e := c.Encode(m)
	if e != nil {
		t.Error(e)
	}
	
	tbuf := make([]byte, len(buf))
	copy(tbuf, buf)
	copy(tbuf, []byte(str))

	l2, e := c.Redecode(tbuf)
	if e != nil {
		t.Error(e)
	}
	for _, d := range l2 {
		t.Log("l2: ", d.(*Message).Size, d.(*Message).Version, d.(*Message).Server, d.(*Message).TimeStamp, d.(*Message).Body)
	}
	
	t.Log("coding byte", c.Bytes())

	str2 := "iiiisssssssssssssssssssssssssssssssssssssssssssss"
	m2 := &Message{
		Version: 2,
		Server: 2,
		Body: []byte(str2),
	}

	buf2, e := c.Encode(m2)
	if e != nil {
		t.Error(e)
	}

	b := make([]byte, len(buf) + len(buf2))
	copy(b, buf)
	copy(b[len(buf):], buf2)

	// 解码
	l, e := c.Redecode(b)
	if e != nil {
		t.Error(e)
	}
	for _, d := range l {
		v := d.(*Message)
		t.Log("Size: ", v.Size, "Version: ", v.Version, "Server: ", v.Server, "TimeStamp: ", v.TimeStamp, "Body: ", v.Body)
	}

	t.Log("coding byte: ", c.Bytes())
}