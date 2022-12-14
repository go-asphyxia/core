package conversion

import (
	"unsafe"

	"github.com/go-asphyxia/core"
	"github.com/go-asphyxia/core/runtime"
)

func BytesToStringNoCopy(source []byte) (target string) {
	target = BytesToStringNoCopyOf[string](source)
	return
}

func BytesToStringNoCopyOf[S core.String, B core.Byte](source []B) (target S) {
	sh := (*runtime.SliceOf[B])(unsafe.Pointer(&source))
	th := (*runtime.StringOf[B])(unsafe.Pointer(&target))

	th.Data = sh.Data
	th.Length = sh.Length
	return
}
