package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "capstone/helpers"
	"capstone/models"
	"capstone/services"
)

type EventController interface {
	GetEventsController(c echo.Context) error
	GetEventController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type eventController struct {
	EventS services.EventService
}

func NewEventController(EventS services.EventService) EventController {
	return &eventController{
		EventS: EventS,
	}
}

func (e *eventController) GetEventsController(c echo.Context) error {
	Events, err := e.EventS.GetEventsService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Events,
		Message: "Get all Events success",
		Status:  true,
	})
}

func (e *eventController) GetEventController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Event *models.Event

	Event, err = e.EventS.GetEventService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Event,
		Message: "Get Event success",
		Status:  true,
	})
}

func (e *eventController) CreateController(c echo.Context) error {
	var Event *models.Event

	err := c.Bind(&Event)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Event, err = e.EventS.CreateService(*Event)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Event,
		Message: "Create Event success",
		Status:  true,
	})
}

func (e *eventController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Event *models.Event

	err = c.Bind(&Event)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Event, err = e.EventS.UpdateService(id, *Event)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Event,
		Message: "Update Event success",
		Status:  true,
	})
}

func (e *eventController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = e.EventS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Event success",
		Status:  true,
	})
}
