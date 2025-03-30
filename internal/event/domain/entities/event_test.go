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

func TestPublish(t *testing.T) {
	partnerId, _ := common.CreateUUID()
	event, err := entities.CreateEvent(entities.CreateEventCommand{
		Name:        "Testing Event",
		Description: "A testing event",
		Date:        time.Date(2025, 3, 24, 13, 30, 0, 0, time.UTC),
		IsPublished: false,
		PartnerId:   string(partnerId),
	})

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	event.AddSection(entities.CreateEventSectionCommand{
		Name:        "Testing Section",
		Description: "A testing section",
		TotalSpots:  20,
		Price:       20,
	})

	if event.GetTotalSpots() != 20 {
		t.Fatalf("expected 20 total spots, received %d", event.GetTotalSpots())
	}

	if len(event.GetSections()) < 1 {
		t.Fatalf("expected 1 section, received %d", len(event.GetSections()))
	}

	shouldBePublished(t, event, false)

	event.Publish()

	if !event.IsPublished() {
		t.Fatal("event should be published")
	}

	shouldBePublished(t, event, true)
}

func shouldBePublished(t *testing.T, event *entities.Event, should bool) {
	t.Helper()
	for _, section := range event.GetSections() {
		if section.IsPublished() != should {
			t.Fatalf("section %q IsPublished() should return %v", section.GetID(), should)
		}

		for _, spot := range section.GetSpots() {
			if spot.IsPublished() != should {
				t.Fatalf("spot %q IsPublished() should return %v", spot.GetID(), should)
			}

		}
	}
}
