package logic

import (
	"fmt"
	"strings"

	"github.com/Funmi4194/myMod/helper"
	"github.com/Funmi4194/myMod/repo"
)

//CreateWord  creates new wordCount document
func CreateWord(wordName string, content string) (*repo.WordCount, error) {

	documents := repo.WordCount{
		DocumentName:     strings.ToLower(wordName),
		Content:          content,
		Words:            helper.Content(content),
		Characters:       helper.TotalCharacters(content),
		CharWithoutSpace: helper.CharWithoutSpace(content),
		Sentence:         helper.Sentence(content),
		Paragraphs:       helper.Paragraphs(content),
		Lines:            helper.Line(content),
	}
	// check if document exists
	if err := documents.FindDocument(); err == nil {
		return nil, fmt.Errorf("document already exists")
	}

	// if user does not exist
	if err := documents.Create(); err != nil {
		return nil, err
	}
	return &documents, nil
}
func GetDocument(wordName string) (*repo.WordCount, error) {

	document := repo.WordCount{
		DocumentName: wordName,
	}
	if err := document.FindDocument(); err != nil {
		return nil, fmt.Errorf("document not found")
	}
	return &document, nil
}

// func GetDocuments() (*repo.WordCounts, error) {
// 	var documents repo.WordCounts
// 	if err := documents.FindDocuments(); err != nil {
// 		return nil, err
// 	}

// }
