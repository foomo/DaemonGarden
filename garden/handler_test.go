package garden

import (
	"testing"
)

func TestDecodeURINoErrorLength(t *testing.T) {
	decoded, err := decodeURI("/cmd/test/bla")
	if err != nil {
		t.Fatal("unexpected decoding error", err)
	}
	expected := []string{"cmd", "test", "bla"}
	if len(expected) != len(decoded) {
		t.Fatal("length mismatch", len(expected), len(decoded))
	}
	for i := 0; i < len(expected); i++ {
		if decoded[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestDecodeURIErrorDecode(t *testing.T) {
	decoded, err := decodeURI("/cmd/test/bla%2")
	if err == nil {
		t.Fatal("there should have been a decoding error, but we got decoded instead", decoded)
	}
}
