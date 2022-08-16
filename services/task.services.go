package services

import "example/taskmanagement/models"

type TaskService interface {
	CreateTask(*models.Task) error
	GetTask(*string) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	UpdateTask(*models.Task) error
	DeleteTask(*string) error
}
