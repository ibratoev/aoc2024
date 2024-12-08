package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibratoev/aoc2024/challenge"
)

const testData = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestA(t *testing.T) {
	input := strings.NewReader(testData)

	result := partA(input)

	require.Equal(t, 143, result)
}

func TestAFull(t *testing.T) {
	input := challenge.InputFile()

	result := partA(input)

	require.Equal(t, 4924, result)
}
