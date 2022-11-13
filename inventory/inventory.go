package inventory

import "errors"

func New() Inventory {
	return &inventory{
		stock: map[Item]Stock{},
	}
}

type Inventory interface {
	AddStock(Item, Amount) error
	GetStock(Item) Amount
}

type inventory struct {
	stock map[Item]Stock
}

func (i *inventory) GetStock(item Item) Amount {
	return i.stock[item].Amount
}

func (i *inventory) AddStock(item Item, a Amount) error {
	if a < 0 {
		return negativeAmountError
	}
	i.stock[item] = Stock{
		Amount: i.stock[item].Amount + a,
	}
	return nil
}

var negativeAmountError = errors.New("cannot add a negative amount to a stock")
