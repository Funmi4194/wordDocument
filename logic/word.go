package logic

import (
	"fmt"
	"strings"

	"github.com/Funmi4194/myMod/helper"
	repo "github.com/Funmi4194/myMod/repository"
)

//CreateWord  creates new wordCount document
func CreateWord(w repo.WordCount) (*repo.WordCount, error) {

	documents := repo.WordCount{
		DocumentName:     strings.ToLower(w.DocumentName),
		Content:          w.Content,
		Words:            helper.Content(w.Content),
		Characters:       helper.TotalCharacters(w.Content),
		CharWithoutSpace: helper.CharWithoutSpace(w.Content),
		Sentence:         helper.Sentence(w.Content),
		Paragraphs:       helper.Paragraphs(w.Content),
		Lines:            helper.Line(w.Content),
		Pages:            helper.Pages(w.Content),
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

//Get a document by its name
func GetDocument(wordName string) (*repo.WordCount, error) {

	document := repo.WordCount{
		DocumentName: strings.ToLower(wordName),
	}

	if err := document.FindDocument(); err != nil {
		return nil, fmt.Errorf("document not found")
	}
	return &document, nil
}

func GetDocuments() (*repo.WordCounts, error) {
	var documents repo.WordCounts
	if err := documents.Documents(); err != nil {
		return nil, fmt.Errorf("no document was found")
	}
	return &documents, nil
}
