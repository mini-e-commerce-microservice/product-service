package model

type ProductMedia struct {
	ID             int64  `db:"id"`
	ProductID      int64  `db:"product_id"`
	Media          string `db:"media"`
	MediaType      string `db:"media_type"`
	IsPrimaryMedia bool   `db:"is_primary_media"`
}
