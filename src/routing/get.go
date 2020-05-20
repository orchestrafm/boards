package routing

import (
	"net/http"
	"strconv"

	"github.com/orchestrafm/boards/src/database"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func getBoard(c echo.Context) error {
	i, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed id parameter (%s) was not a valid number", c.Param("id"))

		return c.JSON(http.StatusBadRequest, nil)
	}

	b, err := database.SelectID(i)
	if err != nil {
		logger.Error().
			Err(err).
			Msg(".")

		c.JSON(http.StatusNotFound, ErrGeneric)
	}

	return c.JSON(http.StatusOK, b)
}

func getBoardsFromTrack(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Passed id parameter (%s) was not a valid number", c.Param("id"))

		return c.JSON(http.StatusBadRequest, nil)
	}

	bs, err := database.SelectTrackID(id)
	if err != nil {
		logger.Error().
			Err(err).
			Msg(".")

		c.JSON(http.StatusNotFound, bs)
	}

	return c.JSON(http.StatusOK, bs)
}

func getBoardFromHash(c echo.Context) error {
	hash := c.Param("hash")

	b, err := database.SelectByHash(hash)
	if err != nil {
		logger.Error().
			Err(err).
			Msg(".")

		c.JSON(http.StatusNotFound, ErrGeneric)
	}

	return c.JSON(http.StatusOK, b)
}
