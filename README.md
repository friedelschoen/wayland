# Wayland for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/friedelschoen/wayland/wayland.svg)](https://pkg.go.dev/github.com/friedelschoen/wayland)

This project provides a **Go runtime library** and a **code generator** for the [Wayland](https://wayland.freedesktop.org/) protocol.
It enables you to write native Wayland clients directly in Go, without depending on C bindings.

## Features

* Low-level **Wayland connection handling** (`Conn`), including:
  * Object registration/unregistration
  * Event dispatching
  * File descriptor passing
  * Shared memory (mmap) support
* Interfaces for Wayland protocol objects (`Proxy`, `Event`, `EventHandlerFunc`).
* A generator tool `gowls`, which:
  * Reads official Wayland XML protocol files
  * Produces Go bindings for requests, events, and enums
  * Generates strongly typed handlers and message marshalling
  * Automatically formats code using `gofmt`

## Components

### Runtime Library (`package wayland`)

Implements the building blocks of a Wayland client in Go:

* **Connection Management**
  `Connect()` opens a Wayland socket (`WAYLAND_DISPLAY`) and starts the event loop.

* **Proxies & Events**
  Each protocol object implements the `Proxy` interface. Events are dispatched to registered handlers or an event drain channel.

* **Message I/O**
  Encoding (`MessageWriter`) and decoding (`MessageReader`) of Wayland messages, with support for:

  * integers, strings, arrays, fixed-point values
  * proxy references
  * file descriptors (SCM\_RIGHTS)

* **Shared Memory (shm)**
  Helpers for mapping/unmapping buffers with `syscall.Mmap`.

### Scanner Tool (`cmd/gowls`)

The generator replaces `wayland-scanner` from the C ecosystem.

```bash
go run ./cmd/gowls \
  -p myproto \
  path/to/protocol.xml > myproto.go
```

## Installation

```bash
go get github.com/friedelschoen/wayland
```

To build the scanner:

```bash
cd cmd/gowls
go install
```

## Status

* Some advanced features (error handling, protocol extensions) are still being fleshed out.
* API stability is not guaranteed yet.

## License

MIT — see [LICENSE](./LICENSE).
