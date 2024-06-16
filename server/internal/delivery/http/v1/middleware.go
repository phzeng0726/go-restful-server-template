package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/phzeng0726/go-server-template/pkg/auth"

	"github.com/gin-gonic/gin"
)

// NOTE: 本檔案是驗證用的，有過的話代表是受驗證
const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

// 接受的格式是 Bearer jwt
func (h *Handler) parseAuthHeader(c *gin.Context) (auth.CustomMapClaims, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return auth.CustomMapClaims{}, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
		return auth.CustomMapClaims{}, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return auth.CustomMapClaims{}, errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

// 驗證，有過才會進到下一層
func userIdentityMiddleware(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := h.parseAuthHeader(c)
		if err != nil {
			newResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set(userCtx, claims.UserId)
		c.Next()
	}
}
