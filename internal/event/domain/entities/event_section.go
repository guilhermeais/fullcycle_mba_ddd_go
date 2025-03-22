package entities

import "ingressos/internal/common"

type EventSectionId = common.UUID
type EventSection struct {
	id                 EventSectionId
	name               string
	description        string
	totalSpots         int
	totalSpotsReserved int
	price              int
	isPublished        bool
	spots              []*EventSpot
}

type CreateEventSectionCommand struct {
	Name        string
	Description string
	TotalSpots  int
	Price       int
}

func CreateEventSection(c CreateEventSectionCommand) (*EventSection, error) {
	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	var spots = make([]*EventSpot, c.TotalSpots)
	for range c.TotalSpots {
		spot, err := CreateEventSpot(CreateEventSpotCommand{})
		if err != nil {
			return nil, err
		}

		spots = append(spots, spot)
	}

	return &EventSection{
		id:                 uuid,
		name:               c.Name,
		description:        c.Description,
		totalSpots:         c.TotalSpots,
		price:              c.Price,
		totalSpotsReserved: 0,
		isPublished:        false,
		spots:              spots,
	}, nil
}
