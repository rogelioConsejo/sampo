package location

import (
	"errors"
	"testing"
)

func TestNew_NameCannotBeEmpty(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Fatalf("no error thrown on empty name")
	}
	if !errors.Is(err, emptyNameError) {
		t.Fatalf("invalid error thrown on empty name")
	}
}
