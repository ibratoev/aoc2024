package day6

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
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.InputFile()))
		},
	}
}

type Direction rune
type Puzzle [][]rune

func (p Puzzle) at(pos Position) rune {
	return p[pos.y][pos.x]
}

func (p Puzzle) set(pos Position, r rune) {
	p[pos.y][pos.x] = r
}

const (
	Up    Direction = '^'
	Right           = '>'
	Down            = 'v'
	Left            = '<'
)

const Obstruction = '#'

func changeDirection(d Direction) Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("Unexpected direction.")
	}
}

type Position struct {
	x int
	y int
}

func calculatePath(puzzle Puzzle, pos Position, visited map[Position]bool) int {
	lenY := len(puzzle)
	lenX := len(puzzle[0])
	direction := Direction(puzzle.at(pos))

	visited[pos] = true

	var nextPos Position = pos
	switch direction {
	case Up:
		nextPos.y--
	case Right:
		nextPos.x++
	case Down:
		nextPos.y++
	case Left:
		nextPos.x--
	default:
		panic("Unexpected direction.")
	}

	if nextPos.x < 0 || nextPos.x == lenX || nextPos.y < 0 || nextPos.y == lenY {
		return len(visited)
	}

	if puzzle.at(nextPos) == Obstruction {
		puzzle.set(pos, rune(changeDirection(direction)))
		return calculatePath(puzzle, pos, visited)
	}

	puzzle.set(pos, '.')
	puzzle.set(nextPos, rune(direction))
	return calculatePath(puzzle, nextPos, visited)
}

func partA(input io.Reader) int {
	puzzle := make([][]rune, 0, 150)
	var position Position

	y := 0
	for l := range challenge.Lines(input) {
		if l == "" {
			break
		}

		r := []rune(l)
		puzzle = append(puzzle, r)
		if x := slices.Index(r, rune(Up)); x != -1 {
			position = Position{x, y}
		}

		y++
	}

	return calculatePath(puzzle, position, make(map[Position]bool))
}
