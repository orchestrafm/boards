package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func editBoard(c echo.Context) error {
	b := new(database.Board)
	if err := c.Bind(b); err != nil {
		logger.Error().
			Err(err).
			Msg("Request Data could not be binded to datastructure.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for editing a track.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}
	err = b.Edit(id)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed music track data."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Track edited successfully."})
}
