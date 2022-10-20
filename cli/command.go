package cli

import (
	"strings"
	"wordle-helper/helper"
)

type command struct {
	triggers []string
	args     []string
	handler  func(*helper.Helper, []string) error
	about    string
}

func (c *command) String() string {
	var b strings.Builder
	b.Grow(16)

	for i, t := range c.triggers {
		if i != 0 {
			b.WriteString(" | ")
		}
		b.WriteString(t)
	}

	for _, a := range c.args {
		b.WriteByte(' ')
		b.WriteString("[" + a + "]")
	}

	b.WriteString(" // ")
	b.WriteString(c.about)

	return b.String()

}

func (c *command) containsTrigger(trigger string) bool {
	for _, t := range c.triggers {
		if t == trigger {
			return true
		}
	}

	return false
}
