package helper

const (
	GREEN  = 'g'
	YELLOW = 'y'
	BLACK  = 'b'
)

type Info struct {
	Word   string
	Colors string
}

func (g *Info) Validate(word string) bool {
	return g.Colors == GetColors(word, g.Word)
}

func GetColors(answer, guess string) string {
	colors := make([]byte, 5)
	answerChars := []byte(answer)

	for i := range guess {
		if guess[i] == answer[i] {
			colors[i] = GREEN
			answerChars[i] = 0
		}
	}

	for i := range guess {
		if colors[i] == GREEN {
			continue
		}

		for j := range answerChars {
			if answerChars[j] == guess[i] {
				colors[i] = YELLOW
				answerChars[j] = 0
				break
			}
		}
	}

	for i, c := range colors {
		if c != GREEN && c != YELLOW {
			colors[i] = BLACK
		}
	}

	return string(colors)
}
