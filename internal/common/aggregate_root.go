package common

type EventName string
type DomainEvent interface {
	GetEventName() EventName
}

type AggregateRoot struct {
	domainEvents []DomainEvent
	id           UUID
}

func NewAggregateRoot(id UUID) AggregateRoot {
	return AggregateRoot{[]DomainEvent{}, id}
}

func (ar AggregateRoot) GetID() UUID {
	return ar.id
}

func (ar *AggregateRoot) AddDomainEvent(event DomainEvent) {
	ar.domainEvents = append(ar.domainEvents, event)
}

func (ar *AggregateRoot) GetDomainEvents() []DomainEvent {
	return ar.domainEvents
}

func (ar *AggregateRoot) ClearDomainEvents() {
	ar.domainEvents = nil
}
