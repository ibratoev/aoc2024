package day4

import (
	"strings"
	"testing"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	result := partA(input)

	require.Equal(t, 18, result)
}

func TestAFull(t *testing.T) {
	input := challenge.InputFile()

	result := partA(input)

	require.Equal(t, 2468, result)
}
