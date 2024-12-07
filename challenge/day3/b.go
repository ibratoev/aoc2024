package day3

import (
	"fmt"
	"io"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

func partB(input io.Reader) int {
	sum := 0
	enabled := true
	mulRegex := regexp.MustCompile(`do\(\)|don't\(\)|mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	for _, res := range mulRegex.FindAllStringSubmatch(challenge.Raw(input), -1) {
		if res[0] == "do()" {
			enabled = true
		} else if res[0] == "don't()" {
			enabled = false
		} else if enabled {
			sum += util.MustAtoI(res[1]) * util.MustAtoI(res[2])
		}

	}
	return sum
}
