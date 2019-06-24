package database

import (
	"time"
)

type Board struct {
	TrackID          uint64    `db:"track_id" json:"track_id"`
	ID               uint64    `db:"id" json:"id"`
	DateCreated      time.Time `db:"date_created" json:"date_created,omitempty"`
	DateModified     time.Time `db:"date_modified" json:"date_modified,omitempty"`
	SHA3             [512]byte `db:"sha3" json:"sha3"`
	Jacket           []byte    `db:"jacket" json:"jacket"`
	Charters         string    `db:"charters" json:"charters"`
	DifficultyName   uint64    `db:"difficulty_name" json:"difficulty_name"`
	DifficultyRating uint8     `db:"difficulty_rating" json:"difficulty_rating"`
}
