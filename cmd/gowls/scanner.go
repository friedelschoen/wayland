package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
	"unicode"

	"github.com/spf13/pflag"
)

var (
	outputFile    = pflag.StringP("output", "o", "", "Writes the generated code to `path` instead stdout")
	packageName   = pflag.StringP("package", "p", "wayland_proto", "Package `name` to use")
	runtime       = pflag.StringP("runtime", "r", "github.com/friedelschoen/wayland", "Name of wayland-rumtime `package`")
	trimPrefix    = pflag.StringSliceP("strip-prefix", "P", nil, "Strip `prefix` from all identifiers, comma-seperated and repeatable")
	trimSuffix    = pflag.StringSliceP("strip-suffix", "S", nil, "Strip `suffix` from all identifiers, comma-seperated and repeatable")
	trimExcept    = pflag.StringSliceP("duplicates", "D", nil, "Exclude `identifiers` from stipping to avoid duplicates, comma-seperated and repeatable")
	foreignObject = pflag.StringSliceP("foreign", "f", nil, "Mark `object` as foreign and replace occurences with general proxies, comma-seperated and repeatable")
	noCleanup     = pflag.Bool("no-cleanup", false, "Do not use garbage-collection and relay on manual destroy")
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
	noWrite     bool        /* don't actually write a message, used as default destructor */
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
	s = strings.Join(parts, "")
	if strings.HasSuffix(s, "Id") {
		s = s[:len(s)-2] + "ID"
	}
	return s
}

func toPascalCase(s string) string {
	return toCase(s, true)
}

func toCamelCase(s string) string {
	s = toCase(s, false)
	if s == "map" {
		s = "_map"
	} else if s == "interface" {
		s = "_interface"
	}
	return s
}

func writeComment(w io.Writer, s string) {
	for line := range strings.SplitSeq(s, "\n") {
		fmt.Fprintf(w, "// %s\n", strings.TrimSpace(line))
	}
}

func getDestructor(v Interface) string {
	for _, r := range v.Requests {
		if r.Type == "destructor" {
			return r.Name
		}
	}
	return ""
}

func writeInterface(w io.Writer, interfaces map[string]*Interface, imports map[string]bool, v Interface) {
	ifaceName := toPascalCase(v.Name)

	// Interface struct
	fmt.Fprintf(w, "// %s : %s\n", ifaceName, strings.Replace(v.Description.Summary, "\n", " ", -1))
	writeComment(w, v.Description.Text)
	fmt.Fprintf(w, "type %s struct {\n", ifaceName)
	fmt.Fprintf(w, "wayland.BaseProxy\n")
	fmt.Fprintf(w, "children []wayland.Proxy\n")
	if len(v.Events) != 0 {
		fmt.Fprintf(w, "Handlers *%sHandlers\n", ifaceName)
	}
	fmt.Fprintf(w, "}\n")

	if len(v.Events) != 0 {
		fmt.Fprintf(w, "type %sHandlers struct {\n", ifaceName)
		for _, event := range v.Events {
			fmt.Fprintf(w, "On%s wayland.EventHandlerFunc\n", toPascalCase(event.Name))
		}
		fmt.Fprintf(w, "}\n")
	}

	// Constructor
	fmt.Fprintf(w, "// New%s : %s\n", ifaceName, strings.Replace(v.Description.Summary, "\n", " ", -1))
	writeComment(w, v.Description.Text)
	if len(v.Events) != 0 {
		fmt.Fprintf(w, "func New%s(handlers *%sHandlers) *%s {\n", ifaceName, ifaceName, ifaceName)
		fmt.Fprintf(w, "return &%s{Handlers: handlers}\n", ifaceName)
	} else {
		fmt.Fprintf(w, "func New%s() *%s {\n", ifaceName, ifaceName)
		fmt.Fprintf(w, "return &%s{}\n", ifaceName)
	}
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "func (i *%s) Name() string {\n", ifaceName)
	fmt.Fprintf(w, "return \"%s\"\n", v.Name)
	fmt.Fprintf(w, "}\n")

	if getDestructor(v) == "" {
		v.Requests = append(v.Requests, Request{
			Name:    "destroy",
			Type:    "destructor",
			noWrite: true,
		})
	}

	destructorEvent := false
	for _, e := range v.Events {
		if e.Type == "destructor" {
			destructorEvent = true
			break
		}
	}

	// Requests
	for i, r := range v.Requests {
		writeRequest(w, interfaces, ifaceName, destructorEvent, i, r)
	}

	if !*noCleanup && !destructorEvent {
		imports["runtime"] = true
		fmt.Fprintf(w, "func (i *%s) Register(conn *wayland.Conn, id uint32) {\n", ifaceName)
		fmt.Fprintf(w, "i.BaseProxy.Register(conn, id)\n")
		fmt.Fprintf(w, "if conn != nil && id != 0 {\n")
		fmt.Fprintf(w, "runtime.AddCleanup(i, destroy%s, i.BaseProxy)\n", ifaceName)
		fmt.Fprintf(w, "}\n")
		fmt.Fprintf(w, "}\n")
	}

	// Enums
	if len(v.Enums) > 0 {
		imports["fmt"] = true
	}
	for _, e := range v.Enums {
		writeEnum(w, ifaceName, e)
	}

	// Events
	for _, e := range v.Events {
		writeEvent(w, ifaceName, e)
	}

	// Event dispatcher
	writeEventDispatcher(w, imports, ifaceName, v, getDestructor(v))
}

