package common

type AggregateRoot struct {
	domainEvents []any
}

func (ar *AggregateRoot) AddDomainEvent(event any) {
	ar.domainEvents = append(ar.domainEvents, event)
}

func (ar *AggregateRoot) GetDomainEvents() []any {
	return ar.domainEvents
}

func (ar *AggregateRoot) ClearDomainEvents() {
	ar.domainEvents = nil
}
