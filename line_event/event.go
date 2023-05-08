package line_event

import (
	"sync"
)

type Event struct {
	logChannel       chan func()
	customLogChannel chan func()
	mutex            sync.Mutex
	customMutex      sync.Mutex
}

func New() *Event {
	line := &Event{
		logChannel:       make(chan func()),
		customLogChannel: make(chan func()),
	}

	go line.listen()
	go line.listenCustom()

	return line
}
