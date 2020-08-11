package routing

import (
	"net/http"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createBoard(c echo.Context) error {
	if authorized := HasRole(c, "create-board"); authorized != true {
		logger.Info().
			Msg("user intent to create a new board, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: ErrPermissions.Error()})
	}

	b := new(database.Board)
	if err := c.Bind(b); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed music board data.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music board data was invalid or malformed."})
	}

	err := b.New()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music board data did not get submitted to the database."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "OK."})
}
