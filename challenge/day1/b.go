package day1

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	var left []int
	var rightMap = make(map[int]int)

	for l := range challenge.Lines(input) {
		fields := strings.Fields(l)
		left = append(left, util.MustAtoI(fields[0]))
		rightMap[util.MustAtoI(fields[1])]++
	}

	sum := 0
	for _, val := range left {
		sum += rightMap[val] * val
	}
	return sum
}
