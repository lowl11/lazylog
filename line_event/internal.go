package line_event

func (event *Event) listen() {
	for logMessageFunc := range event.logChannel {
		logMessageFunc()
	}
}

func (event *Event) listenCustom() {
	for logMessageFunc := range event.customLogChannel {
		logMessageFunc()
	}
}
