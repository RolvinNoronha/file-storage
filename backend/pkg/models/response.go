package models

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"Message"`
	Data    any    `json:"data"`
	Errors  any    `json:"errors"`
}
