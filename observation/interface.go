package observation

import "context"

type Manager interface {
	Sink(name string) Sink
	Register(o Observer)
	Run(ctx context.Context) error
}

type Observer interface {
	Submit(schema interface{}, value interface{})
}

type Sink interface {
	Submit(schema interface{}, value interface{})
}
