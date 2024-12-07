package day4

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	lines := slices.Collect(challenge.Lines(input))
	sum := 0

	for y, line := range lines {
		for x, char := range line {
			xlen := len(line)
			ylen := len(lines)

			// check all the X-es if they start an "XMAS"
			if char != 'A' || x == 0 || y == 0 || x == xlen-1 || y == ylen-1 {
				continue
			}

			if (lines[y-1][x-1] == 'M' && lines[y+1][x+1] == 'S' || lines[y-1][x-1] == 'S' && lines[y+1][x+1] == 'M') &&
				(lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S' || lines[y+1][x-1] == 'S' && lines[y-1][x+1] == 'M') {
				sum++
			}
		}
	}

	return sum
}
