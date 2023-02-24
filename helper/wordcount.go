package helper

import (
	"strings"
)

//Counts the numbers of words
func Content(content string) int {
	word := strings.Join(strings.Fields(content), " ")
	count := len(strings.Split(word, " "))

	return count
}

//Counts the numbers of charcters
func TotalCharacters(content string) int {
	count := len(strings.Split(content, ""))

	return count
}

//Counts the numbers of charcaters excluding spaces
func CharWithoutSpace(content string) int {

	content2 := strings.Replace(content, " ", "", -1)
	count := len(strings.Split(content2, ""))

	return count
}

//Counts the numbers of sentences
func Sentence(content string) int {
	count := len(strings.Split(content, "."))

	return count
}

//counts the numbers of paragraphs
func Paragraphs(content string) int {
	count := len(strings.Split(content, "\n\n"))

	return count

}

//Counts the numbers of lines
func Line(content string) int {
	lines := strings.SplitAfterN(content, "\n", -1)
	count := len(lines)

	return count
}

// func Pages(content string)int{
// 	count := len(strings.Split(content, ""))
// 	page := 1800
// 	if count >= page{
// 		pa
// 	}
// }
