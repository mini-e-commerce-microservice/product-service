package primitive

type PaginationInput struct {
	Page     int64
	PageSize int64
}

type PaginationOutput struct {
	Page      int64
	PageSize  int64
	PageCount int64
	TotalData int64
}
