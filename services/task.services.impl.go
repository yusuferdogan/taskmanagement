package services

import (
	"context"
	"errors"
	"example/taskmanagement/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskServiceImpl struct {
	taskcollection *mongo.Collection
	ctx            context.Context
}

func NewTaskService(taskcollection *mongo.Collection, ctx context.Context) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskcollection: taskcollection,
		ctx:            ctx,
	}
}

func (u *TaskServiceImpl) CreateTask(task *models.Task) error {

	if task.Group < models.ToDo || task.Group > models.Done {
		return errors.New("invalid group")
	}
	var mtask *models.Task
	query := bson.D{bson.E{Key: "entry_code", Value: task.EntryCode}}
	u.taskcollection.FindOne(u.ctx, query).Decode(&mtask)
	if mtask != nil {
		fmt.Println(mtask)
		return errors.New("key code already exists")
	}

	task.ViewDate = time.Now()
	task.ID = primitive.NewObjectID().Hex()
	_, err := u.taskcollection.InsertOne(u.ctx, task)

	return err
}

func (u TaskServiceImpl) GetTask(id *string) (*models.Task, error) {
	var task *models.Task
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := u.taskcollection.FindOne(u.ctx, query).Decode(&task)
	if err == nil {
		task.ViewDate = time.Now()
		u.UpdateTask(task)
	}
	return task, err
}

func (u TaskServiceImpl) GetAll() ([]*models.Task, error) {
	var tasks []*models.Task
	sortOptions := options.Find()
	sortOptions.SetSort(bson.D{{Key: "view_date", Value: -1}})

	cursor, err := u.taskcollection.Find(u.ctx, bson.D{}, sortOptions)

	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var task models.Task
		err := cursor.Decode(&task)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tasks = append(tasks, &task)
		if err := cursor.Err(); err != nil {
			return nil, err
		}
	}
	cursor.Close(u.ctx)

	if len(tasks) == 0 {
		return nil, errors.New("document not found")
	}
	return tasks, nil

}

func (u TaskServiceImpl) UpdateTask(task *models.Task) error {
	filter := bson.D{bson.E{Key: "id", Value: task.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "entry_code", Value: task.EntryCode}, bson.E{Key: "assignee", Value: task.Assignee},
		bson.E{Key: "tags", Value: task.Tags}, bson.E{Key: "due_date", Value: task.DueDate}, bson.E{Key: "creation_date", Value: task.CreationDate},
		bson.E{Key: "update_date", Value: task.UpdateDate}, bson.E{Key: "view_date", Value: task.ViewDate}}}}
	result, _ := u.taskcollection.UpdateOne(u.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (u *TaskServiceImpl) DeleteTask(id *string) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := u.taskcollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no mathced document found for update")
	}

	return nil
}
