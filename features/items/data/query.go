package data

import (
	"duaz/features/items"
	"errors"

	"gorm.io/gorm"
)

type ItemQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) items.ItemDataInterface {
	return &ItemQuery{
		db: db,
	}
}

func (repo *ItemQuery) Insert(UserID uint, input items.ItemCore) error {

	var itemModel = ItemCoreToModel(input)
	itemModel.UserID = UserID

	tx := repo.db.Create(&itemModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ItemQuery) ReadAll(page, itemPerPage uint, searchName string) ([]items.ItemCore, int64, error) {
	var itemData []Item
	var totalCount int64

	if page == 0 && itemPerPage == 0 {
		tx := repo.db

		if searchName != "" {
			tx = tx.Where("name LIKE ?", "%"+searchName+"%")
		}
		tx.Find(&itemData)
	} else {

		offset := int((page - 1) * itemPerPage)

		query := repo.db.Offset(offset).Limit(int(itemPerPage))

		if searchName != "" {
			query = query.Where("name LIKE ?", "%"+searchName+"%")
		}

		tx := query.Find(&itemData)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
	}

	var itemCore []items.ItemCore
	for _, value := range itemData {
		itemCore = append(itemCore, items.ItemCore{
			ID:               value.ID,
			Name:             value.Name,
			Description_Item: value.Description_Item,
			Price:            value.Price,
			Stock:            value.Stock,
		})
	}

	repo.db.Model(&Item{}).Count(&totalCount)

	return itemCore, totalCount, nil
}

func (repo *ItemQuery) UpdateDataItem(UserID uint, id uint, input items.ItemCore) error {
	var item Item
	tx := repo.db.Where("id = ? AND user_id = ?", id, UserID).First(&item)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	updatedItem := ItemCoreToModel(input)

	tx = repo.db.Model(&item).Updates(updatedItem)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}

func (repo *ItemQuery) Delete(UserID, id uint) error {
	tx := repo.db.Where("id = ? AND user_id = ?", id, UserID).Delete(&Item{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}
