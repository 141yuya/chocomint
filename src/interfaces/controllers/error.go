package controllers

type Error struct {
	Data interface{} `json:"data"`
}

func NewError(data interface{}) *Error {
	error := new(Error)
	error.Data = data
	return error
}
