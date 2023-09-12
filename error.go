package panik

import (
	"errors"
)

var (
	ErrInvalidNIK    = errors.New("NIK is invalid")
	ErrNIKMustNumber = errors.New("NIK must be number")
	ErrNIKLength     = errors.New("NIK length must be 16")
)
