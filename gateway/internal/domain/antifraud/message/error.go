package message

import "errors"

var (
	ErrBadRequest    = errors.New("error bad request")
	ErrInternalError = errors.New("error internal")

	ErrFormingResponse = errors.New("error forming response")

	ErrFetching = errors.New("error fetching")
)
