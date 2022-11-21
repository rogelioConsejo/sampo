package tagger

import (
	"errors"
	"os"
	"sampo/asset/item"
	"testing"
)

func TestTagger(t *testing.T) {
	setupTest(t)
	var testTagger Tagger = New()
	id := getIdFromTagger(testTagger.(tagger))
	validateIdAsUUID(t, id)
	err := testTagger.MakeQR(testQRFileLocation)
	checkIfQRWasCorrectlyCreated(t, err)
}

func checkIfQRWasCorrectlyCreated(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("could not get QR: %s", err.Error())
	}
	fileNotFound := checkIfFileIsNotFound(err, testQRFileLocation)
	if fileNotFound {
		t.Fatalf("qr file not found")
	}
}

func validateIdAsUUID(t *testing.T, id item.Id) {
	idLength, idLengthIsNotValid := checkIfLengthIsNotValid(id)
	formatIsNotValid := checkIfFormatIsNotValid(id)
	if idLengthIsNotValid {
		t.Fatalf("invalid item Id length, expected %d, got %d", expectedIdLength, idLength)
	}
	if formatIsNotValid {
		t.Fatalf("invalid format")
	}
}

func checkIfLengthIsNotValid(id item.Id) (int, bool) {
	tagLength := len(id)
	tagLengthIsNotValid := tagLength != expectedIdLength
	return tagLength, tagLengthIsNotValid
}

func getIdFromTagger(tagger tagger) item.Id {
	tag := tagger.getTag()
	id := tag.Item().Id
	return id
}

func setupTest(t *testing.T) {
	t.Parallel()
	t.Cleanup(func() {
		//_ = os.Remove(testQRFileLocation)
	})
}

func checkIfFileIsNotFound(err error, qrLocation FilePath) bool {
	_, err = os.Stat(string(qrLocation))
	return errors.Is(err, os.ErrNotExist)
}

func checkIfFormatIsNotValid(id item.Id) bool {
	return id[8] != '-' || id[13] != '-' || id[18] != '-'
}

const testQRFileLocation = "./test-qr.png"
const expectedIdLength = 36
