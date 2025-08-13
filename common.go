package wayland

type Event interface {
	Proxy() Proxy
}

type Proxy interface {
	Register(*Conn, uint32)
	Conn() *Conn
	ID() uint32

	Dispatch(opcode uint32, fd int, data []byte)
	Name() string
}

type BaseProxy struct {
	conn *Conn
	id   uint32
}

func (p *BaseProxy) Register(conn *Conn, id uint32) {
	p.conn = conn
	p.id = id
}

func (p *BaseProxy) Conn() *Conn {
	return p.conn
}

func (p *BaseProxy) ID() uint32 {
	return p.id
}
