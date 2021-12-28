package handler

import (
	"collar-service-api/helper"
	"collar-service-api/reasons"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReasonsHandler struct {
	ReasonsService reasons.Service
}

func ListReasonsHandler(listReasonsService reasons.Service) *ReasonsHandler {
	return &ReasonsHandler{listReasonsService}

}
func (rs *ReasonsHandler) GetAllListReasons(c *gin.Context) {

	listReasons, err := rs.ReasonsService.GetAllListReason()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listReasons)

	c.JSON(http.StatusOK, response)
}
