package v1

import (
	"errors"

	"github.com/phzeng0726/go-server-template/internal/service"
	"github.com/phzeng0726/go-server-template/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

// 獲取呼叫api的用戶id並轉成字串
func (h *Handler) getUserIdByCtx(c *gin.Context) (string, error) {
	userId := c.Value("userId")

	userIdStr, ok := userId.(string)
	if !ok {
		return "", errors.New("failed to convert value to string")
	}

	return userIdStr, nil
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAutomationRoutes(v1)
	}
}
