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
	if nil != self.events {
		length := len(self.events[event])
		for i := 0; i < length; i++ {
			err := self.events[event][i]()
			if nil != err {
				return err
			}
		}
	}

	return nil
}
