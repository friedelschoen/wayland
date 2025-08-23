package wayland

import (
	"fmt"
	"math"
	"syscall"
	"unsafe"
)

type MessageWriter struct {
	sender  Proxy
	opcode  uint16
	content []byte
	fds     []int
}

func NewMessageWriter(sender Proxy, opcode int) *MessageWriter {
	w := &MessageWriter{sender: sender, opcode: uint16(opcode)}
	w.WriteUint(sender.ID()) // first word: object id
	w.WriteUint(0)           // placeholder; later be set to opcode&length
	return w
}

// WriteUint appends v in native-endian without branches.
func (w *MessageWriter) WriteUint(v uint32) {
	b := *(*[4]byte)(unsafe.Pointer(&v)) // copy value as [4]byte
	w.content = append(w.content, b[:]...)
}

func (w *MessageWriter) WriteInt(v int32)    { w.WriteUint(uint32(v)) }
func (w *MessageWriter) WriteObject(v Proxy) { w.WriteUint(v.ID()) }

// Overwrite 4 bytes at offset off in native-endian.
func putU32Native(dst []byte, off int, v uint32) {
	b := *(*[4]byte)(unsafe.Pointer(&v))
	copy(dst[off:off+4], b[:])
}

func pad4(n int) int { return (n + 3) &^ 3 }

// WriteArray appends Arrays: length + bytes + pad (no NUL)
func (w *MessageWriter) WriteArray(v []byte) {
	l := len(v)
	w.WriteUint(uint32(l))
	w.content = append(w.content, v...)
	pad := pad4(l) - l
	if pad > 0 {
		w.content = append(w.content, make([]byte, pad)...)
	}
}

// WriteString appends Strings: length incl NUL + bytes incl NUL + pad
func (w *MessageWriter) WriteString(s string) {
	if s == "" {
		w.WriteUint(1) // length incl NUL
		w.content = append(w.content, 0)
		w.content = append(w.content, 0, 0, 0) // pad to 4
		return
	}
	b := append([]byte(s), 0)
	w.WriteUint(uint32(len(b)))
	w.content = append(w.content, b...)
	pad := pad4(len(b)) - len(b)
	if pad > 0 {
		w.content = append(w.content, make([]byte, pad)...)
	}
}

// WriteFixed appends Signed 24.8 fixed point
func (w *MessageWriter) WriteFixed(f float64) {
	fp := int32(math.Round(f * 256.0))
	w.WriteInt(fp)
}

// WriteFd appends filedescriptor which can be passed to the compositor
func (w *MessageWriter) WriteFd(fd int) { w.fds = append(w.fds, fd) }

func (w *MessageWriter) Finish() error {
	if len(w.content) < 8 || len(w.content) > 0xffff {
		return fmt.Errorf("length of message out of bounds: 8 <= %d <= 0xffff", len(w.content))
	}

	second := (uint32(len(w.content)) << 16) | uint32(w.opcode)
	putU32Native(w.content, 4, second)

	var oob []byte
	if len(w.fds) > 0 {
		oob = syscall.UnixRights(w.fds...)
	}

	conn := w.sender.Conn()

	conn.writeLock.Lock()
	defer conn.writeLock.Unlock()
	n, oobn, err := conn.WriteMsgUnix(w.content, oob, nil)
	if err != nil {
		return err
	}
	if n != len(w.content) || oobn != len(oob) {
		return fmt.Errorf("WriteMsgUnix short write: n=%d/%d oob=%d/%d", n, len(w.content), oobn, len(oob))
	}
	/* reset state */
	w.content = nil
	w.fds = nil
	return nil
}
