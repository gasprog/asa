//Package sales provide functionality to manage sales transaction.
package sales

import "github.com/gasprog/asa/dbi"

type Order struct {
	ID string
}

type OrderDetail struct {
	ID string
}

func (o *Order) Insert(ctx dbi.DBContext) error {
	return nil
}

func (o *Order) Delete(ctx dbi.DBContext) error {
	return nil
}

func (o *Order) Update(ctx dbi.DBContext) error {
	return nil
}

func (o *Order) Get(ctx dbi.DBContext) error {
	return nil
}
