package client

import (
	"encoding/binary"
	"errors"
	"io"
)

// Detecteer native endianness (Wayland gebruikt host byte order).
var nativeEndian = binary.LittleEndian

type WireReader struct {
	io.Reader

	// Optioneel: queue met ontvangen file descriptors (out-of-band).
	// Vul deze zelf aan vanuit je UnixConn/ReadMsgUnix code.
	fdQueue []int
}

// -------- 32-bit primitives --------

func (wr *WireReader) ReadInt() (int32, error) {
	var v int32
	if err := binary.Read(wr.Reader, nativeEndian, &v); err != nil {
		return 0, err
	}
	return v, nil
}

func (wr *WireReader) ReadUint() (uint32, error) {
	var v uint32
	if err := binary.Read(wr.Reader, nativeEndian, &v); err != nil {
		return 0, err
	}
	return v, nil
}

// Wayland fixed: 24.8 signed fixed-point.
// Je krijgt zowel de ruwe 32-bit waarde als een float64 helper.
func (wr *WireReader) ReadFixedRaw() (int32, error) {
	return wr.ReadInt()
}

func (wr *WireReader) ReadFixed() (float64, error) {
	raw, err := wr.ReadFixedRaw()
	if err != nil {
		return 0, err
	}
	return float64(raw) / 256.0, nil // 24.8 => / 2^8
}

// object/new_id/enums zijn 32-bit ints; we geven semantische helpers.
func (wr *WireReader) ReadObjectID() (uint32, error) { return wr.ReadUint() }
func (wr *WireReader) ReadNewID() (uint32, error)    { return wr.ReadUint() }
func (wr *WireReader) ReadEnum() (int32, error)      { return wr.ReadInt() }

// -------- string & array (met 32-bit length + 4-byte padding) --------

func (wr *WireReader) ReadString() (string, error) {
	n, err := wr.ReadUint()
	if err != nil {
		return "", err
	}
	if n == 0 {
		// Lege string; nog steeds padding skippen (0 is al 4-aligned).
		return "", nil
	}
	// n bytes incl. NUL-terminator
	buf := make([]byte, n)
	if err := readFull(wr.Reader, buf); err != nil {
		return "", err
	}
	// Wayland specificeert NUL-terminator; check defensief.
	if buf[len(buf)-1] != 0x00 {
		return "", errors.New("wayland string: missing NUL terminator")
	}
	// Strip NUL.
	s := string(buf[:len(buf)-1])

	// Skip padding naar 4-byte alignment.
	if err := wr.skipPad(uint32(n)); err != nil {
		return "", err
	}
	return s, nil
}

func (wr *WireReader) ReadArray() ([]byte, error) {
	n, err := wr.ReadUint()
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, nil
	}
	buf := make([]byte, n)
	if err := readFull(wr.Reader, buf); err != nil {
		return nil, err
	}
	if err := wr.skipPad(n); err != nil {
		return nil, err
	}
	return buf, nil
}

// -------- file descriptors (out-of-band) --------

// NextFD haalt het eerstvolgende ontvangen FD uit de queue.
// Vul wr.fdQueue zelf vanuit je UnixConn ReadMsgUnix code.
func (wr *WireReader) NextFD() (int, error) {
	if len(wr.fdQueue) == 0 {
		return -1, io.EOF
	}
	fd := wr.fdQueue[0]
	wr.fdQueue = wr.fdQueue[1:]
	return fd, nil
}

// Helper om FDs toe te voegen (aanroepen vanuit je socket-ontvangst).
func (wr *WireReader) enqueueFDs(fds []int) { // veilig publiek maken indien gewenst
	wr.fdQueue = append(wr.fdQueue, fds...)
}

// -------- interne helpers --------

func (wr *WireReader) skipPad(n uint32) error {
	// Pad tot veelvoud van 4: (4 - (n % 4)) % 4
	pad := int((4 - (n & 3)) & 3)
	if pad == 0 {
		return nil
	}
	var junk [3]byte
	_, err := io.ReadFull(wr.Reader, junk[:pad])
	return err
}

func readFull(r io.Reader, buf []byte) error {
	_, err := io.ReadFull(r, buf)
	return err
}
