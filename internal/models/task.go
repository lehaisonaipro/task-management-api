package models

type Task struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"` // "Pending", "In Progress", "Completed"
	AssignedTo  string `json:"assigned_to" bson:"assigned_to"`
	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	DueDate     int64  `json:"due_date" bson:"due_date"`
}

type TaskStatus string

const (
	Pending    TaskStatus = "Pending"
	InProgress TaskStatus = "In Progress"
	Completed  TaskStatus = "Completed"
	Any        TaskStatus = ""
)

type TaskSummary struct {
	EmployeeID string `json:"employee_id" bson:"employee_id"`
	TotalTasks int    `json:"total_tasks" bson:"total_tasks"`
	Pending    int    `json:"pending" bson:"pending"`
	InProgress int    `json:"in_progress" bson:"in_progress"`
	Completed  int    `json:"completed" bson:"completed"`
}
