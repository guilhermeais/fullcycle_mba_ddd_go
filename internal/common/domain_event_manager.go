package common

type Handler interface {
	Handle(event DomainEvent) bool
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
		handlers := d.handlers[event.GetEventName()]
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	aggregateRoot.ClearDomainEvents()
}
func (d *DomainEventManager) Register(name EventName, subscriber Handler) {
	d.handlers[name] = append(d.handlers[name], subscriber)
}
