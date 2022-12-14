package runtime

import (
	"unsafe"

	"github.com/go-asphyxia/core"
)

type (
	String struct {
		Data   unsafe.Pointer
		Length int
	}

	StringOf[T core.Byte] struct {
		Data   *T
		Length int
	}

	Slice struct {
		Data     unsafe.Pointer
		Length   int
		Capacity int
	}

	SliceOf[T any] struct {
		Data     *T
		Length   int
		Capacity int
	}

	TypeFlag uint8

	NameOff int32
	TypeOff int32
	TextOff int32

	Type struct {
		Size                  uintptr
		PointerData           uintptr
		Hash                  uint32
		TypeFlag              TypeFlag
		Align                 uint8
		FieldAlign            uint8
		Kind                  uint8
		Equal                 func(unsafe.Pointer, unsafe.Pointer) bool
		GarbageCollectionData *byte
		String                NameOff
		PointerToThis         TypeOff
	}

	Interface struct {
		Type *Type
		Data unsafe.Pointer
	}

	InterfaceOf[T any] struct {
		Type *Type
		Data *T
	}
)

const (
	TypeFlagUncommon      TypeFlag = 1 << 0
	TypeFlagExtraStar     TypeFlag = 1 << 1
	TypeFlagNamed         TypeFlag = 1 << 2
	TypeFlagRegularMemory TypeFlag = 1 << 3
)
