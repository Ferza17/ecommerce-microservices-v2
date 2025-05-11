package bson

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationTemplate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Type         string             `bson:"type"`
	Template     string             `bson:"template"`
	TemplateVars map[string]any     `bson:"templateVars"`
}
