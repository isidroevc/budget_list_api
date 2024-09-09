package models

import add_item_use_case "github.com/isidroevc/blist_api/domain/usecase"

type CreateItemInput struct {
	Description  string `json:"description"`
	UnitaryPrice string `json:"unitaryPrice"`
	Qty          uint64 `json:"qty"`
}

type UpdateItemInput struct {
	Id           int64
	Description  string
	UnitaryPrice string
	Qty          uint64
}

func NewCreateItemInput(description string, unitaryPrice string, qty uint64) (*CreateItemInput, error) {

	return &CreateItemInput{
		Description:  description,
		UnitaryPrice: unitaryPrice,
		Qty:          qty,
	}, nil
}

func (c *CreateItemInput) ToCreateItemInput() (*add_item_use_case.CreateItemInput, error) {
	ci, err := add_item_use_case.NewCreateItemInput(
		c.Description,
		c.UnitaryPrice,
		c.Qty,
	)
	return ci, err
}
