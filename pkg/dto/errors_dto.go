package dto

type ApiError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ResponseError struct {
	ApiError ApiError `json:"apiError"`
}
