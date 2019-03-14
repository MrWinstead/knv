package cmd

import (
	"io"
	"math/rand"

	"github.com/oklog/ulid"
)

func newULIDMonotonicSource() io.Reader {
	return ulid.Monotonic(rand.New(rand.NewSource(rand.Int63())), 0)
}
