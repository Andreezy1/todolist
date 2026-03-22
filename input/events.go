package input

import (
	"time"

	"github.com/k0kubun/pp"
)

type Events struct {
	event         string
	errorOpisanie string
	Eventstime    time.Time
}

var arraiEvents = []Events{}

func newEvents(text string, error string) {
	arraiEvents = append(arraiEvents, Events{event: text,
		errorOpisanie: error,
		Eventstime:    time.Now(),
	})
}

func printEvents() {
	pp.Println(arraiEvents)
}
