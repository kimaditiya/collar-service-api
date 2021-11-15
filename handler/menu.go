package handler

import (
	"collar-service-api/helper"
	"collar-service-api/menu"
	"net/http"

	"github.com/gin-gonic/gin"
)

type menuHandler struct {
	menuService menu.Service
}

func NewMenuHandler(menuService menu.Service) *menuHandler {
	return &menuHandler{menuService}
}

func (h *menuHandler) MenuNew(c *gin.Context) {

	var input menu.MenuInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Created Menu", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newMenu, err := h.menuService.MenuInput(input)

	if err != nil {
		response := helper.APIResponse("Failed Created Account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := menu.FormatMenu(newMenu)

	response := helper.APIResponse("Account Created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
