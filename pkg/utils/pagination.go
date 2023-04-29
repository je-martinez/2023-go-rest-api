package utils

import "github.com/je-martinez/2023-go-rest-api/pkg/types"

const pageSize = 100

func GetPaginationParams(pageNumber int) (params *types.PaginationOptions) {

	var page_number int

	if pageNumber < 0 {
		page_number = 1
	}

	return &types.PaginationOptions{
		PageNumber: page_number,
		PageSize:   pageSize,
		Skip:       (page_number - 1*pageSize),
		Take:       pageSize,
	}
}
