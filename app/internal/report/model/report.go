package model

type Report struct {
	ID         int64  `db:"id"`
	Definition string `db:"definition"`
}
