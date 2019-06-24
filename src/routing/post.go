package routing

import (
	"net/http"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createBoard(c echo.Context) error {
	b := new(database.Board)
	if err := c.Bind(b); err != nil {
		logger.Error().
			Err(err).
			Msg("Invalid or malformed music track data.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music track data was invalid or malformed."})
	}

	err := b.New()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music track data did not get submitted to the database."})
	}

	return c.JSON(http.StatusOK, &struct {
		Message string
	}{
		Message: "OK."})
}
