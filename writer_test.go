package coding

import (
	"testing"
)

func TestWriter(t *testing.T) {
	writer := NewWriter()

	s := "new writer sakldfshfjakshfjkahslakhsdjfklashflakshjkalsgdkalsgklahsf"
	writer.WriteString(s)

	bs := []byte(s)
	writer.Write(bs)

	bt := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	writer.Write(bt)

	var x int16 = 100
	var y int32 = 88888
	var z int64 = 1000000

	writer.WriteInt16(x)
	writer.WriteInt32(y)
	writer.WriteInt64(z)

	t.Log(writer.Bytes())
}

func BenchmarkWriter(b *testing.B) {
	
}