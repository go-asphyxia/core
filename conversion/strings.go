package conversion

import (
	"unsafe"

	"github.com/go-asphyxia/core"
	"github.com/go-asphyxia/core/runtime"
)

func StringToBytesNoCopy(source string) (target []byte) {
	target = StringToBytesNoCopyOf[byte](source)
	return
}

func StringToBytesNoCopyOf[B core.Byte, S core.String](source S) (target []B) {
	sh := (*runtime.String)(unsafe.Pointer(&source))
	th := (*runtime.Slice)(unsafe.Pointer(&target))

	th.Data = sh.Data
	th.Length = sh.Length
	th.Capacity = sh.Length
	return
}
