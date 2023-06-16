package line_event

type InfoFunc func(args ...any)
type ErrorFunc func(err error, args ...any)

func (event *Event) AddInfo(logMessageFunc InfoFunc, args ...any) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	event.logChannel <- func() {
		logMessageFunc(args...)
	}
}

func (event *Event) AddInfoCustom(logMessageFunc InfoFunc, args ...any) {
	event.customMutex.Lock()
	defer event.customMutex.Unlock()

	event.customLogChannel <- func() {
		logMessageFunc(args...)
	}
}

func (event *Event) AddError(logMessageFunc ErrorFunc, err error, args ...any) {
	event.mutex.Lock()
	defer event.mutex.Unlock()

	event.logChannel <- func() {
		logMessageFunc(err, args...)
	}
}

func (event *Event) AddErrorCustom(logMessageFunc ErrorFunc, err error, args ...any) {
	event.customMutex.Lock()
	defer event.customMutex.Unlock()

	event.customLogChannel <- func() {
		logMessageFunc(err, args...)
	}
}
