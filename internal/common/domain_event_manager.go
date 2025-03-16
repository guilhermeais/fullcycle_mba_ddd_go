package common

type DomainEventHandler interface {
	Handle(event DomainEvent) bool
}
type handlers map[EventName][]DomainEventHandler
type DomainEventManager struct {
	handlers handlers
}

func NewDomainEventManager() *DomainEventManager {
	return &DomainEventManager{handlers: handlers{}}
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
func (d *DomainEventManager) Register(name EventName, handler DomainEventHandler) {
	d.handlers[name] = append(d.handlers[name], handler)
}
