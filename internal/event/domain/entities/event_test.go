package entities_test

import (
	"ingressos/internal/common"
	"ingressos/internal/event/domain/entities"
	"ingressos/internal/event/domain/events"
	"reflect"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	t.Run("should create an event", func(t *testing.T) {
		partnerId, _ := common.CreateUUID()
		event, err := entities.CreateEvent(entities.CreateEventCommand{
			Name:        "Testing Event",
			Description: "A testing event",
			Date:        time.Date(2025, 3, 24, 13, 30, 0, 0, time.UTC),
			IsPublished: false,
			TotalSpots:  20,
			PartnerId:   string(partnerId),
		})

		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}

		expectedDomainEvent := events.EventCreatedEvent{
			Id:                 event.GetID(),
			Name:               event.GetName(),
			Description:        event.GetDescription(),
			Date:               event.GetDate(),
			IsPublished:        event.IsPublished(),
			TotalSpots:         event.GetTotalSpots(),
			TotalSpotsReserved: event.GetTotalSpotsReserved(),
			PartnerId:          event.GetPartnerId(),
		}

		if len(event.GetDomainEvents()) < 1 {
			t.Fatal("should have emmited one domain event")
		}

		domainEvent := event.GetDomainEvents()[0]
		if !reflect.DeepEqual(expectedDomainEvent, domainEvent) {
			t.Fatalf("should have emitted one domain event of %T but received %T", expectedDomainEvent, domainEvent)
		}
	})
}
