package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibratoev/aoc2024/challenge"
)

func TestB(t *testing.T) {
	input := strings.NewReader(testInput)

	result := partB(input)

	require.Equal(t, 9, result)
}

func TestBFull(t *testing.T) {
	input := challenge.InputFile()

	result := partB(input)

	require.Equal(t, 1864, result)
}
