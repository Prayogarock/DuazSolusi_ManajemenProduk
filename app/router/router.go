package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"duaz/app/middlewares"
	_userData "duaz/features/users/data"
	_userHandler "duaz/features/users/handler"
	_userService "duaz/features/users/service"

	_itemData "duaz/features/items/data"
	_itemHandler "duaz/features/items/handler"
	_itemService "duaz/features/items/service"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	ItemData := _itemData.New(db)
	ItemService := _itemService.New(ItemData)
	ItemHandlerAPI := _itemHandler.New(ItemService)

	//users
	c.POST("/users", UserHandlerAPI.Add)
	c.POST("/login", UserHandlerAPI.Login)
	c.GET("/users", UserHandlerAPI.GetUserData, middlewares.JWTMiddleware())
	c.PUT("/users", UserHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	//items
	c.POST("/items", ItemHandlerAPI.CreateItem, middlewares.JWTMiddleware())
	c.GET("/items", ItemHandlerAPI.GetAll)
	c.PUT("/items/:item_id", ItemHandlerAPI.UpdateItemByID, middlewares.JWTMiddleware())
	c.DELETE("/items/:item_id", ItemHandlerAPI.DeleteItem, middlewares.JWTMiddleware())
}
