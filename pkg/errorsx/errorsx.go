package errorsx

import "errors"

var (
	ErrRoleNotFound = errors.New("role not found")
	ErrUserNotFound = errors.New("user not found")
	ErrInternal     = errors.New("internal error")
	ErrBadRequest   = errors.New("bad request")
	ErrUserExists   = errors.New("user already exists")
	ErrPasswordHash = errors.New("failed to hash password")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidTimeFormat = errors.New("invalid time format")
	ErrCityNotFound = errors.New("cannot register in this city")
	ErrPvzExists = errors.New("pvz already exists")
	ErrInvalidUUID = errors.New("invalid uuid")
	ErrCityIsNotExists = errors.New("city is not exists")
	ErrAuthHeaderIsEmpty = errors.New("authorization header is empty")
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
	ErrUnexpSignedMetod = errors.New("unsupported signed method")
	ErrInvUserUUIDInToken = errors.New("invalid user uuid in token")
	ErrInvUserRoleInToken = errors.New("invalid user role in token")
	ErrNoPermission = errors.New("no permission")
	ErrPVZNotFound = errors.New("pvz not found")
	ErrReceptionNotFound = errors.New("reception not found")
	ErrReceptionIsNotClosed = errors.New("last reception is not closed")
	ErrReceptionIsClosed = errors.New("last reception is closed")
	ErrProductNotFound = errors.New("product not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrStartDateAfterEndDate = errors.New("start date must be before end date")
	ErrNotFound = errors.New("not found")
	ErrInvalidEmail = errors.New("invalid email forma")
	ErrEmptyField = errors.New("one or several fields are empty")
	ErrEmptyEmail     = errors.New("email cannot be empty")
    ErrEmailTooLong   = errors.New("email is too long")
    ErrInvalidDomain  = errors.New("invalid email domain")
	ErrInvalidLimit = errors.New("invalid limit")
)
