package day3

import (
	"fmt"
	"io"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

func partA(input io.Reader) int {
	sum := 0
	mulRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	for _, res := range mulRegex.FindAllStringSubmatch(challenge.Raw(input), -1) {
		sum += util.MustAtoI(res[1]) * util.MustAtoI(res[2])
	}
	return sum
}
