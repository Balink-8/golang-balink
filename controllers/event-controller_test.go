package controllers

// import (
// 	"fmt"
// 	"net/http"
// 	"regexp"
// 	"strconv"

// 	"github.com/labstack/echo/v4"

// 	h "capstone/helpers"
// 	"capstone/models"
// 	"capstone/services"
// )

// type EventController interface {
// 	GetEventsController(c echo.Context) error
// 	GetEventController(c echo.Context) error
// 	CreateController(c echo.Context) error
// 	UpdateController(c echo.Context) error
// 	DeleteController(c echo.Context) error
// }

// type eventController struct {
// 	EventS services.EventService
// }

// func NewEventController(EventS services.EventService) EventController {
// 	return &eventController{
// 		EventS: EventS,
// 	}
// }

// func (e *eventController) GetEventsController(c echo.Context) error {
// 	page, err := strconv.Atoi(c.QueryParam("page"))
// 	if err != nil || page < 1 {
// 		page = 1
// 	}

// 	limit, err := strconv.Atoi(c.QueryParam("limit"))
// 	if err != nil || limit < 1 {
// 		limit = 10
// 	}

// 	order := c.QueryParam("order")
// 	search := c.QueryParam("search")

// 	Events, totalData, err := e.EventS.GetEventsService(page, limit, order, search)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	responseData := map[string]interface{}{
// 		"data":       Events,
// 		"page":       page,
// 		"data_shown": len(Events),
// 		"total_data": totalData,
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    responseData,
// 		Message: "Get all Event success",
// 		Status:  true,
// 	})
// }

// func (e *eventController) GetEventController(c echo.Context) error {
// 	id := c.Param("id")

// 	err := h.IsNumber(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	var Event *models.Event

// 	Event, err = e.EventS.GetEventService(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusNotFound, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    Event,
// 		Message: "Get Event success",
// 		Status:  true,
// 	})
// }

// func (e *eventController) CreateController(c echo.Context) error {
// 	var Event models.Event

// 	fmt.Println("Data :", &Event)

// 	err := c.Bind(&Event)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	file, err := c.FormFile("image") // Mengubah ctx menjadi c pada bagian ini

// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: "Image cannot be empty", // Mengubah pesan error menjadi string statis
// 			Status:  false,
// 		})
// 	}

// 	src, err := file.Open()
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: "Failed to open file", // Mengubah pesan error menjadi string statis
// 			Status:  false,
// 		})
// 	}

// 	re := regexp.MustCompile(`.png|.jpeg|.jpg`)

// 	if !re.MatchString(file.Filename) {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: "The provided file format is not allowed. Please upload a JPEG or PNG image", // Mengubah pesan error menjadi string statis
// 			Status:  false,
// 		})
// 	}

// 	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: src})
// 	if err != nil {
// 		return h.Response(c, http.StatusInternalServerError, h.ResponseModel{
// 			Data:    nil,
// 			Message: "Error uploading photo", // Mengubah pesan error menjadi string statis
// 			Status:  false,
// 		})
// 	}
// 	Event.Image = uploadUrl // Mengubah artikelInput menjadi Produk

// 	Event, err = e.EventS.CreateService(Event)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    Event,
// 		Message: "Create Event success",
// 		Status:  true,
// 	})
// }

// func (e *eventController) UpdateController(c echo.Context) error {
// 	id := c.Param("id")

// 	err := h.IsNumber(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	var Event *models.Event

// 	err = c.Bind(&Event)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	Event, err = e.EventS.UpdateService(id, *Event)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    Event,
// 		Message: "Update Event success",
// 		Status:  true,
// 	})
// }

// func (e *eventController) DeleteController(c echo.Context) error {
// 	id := c.Param("id")

// 	err := h.IsNumber(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	err = e.EventS.DeleteService(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    nil,
// 		Message: "Delete Event success",
// 		Status:  true,
// 	})
// }
