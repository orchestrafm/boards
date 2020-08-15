package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/orchestrafm/boards/src/database"
	"github.com/orchestrafm/boards/src/objstore"
	"github.com/spidernest-go/logger"
	"github.com/spidernest-go/mux"
)

func createBoard(c echo.Context) error {
	// Check Authorization
	if authorized := HasRole(c, "create-board"); authorized != true {
		logger.Info().
			Msg("user intent to create a new board, but was unauthorized.")

		return c.JSON(http.StatusUnauthorized, &struct {
			Message string
		}{
			Message: ErrPermissions.Error()})
	}

	// Data Binding
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

	// Jacket Contents
	f, err := decodeJacket(b.Jacket)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Upload failed at jacket decoding step.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music board data was invalid or malformed."})
	}
	defer f.Close()
	defer os.Remove(f.Name())

	ff, err := os.Open(f.Name()) //HACK: Go is fucking stupid and won't let me reuse
	// the file pointer for what reason. Introducing yet
	// another vector that could fail for any reason.
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Opening image on local disk failed.")

		return c.JSON(http.StatusInternalServerError, &struct {
			Message string
		}{
			Message: "Image on disk was unreadable or unopenable."})
	}
	defer ff.Close()

	// Push Database Entry
	b.Jacket = "" // clear this because a URL will fill it's spot
	err = b.New()
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music board data did not get submitted to the database."})
	}

	// Upload to Object Storage
	fname := fmt.Sprint(b.ID) + ".webp"
	url, err := objstore.Upload(ff, "/Images/Effective/Pre-Season/"+fname, "public-read", true)
	if err != nil {
		defer database.Remove(b.ID, b.TrackID)
		logger.Error().
			Err(err).
			Msg("Object Store rejected putting the object.")

		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "File could not be commited to disk."})
	}
	b.Jacket = url

	// Update Database Entry
	err = b.Update()
	if err != nil {
		defer database.Remove(b.ID, b.TrackID)
		return c.JSON(http.StatusNotAcceptable, &struct {
			Message string
		}{
			Message: "Music board data did not get submitted to the database."})
	}

	return c.JSON(http.StatusCreated, &struct {
		Message string
	}{
		Message: "OK."})
}
