package data

import (
	"duaz/features/items"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name             string `gorm:"name;not null"`
	Description_Item string `gorm:"description_item;not null"`
	Price            uint   `gorm:"price;not null"`
	Stock            uint   `gorm:"stock;not null"`
	UserID           uint
}

func ItemCoreToModel(input items.ItemCore) Item {
	var itemModel = Item{
		Model:            gorm.Model{},
		Name:             input.Name,
		Description_Item: input.Description_Item,
		Price:            input.Price,
		Stock:            input.Stock,
	}
	return itemModel
}

func ItemModelToCore(input Item) items.ItemCore {
	var itemCore = items.ItemCore{
		ID:               input.ID,
		Name:             input.Name,
		Description_Item: input.Description_Item,
		Price:            input.Price,
		Stock:            input.Stock,
	}
	return itemCore
}
