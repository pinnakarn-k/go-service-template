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
