package day5

import (
	"strings"
	"testing"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := strings.NewReader(testData)

	result := partB(input)

	require.Equal(t, 123, result)
}

func TestBFull(t *testing.T) {
	input := challenge.InputFile()

	result := partB(input)

	require.Equal(t, 6085, result)
}
