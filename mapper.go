package wayland

import (
	"runtime"
	"syscall"
)

type MappingProtection int

const (
	ProtNone  MappingProtection = syscall.PROT_NONE
	ProtRead  MappingProtection = syscall.PROT_READ
	ProtWrite MappingProtection = syscall.PROT_WRITE
	ProtExec  MappingProtection = syscall.PROT_EXEC
)

type MappingFlags int

const (
	Map32bit      MappingFlags = syscall.MAP_32BIT
	MapAnon       MappingFlags = syscall.MAP_ANON
	MapAnonymous  MappingFlags = syscall.MAP_ANONYMOUS
	MapDenywrite  MappingFlags = syscall.MAP_DENYWRITE
	MapExecutable MappingFlags = syscall.MAP_EXECUTABLE
	MapFile       MappingFlags = syscall.MAP_FILE
	MapFixed      MappingFlags = syscall.MAP_FIXED
	MapGrowsdown  MappingFlags = syscall.MAP_GROWSDOWN
	MapHugetlb    MappingFlags = syscall.MAP_HUGETLB
	MapLocked     MappingFlags = syscall.MAP_LOCKED
	MapNonblock   MappingFlags = syscall.MAP_NONBLOCK
	MapNoreserve  MappingFlags = syscall.MAP_NORESERVE
	MapPopulate   MappingFlags = syscall.MAP_POPULATE
	MapPrivate    MappingFlags = syscall.MAP_PRIVATE
	MapShared     MappingFlags = syscall.MAP_SHARED
	MapStack      MappingFlags = syscall.MAP_STACK
	MapType       MappingFlags = syscall.MAP_TYPE
)

type MappingEvent interface {
	Fd() int
	Size() uint32
}

// MapMemory maps a block of memory given by the Compositor.
// The result has a finializer on it which unmaps the block if out of range.
func MapMemory(ev MappingEvent, prot MappingProtection, flags MappingFlags) ([]byte, error) {
	data, err := syscall.Mmap(ev.Fd(), 0, int(ev.Size()), int(prot), int(flags))
	runtime.SetFinalizer(&data, func(ptr *[]byte) { syscall.Munmap(*ptr) })
	return data, err
}
