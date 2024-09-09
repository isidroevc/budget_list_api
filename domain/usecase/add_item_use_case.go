package add_item_use_case

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

func NewCreateItemInput(description string, unitaryPrice string, qty uint64) (*CreateItemInput, error) {
	parsedUnitaryPrice, err := decimal.NewFromString(unitaryPrice)
	if err != nil {
		return nil, fmt.Errorf("error while parsing UnitaryPrice %s", err.Error())
	}

	return &CreateItemInput{Description: description, UnitaryPrice: parsedUnitaryPrice, Qty: qty}, nil
}

type ToCreateItemInput interface {
	ToItem() *CreateItemInput
}

func (c *PartialItemBeforeCreation) ToCreateItemInput() *item_repository.CreateItemInput {
	return &item_repository.CreateItemInput{
		Description:  c.Description,
		UnitaryPrice: c.UnitaryPrice.String(),
		Qty:          c.Qty,
		Subtotal:     c.Subtotal.String(),
	}
}

type PartialItemBeforeCreation struct {
	Description  string
	UnitaryPrice decimal.Decimal
	Qty          uint64
	Subtotal     decimal.Decimal
}

func FromCreatedItemOutput(pi *item_repository.CreateItemOutput) (*CreatedItem, error) {
	parsedSubtotal, err := decimal.NewFromString(pi.Subtotal)
	if err != nil {
		return nil, fmt.Errorf("could not create CreatedItem, subtotal could not be parsed due to error %s", err.Error())
	}
	parsedUnitaryPrice, err := decimal.NewFromString(pi.UnitaryPrice)
	if err != nil {
		return nil, fmt.Errorf("could not create CreatedItem, subtotal could not be parsed due to error %s", err.Error())
	}
	return &CreatedItem{
		Id:           pi.Id,
		Description:  pi.Description,
		Qty:          pi.Qty,
		Subtotal:     parsedSubtotal,
		UnitaryPrice: parsedUnitaryPrice,
	}, nil
}

type CreatedItem struct {
	Id           int64
	Description  string
	UnitaryPrice decimal.Decimal
	Qty          uint64
	Subtotal     decimal.Decimal
}

type CreateItemResult struct {
	Item *CreatedItem
}

func CreateItem(input *CreateItemInput) (*CreateItemResult, error) {
	result := new(CreateItemResult)

	item := new(PartialItemBeforeCreation)
	item.Description = input.Description
	item.UnitaryPrice = input.UnitaryPrice
	item.Qty = input.Qty
	qty := decimal.NewFromUint64(input.Qty)
	item.Subtotal = qty.Mul(item.UnitaryPrice)
	item_repository, err := item_repository.New()

	if err != nil {
		return nil, fmt.Errorf("could not create item due to error %s", err.Error())
	}

	creationResult, err := item_repository.CreateItem(item.ToCreateItemInput())
	if err != nil {
		return nil, fmt.Errorf("could not create item due to error %s", err.Error())
	}

	createdItem, err := FromCreatedItemOutput(creationResult)
	if err != nil {
		return nil, fmt.Errorf("could not create item due to error %s", err.Error())
	}
	result.Item = createdItem
	return result, nil
}
