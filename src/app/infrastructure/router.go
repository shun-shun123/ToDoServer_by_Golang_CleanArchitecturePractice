package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/shun-shun123/clean_architecture_todo/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	todoController := controllers.NewToDoController(NewDBHandler())

	router.POST("/todo/create", func(c *gin.Context) {todoController.Create(c)})
	router.POST("/todo/read", func(c *gin.Context) {todoController.Read(c)})
	router.POST("/todo/update", func(c *gin.Context) {todoController.Update(c)})
	router.POST("/todo/delete", func(c *gin.Context) {todoController.Delete(c)})

	Router = router
}
