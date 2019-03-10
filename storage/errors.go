package storage

import "fmt"

var _ error = &ErrNotFound{}

type ErrNotFound struct {
	Table, Index, Key string
}

func (e *ErrNotFound) Error() string {
	msg := fmt.Sprintf("Key %v not found on Table %v under Index %v",
		e.Key, e.Table, e.Index)
	return msg
}

func IsErrNotFound(e error) bool {
	_, isType := e.(*ErrNotFound)
	return isType
}
