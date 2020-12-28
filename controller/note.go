package controller

import (
	"golang-crud/model"
	"net/http"
	"strconv"
	"time"

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
			c.JSON(http.StatusNotFound, model.StatusMessage{
				Message: "Không tìm thấy note",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, model.StatusMessage{
			Message: "Lỗi server",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (env *env) CreateNote(c *gin.Context) {
	var newNote model.NewNote
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(http.StatusBadRequest, model.StatusMessage{
			Message: "Bad request",
		})
		return
	}

	result, err := model.InsertNote(env.db, newNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.StatusMessage{
			Message: "Không thể tạo note mới",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (env *env) UpdateNote(c *gin.Context) {
	var updateNote model.UpdatedNote
	if err := c.ShouldBindJSON(&updateNote); err != nil {
		c.JSON(http.StatusBadRequest, model.StatusMessage{
			Message: "Bad request",
		})
		return
	}

	updateNote.UpdatedAt = time.Now()
	err := model.UpdateNote(env.db, updateNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.StatusMessage{
			Message: "Không thể cập nhật note",
		})
		return
	}

	c.JSON(http.StatusOK, model.StatusMessage{
		Message: "Cập nhật thành công",
	})
}

func (env *env) DeleteNote(c *gin.Context) {
	var deleteNote model.DeletedNote
	if err := c.ShouldBindJSON(&deleteNote); err != nil {
		c.JSON(http.StatusBadRequest, model.StatusMessage{
			Message: "Bad request",
		})
		return
	}
	if deleteNote.Id == 0 {
		c.JSON(http.StatusBadRequest, model.StatusMessage{
			Message: "Bad request",
		})
		return
	}

	err := model.DeleteNote(env.db, deleteNote)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.StatusMessage{
			Message: "Không thể xóa note",
		})
		return
	}

	c.JSON(http.StatusOK, model.StatusMessage{
		Message: "Xóa thành công",
	})
}
