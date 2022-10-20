package cli

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
	"wordle-helper/helper"
	"wordle-helper/loader"
	"wordle-helper/sort"
)

const (
	DefaultThreshold = 5
)

var commands = []*command{
	{
		[]string{"h", "help"},
		[]string{},
		nil,
		"Prints this help message.",
	},
	{
		[]string{"addGuess", "ag"},
		[]string{"word", "colors"},
		handleAddGuess,
		"Adds a new guess to previous guesses. Colors should be a string such as \"byggb\"",
	},
	{
		[]string{"all"},
		[]string{},
		handleAll,
		"Prints all possible answers.",
	},
	{
		[]string{"clear", "clr", "reset", "rst"},
		[]string{},
		handleReset,
		"Clears the guess list.",
	},
	{
		[]string{"top"},
		[]string{"n"},
		handleTop,
		"Prints the top n possible answers.",
	},
	{
		[]string{"sort"},
		[]string{},
		handleSort,
		"Sorts the possible answers based on their most used letters in alphabet.",
	},
}

func RunCLI() {
	allWords, err := loader.LoadFiveLetterWords()

	if err != nil {
		fmt.Println(err)
		return
	}

	letterScores, err := loader.LoadLetterScores()

	if err != nil {
		fmt.Println(err)
		return
	}

	h := helper.New(allWords, letterScores)

	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')

	for err == nil && command != "end" && command != "" {
		if err := handleCommand(h, strings.Fields(command)); err != nil {
			fmt.Println(err)
			fmt.Println()
		}
		command, err = reader.ReadString('\n')
	}
}

func handleCommand(h *helper.Helper, parts []string) error {
	for _, command := range commands {
		if command.containsTrigger(parts[0]) {
			if err := checkArgsCount(len(command.args), len(parts)-1); err != nil {
				return err
			}
			if command.handler == nil {
				return handleHelp(h, parts)
			}
			return command.handler(h, parts)
		}
	}
	return fmt.Errorf("invalid command")
}

func handleReset(h *helper.Helper, parts []string) error {
	h.Reset()
	fmt.Printf("Reset applied.\n\n")
	return nil
}

func checkArgsCount(needed, provided int) error {
	if provided < needed {
		return fmt.Errorf("needed %d args for this command, got %d", needed, provided)
	}
	return nil
}

func handleTop(h *helper.Helper, parts []string) error {

	threshold, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	printAllAnswers(h, threshold)
	return nil
}

func handleAll(h *helper.Helper, parts []string) error {
	printAllAnswers(h, h.Answers.Len())
	return nil
}

func handleAddGuess(h *helper.Helper, parts []string) error {

	g := &helper.Info{
		Word:   strings.ToLower(parts[1]),
		Colors: strings.ToLower(parts[2]),
	}

	if len(g.Word) != len(g.Colors) {
		return fmt.Errorf("invalid lengths: %v (%d) vs %v (%d)", g.Word, len(g.Word), g.Colors, len(g.Colors))
	}

	h.AddGuess(g)

	fmt.Printf("Guess added.\n\n")
	printAllAnswers(h, DefaultThreshold)

	return nil
}

func printAllAnswers(h *helper.Helper, threshold int) {
	if h.Answers.Len() == 0 {
		fmt.Println("No answers found.")
		return
	}

	fmt.Printf("%d answers found.\n", h.Answers.Len())

	count := 0
	for e := h.Answers.Front(); e != nil && count < threshold; e = e.Next() {
		fmt.Println(e.Value)
		count++
	}

	if threshold < h.Answers.Len() {
		fmt.Println("...")
	}
	fmt.Println()
}

func handleHelp(h *helper.Helper, _ []string) error {
	for _, c := range commands {
		fmt.Println(c)
	}
	fmt.Println()
	return nil
}

func handleSort(h *helper.Helper, _ []string) error {
	h.Answers = sort.Sort(h.Answers, func(a, b *list.Element) bool {
		// return int(a.Value.(string)[4])-int(b.Value.(string)[4]) >= 0
		return h.WordScoreOf(a.Value.(string)) >= h.WordScoreOf(b.Value.(string))
	})
	fmt.Printf("Answers sorted.\n\n")
	return nil
}
