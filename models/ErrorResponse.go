package models

type ErrorResponse struct {
	Message string `json:"message"`
}

type DatabaseError struct {
	Type string
	Code int
}

type DatabaseErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}
