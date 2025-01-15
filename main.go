package main

import (
	"log"
	"task-management/internal/database"

	"github.com/gin-gonic/gin"
	"task-management/internal/task"
)

func main() {
	db := database.NewDatabase()

	taskTable, err := db.CreateTable("tasks", task.TaskTableStruct)
	if err != nil {
		log.Fatalf("Error creating table 'tasks': %s", err.Error())
	}

	task.TaskTable = taskTable

	router := gin.New()

	task.RouteTasks(router)

	err = router.Run()
	if err != nil {
		log.Fatalf("[ERROR] running router: %v", err)
	}
}
