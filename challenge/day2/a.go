package day2

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

// Checks if a report is safe. If the report is not safe, the index of the first failing item is returned.
func isSafe(report []int) (bool, int) {
	increasing := report[0] < report[1]

	for i := 0; i < len(report)-1; i++ {
		cur := report[i]
		next := report[i+1]
		diff := next - cur
		if increasing {
			if diff < 1 || diff > 3 {
				return false, i
			}
		} else {
			if diff < -3 || diff > -1 {
				return false, i
			}
		}
	}

	return true, len(report) - 1
}

func partA(input io.Reader) int {
	sum := 0

	for l := range challenge.Lines(input) {
		fields := strings.Fields(l)

		report := make([]int, len(fields))

		for i, f := range fields {
			report[i] = util.MustAtoI(f)
		}

		if isSafe, _ := isSafe(report); isSafe {
			sum++
		}
	}

	return sum
}
