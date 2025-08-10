package client

type Event interface {
}

type Proxy interface {
	Context() *Context
	ID() uint32
	Dispatch(opcode uint32, fd int, data []byte) Event
}
