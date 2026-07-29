package transaction

type SearchTransactionsRequest struct {
	AccountNo string `json:"accountNo" validate:"omitempty,max=20"`
	FromDate  string `json:"fromDate" validate:"omitempty,datetime=2006-01-02"`
	ToDate    string `json:"toDate" validate:"omitempty,datetime=2006-01-02"`
	Page      int    `json:"page" validate:"required,min=1"`
	PageSize  int    `json:"pageSize" validate:"required,min=1,max=100"`
}
