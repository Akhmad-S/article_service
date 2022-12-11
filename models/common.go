package models

type JSONResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONError struct {
	Error string `json:"error"`
}