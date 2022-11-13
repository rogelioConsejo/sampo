package tagger

import (
	"errors"
	"os"
	"testing"
)

func TestTagger(t *testing.T) {
	t.Parallel()
	const testQRFileLocation = "./test-qr.png"
	t.Cleanup(func() {
		_ = os.Remove(testQRFileLocation)
	})

	var testTagger Tagger = New()
	tag := testTagger.Tag()
	id := tag.Item().Id

	if tagLength := len(id); tagLength != 36 {
		t.Fatalf("invalid item Id length, expected 10, got %d", tagLength)
	}
	if id[8] != '-' || id[13] != '-' || id[18] != '-' {
		t.Fatalf("invalid format")
	}

	var qrLocation FilePath = testQRFileLocation
	err := testTagger.MakeQR(qrLocation)
	if err != nil {
		t.Fatalf("could not get QR: %s", err.Error())
	}
	
	if _, err := os.Stat(string(qrLocation)); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("qr file not found")
	}
}
