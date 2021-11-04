package handler

import (
	"collar-service-api/customers"
	"collar-service-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerService customers.Service
}

func NewCustomerHandler(customerService customers.Service) *customerHandler {
	return &customerHandler{customerService}
}

func (h *customerHandler) ListOfCustomer(c *gin.Context) {
	var input customers.FindCustomer

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	listCustomer, err := h.customerService.ListOfCustomer(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listCustomer)

	c.JSON(http.StatusOK, response)
}
