package database

import (
	"database/sql"
	"strconv"

	"github.com/spidernest-go/logger"
)

func SelectID(id uint64) (*Board, error) {
	boards := db.Collection("boards")
	rs := boards.Find(id)
	b := *new(Board)
	err := rs.One(&b)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return &b, nil
	}
}

func SelectTrackID(id uint64) ([]*Board, error) {
	var bs []*Board
	boards := db.Collection("boards")
	rs := boards.Find().Where("track_id = ", id)

	err := rs.All(&bs)
	if err != nil && err != sql.ErrNoRows {
		logger.Error().
			Err(err).
			Msg("Bad parameters or database error.")
	}

	if err == sql.ErrNoRows {
		return nil, err
	} else {
		return bs, nil
	}
}
