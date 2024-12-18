package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	s := "47|53\n" +
		"97|13\n" +
		"97|61\n" +
		"97|47\n" +
		"75|29\n" +
		"61|13\n" +
		"75|53\n" +
		"29|13\n" +
		"97|29\n" +
		"53|29\n" +
		"61|53\n" +
		"97|53\n" +
		"61|29\n" +
		"47|13\n" +
		"75|47\n" +
		"97|75\n" +
		"47|61\n" +
		"75|61\n" +
		"47|29\n" +
		"75|13\n" +
		"53|13"
	s2 := "75,47,61,53,29\n" +
		"97,61,53,29,13\n" +
		"75,29,13\n" +
		"75,97,47,61,53\n" +
		"61,13,29\n" +
		"97,13,75,29,47"
	inp := parse(s, "|")
	inp2 := parse(s2, ",")
	want := 143
	assert.Equal(t, want, countValid(inp, inp2))
}

func TestFixInvalid(t *testing.T) {
	s := "47|53\n" +
		"97|13\n" +
		"97|61\n" +
		"97|47\n" +
		"75|29\n" +
		"61|13\n" +
		"75|53\n" +
		"29|13\n" +
		"97|29\n" +
		"53|29\n" +
		"61|53\n" +
		"97|53\n" +
		"61|29\n" +
		"47|13\n" +
		"75|47\n" +
		"97|75\n" +
		"47|61\n" +
		"75|61\n" +
		"47|29\n" +
		"75|13\n" +
		"53|13"
	s2 := "75,47,61,53,29\n" +
		"97,61,53,29,13\n" +
		"75,29,13\n" +
		"75,97,47,61,53\n" +
		"61,13,29\n" +
		"97,13,75,29,47"
	inp := parse(s, "|")
	inp2 := parse(s2, ",")
	want := 123
	assert.Equal(t, want, countFixInvalid(inp, inp2))
}

func parse(s, div string) [][]string {
	spl := strings.Split(s, "\n")
	ret := make([][]string, 0)
	for i := 0; i < len(spl); i++ {
		p := strings.Split(spl[i], div)
		ret = append(ret, p)
	}
	return ret
}
