//Package product provide functionality to manage a product.
package product

import "github.com/gasprog/asa/dbi"

type Item struct {
	ID string
}

func (i *Item) Insert(ctx dbi.DBContext) error {
	return nil
}

func (i *Item) Update(ctx dbi.DBContext) error {
	return nil
}

func (i *Item) Delete(ctx dbi.DBContext) error {
	return nil
}

func (i *Item) GetByID(ctx dbi.DBContext, id string) error {
	return nil
}
