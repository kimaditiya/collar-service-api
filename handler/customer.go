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
func (h *customerHandler) FindListOfCustomerByProduct(c *gin.Context) {
	var input customers.FilterCustomerByProduct

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("No Data Found", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}
	listCustomerByProduct, err := h.customerService.FindLisOfCustomerByProduct(input)

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listCustomerByProduct)
	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) FindListOfCustomerByOvdDays(c *gin.Context) {
	var input customers.FilterListCustomerByOvd

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("No Data Found", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}
	listCustomerByOvd, err := h.customerService.FindListOfCustomerByOvdDays(input)

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listCustomerByOvd)
	c.JSON(http.StatusOK, response)
}
func (h *customerHandler) FindListOfCustomerByDueDate(c *gin.Context) {
	var input customers.FilterListCustomerByDueDate

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("No Data Found", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}
	listCustomerByDueDate, err := h.customerService.FindListOfCustomerByDueDate(input)

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listCustomerByDueDate)
	c.JSON(http.StatusOK, response)
}
func (h *customerHandler) ListOfCustomer(c *gin.Context) {
	listCustomer, err := h.customerService.ListOfCustomer()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listCustomer)

	c.JSON(http.StatusOK, response)
}

func (h *customerHandler) ListOfCustomerColl(c *gin.Context) {

	listCustomerColl, err := h.customerService.ListOfCustomerColl()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		c.IndentedJSON(http.StatusOK, errorMessage)
		return
	}

	c.IndentedJSON(http.StatusOK, listCustomerColl)
}
