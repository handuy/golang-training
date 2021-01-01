package route

import (
	"github.com/gin-gonic/gin"
	"golang-crud/controller"
)

func noteRoutes(rg *gin.RouterGroup, env *controller.Env) {
	notes := rg.Group("/notes")

	notes.GET("/", env.GetAllNote)
	notes.GET("/:id", env.GetNoteById)
	notes.POST("/new", env.CreateNote)
	notes.POST("/update", env.UpdateNote)
	notes.POST("/delete", env.DeleteNote)
}