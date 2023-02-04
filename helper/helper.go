package helper

import (
	"container/list"
	"strings"
)

type Helper struct {
	all          []string
	guesses      []*Info
	Answers      *list.List
	letterScores map[byte]float32
}

func New(all []string, letterScores map[byte]float32) *Helper {
	return &Helper{
		all:          all,
		guesses:      nil,
		Answers:      nil,
		letterScores: letterScores,
	}
}

func (h *Helper) Reset() {
	h.guesses = nil
	h.Answers = nil
}

func (h *Helper) AddGuess(guess *Info) {

	h.guesses = append(h.guesses, guess)

	if len(h.guesses) == 1 {
		h.initializeAnswers()
		return
	}

	h.filter()
}

func (h *Helper) getLastInfo() *Info {
	return h.guesses[len(h.guesses)-1]
}

func (h *Helper) initializeAnswers() {

	info := h.getLastInfo()
	h.Answers = list.New()

	for _, word := range h.all {
		if info.Validate(word) {
			h.Answers.PushBack(word)
		}
	}
}

func (h *Helper) filter() {
	info := h.getLastInfo()

	for e := h.Answers.Front(); e != nil; {

		if info.Validate(e.Value.(string)) {
			e = e.Next()
			continue
		}

		current := e
		e = e.Next()
		h.Answers.Remove(current)
	}
}

func (h *Helper) WordScoreOf(word string) float32 {
	word = strings.ToUpper(word)
	var sum float32 = 0
	occurrences := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		occurrences[word[i]] += 1
		coefficient := occurrences[word[i]]
		sum += h.letterScores[word[i]] / (float32(coefficient) * 0.25)
	}
	return sum
}
