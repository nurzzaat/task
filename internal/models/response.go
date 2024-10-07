package models

type SuccessResponse struct {
	Result interface{} `json:"result"`
}

type SuccessResponsePagination struct {
	Result interface{} `json:"result"`
	Count  int         `json:"count"`
}

type ErrorResponse struct {
	Result ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Properties struct {
	Group string
	Song  string
	Lyric string
	Link  string
	From  string
	To    string
	Page  int
	Size  int
}
