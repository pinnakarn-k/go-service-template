package apperror

type Response struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []Detail `json:"details,omitempty"`
}

type Detail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
