package handler

import (
	"collar-service-api/collectiongroup"
	"collar-service-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type collectiongroupHandler struct {
	collectiongroupService collectiongroup.Service
}

func NewCollectionGroupHandler(collectiongroupService collectiongroup.Service) *collectiongroupHandler {
	return &collectiongroupHandler{collectiongroupService}
}
func (h *collectiongroupHandler) ListOfPrivelege(c *gin.Context) {
	listPrivelge, err := h.collectiongroupService.ListPrivelege()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listPrivelge)

	c.JSON(http.StatusOK, response)
}

func (h *collectiongroupHandler) CollectionGroupNew(c *gin.Context) {

	var input collectiongroup.CreateCGInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Created Collection Group", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newCG, err := h.collectiongroupService.CreateCG(input)

	if err != nil {
		response := helper.APIResponse("Failed Created Collection Group", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := collectiongroup.FormatCreateCG(newCG)

	response := helper.APIResponse("Collection Group Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *collectiongroupHandler) CollectionGroupUpdate(c *gin.Context) {

	var input collectiongroup.CreateCGInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Collection Group", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newCG, err := h.collectiongroupService.UpdateCG(input)

	if err != nil {
		response := helper.APIResponse("Failed Update Collection Group", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := collectiongroup.FormatCreateCG(newCG)

	response := helper.APIResponse("Collection Group Update", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *collectiongroupHandler) ListCollectionGroup(c *gin.Context) {
	listPrivelge, err := h.collectiongroupService.ListCG()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listPrivelge)

	c.JSON(http.StatusOK, response)
}
