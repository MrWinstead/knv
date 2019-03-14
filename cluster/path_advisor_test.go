package cluster

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestNewEtcdPathAdvisor(t *testing.T) {
	randomPathPrefix := gofakeit.Color()
	created := NewEtcdPathAdvisor(randomPathPrefix)
	assert.NotNil(t, created)
}
