package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/pflag"
)

var (
	outputFile  = pflag.StringP("output", "o", "", "Writes the generated code to `path` instead stdout")
	packageName = pflag.StringP("package", "p", "wayland_proto", "Package `name` to use")
	runtime     = pflag.StringP("runtime", "r", "github.com/friedelschoen/wayland", "Name of wayland-rumtime `package`")
	trimPrefix  = pflag.StringSliceP("strip-prefix", "P", nil, "Strip `prefix` from all identifiers, comma-seperated and repeatable")
	trimSuffix  = pflag.StringSliceP("strip-suffix", "S", nil, "Strip `suffix` from all identifiers, comma-seperated and repeatable")
	trimExcept  = pflag.StringSliceP("strip-except", "V", nil, "Exclude `identifiers` from stipping (e.g. to avoid duplicates), comma-seperated and repeatable")
)

type Protocol struct {
	XMLName    xml.Name    `xml:"protocol"`
	Name       string      `xml:"name,attr"`
	Copyright  string      `xml:"copyright"`
	Interfaces []Interface `xml:"interface"`
}

type Interface struct {
	XMLName     xml.Name    `xml:"interface"`
	Name        string      `xml:"name,attr"`
	Description Description `xml:"description"`
	Requests    []Request   `xml:"request"`
	Events      []Event     `xml:"event"`
	Enums       []Enum      `xml:"enum"`
	Version     int         `xml:"version,attr"`
}

type Request struct {
	XMLName     xml.Name    `xml:"request"`
	Name        string      `xml:"name,attr"`
	Type        string      `xml:"type,attr"`
	Description Description `xml:"description"`
	Args        []Arg       `xml:"arg"`
	Since       int         `xml:"since,attr"`
}

type Event struct {
	XMLName     xml.Name    `xml:"event"`
	Name        string      `xml:"name,attr"`
	Type        string      `xml:"type,attr"`
	Description Description `xml:"description"`
	Args        []Arg       `xml:"arg"`
	Since       int         `xml:"since,attr"`
}

type Enum struct {
	XMLName     xml.Name    `xml:"enum"`
	Name        string      `xml:"name,attr"`
	Description Description `xml:"description"`
	Entries     []Entry     `xml:"entry"`
	Since       int         `xml:"since,attr"`
	Bitfield    bool        `xml:"bitfield,attr"`
}

type Entry struct {
	XMLName     xml.Name    `xml:"entry"`
	Name        string      `xml:"name,attr"`
	Value       string      `xml:"value,attr"`
	Summary     string      `xml:"summary,attr"`
	Description Description `xml:"description"`
	Since       int         `xml:"since,attr"`
}

type Arg struct {
	XMLName     xml.Name    `xml:"arg"`
	Name        string      `xml:"name,attr"`
	Type        string      `xml:"type,attr"`
	Summary     string      `xml:"summary,attr"`
	Interface   string      `xml:"interface,attr"`
	Enum        string      `xml:"enum,attr"`
	Description Description `xml:"description"`
	AllowNull   bool        `xml:"allow-null,attr"`
}

type Description struct {
	XMLName xml.Name `xml:"description"`
	Text    string   `xml:",chardata"`
	Summary string   `xml:"summary,attr"`
}

var typeToGoTypeMap = map[string]string{
	"int":    "int32",
	"uint":   "uint32",
	"fixed":  "float64",
	"string": "string",
	"object": "Proxy",
	"array":  "[]byte",
	"fd":     "int",
}

func getInputFile(file string) (io.ReadCloser, error) {
	if strings.HasPrefix(file, "http") {
		resp, err := http.Get(file)
		if err != nil {
			return nil, err
		}

		return resp.Body, nil
	}

	return os.Open(file)
}

func trim(s string) string {
	for _, exc := range *trimExcept {
		if strings.Contains(s, exc) {
			return s
		}
	}

	for _, pre := range *trimPrefix {
		s = strings.TrimPrefix(s, pre)
	}
	for _, suf := range *trimSuffix {
		s = strings.TrimSuffix(s, suf)
	}

	return s
}

