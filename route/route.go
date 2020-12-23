package route

import (
	"github.com/gin-gonic/gin"
	"golang-crud/controller"
)

func SetupRoute() *gin.Engine {
	env, err := controller.NewEnv()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	r.GET("/", env.GetAllNote)
	r.GET("/:id", env.GetNoteById)
	r.POST("/new", env.CreateNote)
	r.POST("/update", env.UpdateNote)

	return r
}
