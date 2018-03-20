package main

import  (
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  v1 := router.Group("api/v1/todos") {
    v1.POST("/", createTodo)
  }
}
