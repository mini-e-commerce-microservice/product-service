package models

type Outlet struct {
	ID          int64  `db:"id"`
	UserID      int64  `db:"user_id"`
	Logo        string `db:"logo"`
	Name        string `db:"name"`
	Slogan      string `db:"slogan"`
	Description string `db:"description"`
}
