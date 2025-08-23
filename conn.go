package wayland

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

type Conn struct {
	*net.UnixConn
	regLock   sync.Mutex
	writeLock sync.Mutex
	objects   []Proxy
	drain     chan<- Event
}

func (conn *Conn) SetDrain(ch chan<- Event) {
	conn.drain = ch
}

func (conn *Conn) Register(p Proxy) {
	conn.regLock.Lock()
	defer conn.regLock.Unlock()
	for i := 0; i < len(conn.objects); i++ {
		if conn.objects[i] == nil {
			conn.objects[i] = p
			p.Register(conn, uint32(i+1))
			return
		}
	}
	conn.objects = append(conn.objects, p)
	p.Register(conn, uint32(len(conn.objects)))
}

func (conn *Conn) Unregister(p Proxy) {
	conn.regLock.Lock()
	defer conn.regLock.Unlock()
	conn.objects[p.ID()-1] = nil
}

func (conn *Conn) GetProxy(id uint32) Proxy {
	conn.regLock.Lock()
	defer conn.regLock.Unlock()
	return conn.objects[id-1]
}

func (conn *Conn) pullEvents() {
	for {
		msg, err := conn.ReadMsg()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Printf("conn.Dispatch: unable to read msg: %v", err)
			if errors.Is(err, io.EOF) {
				return
			}
			continue
		}

		conn.regLock.Lock()
		if msg.SenderID > uint32(len(conn.objects)) || conn.objects[msg.SenderID-1] == nil {
			log.Printf("conn.Dispatch: unable find sender (senderID=%d)", msg.SenderID)
			conn.regLock.Unlock()
			continue
		}
		sender := conn.objects[msg.SenderID-1]
		conn.regLock.Unlock()

		sender.Dispatch(msg, conn.drain)
	}
}

func Connect(addr string) (*Conn, error) {
	if addr == "" {
		runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
		if runtimeDir == "" {
			return nil, errors.New("env XDG_RUNTIME_DIR not set")
		}
		addr = os.Getenv("WAYLAND_DISPLAY")
		if addr == "" {
			addr = "wayland-0"
		}
		addr = runtimeDir + "/" + addr
	}

	conn := &Conn{}

	var err error
	conn.UnixConn, err = net.DialUnix("unix", nil, &net.UnixAddr{Name: addr, Net: "unix"})
	if err != nil {
		return nil, err
	}

	go conn.pullEvents()

	return conn, nil
}
