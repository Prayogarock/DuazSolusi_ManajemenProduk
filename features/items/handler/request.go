package handler

import "duaz/features/items"

type ItemRequest struct {
	Name             string `json:"name" form:"name" validate:"required"`
	Description_Item string `json:"description_item" form:"description_item"`
	Price            uint   `json:"price" form:"price" validate:"required"`
	Stock            uint   `json:"stock" form:"stock" validate:"required"`
}

func ItemRequestToCore(input ItemRequest) items.ItemCore {
	return items.ItemCore{
		Name:             input.Name,
		Description_Item: input.Description_Item,
		Price:            input.Price,
		Stock:            input.Stock,
	}
}
