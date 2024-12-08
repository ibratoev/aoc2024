package day5

import (
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ibratoev/aoc2024/challenge"
	"github.com/ibratoev/aoc2024/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

const multiplier = 1000

// checks if the update is correct; if not correct returns the indeces of the pair that breaks the update
func checkUpdate(update []int, rules []int) (bool, int, int) {
	for i, u := range update {
		for j := 0; j < i; j++ {
			if slices.Contains(rules, u*multiplier+update[j]) {
				return false, i, j
			}
		}
	}

	return true, -1, -1
}

func partA(input io.Reader) int {
	rules := make([]int, 0, 1200)

	sum := 0
	gatherRules := true

	for l := range challenge.Lines(input) {
		if l == "" {
			gatherRules = false
			continue
		}

		if gatherRules {
			rule := strings.Split(l, `|`)
			if len(rule) != 2 {
				panic(fmt.Sprintf(`Unexpected rule: %v`, rule))
			}

			rules = append(rules, util.MustAtoI(rule[0])*multiplier+util.MustAtoI(rule[1]))
		} else { // process updates
			updateStr := strings.Split(l, `,`)
			update := make([]int, len(updateStr))

			for i, u := range updateStr {
				update[i] = util.MustAtoI(u)
			}

			if s, _, _ := checkUpdate(update, rules); s {
				if len(update)%2 == 0 {
					panic(fmt.Sprintf("Unexpected update with even page numbers: %v", update))
				}

				sum += update[(len(update)-1)/2]
			}
		}
	}
	return sum
}
