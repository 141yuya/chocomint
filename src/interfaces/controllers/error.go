package controllers

type Error struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewError(message string, data interface{}) *Error {
	H := new(Error)
	H.Message = message
	H.Data = data
	return H
}
