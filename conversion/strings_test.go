package conversion_test

import (
	"bytes"
	"testing"

	"github.com/go-asphyxia/core/conversion"
)

func TestStringToBytesNoCopy(t *testing.T) {
	example := []byte("test")

	s := string(example)
	b := conversion.StringToBytesNoCopy(s)

	if bytes.Equal(b, example) {
		t.Fatalf("%s != %s", b, example)
	}
}

type (
	TestStringToBytesNoCopyOf_byte   byte
	TestStringToBytesNoCopyOf_string string
)

func TestStringToBytesNoCopyOf(t *testing.T) {
	example := []TestStringToBytesNoCopyOf_byte("test")

	s := TestStringToBytesNoCopyOf_string(example)
	b := conversion.StringToBytesNoCopyOf[TestStringToBytesNoCopyOf_byte](s)

	if string(b) != string(example) {
		t.Fatalf("%s != %s", b, example)
	}
}
