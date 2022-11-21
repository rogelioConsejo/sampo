// Package item is used to represent a non-fungible asset
package item

func New(id Id) Item {
	return Item{id}
}

type Item struct {
	Id `json:"id"`
}

type Id string
