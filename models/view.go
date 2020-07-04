package models

type Response struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}
