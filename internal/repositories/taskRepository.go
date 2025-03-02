package repositories

import (
	"context"

	"github.com/lehaisonaipro/task-management-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(taskID string) (*models.Task, error)
	AssignTask(taskID string, employee string) error
	UpdateTaskStatus(taskID string, status models.TaskStatus) error
	ViewTasks(employeeID string, status string, sortDateType string, sortDateAsc bool, sortTastStaus bool, stortTastAsc bool) ([]*models.Task, error)
	TaskSummaryByEmployee(employeeID string) ([]*models.TaskSummary, error)
	AllTaskSummary() (*models.TaskSummary, error)
}

type TaskRepository struct {
	collection *mongo.Collection
}

// AssignTask implements ITaskRepository.
func (repo *TaskRepository) AssignTask(taskID string, employee string) error {
	// Define filter to find the task by its ID
	filter := bson.M{"_id": taskID}

	// Define the update operation to set the assigned employee
	update := bson.M{"$set": bson.M{"assigned_to": employee}}

	// Execute the update operation
	result, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	// Check if the task was found and updated
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments // No task found with the given ID
	}
	return nil
}

// GetTastk implements ITaskRepository.
func (repo *TaskRepository) GetTask(taskID string) (*models.Task, error) {
	var task models.Task

	// Define filter to find the task by its ID
	filter := bson.M{"_id": taskID}

	// Query the collection
	err := repo.collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No task found
		}
		return nil, err
	}

	return &task, nil
}

// Return total number of tasks, pending tasks, tasks in progress, and completed tasks
func (repo *TaskRepository) AllTaskSummary() (*models.TaskSummary, error) {
	var summary models.TaskSummary

	// Define the MongoDB aggregation pipeline
	pipeline := mongo.Pipeline{
		// Group all tasks and count the tasks by status
		{
			{Key: "$group", Value: bson.M{
				"_id":         nil,
				"total_tasks": bson.M{"$sum": 1},
				"pending":     bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", "pending"}}, 1, 0}}},
				"in_progress": bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", "in_progress"}}, 1, 0}}},
				"completed":   bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", "completed"}}, 1, 0}}},
			}},
		},
		// Rename fields
		{
			{Key: "$project", Value: bson.M{
				"total_tasks": 1,
				"pending":     1,
				"in_progress": 1,
				"completed":   1,
			}},
		},
	}

	// Execute the aggregation pipeline
	cursor, err := repo.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the result
	if cursor.Next(context.Background()) {
		if err := cursor.Decode(&summary); err != nil {
			return nil, err
		}
	} else {
		// If no tasks exist, return a default summary with zeros
		summary = models.TaskSummary{
			TotalTasks: 0,
			Pending:    0,
			InProgress: 0,
			Completed:  0,
		}
	}

	return &summary, nil
}

func (repo *TaskRepository) TaskSummaryByEmployee(employeeID string) ([]*models.TaskSummary, error) {
	var summaries []*models.TaskSummary

	// Define the match filter
	filter := bson.M{}
	if employeeID != "" {
		filter["assigned_to"] = employeeID
	}

	// MongoDB aggregation pipeline to summarize tasks
	pipeline := mongo.Pipeline{
		// Match tasks based on employeeID (if provided)
		{{Key: "$match", Value: filter}},
		// Group tasks by employee and count the tasks by status
		{
			{Key: "$group", Value: bson.M{
				"_id":         "$assigned_to",
				"total_tasks": bson.M{"$sum": 1},
				"pending":     bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", "pending"}}, 1, 0}}},
				"in_progress": bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", "in_progress"}}, 1, 0}}},
				"completed":   bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", "completed"}}, 1, 0}}},
			}},
		},
		// Rename fields
		{
			{Key: "$project", Value: bson.M{
				"employee_id": "$_id",
				"total_tasks": 1,
				"pending":     1,
				"in_progress": 1,
				"completed":   1,
			}},
		},
	}

	// Run aggregation
	cursor, err := repo.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode results
	for cursor.Next(context.Background()) {
		var summary models.TaskSummary
		if err := cursor.Decode(&summary); err != nil {
			return nil, err
		}
		summaries = append(summaries, &summary)
	}
	return summaries, nil
}

func (repo *TaskRepository) ViewTasks(employeeID string, status string, sortDateType string, sortDateAsc bool, sortTastStatus bool, sortTastAsc bool) ([]*models.Task, error) {
	var tasks []*models.Task

	// Define filter conditions
	filter := bson.M{}

	// Filter by employee ID if provided
	if employeeID != "" {
		filter["assigned_to"] = employeeID
	}

	// Filter by task status if provided
	if status != "" {
		filter["status"] = status
	}

	// Define sorting options
	sortOptions := bson.D{}

	// Sorting by date if sortDateType is provided
	if sortDateType != "" {
		sortOrder := 1
		if !sortDateAsc {
			sortOrder = -1
		}
		sortOptions = append(sortOptions, bson.E{Key: sortDateType, Value: sortOrder})
	}

	// Sorting by task status if requested
	if sortTastStatus {
		taskSortOrder := 1
		if !sortTastAsc {
			taskSortOrder = -1
		}
		sortOptions = append(sortOptions, bson.E{Key: "status", Value: taskSortOrder})
	}

	// Find tasks with the defined filter and sorting
	findOptions := options.Find().SetSort(sortOptions)
	cursor, err := repo.collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the tasks
	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (repo *TaskRepository) UpdateTaskStatus(taskID string, status models.TaskStatus) error {
	filter := bson.M{"_id": taskID}
	update := bson.M{"$set": bson.M{"status": string(status)}}
	_, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func NewTaskRepository(collection *mongo.Collection) ITaskRepository {
	return &TaskRepository{collection: collection}
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	_, err := repo.collection.InsertOne(context.Background(), task)
	if err != nil {
		return err
	}
	return err
}

func (repo *TaskRepository) GetTasksByEmloyee(employeeID string) ([]models.Task, error) {
	var tasks []models.Task

	filter := bson.M{"assigned_to": employeeID} // Filter by assigned_to field

	cursor, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
