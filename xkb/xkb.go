// Package xkb is an implementation of libxkbcommon for Go using CGo bindings.
// It is made for wayland-library and only implements the required functionallity
package xkb

/*
#cgo pkg-config: xkbcommon
#include <stdlib.h>
#include <xkbcommon/xkbcommon.h>
#include <xkbcommon/xkbcommon-compose.h>
*/
import "C"

import (
	"errors"
	"log"
	"os"
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

	// Cached modifier indices for common modifiers (may be XKB_MOD_INVALID)
	modShift C.xkb_mod_index_t
	modCtrl  C.xkb_mod_index_t
	modAlt   C.xkb_mod_index_t // Alt or Mod1
	modLogo  C.xkb_mod_index_t // Logo or Mod4 (Super/Win)
	modAltGr C.xkb_mod_index_t // Level3 / ISO_Level3_Shift
	modCaps  C.xkb_mod_index_t // Lock
	modNum   C.xkb_mod_index_t // NumLock
}

// NewFromKeymapText creates a keyboard state from an XKB text v1 keymap.
// locale: leave empty ("") to use the current process locale (LC_ALL/LC_CTYPE/LANG).
func NewFromKeymapText(keymapText []byte, locale string) (*Keyboard, error) {
	if len(keymapText) == 0 {
		return nil, errors.New("empty keymap")
	}

	kb := &Keyboard{}

	// Context
	kb.ctx = C.xkb_context_new(0)
	if kb.ctx == nil {
		return nil, errors.New("xkb_context_new failed")
	}

	// Keymap (TEXT_V1)
	kb.keymap = C.xkb_keymap_new_from_string(
		kb.ctx,
		(*C.char)(unsafe.Pointer(&keymapText[0])),
		C.enum_xkb_keymap_format(C.XKB_KEYMAP_FORMAT_TEXT_V1),
		0,
	)
	if kb.keymap == nil {
		kb.Close()
		return nil, errors.New("xkb_keymap_new_from_string failed (is this text v1?)")
	}

	// State
	kb.state = C.xkb_state_new(kb.keymap)
	if kb.state == nil {
		kb.Close()
		return nil, errors.New("xkb_state_new failed")
	}

	// Compose: optional but commonly desired
	loc := locale
	if loc == "" {
		// prefer LC_ALL, then LC_CTYPE, then LANG
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

	// Resolve commonly-used modifier indices (best-effort)
	kb.modShift = kb.findMod("Shift")
	kb.modCtrl = kb.findMod("Control")
	kb.modAlt = kb.findMod("Alt", "Mod1")
	kb.modLogo = kb.findMod("Logo", "Mod4", "Super") // different keymaps use different names
	kb.modAltGr = kb.findMod("Level3", "ISO_Level3_Shift")
	kb.modCaps = kb.findMod("Lock", "Caps", "Caps_Lock", "CapsLock")
	kb.modNum = kb.findMod("NumLock")

	return kb, nil
}

func (k *Keyboard) findMod(names ...string) C.xkb_mod_index_t {
	for _, n := range names {
		cs := C.CString(n)
		idx := C.xkb_keymap_mod_get_index(k.keymap, cs)
		C.free(unsafe.Pointer(cs))
		if idx != C.XKB_MOD_INVALID {
			return idx
		}
	}
	return C.XKB_MOD_INVALID
}

func firstNonEmpty(s ...string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

// Close releases all XKB objects.
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

// update processes a key up/down and keeps modifiers/levels in sync.
// keycode must be an XKB keycode (on Linux: evdev scancode + 8).
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

// Mod holds both raw XKB modifier masks and normalized booleans for common mods.
type Mod int

const (
	// Normalized flags (derived from Effective)
	ModCtrl Mod = 1 << iota
	ModShift
	ModAlt
	ModLogo
	ModAltGr
	ModCaps
	ModNum
)

// Result of Translate.
type Result struct {
	Sym      Key    // XKB keysym (e.g., XKB_KEY_a)
	Name     string // "a", "Left", etc.
	Char     rune   // produced Unicode rune (zero for non-printables)
	Composed bool   // true if produced via Compose
	Mod      Mod    // modifiers snapshot
}

// Translate combines state update and retrieval of keysym/text.
// Assumption: call this on key *press*. On release, call update(..., false).
func (k *Keyboard) Translate(keycode uint32, pressed bool) (Result, error) {
	if k == nil || k.state == nil {
		return Result{}, errors.New("keyboard not initialized")
	}

	// Update state first.
	k.update(keycode, pressed)

	// Build modifiers snapshot after the update.
	var m Mod
	if k.modActive(k.modCtrl) {
		m |= ModCtrl
	}
	if k.modActive(k.modShift) {
		m |= ModShift
	}
	if k.modActive(k.modAlt) {
		m |= ModAlt
	}
	if k.modActive(k.modLogo) {
		m |= ModLogo
	}
	if k.modActive(k.modAltGr) {
		m |= ModAltGr
	}
	if k.modActive(k.modCaps) {
		m |= ModCaps
	}
	if k.modActive(k.modNum) {
		m |= ModNum
	}

	// Only produce text/keysyms on press.
	if !pressed {
		return Result{Mod: m}, nil
	}

	// "Raw" keysym for this key/level/modifiers:
	sym := C.xkb_state_key_get_one_sym(k.state, C.uint32_t(keycode))

	// Compose (optional)
	if k.useCompose {
		C.xkb_compose_state_feed(k.cstate, sym)
		switch C.xkb_compose_state_get_status(k.cstate) {
		case C.XKB_COMPOSE_COMPOSED:
			// Fetch composed keysym and UTF-8
			csym := C.xkb_compose_state_get_one_sym(k.cstate)

			// Name
			name := make([]byte, 64)
			nn := C.xkb_keysym_get_name(csym, (*C.char)(unsafe.Pointer(&name[0])), C.size_t(len(name)))
			keyname := string(name[:nn])

			// UTF-8 -> first rune (if any)
			buf := make([]byte, 128)
			n := C.xkb_compose_state_get_utf8(k.cstate, (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
			var char rune
			if n > 0 {
				b := buf[:int(n)]
				r, used := utf8.DecodeRune(b)
				char = r
				if utf8.RuneCount(b) > 1 {
					log.Printf("possible lost keystroke characters: %s\n", b[used:])
				}
			}

			// Reset for next sequence
			C.xkb_compose_state_reset(k.cstate)

			return Result{
				Sym:      Key(csym),
				Name:     keyname,
				Char:     char,
				Composed: true,
				Mod:      m,
			}, nil

		case C.XKB_COMPOSE_COMPOSING:
			// In-progress compose; nothing to output yet.
			return Result{Mod: m}, nil

		case C.XKB_COMPOSE_CANCELLED:
			// Sequence cancelled; reset and fall back to "raw".
			C.xkb_compose_state_reset(k.cstate)
			// continue to "raw"
		case C.XKB_COMPOSE_NOTHING:
			// No compose: fall through to "raw"
		}
	}

	// No (completed) compose -> query state directly.
	// Name
	name := make([]byte, 64)
	nn := C.xkb_keysym_get_name(sym, (*C.char)(unsafe.Pointer(&name[0])), C.size_t(len(name)))
	keyname := string(name[:nn])

	// UTF-8 text (may be empty for non-printables)
	buf := make([]byte, 128)
	n := C.xkb_state_key_get_utf8(k.state, C.uint32_t(keycode), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	var char rune
	if int(n) > 0 {
		b := buf[:int(n)]
		r, used := utf8.DecodeRune(b)
		char = r
		if utf8.RuneCount(b) > 1 {
			log.Printf("possible lost keystroke characters: %s\n", b[used:])
		}
	}

	return Result{
		Sym:      Key(sym),
		Name:     keyname,
		Char:     char,
		Composed: false,
		Mod:      m,
	}, nil
}

func (k *Keyboard) modActive(idx C.xkb_mod_index_t) bool {
	if idx == C.XKB_MOD_INVALID {
		return false
	}
	on := C.xkb_state_mod_index_is_active(
		k.state,
		idx,
		C.enum_xkb_state_component(C.XKB_STATE_MODS_EFFECTIVE),
	)
	return on == 1
}
