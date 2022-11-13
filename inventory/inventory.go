package inventory

import (
	"errors"
	"fmt"
)

func New() Inventory {
	return &inventory{
		stock: map[Item]Stock{},
	}
}

type Inventory interface {
	AddStock(Item, Amount)
	GetStock(Item) Amount
	RemoveStock(item Item, amount Amount) error
}

type inventory struct {
	stock map[Item]Stock
}

func (i *inventory) RemoveStock(item Item, amount Amount) error {
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

func (i *inventory) GetStock(item Item) Amount {
	return i.stock[item].Amount
}

func (i *inventory) AddStock(item Item, a Amount) {
	i.stock[item] = Stock{
		Amount: i.stock[item].Amount + a,
	}
}

var insufficientAmountError = errors.New("insufficient amount, cannot remove")
