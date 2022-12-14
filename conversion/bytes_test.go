package conversion_test

import (
	"testing"

	"github.com/go-asphyxia/core/conversion"
)

func TestBytesToStringNoCopy(t *testing.T) {
	example := string("test")

	b := []byte(example)
	s := conversion.BytesToStringNoCopy(b)

	if s != example {
		t.Fatalf("%s != %s", s, example)
	}
}

type (
	TestBytesToStringNoCopyOf_string string
	TestBytesToStringNoCopyOf_byte   byte
)

func TestBytesToStringNoCopyOf(t *testing.T) {
	example := TestBytesToStringNoCopyOf_string("test")

	b := []TestBytesToStringNoCopyOf_byte(example)
	s := conversion.BytesToStringNoCopyOf[TestBytesToStringNoCopyOf_string](b)

	if s != example {
		t.Fatalf("%s != %s", s, example)
	}
}
