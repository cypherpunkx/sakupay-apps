package exception

import "errors"

const (
	StatusInternalServer = "Internal Server Error"
	StatusBadRequest     = "Bad Request"
	StatusSuccess        = "Success"
)

var (
	ErrNotFound          = errors.New("data not found")
	ErrCodeAlreadyExist  = errors.New("code already exist")
	ErrFailedCreate      = errors.New("failed to create data")
	ErrFailedUpdate      = errors.New("failed to update data")
	ErrFailedDelete      = errors.New("failed to delete data")
	ErrTitleAlreadyExist = errors.New("title already exist")
	ErrInvalidPage       = errors.New("invalid page")
	ErrInvalidPerPage    = errors.New("invalid per page")
)