func writeRequest(w io.Writer, interfaces map[string]*Interface, ifaceName string, destructorEvent bool, opcode int, r Request) {
	requestName := toPascalCase(r.Name)

	// Generate param & returns types
	var (
		params        []string
		paramHandlers []string
		results       []string
		resultTypes   []string
	)
	for _, arg := range r.Args {
		argNameLower := toCamelCase(arg.Name)
		argIface := toPascalCase(arg.Interface)

		switch arg.Type {
		case "new_id":
			if arg.Interface != "" {
				if slices.Contains(*foreignObject, arg.Interface) {
					fmt.Fprintf(os.Stderr, "warn: Request creates new object `%s` which is explicitly foreign\n      Marking this object foreign may lead to unhandled events.\n", arg.Interface)
					resultTypes = append(resultTypes, "wayland.Proxy")
				} else {
					if fi, ok := interfaces[arg.Interface]; !ok || len(fi.Events) > 0 {
						paramHandlers = append(paramHandlers, argNameLower+"Handlers *"+argIface+"Handlers")
					}
					resultTypes = append(resultTypes, "*"+argIface)
				}
				results = append(results, argNameLower)
			} else {
				// Special for wl_registry.bind
				params = append(params, "iface string", "version uint32", "id wayland.Proxy")
			}

		case "object":
			if slices.Contains(*foreignObject, arg.Interface) {
				params = append(params, argNameLower+" wayland.Proxy")
			} else {
				params = append(params, argNameLower+" *"+argIface)
			}

		case "uint":
			if arg.Enum != "" {
				enumIface, enumName, ok := strings.Cut(arg.Enum, ".")
				if !ok {
					enumName = toPascalCase(enumIface)
					enumIface = ifaceName
				} else {
					enumName = toPascalCase(enumName)
					enumIface = toPascalCase(enumIface)
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

	fmt.Fprintf(w, "// %s : %s\n", requestName, strings.Replace(r.Description.Summary, "\n", " ", -1))
	writeComment(w, r.Description.Text)
	fmt.Fprintf(w, "//\n")
	for _, arg := range r.Args {
		argNameLower := toCamelCase(arg.Name)

		if arg.Summary != "" && arg.Type != "new_id" {
			fmt.Fprintf(w, "// %s : %s\n", argNameLower, strings.Replace(arg.Summary, "\n", " ", -1))
		}
	}
	fmt.Fprintf(w, "func (i *%s) %s(%s) (%s) {\n", ifaceName, requestName, strings.Join(append(params, paramHandlers...), ","), strings.Join(resultTypes, ","))
	if r.Type == "destructor" {
		if !*noCleanup && !destructorEvent {
			fmt.Fprintf(w, "destroy%s(i.BaseProxy)\n", ifaceName)
			fmt.Fprintf(w, "}\n")
			fmt.Fprintf(w, "func destroy%s(p wayland.BaseProxy) {\n", ifaceName)
			fmt.Fprintf(w, "i := &p\n")
		}
	}

	fmt.Fprintf(w, "if !i.Valid() {\n")
	fmt.Fprintf(w, "return %s\n", strings.Join(slices.Repeat([]string{"nil"}, len(results)), ","))
	fmt.Fprintf(w, "}\n")

	for _, arg := range r.Args {
		argNameLower := toCamelCase(arg.Name)
		if arg.Type == "new_id" && arg.Interface != "" {
			if slices.Contains(*foreignObject, arg.Interface) {
				fmt.Fprintf(w, "var %s wayland.BaseProxy // %s\n", argNameLower, arg.Interface)
			} else {
				if fi, ok := interfaces[arg.Interface]; !ok || len(fi.Events) > 0 {
					fmt.Fprintf(w, "%s := New%s(%sHandlers)\n", argNameLower, toPascalCase(arg.Interface), argNameLower)
				} else {
					fmt.Fprintf(w, "%s := New%s()\n", argNameLower, toPascalCase(arg.Interface))
				}
			}
			fmt.Fprintf(w, "i.Conn().Register(%s)\n", argNameLower)
			fmt.Fprintf(w, "i.children = append(i.children, %s)\n", argNameLower)
		}
	}

	if !r.noWrite {
		fmt.Fprintf(w, "w := wayland.NewMessageWriter(i, %d)\n", opcode)
		for _, arg := range r.Args {
			argNameLower := toCamelCase(arg.Name)

			switch arg.Type {
			case "object":
				if arg.AllowNull {
					fmt.Fprintf(w, "if %s == nil {\n", argNameLower)
					fmt.Fprintf(w, "w.WriteUint(0)\n")
					fmt.Fprintf(w, "} else {\n")
					fmt.Fprintf(w, "w.WriteObject(%s)\n", argNameLower)
					fmt.Fprintf(w, "}\n")
				} else {
					fmt.Fprintf(w, "w.WriteObject(%s)\n", argNameLower)
				}

			case "new_id":
				if arg.Interface != "" {
					fmt.Fprintf(w, "w.WriteObject(%s)\n", argNameLower)
				} else {
					fmt.Fprintf(w, "w.WriteString(iface)\n")
					fmt.Fprintf(w, "w.WriteUint(version)\n")
					fmt.Fprintf(w, "w.WriteObject(id)\n")
				}

			case "int":
				fmt.Fprintf(w, "w.WriteInt(%s)\n", argNameLower)
			case "uint":
				fmt.Fprintf(w, "w.WriteUint(uint32(%s))\n", argNameLower)

			case "fixed":
				fmt.Fprintf(w, "w.WriteFixed(%s)\n", argNameLower)

			case "string":
				if arg.AllowNull {
					fmt.Fprintf(w, "if %s == \"\" {\n", argNameLower)
					fmt.Fprintf(w, "w.WriteUint(0)\n")
					fmt.Fprintf(w, "} else {\n")
					fmt.Fprintf(w, "w.WriteString(%s)\n", argNameLower)
					fmt.Fprintf(w, "}\n")
				} else {
					fmt.Fprintf(w, "w.WriteString(%s)\n", argNameLower)
				}

			case "array":
				fmt.Fprintf(w, "w.WriteArray(%s)\n", argNameLower)

			case "fd":
				fmt.Fprintf(w, "w.WriteFd(%s)\n", argNameLower)
			}
		}
		fmt.Fprintf(w, "if err := w.Finish(); err != nil {\n")
		fmt.Fprintf(w, "panic(err)\n")
		fmt.Fprintf(w, "}\n")
	}
	if len(results) > 0 {
		fmt.Fprintf(w, "return %s\n", strings.Join(results, ","))
	}
	fmt.Fprintf(w, "}\n")
}

func writeEnum(w io.Writer, ifaceName string, e Enum) {
	enumName := toPascalCase(e.Name)

	fmt.Fprintf(w, "type %s%s uint32\n", ifaceName, enumName)

	fmt.Fprintf(w, "// %s%s : %s\n", ifaceName, enumName, e.Description.Summary)
	writeComment(w, e.Description.Text)
	fmt.Fprintf(w, "const (\n")
	for _, entry := range e.Entries {
		entryName := toPascalCase(entry.Name)

		if entry.Summary != "" {
			fmt.Fprintf(w, "// %s%s%s : %s\n", ifaceName, enumName, entryName, strings.Replace(entry.Summary, "\n", " ", -1))
		}
		fmt.Fprintf(w, "%s%s%s %s%s = %s\n", ifaceName, enumName, entryName, ifaceName, enumName, entry.Value)
	}
	fmt.Fprintf(w, ")\n")

	fmt.Fprintf(w, "func (e %s%s) Name() string {\n", ifaceName, enumName)
	fmt.Fprintf(w, "switch e {\n")
	for _, entry := range e.Entries {
		entryName := toPascalCase(entry.Name)

		fmt.Fprintf(w, "case %s%s%s:\n", ifaceName, enumName, entryName)
		fmt.Fprintf(w, "return \"%s\"\n", entry.Name)
	}
	fmt.Fprintf(w, "default:\n")
	fmt.Fprintf(w, "return \"\"\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")

	fmt.Fprintf(w, "func (e %s%s) String() string {\n", ifaceName, enumName)
	fmt.Fprintf(w, "return fmt.Sprintf(\"%%s=%%d\", e.Name(), e)\n")
	fmt.Fprintf(w, "}\n")
}

func writeEvent(w io.Writer, ifaceName string, e Event) {
	eventName := toPascalCase(e.Name)
	// eventNameLower := toLowerCamel(e.Name)

	e.Args = append(e.Args, Arg{
		Name: "proxy",
		Type: "object",
	})

	var gw bytes.Buffer

	// Event struct
	fmt.Fprintf(w, "// %s%sEvent : %s\n", ifaceName, eventName, strings.Replace(e.Description.Summary, "\n", " ", -1))
	writeComment(w, e.Description.Text)
	fmt.Fprintf(w, "type %s%sEvent struct {\n", ifaceName, eventName)
	for _, arg := range e.Args {
		argName := toPascalCase(arg.Name)

		var argType string
		switch arg.Type {
		case "object", "new_id":
			if arg.Interface != "" && !slices.Contains(*foreignObject, arg.Interface) {
				argType = "*" + toPascalCase(arg.Interface)
			} else {
				argType = "wayland.Proxy"
			}
		case "uint":
			if arg.Enum != "" {
				enumIface, enumName, ok := strings.Cut(arg.Enum, ".")
				if !ok {
					enumName = toPascalCase(enumIface)
					enumIface = ifaceName
				} else {
					enumName = toPascalCase(enumName)
					enumIface = toPascalCase(enumIface)
				}
				argType = enumIface + enumName
				break
			}
			fallthrough
		case "int", "fixed",
			"string", "array", "fd":
			argType = typeToGoTypeMap[arg.Type]
		}
		fmt.Fprintf(w, "%s %s\n", toCamelCase(arg.Name), argType)

		if arg.Summary != "" {
			fmt.Fprintf(&gw, "// %s %s\n", argName, strings.Replace(arg.Summary, "\n", " ", -1))
		} else if arg.Description.Summary != "" {
			fmt.Fprintf(&gw, "// %s %s\n", argName, strings.Replace(arg.Description.Summary, "\n", " ", -1))
		}
		writeComment(&gw, arg.Description.Text)
		fmt.Fprintf(&gw, "func (e *%s%sEvent) %s() %s {\n", ifaceName, eventName, argName, argType)
		fmt.Fprintf(&gw, "return e.%s\n", toCamelCase(arg.Name))
		fmt.Fprintf(&gw, "}\n")
	}
	fmt.Fprintf(w, "}\n")

	io.Copy(w, &gw)
}

func writeEventDispatcher(w io.Writer, imports map[string]bool, ifaceName string, v Interface, destructor string) {
	destructor = toPascalCase(destructor)
	fmt.Fprintf(w, "func (i *%s) Dispatch(msg *wayland.Message, drain chan<- wayland.Event) {\n", ifaceName)
	if len(v.Events) > 0 {
		fmt.Fprintf(w, "switch msg.Opcode {\n")
		for i, e := range v.Events {
			eventName := toPascalCase(e.Name)

			hasFd := false
			for _, arg := range e.Args {
				if arg.Type == "fd" {
					hasFd = true
					break
				}
			}

			fmt.Fprintf(w, "case %d:\n", i)
			if e.Type == "destructor" {
				fmt.Fprintf(w, "defer i.%s()\n", destructor)
			}
			fmt.Fprintf(w, "if (i.Handlers == nil || i.Handlers.On%s == nil) && drain == nil {\n", eventName)
			if hasFd {
				imports["syscall"] = true
				fmt.Fprintf(w, "for _, fd := range msg.FDs {\n")
				fmt.Fprintf(w, "syscall.Close(fd)\n")
				fmt.Fprintf(w, "}\n")
			}
			fmt.Fprintf(w, "return\n")
			fmt.Fprintf(w, "}\n")

			fmt.Fprintf(w, "e := &%s%sEvent{}\n", ifaceName, eventName)
			fmt.Fprintf(w, "e.proxy = i\n")

			if len(e.Args) > 0 {
				fmt.Fprintf(w, "r := wayland.NewMessageReader(i.Conn(), msg)\n")
			}

			for _, arg := range e.Args {
				argNameLower := toCamelCase(arg.Name)

				switch arg.Type {
				case "object", "new_id":
					if arg.Interface != "" && !slices.Contains(*foreignObject, arg.Interface) {
						argIface := toPascalCase(arg.Interface)

						fmt.Fprintf(w, "%s := r.ReadObject()\n", argNameLower)
						fmt.Fprintf(w, "if %s != nil {\n", argNameLower)
						fmt.Fprintf(w, "e.%s = %s.(*%s)\n", argNameLower, argNameLower, argIface)
						fmt.Fprintf(w, "}\n")
					} else {
						fmt.Fprintf(w, "e.%s = r.ReadObject()\n", argNameLower)
					}

				case "fd":
					fmt.Fprintf(w, "e.%s = r.ReadFd()\n", argNameLower)

				case "uint":
					if arg.Enum != "" {
						enumIface, enumName, ok := strings.Cut(arg.Enum, ".")
						if !ok {
							enumName = toPascalCase(enumIface)
							enumIface = ifaceName
						} else {
							enumName = toPascalCase(enumName)
							enumIface = toPascalCase(enumIface)
						}
						fmt.Fprintf(w, "e.%s = %s%s(r.ReadUint())\n", argNameLower, enumIface, enumName)
					} else {
						fmt.Fprintf(w, "e.%s = r.ReadUint()\n", argNameLower)
					}

				case "int":
					fmt.Fprintf(w, "e.%s = r.ReadInt()\n", argNameLower)

				case "fixed":
					fmt.Fprintf(w, "e.%s = r.ReadFixed()\n", argNameLower)

				case "string":
					fmt.Fprintf(w, "e.%s = r.ReadString()\n", argNameLower)

				case "array":
					fmt.Fprintf(w, "e.%s = r.ReadArray()\n", argNameLower)
				}
			}

			fmt.Fprintf(w, "if i.Handlers != nil && i.Handlers.On%s != nil && i.Handlers.On%s(e) {\n", eventName, eventName)
			fmt.Fprintf(w, "return\n")
			fmt.Fprintf(w, "}\n")
			fmt.Fprintf(w, "if drain != nil {\n")
			fmt.Fprintf(w, "drain <- e\n")
			fmt.Fprintf(w, "return\n")
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
		if *outputFile == "" {
			*outputFile = strings.TrimSuffix(srcname, ".xml") + ".go"
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

	interfaces := make(map[string]*Interface)
	for _, v := range protocol.Interfaces {
		interfaces[v.Name] = &v
	}

	// Header
	fmt.Fprintf(w, "// Generated by gowls\n")
	fmt.Fprintf(w, "// https://github.com/friedelschoen/wayland/cmd/gowls\n")
	if srcname != "" {
		fmt.Fprintf(w, "// XML file : %s\n", srcname)
	}
	fmt.Fprintf(w, "//\n")
	fmt.Fprintf(w, "// %s Protocol Copyright: \n", protocol.Name)
	writeComment(w, protocol.Copyright)
	fmt.Fprintf(w, "\n\n")
	fmt.Fprintf(w, "// Package %s contains wayland-protocol %s\n", *packageName, protocol.Name)
	fmt.Fprintf(w, "package %s\n", *packageName)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "import wayland \"%s\"\n", *runtime)

	// Interfaces
	var wb bytes.Buffer
	imports := make(map[string]bool)
	for _, v := range protocol.Interfaces {
		writeInterface(&wb, interfaces, imports, v)
	}

	for name, ok := range imports {
		if !ok {
			continue
		}
		fmt.Fprintf(w, "import \"%s\"\n", name)
	}
	io.Copy(w, &wb)

	fmt.Fprintf(w, "func Get%sInterface(name string) wayland.Proxy {\n", toPascalCase(protocol.Name))
	fmt.Fprintf(w, "switch (name) {\n")
	for _, v := range protocol.Interfaces {
		fmt.Fprintf(w, "case \"%s\":\n", v.Name)
		if len(v.Events) == 0 {
			fmt.Fprintf(w, "return New%s()\n", toPascalCase(v.Name))
		} else {
			fmt.Fprintf(w, "return New%s(nil)\n", toPascalCase(v.Name))
		}
	}
	fmt.Fprintf(w, "default:\n")
	fmt.Fprintf(w, "return nil\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")
}
