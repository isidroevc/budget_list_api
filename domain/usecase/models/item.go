package models

import (
	"fmt"

	item_repository "github.com/isidroevc/blist_api/action/data/repositories"
	"github.com/shopspring/decimal"
)

type CreateItemInput struct {
	Description  string
	UnitaryPrice decimal.Decimal
	Qty          uint64
}

type ToCreateItemInput interface {
	ToItem() *CreateItemInput
}

func NewCreateItemInput(description string, unitaryPrice string, qty uint64) (*CreateItemInput, error) {
	parsedUnitaryPrice, err := decimal.NewFromString(unitaryPrice)

	if err != nil {
		return nil, fmt.Errorf("string '%s' is not valid for unitary price", unitaryPrice)
	}

	return &CreateItemInput{
		Description:  description,
		UnitaryPrice: parsedUnitaryPrice,
		Qty:          qty,
	}, nil
}
func (c *Item) ToCreateItemInput() *item_repository.CreateItemInput {
	return &item_repository.CreateItemInput{
		Description:  c.Description,
		UnitaryPrice: c.UnitaryPrice.String(),
		Qty:          c.Qty,
		Subtotal:     c.Subtotal.String(),
	}
}

type Item struct {
	Id           int64
	Description  string
	UnitaryPrice decimal.Decimal
	Qty          uint64
	Subtotal     decimal.Decimal
}
