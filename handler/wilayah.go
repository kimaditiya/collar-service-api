package handler

import (
	"collar-service-api/helper"
	"collar-service-api/wilayah"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WilayahHandler struct {
	WilayahService wilayah.Service
}

func ListWilayahHandler(listWilayahService wilayah.Service) *WilayahHandler {
	return &WilayahHandler{listWilayahService}

}
func (w *WilayahHandler) GetAllListWilayah(c *gin.Context) {

	listWilayah, err := w.WilayahService.GetAllListWilayah()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listWilayah)

	c.JSON(http.StatusOK, response)
}

func (w *WilayahHandler) GetAllListWilayahAssign(c *gin.Context) {

	listWilayah, err := w.WilayahService.GetAllListWilayahAssign()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listWilayah)

	c.JSON(http.StatusOK, response)
}

func (w *WilayahHandler) WilayahTagihNew(c *gin.Context) {

	var input wilayah.CreateWilayah
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Created Wilayah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newWilayah, err := w.WilayahService.CreateWilayah(input)

	if err != nil {
		response := helper.APIResponse("Failed Created Wilayah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := wilayah.FormatCreateWilayah(newWilayah)

	response := helper.APIResponse("Wilayah Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
