package jemalloc

// #cgo         CFLAGS: -std=gnu99
// #cgo       CPPFLAGS: -D_REENTRANT
// #cgo linux CPPFLAGS: -D_GNU_SOURCE
// #cgo        LDFLAGS: -lm
// #cgo linux  LDFLAGS: -lrt
// #include <jemalloc/jemalloc.h>
// void *zmalloc(size_t size) {
//   return malloc(size);
// }
import "C"

import (
	"unsafe"

	"github.com/choleraehyq/fastercgo"
)

func Calloc(count, size int) unsafe.Pointer {
	return C.calloc(C.size_t(count), C.size_t(size))
}

func Malloc(size int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(fastercgo.FasterUnsafeCall1WithRet1(C.zmalloc, uint64(size))))
}

func OriginMalloc(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

func Valloc(size int) unsafe.Pointer {
	return C.valloc(C.size_t(size))
}

func Realloc(ptr unsafe.Pointer, size int) unsafe.Pointer {
	return C.realloc(ptr, C.size_t(size))
}

func Free(ptr unsafe.Pointer) {
	fastercgo.FasterUnsafeCall1WithRet1(C.free, uint64(uintptr(ptr)))
}
