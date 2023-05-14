package pokemon

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pokemon struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Type        []string           `json:"type" bson:"type,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Height      float32            `json:"height" bson:"height,omitempty"`
	Weight      float32            `json:"weight" bson:"weight,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
}
