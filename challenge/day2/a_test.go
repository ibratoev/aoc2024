package day2

import (
	"strings"
	"testing"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/stretchr/testify/require"
)

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestA(t *testing.T) {
	input := strings.NewReader(testInput)

	result := partA(input)

	require.Equal(t, 2, result)
}

func TestAFull(t *testing.T) {
	input := challenge.InputFile()

	result := partA(input)

	require.Equal(t, 564, result)
}
