package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-management/internal/database"
	"time"
)

var TaskTable *database.Table

func GetAllTasks(c *gin.Context) {
	allTasks := TaskTable.Select(database.ExpSelectAll)

	c.IndentedJSON(200, allTasks)
}

func GetSingleTask(c *gin.Context) {
	id := c.Param("id")

	cond := []database.Condition{{"id", database.ConditionEqual, id}}
	task := TaskTable.Select(database.ExpSelectAll, cond...)
	c.IndentedJSON(http.StatusCreated, task)
}

func CreateNewTask(c *gin.Context) {
	var newTask NewTaskRequest

	if err := c.Bind(&newTask); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")

	id := database.GetRandomID()
	TaskTable.Insert(id, newTask.Title, newTask.Description, newTask.DueDate, formattedTime, formattedTime)

	cond := []database.Condition{{"id", database.ConditionEqual, id}}
	inserted := TaskTable.Select(database.ExpSelectAll, cond...)
	c.IndentedJSON(http.StatusCreated, inserted)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updateTask NewTaskRequest

	if err := c.Bind(&updateTask); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	set := []database.Set{
		{"title", updateTask.Title},
		{"description", updateTask.Description},
		{"due_date", updateTask.DueDate},
	}
	cond := []database.Condition{{"id", database.ConditionEqual, id}}
	TaskTable.Update(set, cond)

	updated := TaskTable.Select(database.ExpSelectAll, cond...)
	c.IndentedJSON(http.StatusCreated, updated)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	cond := []database.Condition{{"id", database.ConditionEqual, id}}

	TaskTable.Delete(cond)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Entry deleted",
	})
}