func toCase(s string, leadingUpper bool) string {
	if s == "" {
		return ""
	}
	s = trim(s)

	parts := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for i, p := range parts {
		if i == 0 && !leadingUpper {
			parts[i] = strings.ToLower(p)
		} else {
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return strings.Join(parts, "")
}

func toCamel(s string) string {
	return toCase(s, true)
}

func toLowerCamel(s string) string {
	s = toCase(s, false)
	if s == "map" {
		s = "_map"
	}
	return s
}

func writeComment(w io.Writer, s string) {
	for line := range strings.SplitSeq(s, "\n") {
		fmt.Fprintf(w, "// %s\n", line)
	}
}

func hasDestructor(v Interface) bool {
	for _, r := range v.Requests {
		if r.Type == "destructor" {
			return true
		}
	}

	return false
}

func writeInterface(w io.Writer, imports map[string]bool, v Interface) {
	ifaceName := toCamel(v.Name)

	// Interface struct
	fmt.Fprintf(w, "// %s: %s\n", ifaceName, strings.Replace(v.Description.Summary, "\n", " ", -1))
	writeComment(w, v.Description.Text)
	fmt.Fprintf(w, "type %s struct {\n", ifaceName)
	fmt.Fprintf(w, "runtime.BaseProxy\n")
	if len(v.Events) != 0 {
		fmt.Fprintf(w, "handlers *%sHandlers\n", ifaceName)
	}
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "type %sHandlers struct {\n", ifaceName)
	for _, event := range v.Events {
		fmt.Fprintf(w, "On%s runtime.EventHandlerFunc\n", toCamel(event.Name))
	}
	fmt.Fprintf(w, "}\n")

	// Constructor
	fmt.Fprintf(w, "// New%s: %s\n", ifaceName, strings.Replace(v.Description.Summary, "\n", " ", -1))
	writeComment(w, v.Description.Text)
	fmt.Fprintf(w, "func New%s(handlers *%sHandlers) *%s {\n", ifaceName, ifaceName, ifaceName)
	fmt.Fprintf(w, "i := &%s{}\n", ifaceName)
	if len(v.Events) != 0 {
		fmt.Fprintf(w, "i.handlers = handlers\n")
	}
	fmt.Fprintf(w, "return i\n")
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "func (i *%s) Name() string {\n", ifaceName)
	fmt.Fprintf(w, "return \"%s\"\n", v.Name)
	fmt.Fprintf(w, "}\n")

	// Requests
	for i, r := range v.Requests {
		writeRequest(w, imports, ifaceName, i, r)
	}

	if !hasDestructor(v) {
		fmt.Fprintf(w, "func (i *%s) Destroy() error {\n", ifaceName)
		fmt.Fprintf(w, "i.Conn().Unregister(i)\n")
		fmt.Fprintf(w, "return nil\n")
		fmt.Fprintf(w, "}\n")
	}

	// Enums
	for _, e := range v.Enums {
		writeEnum(w, ifaceName, e)
	}

	// Events
	for _, e := range v.Events {
		writeEvent(w, ifaceName, e)
	}

	// Event dispatcher
	writeEventDispatcher(w, imports, ifaceName, v)
}

