package observation

import (
	"context"
	"sync"
)

var (
	_ Manager = &channelManager{}
	_ Sink    = &channelSink{}
)

type observable struct {
	sinkName      string
	schema, value interface{}
}

type channelManager struct {
	observables chan observable

	observers     []Observer
	observersLock sync.RWMutex
}

type channelSink struct {
	name        string
	observables chan observable
}

func newChannelSink(name string, observables chan observable) Sink {
	created := &channelSink{
		name:        name,
		observables: observables,
	}
	return created
}

func NewChannelManager(backlog uint) Manager {
	created := &channelManager{
		observables: make(chan observable, backlog),
		observers:   make([]Observer, 0),
	}
	return created
}

func (s *channelSink) Submit(schema interface{}, value interface{}) {
	evt := observable{
		sinkName: s.name,
		schema:   schema,
		value:    value,
	}
	s.observables <- evt
}

func (m *channelManager) Sink(name string) Sink {
	return newChannelSink(name, m.observables)
}

func (m *channelManager) Register(o Observer) {
	m.observersLock.Lock()
	defer m.observersLock.Unlock()
	m.observers = append(m.observers, o)
}

func (m *channelManager) Run(ctx context.Context) error {
	continueRunning := true
	for continueRunning {
		select {
		case o := <-m.observables:
			for _, observer := range m.observers {
				observer.Submit(o.schema, o.value)
			}
		case <-ctx.Done():
			continueRunning = false
			break
		}
	}
	return nil
}
