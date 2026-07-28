package pagination

type Meta struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

type Response[T any] struct {
	Data       []T  `json:"data"`
	Pagination Meta `json:"pagination"`
}

func Offset(page, pageSize int) int {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 1
	}

	return (page - 1) * pageSize
}

func NewMeta(page, pageSize, totalItems int) Meta {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 1
	}

	if totalItems < 0 {
		totalItems = 0
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	return Meta{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}
