package jsext

import (
	"syscall/js"
	"unsafe"
)

func Module() js.Value {
	v := js.Global()
	p := (*uint64)(unsafe.Pointer(&v))
	if *p != 0x7FF80000<<32|5 {
		panic("internal error")
	}
	*p = 0x7FF80000<<32 | 8
	return v
}
