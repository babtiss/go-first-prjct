package handler

import (
	"github.com/gin-gonic/gin"
	todo "go-application/model"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}
func (h *Handler) signIn(c *gin.Context) {

}
