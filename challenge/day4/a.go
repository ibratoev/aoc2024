package day4

import (
	"fmt"
	"io"
	"slices"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

//nolint:gocyclo
func partA(input io.Reader) int {
	lines := slices.Collect(challenge.Lines(input))
	sum := 0

	for y, line := range lines {
		for x, char := range line {
			xlen := len(line)
			ylen := len(lines)

			// check all the X-es if they start an "XMAS"
			if char != 'X' {
				continue
			}

			// check horizontally forward
			if x < xlen-3 && line[x+1] == 'M' && line[x+2] == 'A' && line[x+3] == 'S' {
				sum++
			}

			// check horizontally backwards
			if x > 2 && line[x-1] == 'M' && line[x-2] == 'A' && line[x-3] == 'S' {
				sum++
			}

			// check vertical downwards
			if y < ylen-3 && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
				sum++
			}

			// check vertical upwards
			if y > 2 && lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S' {
				sum++
			}

			// check diagonal forward downwards
			if x < xlen-3 && y < ylen-3 && lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
				sum++
			}

			// check diagonal forward upwards
			if x < xlen-3 && y > 2 && lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
				sum++
			}

			// check diagonal backwards downwards
			if x > 2 && y < ylen-3 && lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S' {
				sum++
			}

			// check diagonal backwards upwards
			if x > 2 && y > 2 && lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S' {
				sum++
			}
		}
	}

	return sum
}
