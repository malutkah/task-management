package task

import "task-management/internal/database"

type Tasks = []Task

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type NewTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

var TaskTableStruct = []database.Column{
	{"id", database.ColumnVarchar},
	{"title", database.ColumnVarchar},
	{"description", database.ColumnVarchar},
	{"due_date", database.ColumnVarchar},
	{"created_at", database.ColumnVarchar},
	{"updated_at", database.ColumnVarchar},
}
