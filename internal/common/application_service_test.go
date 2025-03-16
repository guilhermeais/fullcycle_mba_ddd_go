package common_test

import (
	"ingressos/internal/common"
	"testing"
)

func TestRun(t *testing.T) {
	uow := &common.UnitOfWork{}
	domainEventMagger := common.NewDomainEventManager()
	mockedAggregate := common.AggregateRoot{}
	mockedDomainEvent := MockedDomainEvent{
		SomeData: "some_data",
	}
	var emittedEvent common.DomainEvent
	domainEventMagger.Register(MockedDomainEventName, &MockedDomainEventHandler{
		handle: func(event common.DomainEvent) bool {
			emittedEvent = event
			return true
		},
	})
	res := common.Run(func() common.AggregateRoot {
		mockedAggregate.AddDomainEvent(mockedDomainEvent)
		uow.RegisterAggregate(mockedAggregate)
		return mockedAggregate
	}, uow, domainEventMagger)

	if !res.GetID().IsEqual(mockedAggregate.GetID()) {
		t.Fatalf("invalid return, expected %q but receive %q", mockedAggregate.GetID(), res.GetID())
	}

	emittedEventValidated, ok := emittedEvent.(MockedDomainEvent)

	if !ok {
		t.Fatal("invalid emitted event")
	}

	if emittedEventValidated.GetEventName() != MockedDomainEventName {
		t.Fatalf("should have emitted %q but emit %q", MockedDomainEventName, emittedEventValidated.GetEventName())
	}

	if emittedEventValidated.SomeData != "some_data" {
		t.Fatal("invalid emitted data")
	}
}
