package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	BearerToken  = "Bearer"
	GetHeaderKey = "Authorization"
)

func (h *Handler) userID(c *gin.Context) {
	header := c.GetHeader(GetHeaderKey)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "header is empty")
		return
	}

	headerArray := strings.Split(header, " ")
	if len(headerArray[1]) == 0 || len(headerArray) != 2 || headerArray[0] != BearerToken {
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

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("error user id")
	}
	return idInt, nil
}
