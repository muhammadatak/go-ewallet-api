package dto

type Response[T any] struct {
	Data T     `json:"data"`
	Meta *Meta `json:"meta,omitempty"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorResponse struct {
	Code    string            `json:"code,omitempty"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func CreateResponseSuccess[T any](data T) Response[T] {
	return Response[T]{Data: data}
}

func CreateResponsePaginated[T any](data T, meta Meta) Response[T] {
	return Response[T]{Data: data, Meta: &meta}
}

func CreateResponseError(message string) ErrorResponse {
	return ErrorResponse{Message: message}
}

func CreateResponseErrorData(message string, errors map[string]string) ErrorResponse {
	return ErrorResponse{Message: message, Errors: errors}
}
