package ecs

import (
	"fmt"
)

type ecsError struct {
	reason *string
}

func NewError(args ...interface{}) error {
	value := fmt.Sprint("", args)
	return ecsError{&value}
}

func (e ecsError) Error() string {
	return *e.reason
}
