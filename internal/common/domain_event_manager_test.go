package common_test

import (
	"ingressos/internal/common"
	"testing"
)

type MockedDomainEvent struct {
	SomeData string
}

const MockedDomainEventName = common.EventName("MockedDomainGetEventName")

func (md MockedDomainEvent) GetEventName() common.EventName {
	return MockedDomainEventName
}

type MockedDomainEventHandler struct {
	handle func(event common.DomainEvent) bool
}

func (d *MockedDomainEventHandler) Handle(event common.DomainEvent) bool {
	return d.handle(event)
}

func TestDomainEventManager(t *testing.T) {
	t.Run("Publish()", func(t *testing.T) {
		t.Run("shoud publish correclty to all subscribers", func(t *testing.T) {
			domainEventManager := common.NewDomainEventManager()
			theEvent := MockedDomainEvent{SomeData: "teste"}
			var publishedEvent MockedDomainEvent
			domainEventManager.Register(MockedDomainEventName, &MockedDomainEventHandler{
				handle: func(e common.DomainEvent) bool {
					event, ok := e.(MockedDomainEvent)
					if ok {
						publishedEvent = event
					}
					return true
				},
			})
			anAggregate := common.AggregateRoot{}
			anAggregate.AddDomainEvent(theEvent)
			domainEventManager.Publish(anAggregate)

			if publishedEvent.GetEventName() != MockedDomainEventName {
				t.Fatalf("should have published the event %q but receibed %s", MockedDomainEventName, publishedEvent.GetEventName())
			}
			if publishedEvent.SomeData != "teste" {
				t.Fatal("should have published the event correctly")
			}
		})
	})
}
