package routers

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func editBoard(c echo.Context) error {
	// auth check
	admin := HasRole(c, "manage-boards")
	auth := HasRole(c, "create-board")
	if auth != true {
		logger.Info().
			Msg("user intent to create a update a board, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Insufficient Permissions."})
	}

	// data binding
	b := new(database.Board)
	if err := c.Bind(b); err != nil {
		logger.Error().
			Err(err).
			Msg("Request Data could not be binded to datastructure.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed board data."})
	}

	// validation
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid Parameters for editing a board.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed board data."})
	}

	// update resource
	if !admin {
		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: "Not an Administrator."})
	}
	err = b.Edit(id)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Board could not be updated.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Invalid or malformed board data."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "Board edited successfully."})
}
