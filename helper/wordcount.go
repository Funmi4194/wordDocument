package helper

import "strings"

func Content(content string) int {
	word := strings.Join(strings.Fields(content), " ")
	count := len(strings.Split(word, " "))

	return count
}

func TotalCharacters(content string) int {
	count := len(strings.Split(content, ""))

	return count
}

func CharWithoutSpace(content string) int {

	content2 := strings.Replace(content, " ", "", -1)
	count := len(strings.Split(content2, ""))

	return count
}

func Sentence(content string) int {
	count := len(strings.Split(content, "."))

	return count
}

func Paragraphs(content string) int {
	count := len(strings.Split(content, "\n\n"))

	return count
}

func Line(content string) int {
	lines := strings.SplitAfterN(content, "\n", -1)
	count := len(lines)

	return count
}
