package routers

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func deleteBoard(c echo.Context) error {
	// auth check
	admin := HasRole(c, "manage-boards")
	auth := HasRole(c, "create-board")
	if auth != true {
		logger.Info().
			Msg("user intent to create a delete a board, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Insufficient Permissions."})
	}

	// board search
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

	// remove board
	if !admin {
		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Not an Administrator."})
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
