package services

import (
	"capstone/models"
	"capstone/repositories"
)

type EventService interface {
	GetEventsService(page int, limit int, order string, search string) ([]*models.Event, int, error)
	GetEventService(id string) (*models.Event, error)
	CreateService(Event models.Event) (models.Event, error)
	UpdateService(id string, EventBody models.Event) (*models.Event, error)
	DeleteService(id string) error
}

type eventService struct {
	EventR repositories.EventRepository
}

func NewEventService(EventR repositories.EventRepository) EventService {
	return &eventService{
		EventR: EventR,
	}
}

func (e *eventService) GetEventsService(page int, limit int, order string, search string) ([]*models.Event, int, error) {
	Events, totalData, err := e.EventR.GetEventsRepository(page, limit, order, search)
	if err != nil {
		return nil, 0, err
	}

	return Events, totalData, nil
}

func (e *eventService) GetEventService(id string) (*models.Event, error) {
	Event, err := e.EventR.GetEventRepository(id)
	if err != nil {
		return nil, err
	}

	return Event, nil
}

func (e *eventService) CreateService(Event models.Event) (models.Event, error) {
	EventR, err := e.EventR.CreateRepository(Event)
	if err != nil {
		return models.Event{}, err
	}

	return EventR, nil
}

func (e *eventService) UpdateService(id string, EventBody models.Event) (*models.Event, error) {
	Event, err := e.EventR.UpdateRepository(id, EventBody)
	if err != nil {
		return Event, err
	}

	return Event, nil
}

func (e *eventService) DeleteService(id string) error {
	err := e.EventR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}