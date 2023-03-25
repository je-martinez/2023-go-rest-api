package types

type QueryOptions struct {
	Pagination PaginationOptions
}

type PaginationOptions struct {
	PageNumber int
	PageSize   int
	Skip       int
	Take       int
}
