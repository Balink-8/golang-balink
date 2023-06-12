package controllers

import (
	"net/http"

	"capstone/models"
	"capstone/services"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type TransactionController interface {
	SaveTransactionProduct(c echo.Context) error
	DeleteTransactionProduct(c echo.Context) error
	GetTransactionProducts(c echo.Context) error
	SaveTransactionTicket(c echo.Context) error
	DeleteTransactionTicket(c echo.Context) error
	GetTransationTicket(c echo.Context) error
}

type transactionController struct {
	Repository services.TransactionService
}

func SaveTransactionProduct(c echo.Context) error {
	// Get the transaction product data from the request body
	transactionProduct := &models.TransactionProduct{}
	c.BindJSON(transactionProduct)
	// Save the transaction product
	transactionProduct, err := models.SaveTransactionProduct(transactionProduct)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Return the transaction product
	return c.JSON(http.StatusOK, transactionProduct)
}

func GetTransactionProduct(c echo.Context) error {
	// Get the transaction product ID from the request path
	productID := c.Param("id")
	// Get the transaction product from the database
	transactionProduct, err := models.GetTransactionProduct(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	// Return the transaction product
	return c.JSON(http.StatusOK, transactionProduct)
}

func DeleteTransactionProduct(c echo.Context) error {
	// Get the transaction product ID from the request path
	productID := c.Param("id")
	// Delete the transaction product from the database
	err := models.DeleteTransactionProduct(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	// Return a success response
	return c.JSON(http.StatusOK, gin.H{"message": "Transaction product deleted successfully"})
}

func SaveTransactionTicket(c echo.Context) error {
	// Get the transaction ticket data from the request body
	transactionTicket := &models.Ticket{}
	c.BindJSON(transactionTicket)
	// Save the transaction ticket
	transactionTicket, err := models.SaveTransactionTicket(transactionTicket)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Return the transaction ticket
	return c.JSON(http.StatusOK, transactionTicket)
}

func GetTransactionTicket(c echo.Context) error {
	// Get the transaction ticket ID from the request path
	ticketID := c.Param("id")
	// Get the transaction ticket from the database
	transactionTicket, err := models.GetTransactionTicket(ticketID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	// Return the transaction ticket
	return c.JSON(http.StatusOK, transactionTicket)
}

// func (t *transactionController) GetTransactionController(c gin.Context) error {
// 	Transactions, err := t.Repository.GetTransactionService()
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    Transactions,
// 		Message: "Get all Transaction success",
// 		Status:  true,
// 	})
// }

// func (t *transaksiController) GetTransaksiController(c gin.Context) error {
// 	id := c.Param("id")

// 	err := h.IsNumber(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	var Transaksi *models.Transaksi

// 	Transaksi, err = t.TransaksiS.GetTransaksiService(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusNotFound, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    Transaksi,
// 		Message: "Get Transaction success",
// 		Status:  true,
// 	})
// }

// func (t *transaksiController) DeleteController(c echo.Context) error {
// 	id := c.Param("id")

// 	err := h.IsNumber(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	err = t.TransaksiS.DeleteService(id)
// 	if err != nil {
// 		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
// 			Data:    nil,
// 			Message: err.Error(),
// 			Status:  false,
// 		})
// 	}

// 	return h.Response(c, http.StatusOK, h.ResponseModel{
// 		Data:    nil,
// 		Message: "Delete Transaction success",
// 		Status:  true,
// 	})
// }
