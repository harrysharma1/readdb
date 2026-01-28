package handlers

import (
	"readdb/database"
	"strconv"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func GetSeriesByIdHandler(db *bolt.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return ctx.String(400, "invalid series id")
		}
		series, err := database.GetSeriesById(db, uint(idInt))
		if err != nil {
			return ctx.String(404, "series not found")
		}
		return ctx.Render(200, "series", map[string]interface{}{
			"Series": series,
		})
	}
}
