package day3

import (
	"strings"
	"testing"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

	result := partB(input)

	require.Equal(t, 48, result)
}

func TestBFull(t *testing.T) {
	input := challenge.InputFile()

	result := partB(input)

	require.Equal(t, 87163705, result)
}
