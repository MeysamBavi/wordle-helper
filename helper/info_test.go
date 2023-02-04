package helper_test

import (
	"github.com/MeysamBavi/wordle-helper/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type getColorsTestCase struct {
	answer         string
	guess          string
	expectedColors string
}

func (g *getColorsTestCase) Test(t *testing.T) {
	assert.Equal(t, g.expectedColors, helper.GetColors(g.answer, g.guess))
}

func TestGetColorsForFiveLetter(t *testing.T) {
	testCases := []*getColorsTestCase{
		{"kazoo", "loose", "byybb"},
		{"unlit", "trace", "ybbbb"},
		{"unlit", "innit", "bgbgg"},
		{"unlit", "unlit", "ggggg"},
		{"unlit", "unlit", "ggggg"},
		{"cramp", "loose", "bbbbb"},
		{"cramp", "chase", "gbgbb"},
		{"cramp", "porge", "ybybb"},
		{"cramp", "fleet", "bbbbb"},
		{"cramp", "scrap", "byyyg"},
		{"cramp", "crapy", "gggyb"},
		{"cramp", "aaxyz", "ybbbb"},
		{"cramp", "cramp", "ggggg"},
	}

	for _, testCase := range testCases {
		testCase.Test(t)
	}
}
