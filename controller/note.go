package controller

import (
	"golang-crud/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (env *env) GetAllNote(c *gin.Context) {
	result, err := model.GetAllNotes(env.db)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (env *env) GetNoteById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	result, err := model.GetNoteById(env.db, idInt)
	if err != nil {
		if err.Error() == "Không tìm thấy note" {
			c.JSON(http.StatusNotFound, model.ErrorMessage{
				Message: "Không tìm thấy note",
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Lỗi server",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
