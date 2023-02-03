package wrongs

// StatusUnprocessableEntity will return an error when the server cannot process the contained instructions
type StatusUnprocessableEntity string

// StatusBadRequest implements the Error interface
func (e StatusUnprocessableEntity) Error() string {
	return string(e)
}

// StatusBadRequest will return an error when the client makes a mistakes
type StatusBadRequest string

// StatusBadRequest implements the Error interface
func (e StatusBadRequest) Error() string {
	return string(e)
}

// StatusInternalServerError will return an error when the server encounters an error
type StatusInternalServerError string

// StatusInternalServerError implements the Error interface
func (e StatusInternalServerError) Error() string {
	return string(e)
}
