package loader

import (
	"bytes"
	_ "embed"
	"strconv"
)

const (
	FIVE         = "./5words.txt"
	LetterScores = "./letterFrequency.txt"
)

//go:embed 5words.txt
var fiveLetterWordsFile []byte

//go:embed letterFrequency.txt
var letterFrequencyFile []byte

func LoadFiveLetterWords() ([]string, error) {
	return bytesToStrings(splitLines(fiveLetterWordsFile)), nil
}

func LoadLetterScores() (map[byte]float32, error) {
	sep := []byte(" ")
	lines := splitLines(letterFrequencyFile)
	result := make(map[byte]float32)
	for _, line := range lines {
		parts := bytes.Split(line, sep)
		freq, err := strconv.ParseFloat(string(parts[1]), 32)
		if err != nil {
			return nil, err
		}
		result[parts[0][0]] = float32(freq)
	}

	return result, nil
}

func splitLines(file []byte) [][]byte {
	noCR := bytes.ReplaceAll(file, []byte("\r\n"), []byte("\n"))
	lines := bytes.Split(noCR, []byte("\n"))
	return lines
}

func bytesToStrings(lines [][]byte) []string {
	result := make([]string, 0, len(lines))
	for _, l := range lines {
		s := string(l)
		if s == "" {
			continue
		}
		result = append(result, s)
	}
	return result
}
