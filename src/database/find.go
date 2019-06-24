package database

import (
	"database/sql"
	"strconv"

	"github.com/spidernest-go/logger"
)

func SelectID(id uint64) (*Board, error) {
	boards := db.Collection("boards")
	rs := boards.Find(id)
	t := *new(Board)
	err := rs.One(&t)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return &t, nil
	}
}
