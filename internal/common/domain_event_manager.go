package common

import "fmt"

type EventName string

type Handler interface {
	Handle(event any) bool
}
type Handlers map[EventName][]Handler
type DomainEventManager struct {
	handlers Handlers
}

func NewDomainEventManager() *DomainEventManager {
	return &DomainEventManager{handlers: Handlers{}}
}

func (d *DomainEventManager) Publish(aggregateRoot AggregateRoot) {
	for _, event := range aggregateRoot.GetDomainEvents() {
		eventName := EventNameFromClass(event)
		handlers := d.handlers[eventName]
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	aggregateRoot.ClearDomainEvents()
}

func EventNameFromClass(class any) EventName {
	return EventName(fmt.Sprintf("%T", class))
}

func (d *DomainEventManager) Register(name EventName, subscriber Handler) {
	d.handlers[name] = append(d.handlers[name], subscriber)
}
