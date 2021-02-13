package coding

import (
	"testing"
)

func TestReader(t *testing.T) {
	writer := NewWriter()
	s := "new writer sakldfshfjakshfjkahslakhsdjfklashflakshjkalsgdkalsgklahsf"
	writer.WriteString(s)
	bs := []byte(s)
	writer.Write(bs)
	
	var x int16 = -100
	var y int32 = 88888
	var z int64 = -1000000
	writer.WriteInt16(x)
	writer.WriteInt32(y)
	writer.WriteInt64(z)

	bt := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	writer.Write(bt)

	reader := NewReader(writer.Bytes())
	t.Log(reader.ReadString(len(s)))
	t.Log(reader.ReadString(len(s)))
	t.Log(reader.ReadInt16())
	t.Log(reader.ReadInt32())
	t.Log(reader.ReadInt64())
	t.Log(reader.Bytes())
}

func BenchmarkReader(b *testing.B) {
	
}