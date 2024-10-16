package categories

import wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"

type repository struct {
	rdbms wsqlx.Rdbms
}

func New(rdbms wsqlx.Rdbms) *repository {
	return &repository{rdbms: rdbms}
}
