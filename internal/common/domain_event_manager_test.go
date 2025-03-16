package common_test

import (
	"ingressos/internal/common"
	"testing"
)

type MockedDomainEvent struct {
	SomeData string
}
type MockedDomainEventHandler struct {
	handle func(event any) bool
}

func (d *MockedDomainEventHandler) Handle(item any) bool {
	return d.handle(item)
}

func TestDomainEventManager(t *testing.T) {
	t.Run("Publish()", func(t *testing.T) {
		t.Run("shoud publish correclty to all subscribers", func(t *testing.T) {
			domainEventManager := common.NewDomainEventManager()
			theEvent := MockedDomainEvent{SomeData: "teste"}
			var publishedEvent MockedDomainEvent
			domainEventManager.Register(common.EventNameFromClass(MockedDomainEvent{}), &MockedDomainEventHandler{
				handle: func(e any) bool {
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

			if publishedEvent.SomeData != "teste" {
				t.Fatal("should have published the event correctly")
			}
		})
	})
}
