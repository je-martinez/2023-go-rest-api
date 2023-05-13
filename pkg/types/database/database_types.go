package database_types

type Pagination struct {
	Pagination PaginationOptions
}

type PaginationOptions struct {
	PageNumber int
	PageSize   int
	Skip       int
	Take       int
}

type QueryOptions struct {
	Pagination *Pagination
	Query      interface{}
	Args       any
	Preloads   []string
}
