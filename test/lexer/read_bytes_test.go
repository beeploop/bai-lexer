package test

import (
	"bytes"
	"testing"

	"github.com/BeepLoop/bai-lexer/lexer"
)

func TestReadBytes(t *testing.T) {
	result, err := lexer.ReadBytes("read_bytes_input")
	if err != nil {
		t.Fatalf(err.Error())
	}

	expected := []byte("hello\n")

	if bytes.Equal(result, expected) == false {
		t.Fatalf("Expected: %s, Got: %s\n", expected, result)
	}
}
