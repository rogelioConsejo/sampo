package tag

import (
	"sampo/asset/item"
	"testing"
)

func TestTag_Item(t *testing.T) {
	var testItem = item.Item{Id: "1234"}
	var testTag Tag = New(testItem)

	var tagItem item.Item = testTag.Item()

	if testItem != tagItem {
		t.Fatalf("invalid item, expected %+v, retrieve %+v", testItem, tagItem)
	}
}

func TestTag_Price(t *testing.T) {
	var testItem = item.Item{Id: "1234"}
	var testTag Tag = New(testItem)
	var testPrice Price = 100 * Fraction

	testTag.SetPrice(testPrice)

	if retrievedPrice := testTag.Price(); retrievedPrice != testPrice {
		t.Fatalf("invalid amount retrieve, expected `%d`, retrieved `%d`", testPrice, retrievedPrice)
	}
}

func TestTag_Notes(t *testing.T) {
	var testItem = item.Item{Id: "1234"}
	var testTag Tag = New(testItem)
	var noteKey NoteKey = "some-key"
	var noteValue NoteValue = "some-value"

	testTag.AddNote(noteKey, noteValue)

	if retrievedValue, exists := testTag.GetNote(noteKey); !exists || retrievedValue != noteValue {
		if !exists {
			t.Fatal("note not found")
		}
		t.Fatalf("did not retrieve note correctly, expected `%s`, retrieved `%s`", noteValue, retrievedValue)
	}
}
