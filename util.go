// 12 december 2015

package ui

import (
	"unsafe"
)

// #include "pkgui.h"
import "C"

//export pkguiAlloc
func pkguiAlloc(n C.size_t) unsafe.Pointer {
	// cgo turns C.malloc() into a panic-on-OOM version; use it
	// TODO make sure it zero-initializes too
	return C.malloc(n)
}

func freestr(str *C.char) {
	C.free(unsafe.Pointer(str))
}

func tobool(b C.int) bool {
	return b != 0
}

func frombool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
