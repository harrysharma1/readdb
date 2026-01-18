package handlers

import (
	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func GetBooksHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(200, "")
	}
}
