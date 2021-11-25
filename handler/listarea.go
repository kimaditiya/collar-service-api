package handler

import (
	"collar-service-api/helper"
	"collar-service-api/listarea"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listAreaHandler struct {
	listAreaService listarea.Service
}

func NewListAreaHandler(listAreaService listarea.Service) *listAreaHandler {
	return &listAreaHandler{listAreaService}

}
func (h *listAreaHandler) GetAllListArea(c *gin.Context) {

	listArea, err := h.listAreaService.GetAllListArea()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listArea)

	c.JSON(http.StatusOK, response)
}
