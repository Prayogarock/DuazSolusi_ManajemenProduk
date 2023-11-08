package service

import "duaz/features/items"

type ItemService struct {
	itemData items.ItemDataInterface
}

func New(repo items.ItemDataInterface) items.ItemServiceInterface {
	return &ItemService{
		itemData: repo,
	}
}

func (service *ItemService) Create(UserID uint, input items.ItemCore) error {
	return service.itemData.Insert(UserID, input)
}

func (service *ItemService) GetAllItem(page, item uint, search_name string) ([]items.ItemCore, bool, error) {
	result, count, err := service.itemData.ReadAll(page, item, search_name)

	next := true
	var pages int64
	if item != 0 {
		pages = count / int64(item)
		if count%int64(item) != 0 {
			pages += 1
		}
		if page == uint(pages) {
			next = false
		}
	}

	return result, next, err
}

func (service *ItemService) Update(UserID uint, id uint, input items.ItemCore) error {
	return service.itemData.UpdateDataItem(UserID, id, input)
}

func (service *ItemService) Delete(UserID uint, id uint) error {
	return service.itemData.Delete(UserID, id)
}
