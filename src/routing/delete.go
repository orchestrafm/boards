package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func deleteBoard(c echo.Context) error {
	tid, err := strconv.ParseUint(c.Param("tid"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for deleting a board.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "."})
	}
	bid, err := strconv.ParseUint(c.Param("bid"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for deleting a board.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "."})
	}

	err = database.Remove(bid, tid)
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