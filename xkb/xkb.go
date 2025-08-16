package xkb

/*
#cgo pkg-config: xkbcommon
#include <stdlib.h>
#include <locale.h>
#include <xkbcommon/xkbcommon.h>
#include <xkbcommon/xkbcommon-compose.h>
*/
import "C"
import (
	"errors"
	"log"
	"os"
	"runtime"
	"unicode/utf8"
	"unsafe"
)

type Keyboard struct {
	ctx    *C.struct_xkb_context
	keymap *C.struct_xkb_keymap
	state  *C.struct_xkb_state

	ctab   *C.struct_xkb_compose_table
	cstate *C.struct_xkb_compose_state

	useCompose bool
}

// NewFromKeymapText maakt een toetsenbordstate vanuit een XKB text v1 keymap.
// locale: laat leeg ("") om de huidige proces-locale te gebruiken (LC_ALL/LC_CTYPE/LANG).
func NewFromKeymapText(keymapText []byte, locale string) (*Keyboard, error) {
	if len(keymapText) == 0 {
		return nil, errors.New("empty keymap")
	}

	kb := &Keyboard{}

	// Context
	kb.ctx = C.xkb_context_new(C.enum_xkb_context_flags(0))
	if kb.ctx == nil {
		return nil, errors.New("xkb_context_new failed")
	}

	// Keymap (TEXT_V1)
	cstr := C.CBytes(keymapText)
	defer C.free(cstr)
	kb.keymap = C.xkb_keymap_new_from_string(
		kb.ctx,
		(*C.char)(cstr),
		C.enum_xkb_keymap_format(C.XKB_KEYMAP_FORMAT_TEXT_V1),
		C.enum_xkb_keymap_compile_flags(0),
	)
	if kb.keymap == nil {
		kb.Close()
		return nil, errors.New("xkb_keymap_new_from_string failed (is dit echt text v1?)")
	}

	// State
	kb.state = C.xkb_state_new(kb.keymap)
	if kb.state == nil {
		kb.Close()
		return nil, errors.New("xkb_state_new failed")
	}

	// Compose: optioneel maar meestal gewenst
	loc := locale
	if loc == "" {
		// prefer LC_ALL, dan LC_CTYPE, dan LANG
		loc = firstNonEmpty(os.Getenv("LC_ALL"), os.Getenv("LC_CTYPE"), os.Getenv("LANG"))
	}
	if loc != "" {
		cloc := C.CString(loc)
		defer C.free(unsafe.Pointer(cloc))
		kb.ctab = C.xkb_compose_table_new_from_locale(kb.ctx, cloc, C.enum_xkb_compose_compile_flags(0))
		if kb.ctab != nil {
			kb.cstate = C.xkb_compose_state_new(kb.ctab, C.enum_xkb_compose_state_flags(0))
			kb.useCompose = kb.cstate != nil
		}
	}

	runtime.SetFinalizer(kb, func(k *Keyboard) { k.Close() })
	return kb, nil
}

