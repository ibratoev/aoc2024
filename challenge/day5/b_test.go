package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibratoev/aoc2024/challenge"
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
