package client

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Context struct {
	conn      *net.UnixConn
	connMu    sync.RWMutex
	objects   map[uint32]Proxy
	mu        sync.Mutex
	wLock     sync.Mutex
	currentID uint32
	EventC    chan Event
}

func (ctx *Context) Register(p Proxy) uint32 {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.currentID++
	ctx.objects[ctx.currentID] = p
	return ctx.currentID
}

func (ctx *Context) Unregister(p Proxy) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	delete(ctx.objects, p.ID())
}

func (ctx *Context) GetProxy(id uint32) Proxy {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.objects[id]
}

func (ctx *Context) Close() error {
	return ctx.conn.Close()
}

func (ctx *Context) pullEvents() {
	for {
		fmt.Println("pull")
		senderID, opcode, fd, data, err := ctx.ReadMsg()
		if err != nil {
			log.Printf("ctx.Dispatch: unable to read msg: %v", err)
			continue
		}

		fmt.Println("pull a")
		ctx.mu.Lock()
		sender, ok := ctx.objects[senderID]
		ctx.mu.Unlock()
		if !ok {
			log.Printf("ctx.Dispatch: unable find sender (senderID=%d)", senderID)
			continue
		}
		fmt.Println("pull b")

		evt := sender.Dispatch(opcode, fd, data)
		if evt != nil {
			ctx.EventC <- evt
		}
		fmt.Println("pull done")
	}
}

func Connect(addr string) (*Context, error) {
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

	ctx := &Context{
		objects: make(map[uint32]Proxy),
		EventC:  make(chan Event),
	}

	var err error
	ctx.conn, err = net.DialUnix("unix", nil, &net.UnixAddr{Name: addr, Net: "unix"})
	if err != nil {
		return nil, err
	}

	go ctx.pullEvents()

	return ctx, nil
}
