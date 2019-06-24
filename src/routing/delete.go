package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func deleteBoard(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for deleting a board.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "."})
	}

	err = database.Remove(id)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for getting a board.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "."})
	}
	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Board deleted successfully."})
}
