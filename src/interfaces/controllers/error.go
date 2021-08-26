package controllers

type Error struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewError(message string, data interface{}) *Error {
	error := new(Error)
	error.Message = message
	error.Data = data
	return error
}
