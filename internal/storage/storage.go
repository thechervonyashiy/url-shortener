package storage

import (
	"errors"
)

var (
	ErrURLNotFount   = errors.New("url not found")
	ErrURLExists     = errors.New("url exists")
	ErrURLNotDeleted = errors.New("url not deleted")
)
