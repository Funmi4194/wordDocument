package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type WordCount struct {
	ID               primitive.ObjectID `json:"id" bson:"ID,omitempty"`
	DocumentName     string             `json:"document_name" bson:"document_name"`
	Content          string             `json:"content" bson:"content"`
	Words            int                `json:"words" bson:"words"`
	Characters       int                `json:"characters" bson:"characters"`
	CharWithoutSpace int                `json:"char_without_space" bson:"char_without_space"`
	Sentence         int                `json:"sentence" bson:"sentence"`
	Paragraphs       int                `json:"paragraphs" bson:"paragraphs"`
	Lines            int                `json:"lines" bson:"lines"`
	Pages            int                `json:"page" bson:"page"`
}

type WordCounts []WordCount
