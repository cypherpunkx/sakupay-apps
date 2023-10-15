package exception

import "errors"

const (
	StatusInternalServer = "Internal Server Error"
	StatusBadRequest     = "Bad Request"
	StatusSuccess        = "Success"
	StatusUnauthorized   = "Unauthorized"
)

var (
	ErrNotFound                 = errors.New("data not found")
	ErrCodeAlreadyExist         = errors.New("code already exist")
	ErrFailedCreate             = errors.New("failed to create data")
	ErrFailedCreateToken        = errors.New("failed to create token")
	ErrInvalidParseToken        = errors.New("invalid to parse token")
	ErrInvalidTokenMapclaims    = errors.New("invalid token mapclaims ")
	ErrInvalidTokenStringMethod = errors.New("invalid token string method")
	ErrInvalidUsernamePassword  = errors.New("invalid username password")
	ErrTokenNotProvided         = errors.New("token not provided")
	ErrFailedUpdate             = errors.New("failed to update data")
	ErrFailedDelete             = errors.New("failed to delete data")
	ErrTitleAlreadyExist        = errors.New("title already exist")
	ErrInvalidPage              = errors.New("invalid page")
	ErrInvalidPerPage           = errors.New("invalid per page")
)
