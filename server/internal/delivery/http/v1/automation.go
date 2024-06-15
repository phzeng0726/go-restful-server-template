package v1

import (
	"net/http"

	"github.com/phzeng0726/go-server-template/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAutomationRoutes(api *gin.RouterGroup) {
	auth := api.Group("/automations", userIdentityMiddleware(h))
	{
		auth.GET("", h.getPeacockIdByProject)
	}
}

type queryAutomationsInput struct {
	ProjectName string `form:"project" binding:"required"`
}

func (h *Handler) getPeacockIdByProject(c *gin.Context) {
	var inp queryAutomationsInput
	if err := c.BindQuery(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	auto, err := h.services.Automations.GetIdByParam(c, service.QueryAutomationsInput{
		Project: inp.ProjectName,
		Type:    "Peacock",
	})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, auto.Id)
}
