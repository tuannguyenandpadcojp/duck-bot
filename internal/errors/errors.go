package errors

import "errors"

var (
	ErrCommandAlreadyExists  = errors.New("command already exists")
	ErrCommandNotFound       = errors.New("command not found")
	ErrCommandNotImplemented = errors.New("command not implemented")
)
