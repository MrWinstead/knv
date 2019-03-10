package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsErrNotFound_True(t *testing.T) {
	var e error = &ErrNotFound{}
	assert.True(t, IsErrNotFound(e))
}

func TestIsErrNotFound_False(t *testing.T) {
	e := fmt.Errorf("")
	assert.False(t, IsErrNotFound(e))
}
