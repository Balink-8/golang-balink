package repositories

import (
	"capstone/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type EventRepositoryMock interface {
	GetEventsRepository(page int, limit int, order string, search string) ([]*models.Event, int, error)
	GetEventRepository(id string) (*models.Event, error)
	CreateRepository(Event models.Event) (*models.Event, error)
	UpdateRepository(id string, EventBody models.Event) (*models.Event, error)
	DeleteRepository(id string) error
}

type IEventRepositoryMock struct {
	Mock mock.Mock
}

func NewEventRepositoryMock(mock mock.Mock) EventRepositoryMock {
	return &IEventRepositoryMock{
		Mock: mock,
	}
}

func (a *IEventRepositoryMock) GetEventsRepository(page int, limit int, order string, search string) ([]*models.Event, int, error) {
	args := a.Mock.Called()
	if args.Get(0) == nil {
		return nil, 0, args.Get(1).(error)
	}

	Events := args.Get(0).([]*models.Event)

	return Events, 0, nil
}

func (a *IEventRepositoryMock) GetEventRepository(id string) (*models.Event, error) {
	args := a.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Event := args.Get(0).(models.Event)

	return &Event, nil
}

func (a *IEventRepositoryMock) CreateRepository(EventData models.Event) (*models.Event, error) {
	args := a.Mock.Called(EventData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	Event := args.Get(0).(models.Event)

	return &Event, nil
}

func (a *IEventRepositoryMock) UpdateRepository(id string, EventData models.Event) (*models.Event, error) {
	args := a.Mock.Called(id, EventData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	Event := args.Get(0).(models.Event)

	return &Event, nil
}

func (a *IEventRepositoryMock) DeleteRepository(id string) error {
	args := a.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
