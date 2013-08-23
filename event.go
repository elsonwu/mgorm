package mgorm

import (
	"strings"
)

type EventFn func() error

type IEvent interface {
	On(event string, fn EventFn)
	Emit(event string) error
}

type Event struct {
	events map[string][]EventFn
}

func (self *Event) On(event string, fn EventFn) {
	if nil == self.events {
		self.events = make(map[string][]EventFn)
	}

	event = strings.ToLower(event)
	self.events[event] = append(self.events[event], fn)
}

func (self *Event) Emit(event string) error {
	event = strings.ToLower(event)
	if nil == self.events {
		return nil
	}

	events, ok := self.events[event]
	if !ok {
		return nil
	}

	for _, fn := range events {
		err := fn()
		if nil != err {
			return err
		}
	}

	return nil
}
