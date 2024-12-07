package day2

import (
	"strings"
	"testing"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := strings.NewReader(testInput)

	result := partB(input)

	require.Equal(t, 4, result)
}

func TestBFull(t *testing.T) {
	input := challenge.InputFile()

	result := partB(input)

	require.Equal(t, 604, result)
}
