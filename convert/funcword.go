package convert

import (
	"bufio"
	"os"
)

func GetWord(filename string) ([]string, error) {
	var words = []string{}

	file, err := os.Open(filename)
	if err != nil {
		return words, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sentence := scanner.Text()

		words = append(words, sentence)
	}

	err = file.Close()
	if err != nil {
		return words, err
	}
	if scanner.Err() != nil {
		return words, scanner.Err()
	}

	return words, nil
}
