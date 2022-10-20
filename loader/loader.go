package loader

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FIVE          = "./5words.txt"
	LETTER_SCORES = "./letterFrequency.txt"
)

func LoadFiveLetterWords() ([]string, error) {
	file, err := os.Open(FIVE)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	result := make([]string, 0, 16)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func LoadLetterScores() (map[byte]float32, error) {
	file, err := os.Open(LETTER_SCORES)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	result := make(map[byte]float32)
	for scanner.Scan() {
		var letter string
		var score float32

		count, err := fmt.Sscan(scanner.Text(), &letter, &score)

		if count != 2 {
			return nil, err
		}

		result[letter[0]] = score
	}

	return result, nil
}
