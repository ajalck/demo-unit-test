package order

import "github.com/Rhymond/go-money"

type Order struct {
	ID    string
	Items []Item
}
type Item struct {
	ID        string
	Quantity  uint
	UnitPrice *money.Money
}

func (o *Order) ComputeTotal() (*money.Money, error) {

}
