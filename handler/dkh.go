package handler

import (
	"collar-service-api/dkh"
	"collar-service-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DkhHandler struct {
	DkhService dkh.Service
}

func ListDkhHandler(listDkhService dkh.Service) *DkhHandler {
	return &DkhHandler{listDkhService}

}
func (d *DkhHandler) GetAllListDkh(c *gin.Context) {

	listDkh, err := d.DkhService.GetAllListDkh()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listDkh)

	c.JSON(http.StatusOK, response)
}

func (d *DkhHandler) DkhNew(c *gin.Context) {

	var input dkh.CreateDkhInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Created Wilayah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newDkh, err := d.DkhService.CreateDkh(input)

	if err != nil {
		response := helper.APIResponse("Failed Created Wilayah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := dkh.FormatCreateDkh(newDkh)

	response := helper.APIResponse("Wilayah Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
