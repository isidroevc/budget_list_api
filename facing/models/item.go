package models

import "github.com/isidroevc/blist_api/domain/usecase/models"

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

func (c *CreateItemInput) ToCreateItemInput() (*models.CreateItemInput, error) {
	ci, err := models.NewCreateItemInput(
		c.Description,
		c.UnitaryPrice,
		c.Qty,
	)
	return ci, err
}
