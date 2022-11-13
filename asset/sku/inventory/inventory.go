package inventory

import (
	"errors"
	"fmt"
	"sampo/asset/item"
)

func New() Inventory {
	return &inventory{
		stock: map[item.Item]Stock{},
	}
}

type Inventory interface {
	AddStock(item.Item, Amount)
	GetStock(item.Item) Amount
	RemoveStock(item item.Item, amount Amount) error
}

type inventory struct {
	stock map[item.Item]Stock
}

func (i *inventory) RemoveStock(item item.Item, amount Amount) error {
	itemStockAmount := i.stock[item].Amount
	if amount > itemStockAmount {
		return fmt.Errorf("%w: you tried to remove %d from stock, but you only have %d left",
			insufficientAmountError, amount, itemStockAmount)
	}
	i.stock[item] = Stock{
		Amount: itemStockAmount - amount,
	}
	return nil
}

func (i *inventory) GetStock(item item.Item) Amount {
	return i.stock[item].Amount
}

func (i *inventory) AddStock(item item.Item, a Amount) {
	i.stock[item] = Stock{
		Amount: i.stock[item].Amount + a,
	}
}

var insufficientAmountError = errors.New("insufficient amount, cannot remove")
