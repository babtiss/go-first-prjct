package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userID(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "header is empty")
		return
	}

	headerArray := strings.Split(header, " ")
	if len(headerArray[1]) == 0 || len(headerArray) != 2 || headerArray[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "header is invalid")
		return
	}

	userId, err := h.services.Authorization.ParseJWT(headerArray[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}
