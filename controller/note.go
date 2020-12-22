package controller

import (
	"golang-crud/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(env *env) GetAllNote(c *gin.Context) {
	result, err := model.GetAllNotes(env.db)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}