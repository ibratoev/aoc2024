package day2

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func isSafeB(report []int) bool {
	safe, i := isSafe(report)

	if safe {
		return true
	}

	if i > 0 {
		reportPrev := slices.Clone(report)
		if safe, _ := isSafe(slices.Delete(reportPrev, i-1, i)); safe {
			fmt.Println(safe)
			fmt.Println(reportPrev)
			return true
		}
	}

	reportI := slices.Clone(report)
	if safe, _ := isSafe(slices.Delete(reportI, i, i+1)); safe {
		fmt.Println(safe)
		fmt.Println(reportI)
		return true
	}

	reportNext := slices.Clone(report)
	if safe, _ := isSafe(slices.Delete(reportNext, i+1, i+2)); safe {
		fmt.Println(safe)
		fmt.Println(reportNext)
		return true
	}

	return false
}

func partB(input io.Reader) int {
	sum := 0

	for l := range challenge.Lines(input) {
		fields := strings.Fields(l)

		report := make([]int, len(fields))

		for i, f := range fields {
			report[i] = util.MustAtoI(f)
		}

		if isSafeB(report) {
			sum++
		}
	}

	return sum
}
