package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibratoev/aoc2024/challenge"
)

func TestA(t *testing.T) {
	input := strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

	result := partA(input)

	require.Equal(t, 161, result)
}

func TestAFull(t *testing.T) {
	input := challenge.InputFile()

	result := partA(input)

	require.Equal(t, 178886550, result)
}
