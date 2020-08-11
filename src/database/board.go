package database

import (
	"time"
)

type Board struct {
	TrackID          uint64    `db:"track_id" json:"track_id,omitempty"`
	ID               uint64    `db:"id" json:"id,omitempty"`
	DateCreated      time.Time `db:"date_created" json:"date_created,omitempty"`
	DateModified     time.Time `db:"date_modified" json:"date_modified,omitempty"`
	SHA3             string    `db:"sha3" json:"sha3,omitempty"`
	Jacket           string    `db:"jacket" json:"jacket,omitempty"`
	Illustrators     string    `db:"illustrators" json:"illustrators,omitempty"`
	Charters         string    `db:"charters" json:"charters,omitempty"`
	DifficultyName   uint64    `db:"difficulty_name" json:"difficulty_name"`
	DifficultyRating uint8     `db:"difficulty_rating" json:"difficulty_rating"`
}
