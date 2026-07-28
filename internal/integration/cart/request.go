package cart

type GetCartsRequest struct {
	Limit  int `query:"limit" validate:"required,min=1,max=100"`
	Skip   int `query:"skip" validate:"min=0"`
	UserID int `query:"userId" validate:"required,gt=0"`
}
