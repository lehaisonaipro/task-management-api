package models

type Task struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"` // "Pending", "In Progress", "Completed"
	AssignedTo  string `json:"assigned_to" bson:"assigned_to"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	DueDate     string `json:"due_date" bson:"due_date"`
}
