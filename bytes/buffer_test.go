package bytes_test

import (
	"testing"

	"github.com/go-asphyxia/core/bytes"
	"github.com/go-asphyxia/core/conversion"
)

func TestWriteRune(t *testing.T) {
	b := &bytes.Buffer{}

	_, err := bytes.WriteRune(b, 'T')
	if err != nil {
		t.Fatal(err)
	}

	s := conversion.BytesToStringNoCopy(b.Bytes)

	if s != "T" {
		t.Fatalf("%s != T", s)
	}
}
