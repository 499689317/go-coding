package coding

const HEADER_LEN = 10

type Messager interface {
	Type() int
}

type Message struct {
	Size uint16
	Version uint16
	Server uint16
	TimeStamp uint32
	Body []byte
}

func (m *Message) Type() int {
	return 1
}

func (m *Message) totalSize() int {
	return HEADER_LEN + len(m.Body)
}