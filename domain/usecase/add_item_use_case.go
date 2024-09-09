package usecase

import (
	"fmt"

	item_repository "github.com/isidroevc/blist_api/action/data/repositories"
	"github.com/isidroevc/blist_api/domain/usecase/models"
	"github.com/shopspring/decimal"
)

type CreateItemResult struct {
	Item *models.Item
}

func CreateItem(input *models.CreateItemInput) (*CreateItemResult, error) {
	result := new(CreateItemResult)

	item := new(models.Item)
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

	item.Id = creationResult.Id
	result.Item = item
	return result, nil
}
