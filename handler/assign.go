package handler

import (
	"collar-service-api/assign"
	"collar-service-api/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type AssignHandler struct {
	AssignService assign.Service
}

func ListAssignHandler(listAssignService assign.Service) *AssignHandler {
	return &AssignHandler{listAssignService}

}

func (a *AssignHandler) GetAllListAssign(c *gin.Context) {

	filter := &assign.AssignFilter{}

	paramMap := make(map[string]interface{}, 0)
	for k, v := range c.Request.URL.Query() {
		if len(v) == 1 && len(v[0]) != 0 {
			paramMap[k] = v[0]
		} else {
			break
		}
	}

	mapstructure.Decode(paramMap, &filter)

	fmt.Printf("%+v\n", filter)

	listAssign, err := a.AssignService.GetAllListAssign(filter)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("No Data Found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Data Found", http.StatusOK, "Data Found", listAssign)

	c.JSON(http.StatusOK, response)
}

func (a *AssignHandler) AssignNew(c *gin.Context) {

	var input assign.CreateAssign
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Created Assign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newAssign, err := a.AssignService.CreateAssign(input)

	if err != nil {
		response := helper.APIResponse("Failed Created Assign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := assign.FormatCreateAssign(newAssign)

	response := helper.APIResponse("Assign Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
