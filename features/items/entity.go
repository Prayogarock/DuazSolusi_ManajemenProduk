package items

type ItemCore struct {
	ID               uint
	Name             string
	Description_Item string
	Price            uint
	Stock            uint
}

type ItemDataInterface interface {
	Insert(UserID uint, input ItemCore) error
	ReadAll(page, item uint, search_name string) ([]ItemCore, int64, error)
	UpdateDataItem(UserID uint, id uint, input ItemCore) error
	Delete(UserID uint, id uint) error
}

type ItemServiceInterface interface {
	Create(UserID uint, input ItemCore) error
	GetAllItem(page, item uint, search_name string) ([]ItemCore, bool, error)
	Update(UserID uint, id uint, input ItemCore) error
	Delete(UserID uint, id uint) error
}
