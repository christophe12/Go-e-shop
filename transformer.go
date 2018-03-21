package main

//ResponseBuilder struct
type ResponseBuilder struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
	Meta       interface{} `json:"meta"`
}
