package wayland

import (
	"fmt"
	"unsafe"
)

func (conn *Conn) WriteMsg(b []byte, oob []byte) error {
	conn.writeLock.Lock()
	defer conn.writeLock.Unlock()
	n, oobn, err := conn.sock.WriteMsgUnix(b, oob, nil)
	if err != nil {
		return err
	}
	if n != len(b) || oobn != len(oob) {
		return fmt.Errorf("conn.WriteMsg: incorrect number of bytes written (n=%d oobn=%d)", n, oobn)
	}

	return nil
}

func PaddedLen(l int) int {
	if (l & 0x3) != 0 {
		return l + (4 - (l & 0x3))
	}
	return l
}

func PutUint32(dst []byte, v uint32) {
	_ = dst[3]
	*(*uint32)(unsafe.Pointer(&dst[0])) = v
}

func PutFixed(dst []byte, f float64) {
	fx := int32(f * 256)
	_ = dst[3]
	*(*int32)(unsafe.Pointer(&dst[0])) = fx
}

func PutString(dst []byte, v string, l int) {
	PutUint32(dst[:4], uint32(l))
	copy(dst[4:], []byte(v))
}

func PutArray(dst []byte, a []byte) {
	PutUint32(dst[:4], uint32(len(a)))
	copy(dst[4:], a)
}