func writeRequest(w io.Writer, imports map[string]bool, ifaceName string, opcode int, r Request) {
	requestName := toCamel(r.Name)

	// Generate param & returns types
	var (
		params        []string
		paramHandlers []string
		results       []string
		resultTypes   []string
	)
	for _, arg := range r.Args {
		argNameLower := toLowerCamel(arg.Name)
		argIface := toCamel(arg.Interface)

		switch arg.Type {
		case "new_id":
			if arg.Interface != "" {
				paramHandlers = append(paramHandlers, argNameLower+"Handlers *"+argIface+"Handlers")
				resultTypes = append(resultTypes, "*"+argIface)
				results = append(results, argNameLower)
			} else {
				// Special for wl_registry.bind
				params = append(params, "iface string", "version uint32", "id runtime.Proxy")
			}

		case "object":
			params = append(params, argNameLower+" *"+argIface)

		case "uint":
			if arg.Enum != "" {
				enumIface, enumName, ok := strings.Cut(arg.Enum, ".")
				if !ok {
					enumName = toCamel(enumIface)
					enumIface = ifaceName
				} else {
					enumName = toCamel(enumName)
					enumIface = toCamel(enumIface)
				}
				params = append(params, argNameLower+" "+enumIface+enumName)
				break
			}
			fallthrough
		case "int", "fixed",
			"string", "array", "fd":
			params = append(params, argNameLower+" "+typeToGoTypeMap[arg.Type])
		}
	}

	fmt.Fprintf(w, "// %s: %s\n", requestName, strings.Replace(r.Description.Summary, "\n", " ", -1))
	writeComment(w, r.Description.Text)
	fmt.Fprintf(w, "//\n")
	for _, arg := range r.Args {
		argNameLower := toLowerCamel(arg.Name)

		if arg.Summary != "" && arg.Type != "new_id" {
			fmt.Fprintf(w, "//  %s: %s\n", argNameLower, strings.Replace(arg.Summary, "\n", " ", -1))
		}
	}
	fmt.Fprintf(w, "func (i *%s) %s(%s) (%s) {\n", ifaceName, requestName, strings.Join(append(params, paramHandlers...), ","), strings.Join(resultTypes, ","))
	if r.Type == "destructor" {
		fmt.Fprintf(w, "defer i.Conn().Unregister(i)\n")
	}

	for _, arg := range r.Args {
		argNameLower := toLowerCamel(arg.Name)

		if arg.Type == "new_id" && arg.Interface != "" {
			fmt.Fprintf(w, "%s := New%s(%sHandlers)\n", argNameLower, toCamel(arg.Interface), argNameLower)
			fmt.Fprintf(w, "i.Conn().Register(%s)\n", argNameLower)
		}
	}

	// Create request
	fmt.Fprintf(w, "const opcode = %d\n", opcode)

	// Calculate size
	var dysize strings.Builder
	size := 8
	canBeConst := true
	for _, arg := range r.Args {
		argNameLower := toLowerCamel(arg.Name)

		switch arg.Type {
		case "new_id":
			if arg.Interface != "" {
				size += 4
			} else {
				canBeConst = false
				fmt.Fprintf(w, "ifaceLen := runtime.PaddedLen(len(iface)+1)\n")
				fmt.Fprintf(&dysize, "+ ifaceLen")
				size += 12
			}

		case "object", "int", "uint", "fixed":
			size += 4

		case "string":
			canBeConst = false
			fmt.Fprintf(w, "%sLen := runtime.PaddedLen(len(%s)+1)\n", argNameLower, argNameLower)
			fmt.Fprintf(&dysize, "+ %sLen", argNameLower)
			size += 4

		case "array":
			canBeConst = false
			fmt.Fprintf(&dysize, "+ len(%s)", argNameLower)
		}
	}

	writeArgs := []string{"", "nil"}
	if canBeConst {
		fmt.Fprintf(w, "const _reqBufLen = %d%s\n", size, &dysize)
		fmt.Fprintf(w, "var _reqBuf [_reqBufLen]byte\n")
		writeArgs[0] = "_reqBuf[:]"
	} else {
		fmt.Fprintf(w, "_reqBufLen := %d%s\n", size, &dysize)
		fmt.Fprintf(w, "_reqBuf := make([]byte, _reqBufLen)\n")
		writeArgs[0] = "_reqBuf"
	}

	fmt.Fprintf(w, "l := 0\n")
	fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:4], i.ID())\n")
	fmt.Fprintf(w, "l += 4\n")
	fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0xffff))\n")
	fmt.Fprintf(w, "l += 4\n")

	fdIndex := -1
	for i, arg := range r.Args {
		argNameLower := toLowerCamel(arg.Name)

		switch arg.Type {
		case "object":
			if arg.AllowNull {
				fmt.Fprintf(w, "if %s == nil {\n", argNameLower)
				fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], 0)\n")
				fmt.Fprintf(w, "l += 4\n")
				fmt.Fprintf(w, "} else {\n")
				fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], %s.ID())\n", argNameLower)
				fmt.Fprintf(w, "l += 4\n")
				fmt.Fprintf(w, "}\n")
			} else {
				fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], %s.ID())\n", argNameLower)
				fmt.Fprintf(w, "l += 4\n")
			}

		case "new_id":
			if arg.Interface != "" {
				fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], %s.ID())\n", argNameLower)
				fmt.Fprintf(w, "l += 4\n")
			} else {
				fmt.Fprintf(w, "runtime.PutString(_reqBuf[l:l+(4 + ifaceLen)], iface, ifaceLen)\n")
				fmt.Fprintf(w, "l += (4 + ifaceLen)\n")

				fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], uint32(version))\n")
				fmt.Fprintf(w, "l += 4\n")

				fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], id.ID())\n")
				fmt.Fprintf(w, "l += 4\n")
			}

		case "int", "uint":
			fmt.Fprintf(w, "runtime.PutUint32(_reqBuf[l:l+4], uint32(%s))\n", argNameLower)
			fmt.Fprintf(w, "l += 4\n")

		case "fixed":
			fmt.Fprintf(w, "runtime.PutFixed(_reqBuf[l:l+4], %s)\n", argNameLower)
			fmt.Fprintf(w, "l += 4\n")

		case "string":
			fmt.Fprintf(w, "runtime.PutString(_reqBuf[l:l+(4 + %sLen)], %s, %sLen)\n", argNameLower, argNameLower, argNameLower)
			fmt.Fprintf(w, "l += (4 + %sLen)\n", argNameLower)

		case "array":
			fmt.Fprintf(w, "runtime.PutArray(_reqBuf[l:l+(4 + %sLen)], %s)\n", argNameLower, argNameLower)
			fmt.Fprintf(w, "l += %sLen\n", argNameLower)

		case "fd":
			fdIndex = i
		}
	}

	if fdIndex != -1 {
		arg := r.Args[fdIndex]
		argNameLower := toLowerCamel(arg.Name)

		imports["syscall"] = true
		fmt.Fprintf(w, "oob := syscall.UnixRights(%s)\n", argNameLower)
		writeArgs[1] = "oob"
	}
	fmt.Fprintf(w, "if err := i.Conn().WriteMsg(%s, %s); err != nil {\n", writeArgs[0], writeArgs[1])
	fmt.Fprintf(w, "panic(err)\n")
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "return %s\n", strings.Join(results, ","))
	fmt.Fprintf(w, "}\n")
}

