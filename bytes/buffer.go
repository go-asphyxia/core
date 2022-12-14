package bytes

import (
	"io"
	"unicode/utf8"
	"unsafe"

	"github.com/go-asphyxia/core"
	"github.com/go-asphyxia/core/conversion"
)

type (
	Buffer BufferOf[byte]

	BufferOf[B core.Byte] struct {
		Bytes []B
	}
)

func NewBuffer() (b *Buffer) {
	b = (*Buffer)(NewBufferOf[byte]())
	return
}

func NewBufferOf[B core.Byte]() (b *BufferOf[B]) {
	b = &BufferOf[B]{}
	return
}

func Clone(b *Buffer) (target *Buffer) {
	target = (*Buffer)(CloneOf((*BufferOf[byte])(b)))
	return
}

func CloneOf[B core.Byte](b *BufferOf[B]) (target *BufferOf[B]) {
	target = &BufferOf[B]{
		Bytes: make([]B, len(b.Bytes)),
	}

	copy(target.Bytes, b.Bytes)
	return
}

func Grow(b *Buffer, n int) {
	GrowOf((*BufferOf[byte])(b), n)
}

func GrowOf[B core.Byte](b *BufferOf[B], n int) {
	s := len(b.Bytes) + n

	if s <= cap(b.Bytes) {
		b.Bytes = b.Bytes[:s]
		return
	}

	temp := make([]B, s)

	copy(temp, b.Bytes)

	b.Bytes = temp
}

func Clip(b *Buffer, n int) {
	ClipOf((*BufferOf[byte])(b), n)
}

func ClipOf[B core.Byte](b *BufferOf[B], n int) {
	b.Bytes = b.Bytes[:len(b.Bytes)-n]
}

func Reset(b *Buffer) {
	ResetOf((*BufferOf[byte])(b))
}

func ResetOf[B core.Byte](b *BufferOf[B]) {
	b.Bytes = b.Bytes[:0]
}

func Close(b *Buffer) {
	CloseOf((*BufferOf[byte])(b))
}

func CloseOf[B core.Byte](b *BufferOf[B]) {
	b.Bytes = nil
}

func Copy(b *Buffer) (target []byte) {
	target = CopyOf((*BufferOf[byte])(b))
	return
}

func CopyOf[B core.Byte](b *BufferOf[B]) (target []B) {
	target = make([]B, len(b.Bytes))

	copy(target, b.Bytes)
	return
}

func Set(b *Buffer, source []byte) {
	SetOf((*BufferOf[byte])(b), source)
}

func SetOf[B core.Byte](b *BufferOf[B], source []B) {
	b.Bytes = append(b.Bytes[:0], source...)
}

func SetString(b *Buffer, source string) {
	SetStringOf((*BufferOf[byte])(b), source)
}

func SetStringOf[S core.String, B core.Byte](b *BufferOf[B], source S) {
	b.Bytes = append(b.Bytes[:0], conversion.StringToBytesNoCopyOf[B](source)...)
}

func Write(b *Buffer, source []byte) (n int, err error) {
	n, err = WriteOf((*BufferOf[byte])(b), source)
	return
}

func WriteOf[B core.Byte](b *BufferOf[B], source []B) (n int, err error) {
	n = len(source)

	b.Bytes = append(b.Bytes, source...)
	return
}

func WriteByte(b *Buffer, source byte) (err error) {
	err = WriteByteOf((*BufferOf[byte])(b), source)
	return
}

func WriteByteOf[B core.Byte](b *BufferOf[B], source B) (err error) {
	b.Bytes = append(b.Bytes, source)
	return
}

func WriteRune(b *Buffer, source rune) (n int, err error) {
	n, err = WriteRuneOf((*BufferOf[byte])(b), source)
	return
}

func WriteRuneOf[R core.Rune, B core.Byte](b *BufferOf[B], source R) (n int, err error) {
	l := len(b.Bytes)
	c := cap(b.Bytes)

	s := l + utf8.UTFMax

	if s > c {
		temp := make([]B, s)

		copy(temp, b.Bytes)

		b.Bytes = temp
	}

	slice := *(*[]byte)(unsafe.Pointer(&b.Bytes))

	n = utf8.EncodeRune(slice[l:s], rune(source))

	b.Bytes = b.Bytes[:(l + n)]
	return
}

func WriteString(b *Buffer, source string) (n int, err error) {
	n, err = WriteStringOf((*BufferOf[byte])(b), source)
	return
}

func WriteStringOf[S core.String, B core.Byte](b *BufferOf[B], source S) (n int, err error) {
	n = len(source)

	b.Bytes = append(b.Bytes, conversion.StringToBytesNoCopyOf[B](source)...)
	return
}

func ReadFrom(b *Buffer, source core.Reader) (n int64, err error) {
	n, err = ReadFromOf((*BufferOf[byte])(b), (core.ReaderOf[byte])(source))
	return
}

func ReadFromOf[B core.Byte](b *BufferOf[B], source core.ReaderOf[B]) (n int64, err error) {
	l := len(b.Bytes)
	c := cap(b.Bytes)
	r := 0

	for {
		if l == c {
			c = (c + 16) * 2

			temp := make([]B, c)
			copy(temp, b.Bytes)

			b.Bytes = temp
		}

		r, err = source.ReadOf(b.Bytes[l:c])

		n += int64(r)
		l += r

		b.Bytes = b.Bytes[:l]

		if err != nil || l < c {
			if err == io.EOF {
				err = nil
			}

			return
		}
	}
}
