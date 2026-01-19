package handlers

import (
	"readdb/database"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func GetListsHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		lists, err := database.GetAllLists(db)
		if err != nil {
			return ctx.JSON(500, map[string]string{
				"errors": err.Error(),
			})
		}
		if len(lists) > 0 {
			return ctx.JSON(200, lists)
		}
		return ctx.JSON(200, "")
	}
}
