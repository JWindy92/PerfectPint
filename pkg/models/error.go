package models

//TODO: should this be in a different package?
type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

//TODO: make this a method on the struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
