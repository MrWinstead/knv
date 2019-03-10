package observation

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChannelManager(t *testing.T) {
	mgr := NewChannelManager(uint(rand.Int31n(100)))
	assert.NotNil(t, mgr)
}
