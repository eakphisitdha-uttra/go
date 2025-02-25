package helpers

import "math"

type PaginationStruct struct {
	CurrentPage  int64       `json:"current_page" extensions:"x-order=1"`
	NextPage     int64       `json:"next_page" extensions:"x-order=2"`
	PreviousPage int64       `json:"previous_page" extensions:"x-order=3"`
	SizePerPage  int64       `json:"size_per_page" extensions:"x-order=4"`
	TotalPages   int64       `json:"total_pages" extensions:"x-order=5"`
	TotalItems   int64       `json:"total_items" extensions:"x-order=6"`
	Items        interface{} `json:"items" extensions:"x-order=7"`
}

// รับค่า parameter เพื่อ return format paginate
func Pagination(data interface{}, limit, page, total int64) interface{} {
	totalPages := int64(math.Ceil(float64(total) / float64(limit)))
	var pagination PaginationStruct
	pagination.Items = data
	pagination.CurrentPage = int64(page)
	pagination.PreviousPage = getPreviousPage(page)
	pagination.NextPage = getNextPage(page, totalPages)
	pagination.SizePerPage = limit
	pagination.TotalItems = total
	pagination.TotalPages = totalPages
	return pagination
}

func getPreviousPage(page int64) int64 {
	if page == 0 {
		return 1
	}
	if page-1 == 0 {
		return 1
	}
	return page - 1
}

func getNextPage(page, totalPages int64) int64 {
	var nextPage int64
	if page >= totalPages {
		nextPage = totalPages
	} else {
		nextPage = int64(page) + 1
	}
	return nextPage
}
