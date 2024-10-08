package steps

import "go.mongodb.org/mongo-driver/bson/primitive"

type Style struct {
	Type string `json:"type" bson:"type"`
}

type Metadata struct {
	Source string `json:"source" bson:"source"`
}

type Step struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Creator      string             `json:"creator" bson:"creator"`
	DateCreation string             `json:"date-of-creation" bson:"date-of-creation"`
	Region       string             `json:"region" bson:"region"`
	History      string             `json:"history" bson:"history"`
	VideoPath    string             `json:"video-path" bson:"video-path"`
	Style        Style              `json:"style" bson:"style"`
	Meta         *Metadata          `json:"meta,omitempty" bson:"meta,omitempty"`
}