func firstNonEmpty(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

// Close: vrijgeven van alle XKB-objecten.
func (k *Keyboard) Close() {
	if k == nil {
		return
	}
	if k.cstate != nil {
		C.xkb_compose_state_unref(k.cstate)
		k.cstate = nil
	}
	if k.ctab != nil {
		C.xkb_compose_table_unref(k.ctab)
		k.ctab = nil
	}
	if k.state != nil {
		C.xkb_state_unref(k.state)
		k.state = nil
	}
	if k.keymap != nil {
		C.xkb_keymap_unref(k.keymap)
		k.keymap = nil
	}
	if k.ctx != nil {
		C.xkb_context_unref(k.ctx)
		k.ctx = nil
	}
}

// update verwerkt een key up/down en houdt modifiers/levels bij.
// keycode is een **XKB keycode** (dus evdev + 8).
func (k *Keyboard) update(keycode uint32, pressed bool) {
	if k == nil || k.state == nil {
		return
	}
	dir := C.enum_xkb_key_direction(C.XKB_KEY_UP)
	if pressed {
		dir = C.XKB_KEY_DOWN
	}
	C.xkb_state_update_key(k.state, C.uint32_t(keycode), dir)
}

// Result van Translate.
type Result struct {
	KeySym     uint32 // XKB keysym (zoals XKB_KEY_a)
	KeySymName string // "a", "Left", etc.
	Char       rune   // geproduceerde tekst (leeg voor niet-printables)
	Composed   bool   // true als via Compose afgemaakt
}

// Translate combineert Update + ophalen van keysym/tekst.
// Aanname: je roept dit op bij key **press**. Bij release roep je alleen Update(..., false) aan.
func (k *Keyboard) Translate(keycode uint32, pressed bool) (Result, error) {
	if k == nil || k.state == nil {
		return Result{}, errors.New("keyboard not initialized")
	}

	// Werk state bij.
	k.update(keycode, pressed)

	// Alleen op press tekst/keysyms produceren.
	if !pressed {
		return Result{}, nil
	}

	// "Ruwe" keysym voor deze key/level/modifiers:
	sym := C.xkb_state_key_get_one_sym(k.state, C.uint32_t(keycode))

	// Compose (optioneel)
	if k.useCompose {
		C.xkb_compose_state_feed(k.cstate, sym)
		switch C.xkb_compose_state_get_status(k.cstate) {
		case C.XKB_COMPOSE_COMPOSED:
			// Gecomposeerde keysym en utf8 ophalen
			csym := C.xkb_compose_state_get_one_sym(k.cstate)

			// Naam
			name := make([]byte, 64)
			nn := C.xkb_keysym_get_name(csym, (*C.char)(unsafe.Pointer(&name[0])), C.size_t(len(name)))
			keyname := ""
			if nn > 0 {
				keyname = C.GoStringN((*C.char)(unsafe.Pointer(&name[0])), C.int(nn))
			}

			// UTF-8
			buf := make([]byte, 128)
			n := C.xkb_compose_state_get_utf8(k.cstate, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
			buf = buf[:n]
			char, charl := utf8.DecodeRune(buf)
			if utf8.RuneCount(buf) > 1 {
				log.Printf("possible lost keystroke characters: %s\n", buf[charl:])
			}

			// Reset voor de volgende sequence
			C.xkb_compose_state_reset(k.cstate)

			return Result{
				KeySym:     uint32(csym),
				KeySymName: keyname,
				Char:       char,
				Composed:   true,
			}, nil

		case C.XKB_COMPOSE_COMPOSING:
			// Bezig met samenstelling; nog niets produceren.
			return Result{}, nil

		case C.XKB_COMPOSE_CANCELLED:
			// Sequence geannuleerd; reset en val terug op "ruw".
			C.xkb_compose_state_reset(k.cstate)
			// ga door naar "ruw"
		case C.XKB_COMPOSE_NOTHING:
			// geen compose: ga door naar "ruw"
		}
	}

	// Geen (afgeronde) compose -> vraag direct aan state.
	// Naam
	name := make([]byte, 64)
	nn := C.xkb_keysym_get_name(sym, (*C.char)(unsafe.Pointer(&name[0])), C.size_t(len(name)))
	keyname := ""
	if nn > 0 {
		keyname = C.GoStringN((*C.char)(unsafe.Pointer(&name[0])), C.int(nn))
	}

	// UTF-8 tekst (kan leeg zijn bij niet-printables)
	buf := make([]byte, 128)
	n := C.xkb_state_key_get_utf8(k.state, C.uint32_t(keycode), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	buf = buf[:n]
	char, charl := utf8.DecodeRune(buf)
	if utf8.RuneCount(buf) > 1 {
		log.Printf("possible lost keystroke characters: %s\n", buf[charl:])
	}

	return Result{
		KeySym:     uint32(sym),
		KeySymName: keyname,
		Char:       char,
		Composed:   false,
	}, nil
}
