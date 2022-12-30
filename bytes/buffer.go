package bytes

import (
	"io"
	"unicode/utf8"
)

type (
	Buffer struct {
		List []byte
	}
)

func (B *Buffer) Clone() *Buffer {
	return &Buffer{
		List: append([]byte(nil), B.List...),
	}
}

func (B *Buffer) Reset() {
	B.List = B.List[:0]
}

func (B *Buffer) Set(source []byte) {
	B.List = append(B.List[:0], source...)
}

func (B *Buffer) SetString(source string) {
	B.List = append(B.List[:0], source...)
}

func (B *Buffer) Write(source []byte) (n int, err error) {
	n = len(source)

	B.List = append(B.List, source...)
	return
}

func (B *Buffer) WriteByte(source byte) (err error) {
	B.List = append(B.List, source)
	return
}

func (B *Buffer) WriteRune(source rune) (n int, err error) {
	l := len(B.List)
	c := cap(B.List)

	sum := l + utf8.UTFMax

	if sum > c {
		reallocation := make([]byte, sum)
		copy(reallocation, B.List)

		B.List = reallocation
	}

	n = utf8.EncodeRune(B.List[l:sum], source)

	B.List = B.List[:(l + n)]
	return
}

func (B *Buffer) WriteString(source string) (n int, err error) {
	n = len(source)

	B.List = append(B.List, source...)
	return
}

func (buffer *Buffer) ReadFrom(source io.Reader) (n int64, err error) {
	l := len(buffer.List)
	c := cap(buffer.List)
	r := 0

	reallocation := []byte(nil)

cycle:

	r, err = source.Read(buffer.List[l:c])

	n += int64(r)
	l += r

	buffer.List = buffer.List[:l]

	if err == nil {
		if l == c {
			goto cycle
		}

		c = (c + 64) << 2

		reallocation = make([]byte, c)
		copy(reallocation, buffer.List)

		buffer.List = reallocation

		goto cycle
	} else if err == io.EOF {
		return n, nil
	}

	return
}
