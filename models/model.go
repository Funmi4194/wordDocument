package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type WordCount struct {
	ID               primitive.ObjectID `json:"id" bson:"ID,omitempty"`
	DocumentName     string             `json:"documentname" bson:"DocumentName"`
	Content          string             `json:"content" bson:"Content"`
	Words            int                `json:"words" bson:"Words"`
	Characters       int                `json:"characters" bson:"Characters"`
	CharWithoutSpace int                `json:"charWithoutSpace" bson:"WordsCharWithoutSpace"`
	Sentence         int                `json:"sentence" bson:"Sentence"`
	Paragraphs       int                `json:"paragraphs" bson:"Paragraphs"`
	Lines            int                `json:"lines" bson:"Lines"`
}
