package database

import (
	"time"

	"github.com/spidernest-go/logger"
)

func (b *Board) Update() error {
	// TODO: Check if all struct fields are present
	// TODO: Abort if the Board ID isn't already present in the database
	tracks := db.Collection("boards")
	rs := tracks.Find(b.ID)
	// TODO: Check if DateCreated is not equal
	b.DateModified = time.Now()
	err := rs.Update(b)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated from the table.")
	}
	return err
}

func (b *Board) Edit(id uint64) error {
	tracks := db.Collection("boards")
	rs := tracks.Find(id)
	// TODO: Check if DateCreated is not equal
	b.DateModified = time.Now()
	b.ID = id
	err := rs.Update(b)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("Music Track could not be updated from the table.")
	}
	return err
}
