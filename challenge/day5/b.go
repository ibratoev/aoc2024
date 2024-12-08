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

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.InputFile()))
		},
	}
}

// fixes an incorrect update
func fixUpdate(update []int, rules []int, i int, j int) []int {
	fixed := slices.Clone(update)
	fixed[i] = update[j]
	fixed[j] = update[i]

	if correct, x, z := checkUpdate(fixed, rules); correct {
		return fixed
	} else {
		return fixUpdate(fixed, rules, x, z)
	}
}

func partB(input io.Reader) int {
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

			if s, i, j := checkUpdate(update, rules); s {
				continue
			} else {
				fixed := fixUpdate(update, rules, i, j)
				if len(fixed)%2 == 0 {
					panic(fmt.Sprintf("Unexpected update with even page numbers: %v", fixed))
				}

				sum += fixed[(len(fixed)-1)/2]
			}
		}
	}
	return sum
}
