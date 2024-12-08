package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibratoev/aoc2024/challenge"
)

const testInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestA(t *testing.T) {
	input := strings.NewReader(testInput)

	result := partA(input)

	require.Equal(t, 18, result)
}

func TestAFull(t *testing.T) {
	input := challenge.InputFile()

	result := partA(input)

	require.Equal(t, 2468, result)
}
