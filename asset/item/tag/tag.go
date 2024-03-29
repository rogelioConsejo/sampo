// Package tag is a decorator used to add a price or any other annex information to an item.Item
package tag

import (
	"sampo/asset/item"
)

// New creates a new tag from an inventory.Item
func New(i item.Item) Tag {
	return Tag{
		item:  i,
		notes: map[NoteKey]NoteValue{},
	}
}

// Tag contains an inventory.Item and may contain a Price and Notes
type Tag struct {
	item  item.Item
	price Price
	notes Notes
}

// Item is the item that the Tag is tracking
func (t *Tag) Item() item.Item {
	return t.item
}

// SetPrice sets the price of the item that this Tag is tracking
func (t *Tag) SetPrice(p Price) {
	t.price = p
}

// Price is the price of the item that the Tag is tracking
func (t *Tag) Price() Price {
	return t.price
}

// AddNote can be used to add any kind of additional information in the form of a key-value pair of strings.
func (t *Tag) AddNote(key NoteKey, value NoteValue) {
	t.notes[key] = value
}

// GetNote returns the NoteValue associated with the provided NoteKey for this tag and whether the value exists or not.
func (t *Tag) GetNote(key NoteKey) (value NoteValue, exists bool) {
	value, exists = t.notes[key]
	return
}

// Price should be given in "fractions", that is expressed in the smallest fractional value for the used coin
// For example if you use dollars and cents, one cent is the fraction. A Price cannot be smaller than 1 Fraction.
type Price uint

const Fraction Price = 1

type Notes map[NoteKey]NoteValue

// NoteKey is a key for a tag note, which is a key-value pair.
type NoteKey string

// NoteValue is a value for a tag note, which is a key-value pair.
type NoteValue string