func writeEnum(w io.Writer, ifaceName string, e Enum) {
	enumName := toCamel(e.Name)

	fmt.Fprintf(w, "type %s%s uint32\n", ifaceName, enumName)

	fmt.Fprintf(w, "// %s%s: %s\n", ifaceName, enumName, e.Description.Summary)
	writeComment(w, e.Description.Text)
	fmt.Fprintf(w, "const (\n")
	for _, entry := range e.Entries {
		entryName := toCamel(entry.Name)

		if entry.Summary != "" {
			fmt.Fprintf(w, "// %s%s%s: %s\n", ifaceName, enumName, entryName, strings.Replace(entry.Summary, "\n", " ", -1))
		}
		fmt.Fprintf(w, "%s%s%s %s%s = %s\n", ifaceName, enumName, entryName, ifaceName, enumName, entry.Value)
	}
	fmt.Fprintf(w, ")\n")

	fmt.Fprintf(w, "func (e %s%s) Name() string {\n", ifaceName, enumName)
	fmt.Fprintf(w, "switch e {\n")
	for _, entry := range e.Entries {
		entryName := toCamel(entry.Name)

		fmt.Fprintf(w, "case %s%s%s:\n", ifaceName, enumName, entryName)
		fmt.Fprintf(w, "return \"%s\"\n", entry.Name)
	}
	fmt.Fprintf(w, "default:\n")
	fmt.Fprintf(w, "return \"\"\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "func (e %s%s) Value() string {\n", ifaceName, enumName)
	fmt.Fprintf(w, "switch e {\n")
	for _, entry := range e.Entries {
		entryName := toCamel(entry.Name)

		fmt.Fprintf(w, "case %s%s%s:\n", ifaceName, enumName, entryName)
		fmt.Fprintf(w, "return \"%s\"\n", entry.Value)
	}
	fmt.Fprintf(w, "default:\n")
	fmt.Fprintf(w, "return \"\"\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "func (e %s%s) String() string {\n", ifaceName, enumName)
	fmt.Fprintf(w, "return  e.Name() + \"=\" +  e.Value()\n")
	fmt.Fprintf(w, "}\n")
}

func hasParams(params []Arg, wants map[string]string) bool {
	for _, a := range params {
		t, ok := wants[a.Name]
		if ok && t == a.Type {
			delete(wants, a.Name)
		}
	}
	return len(wants) == 0
}

func writeEvent(w io.Writer, ifaceName string, e Event) {
	eventName := toCamel(e.Name)
	// eventNameLower := toLowerCamel(e.Name)

	// Event struct
	fmt.Fprintf(w, "// %s%sEvent : %s\n", ifaceName, eventName, strings.Replace(e.Description.Summary, "\n", " ", -1))
	writeComment(w, e.Description.Text)
	fmt.Fprintf(w, "type %s%sEvent struct {\n", ifaceName, eventName)
	fmt.Fprintf(w, "proxy runtime.Proxy\n")
	for _, arg := range e.Args {
		argName := toCamel(arg.Name)

		if arg.Summary != "" {
			fmt.Fprintf(w, "// %s %s\n", argName, strings.Replace(arg.Summary, "\n", " ", -1))
		} else if arg.Description.Summary != "" {
			fmt.Fprintf(w, "// %s %s\n", argName, strings.Replace(arg.Description.Summary, "\n", " ", -1))
		}
		writeComment(w, arg.Description.Text)
		switch arg.Type {
		case "object", "new_id":
			if arg.Interface != "" {
				argIface := toCamel(arg.Interface)

				fmt.Fprintf(w, "%s *%s\n", argName, argIface)
			} else {
				fmt.Fprintf(w, "%s runtime.Proxy\n", argName)
			}

		case "uint":
			if arg.Enum != "" {
				enumIface, enumName, ok := strings.Cut(arg.Enum, ".")
				if !ok {
					enumName = toCamel(enumIface)
					enumIface = ifaceName
				} else {
					enumName = toCamel(enumName)
					enumIface = toCamel(enumIface)
				}
				fmt.Fprintf(w, "%s %s%s\n", argName, enumIface, enumName)
				break
			}
			fallthrough
		case "int", "fixed",
			"string", "array", "fd":
			fmt.Fprintf(w, "%s %s\n", argName, typeToGoTypeMap[arg.Type])
		}
	}
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "func (e *%s%sEvent) Proxy() runtime.Proxy {\n", ifaceName, eventName)
	fmt.Fprintf(w, "return e.proxy\n")
	fmt.Fprintf(w, "}\n")

	// <arg name="name" type="uint" summary="numeric name of the global object"/>
	// <arg name="interface" type="string" summary="interface implemented by the object"/>
	// <arg name="version" type="uint" summary="interface version"/>
	if hasParams(e.Args, map[string]string{"name": "uint", "interface": "string", "version": "uint"}) {
		fmt.Fprintf(w, "func (e *%s%sEvent) BindInterface() (uint32, string, uint32) {\n", ifaceName, eventName)
		fmt.Fprintf(w, "return e.Name, e.Interface, e.Version\n")
		fmt.Fprintf(w, "}\n")
	}
}

func writeEventDispatcher(w io.Writer, imports map[string]bool, ifaceName string, v Interface) {
	fmt.Fprintf(w, "func (i *%s) Dispatch(opcode uint32, fd int, data []byte, drain chan<- runtime.Event) {\n", ifaceName)
	if len(v.Events) > 0 {
		fmt.Fprintf(w, "if i.handlers == nil && drain == nil {\n")
		fmt.Fprintf(w, "return\n")
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "switch opcode {\n")
		for i, e := range v.Events {
			eventName := toCamel(e.Name)

			hasFd := false
			for _, arg := range e.Args {
				if arg.Type == "fd" {
					hasFd = true
					break
				}
			}

			fmt.Fprintf(w, "case %d:\n", i)
			fmt.Fprintf(w, "if (i.handlers != nil && i.handlers.On%s == nil) && drain == nil {\n", eventName)
			if hasFd {
				imports["syscall"] = true
				fmt.Fprintf(w, "if fd != -1 {\n")
				fmt.Fprintf(w, "syscall.Close(fd)\n")
				fmt.Fprintf(w, "}\n")
			}
			fmt.Fprintf(w, "return\n")
			fmt.Fprintf(w, "}\n")

			fmt.Fprintf(w, "e := &%s%sEvent{}\n", ifaceName, eventName)
			fmt.Fprintf(w, "e.proxy = i\n")

			if len(e.Args) > 0 && (len(e.Args) != 1 || e.Args[0].Type != "fd") {
				fmt.Fprintf(w, "l := 0\n")
			}

			for _, arg := range e.Args {
				argName := toCamel(arg.Name)
				argNameLower := toLowerCamel(arg.Name)

				switch arg.Type {
				case "object", "new_id":
					if arg.Interface != "" {
						argIface := toCamel(arg.Interface)

						fmt.Fprintf(w, "e.%s = i.Conn().GetProxy(runtime.Uint32(data[l :l+4])).(*%s)\n", argName, argIface)
					} else {
						fmt.Fprintf(w, "e.%s = i.Conn().GetProxy(runtime.Uint32(data[l :l+4]))\n", argName)
					}
					fmt.Fprintf(w, "l += 4\n")

				case "fd":
					fmt.Fprintf(w, "e.%s = fd\n", argName)

				case "uint":
					if arg.Enum != "" {
						enumIface, enumName, ok := strings.Cut(arg.Enum, ".")
						if !ok {
							enumName = toCamel(enumIface)
							enumIface = ifaceName
						} else {
							enumName = toCamel(enumName)
							enumIface = toCamel(enumIface)
						}
						fmt.Fprintf(w, "e.%s = %s%s(runtime.Uint32(data[l : l+4]))\n", argName, enumIface, enumName)
					} else {
						fmt.Fprintf(w, "e.%s = runtime.Uint32(data[l : l+4])\n", argName)
					}
					fmt.Fprintf(w, "l += 4\n")

				case "int":
					fmt.Fprintf(w, "e.%s = int32(runtime.Uint32(data[l : l+4]))\n", argName)
					fmt.Fprintf(w, "l += 4\n")

				case "fixed":
					fmt.Fprintf(w, "e.%s = runtime.Fixed(data[l : l+4])\n", argName)
					fmt.Fprintf(w, "l += 4\n")

				case "string":
					fmt.Fprintf(w, "%sLen := runtime.PaddedLen(int(runtime.Uint32(data[l : l+4])))\n", argNameLower)
					fmt.Fprintf(w, "l += 4\n")
					fmt.Fprintf(w, "e.%s = runtime.String(data[l : l+%sLen])\n", argName, argNameLower)
					fmt.Fprintf(w, "l += %sLen\n", argNameLower)

				case "array":
					fmt.Fprintf(w, "%sLen := int(runtime.Uint32(data[l : l+4]))\n", argNameLower)
					fmt.Fprintf(w, "l += 4\n")
					fmt.Fprintf(w, "e.%s = make([]byte, %sLen)\n", argName, argNameLower)
					fmt.Fprintf(w, "copy(e.%s, data[l:l+%sLen])\n", argName, argNameLower)
					fmt.Fprintf(w, "l += %sLen\n", argNameLower)
				}
			}

			fmt.Fprintf(w, "if i.handlers != nil && i.handlers.On%s != nil {\n", eventName)
			fmt.Fprintf(w, "i.handlers.On%s(e)\n", eventName)
			fmt.Fprintf(w, "} else if drain != nil {\n")
			fmt.Fprintf(w, "drain <- e\n")
			fmt.Fprintf(w, "}\n")
		}
		fmt.Fprintf(w, "}\n")
	}
	fmt.Fprintf(w, "}\n")
}

func main() {
	pflag.Parse()

	var src io.ReadCloser = os.Stdin
	var srcname string
	if pflag.NArg() > 0 {
		srcname = pflag.Arg(0)
		var err error
		src, err = getInputFile(srcname)
		if err != nil {
			log.Fatalf("unable to get input file: %v", err)
		}
	}

	var protocol Protocol
	if err := xml.NewDecoder(src).Decode(&protocol); err != nil {
		log.Fatalf("unable to decode protocol xml: %v", err)
	}
	if err := src.Close(); err != nil {
		log.Printf("unable to close input file: %v", err)
	}

	var dst io.WriteCloser = os.Stdout
	if *outputFile != "" {
		var err error
		dst, err = os.Create(*outputFile)
		if err != nil {
			log.Fatalf("unable to create output file: %v", err)
		}
	}
	defer dst.Close()

	w, err := NewFormatter(dst)
	if err != nil {
		log.Printf("unable to use gofmt: %v, continuing without formatter...\n", err)
	}
	defer w.Close()

	// Header
	fmt.Fprintf(w, "// Generated by go-wayland-scanner\n")
	fmt.Fprintf(w, "// https://github.com/friedelschoen/wayland/cmd/go-wl-scanner\n")
	if srcname != "" {
		fmt.Fprintf(w, "// XML file : %s\n", srcname)
	}
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// %s Protocol Copyright: \n", protocol.Name)
	writeComment(w, protocol.Copyright)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "package %s\n", *packageName)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "import runtime \"%s\"\n", *runtime)

	// Interfaces
	var wb bytes.Buffer
	imports := make(map[string]bool)
	for _, v := range protocol.Interfaces {
		writeInterface(&wb, imports, v)
	}

	for name, ok := range imports {
		if !ok {
			continue
		}
		fmt.Fprintf(w, "import \"%s\"\n", name)
	}
	io.Copy(w, &wb)

	fmt.Fprintf(w, "func Get%sInterface(name string) runtime.Proxy {\n", toCamel(protocol.Name))
	fmt.Fprintf(w, "switch (name) {\n")
	for _, v := range protocol.Interfaces {
		fmt.Fprintf(w, "case \"%s\":\n", v.Name)
		fmt.Fprintf(w, "return New%s(nil)\n", toCamel(v.Name))
	}
	fmt.Fprintf(w, "default:\n")
	fmt.Fprintf(w, "return nil\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")
}
