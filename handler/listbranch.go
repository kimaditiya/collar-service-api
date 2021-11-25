package handler

import (
	"collar-service-api/helper"
	"collar-service-api/listbranch"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listBranchHandler struct {
	listBranchService listbranch.Service
}

func NewListBranchHandler(listBranchService listbranch.Service) *listBranchHandler {
	return &listBranchHandler{listBranchService}

}
func (h *listBranchHandler) GetAllListBranch(c *gin.Context) {

	listBranch, err := h.listBranchService.GetAllListBranch()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listBranch)

	c.JSON(http.StatusOK, response)
}
func (h *listBranchHandler) FindListOfBranchByAreaID(c *gin.Context) {
	var input listbranch.BranchByAreaInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("No Data Found", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}
	listBranchByAreaID, err := h.listBranchService.BranchByArea(input)

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listBranchByAreaID)
	c.JSON(http.StatusOK, response)
}
func (h *listBranchHandler) GetAllListBranchPos(c *gin.Context) {

	listBranchPos, err := h.listBranchService.GetAllListBranchPos()

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listBranchPos)

	c.JSON(http.StatusOK, response)
}
func (h *listBranchHandler) FindPostByBranch(c *gin.Context) {
	var input listbranch.BranchByPosInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("No Data Found", http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusOK, response)
		return
	}
	findPostByBranch, err := h.listBranchService.FindPosByBranch(input)

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", findPostByBranch)
	c.JSON(http.StatusOK, response)
}
