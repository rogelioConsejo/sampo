// Package item is used to represent a non-fungible asset
package item

type Item struct {
	Id `json:"id"`
}

type Id string
