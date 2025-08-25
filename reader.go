package wayland

import (
	"errors"
	"fmt"
	"io"
	"syscall"
	"unsafe"
)

var oobSpace = syscall.CmsgSpace(4 * 8)

type Message struct {
	SenderID uint32
	Opcode   uint16
	Size     uint16 // total bytes incl header
	Payload  []byte
	FDs      []int
}

// ReadMsg reads one Wayland message (header + body) and any SCM_RIGHTS FDs.
func (conn *Conn) ReadMsg() (*Message, error) {
	// 1) header (8 bytes)
	header := make([]byte, 8)
	oob := make([]byte, oobSpace)

	n, oobn, flags, _, err := conn.ReadMsgUnix(header, oob)
	if err != nil {
		return nil, err
	}
	if n != 8 {
		return nil, fmt.Errorf("ReadMsg: short header read n=%d", n)
	}
	if (flags & syscall.MSG_CTRUNC) != 0 {
		return nil, errors.New("ReadMsg: control data truncated (MSG_CTRUNC)")
	}

	fds, err := parseRights(oob[:oobn])
	if err != nil {
		return nil, fmt.Errorf("ReadMsg: %w", err)
	}

	// native-endian decode
	_ = header[7]
	word0 := *(*uint32)(unsafe.Pointer(&header[0])) // sender id
	word1 := *(*uint32)(unsafe.Pointer(&header[4])) // (size<<16)|opcode

	size := uint16(word1 >> 16)
	opcode := uint16(word1 & 0xffff)

	if size < 8 {
		return nil, fmt.Errorf("ReadMsg: invalid size=%d (<8)", size)
	}
	bodyLen := int(size) - 8
	msg := &Message{
		SenderID: word0,
		Opcode:   opcode,
		Size:     size,
	}

	if bodyLen == 0 {
		return msg, nil
	}

	// 2) payload in a loop (stream sockets may return partial)
	msg.Payload = make([]byte, bodyLen)
	read := 0
	for read < bodyLen {
		oob = make([]byte, oobSpace) // fresh buffer each recv
		n, oobn, flags, _, err = conn.ReadMsgUnix(msg.Payload[read:], oob)
		if err != nil {
			if errors.Is(err, io.EOF) && read > 0 {
				return nil, fmt.Errorf("ReadMsg: unexpected EOF after %d/%d bytes", read, bodyLen)
			}
			return nil, err
		}
		if (flags & syscall.MSG_CTRUNC) != 0 {
			return nil, errors.New("ReadMsg: control data truncated (MSG_CTRUNC) during payload")
		}
		read += n
		if oobn > 0 {
			more, err := parseRights(oob[:oobn])
			if err != nil {
				return nil, fmt.Errorf("ReadMsg: %w", err)
			}
			fds = append(fds, more...)
		}
	}
	msg.FDs = fds
	return msg, nil
}

func parseRights(oob []byte) ([]int, error) {
	if len(oob) == 0 {
		return nil, nil
	}
	scms, err := syscall.ParseSocketControlMessage(oob)
	if err != nil {
		return nil, fmt.Errorf("parseRights: %w", err)
	}
	var all []int
	for _, scm := range scms {
		fds, err := syscall.ParseUnixRights(&scm)
		if err != nil {
			return nil, fmt.Errorf("parseRights: %w", err)
		}
		all = append(all, fds...)
	}
	return all, nil
}

type MessageReader struct {
	conn     *Conn
	b        []byte
	fds      []int
	overflow bool
}

func NewMessageReader(conn *Conn, m *Message) MessageReader {
	return MessageReader{conn: conn, b: m.Payload, fds: m.FDs}
}

// ReadUint reads one uint32 from the message
func (r *MessageReader) ReadUint() uint32 {
	if len(r.b) < 4 {
		r.overflow = true
		return 0
	}
	v := *(*uint32)(unsafe.Pointer(&r.b[0]))
	r.b = r.b[4:]
	return v
}

// ReadInt reads one int32 from the message
func (r *MessageReader) ReadInt() int32 {
	return int32(r.ReadUint())
}

// ReadFixed reads one fixed-point float64 from the message
func (r *MessageReader) ReadFixed() float64 {
	i := r.ReadInt()
	return float64(i) / 256.0
}

// ReadObject reads an ID registered in Conn and returns the corresponding object
func (r *MessageReader) ReadObject() Proxy {
	id := int(r.ReadUint())
	if id == 0 || id > len(r.conn.objects) {
		return nil
	}
	return r.conn.objects[id-1]
}

// ReadArray reads a bytearray from the message
func (r *MessageReader) ReadArray() []byte {
	l := int(r.ReadUint())
	if len(r.b) < l {
		r.overflow = true
		return nil
	}
	data := r.b[:l] // full slice with capped capacity
	r.b = r.b[l:]
	// pad
	pad := pad4(l) - l
	if pad > 0 {
		if len(r.b) < pad {
			r.overflow = true
			return nil
		}
		r.b = r.b[pad:]
	}
	return data
}

// ReadString reads a string from the message
func (r *MessageReader) ReadString() string {
	l := int(r.ReadUint())
	if len(r.b) < l {
		return ""
	}
	// build Go string without the NUL; this copies once.
	s := string(r.b[:l-1]) // full slice with capped capacity
	r.b = r.b[l:]
	// pad
	pad := pad4(l) - l
	if pad > 0 {
		if len(r.b) < pad {
			r.overflow = true
			return ""
		}
		r.b = r.b[pad:]
	}
	return s
}

// ReadFd reads a file descriptor from the message
func (r *MessageReader) ReadFd() (ret int) {
	if len(r.fds) == 0 {
		r.overflow = true
		return -1
	}
	ret = r.fds[0]
	r.fds = r.fds[1:]
	return
}

// HadOverflow checks if a previous call had a overflow
func (r *MessageReader) HadOverflow() bool {
	return r.overflow
}
