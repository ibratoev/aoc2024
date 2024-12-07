package day1

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
	"github.com/ibratoev/aoc2024/util/gmath"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	var first []int
	var second []int

	for l := range challenge.Lines(input) {
		fields := strings.Fields(l)
		first = append(first, util.MustAtoI(fields[0]))
		second = append(second, util.MustAtoI(fields[1]))
	}

	slices.Sort(first)
	slices.Sort(second)

	sum := 0
	for i, val1 := range first {
		val2 := second[i]
		sum += gmath.Abs(val1 - val2)
	}

	return sum
}
