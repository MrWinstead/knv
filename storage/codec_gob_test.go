package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGobCodec(t *testing.T) {
	f := func() interface{} { return nil }
	created := NewGobCodec(f)
	assert.NotNil(t, created)
}
