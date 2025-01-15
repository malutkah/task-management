package main

import (
	"log"
	
	"github.com/gin-gonic/gin"
	"task-management/internal/task"
)

func main() {
	router := gin.New()
	
	task.RouteTasks(router)
	
	err := router.Run()
	if err != nil {
		log.Fatalf("[ERROR] running router: %v", err)
	}
}
