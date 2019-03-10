package observation

import "sync"

var (
	_ Observer = &RecordingObserver{}
)

type RecordingObserver struct {
	events     []RecordedEvent
	eventsLock sync.RWMutex
}

type RecordedEvent struct {
	Schema interface{}
	Value  interface{}
}

func NewRecordingObserver() *RecordingObserver {
	created := &RecordingObserver{
		events: make([]RecordedEvent, 0),
	}

	return created
}

func (ro *RecordingObserver) Submit(schema interface{}, value interface{}) {
	ro.eventsLock.Lock()
	defer ro.eventsLock.Unlock()
	ro.events = append(ro.events, RecordedEvent{schema, value})
}

func (ro *RecordingObserver) Events() []RecordedEvent {
	ro.eventsLock.RLock()
	defer ro.eventsLock.RUnlock()

	ret := make([]RecordedEvent, len(ro.events))
	copy(ret, ro.events)

	return ret
}
