package item_repository

import (
	"database/sql"
	"fmt"

	supabase "github.com/isidroevc/blist_api/external_services/database"
)

type ItemRepository struct {
	db *sql.DB
}

type CreateItemInput struct {
	Description  string
	UnitaryPrice string
	Qty          uint64
	Subtotal     string
}

type CreateItemOutput struct {
	Id           int64
	Description  string
	UnitaryPrice string
	Qty          uint64
	Subtotal     string
}

type ToCreateItemInput interface {
	ToCreateItemInput() *ToCreateItemInput
}

func (it *ItemRepository) CreateItem(input *CreateItemInput) (*CreateItemOutput, error) {
	var lastInsertedId int64
	err := it.db.QueryRow(`INSERT INTO item
	(
		description,
		unitary_price,
		sub_total,
		qty
	)
		
	VALUES
	(
		$1,
		$2,
		$3,
		$4
	)
	RETURNING id`,
		input.Description,
		input.UnitaryPrice,
		input.Subtotal,
		input.Qty,
	).Scan(&lastInsertedId)

	if err != nil {
		return nil, fmt.Errorf("could not insert item due to error: %s", err.Error())
	}

	output := CreateItemOutput{
		Id:           lastInsertedId,
		Description:  input.Description,
		Qty:          input.Qty,
		UnitaryPrice: input.UnitaryPrice,
		Subtotal:     input.Subtotal,
	}

	return &output, nil
}

func New() (*ItemRepository, error) {
	db, err := supabase.GetConnection()
	if err != nil {
		return nil, fmt.Errorf("could not create item repository due to error: %s", err.Error())
	}
	return &ItemRepository{db: db}, nil
}
