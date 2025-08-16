package wayland

import (
	"errors"
	"log"
	"net"
	"os"
	"sync"
)

type Conn struct {
	sock      *net.UnixConn
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

func (conn *Conn) Close() error {
	return conn.sock.Close()
}

func (conn *Conn) pullEvents() {
	for {
		senderID, opcode, fd, data, err := conn.ReadMsg()
		if err != nil {
			log.Printf("conn.Dispatch: unable to read msg: %v", err)
			continue
		}

		conn.regLock.Lock()
		if senderID > uint32(len(conn.objects)) || conn.objects[senderID-1] == nil {
			log.Printf("conn.Dispatch: unable find sender (senderID=%d)", senderID)
			conn.regLock.Unlock()
			continue
		}
		sender := conn.objects[senderID-1]
		conn.regLock.Unlock()

		sender.Dispatch(opcode, fd, data, conn.drain)
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
	conn.sock, err = net.DialUnix("unix", nil, &net.UnixAddr{Name: addr, Net: "unix"})
	if err != nil {
		return nil, err
	}

	go conn.pullEvents()

	return conn, nil
}
