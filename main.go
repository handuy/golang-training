package main

import (
	"github.com/gin-gonic/gin"
	"golang-crud/controller"
)

func main(){
	env, err := controller.NewEnv()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", env.GetAllNote)
	r.Run("localhost:8181")
}