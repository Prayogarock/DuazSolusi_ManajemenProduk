package handler

import "duaz/features/items"

type ItemResponseAll struct {
	Name             string `json:"name"`
	Description_Item string `json:"description_item"`
	Price            uint   `json:"price"`
	Stock            uint   `json:"stock"`
}

func ItemCoreToResponseAll(input items.ItemCore) ItemResponseAll {
	var itemResp = ItemResponseAll{
		Name:             input.Name,
		Description_Item: input.Description_Item,
		Price:            input.Price,
		Stock:            input.Stock,
	}
	return itemResp
}
