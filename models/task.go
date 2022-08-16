package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID           string             `json:"id" bson:"id"`
	EntryCode    string             `json:"entry_code" bson:"entry_code"`
	Assignee     string             `json:"assignee" bson:"assignee"`
	Tags         []string           `json:"tags" bson:"tags"`
	DueDate      primitive.DateTime `json:"due_date" bson:"due_date"`
	CreationDate primitive.DateTime `json:"creation_date" bson:"creation_date"`
	UpdateDate   primitive.DateTime `json:"update_date" bson:"update_date"`
	ViewDate     primitive.DateTime `json:"view_date" bson:"view_date"`
}
