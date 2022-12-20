package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Choice uint

const (
	Rock Choice = iota
	Paper
	Scissors
)

type RPSRound struct {
	theirs Choice
	mine   Choice
}

func (r *RPSRound) getScore() uint {
	u := r.getMyChoiceScore()
	switch r.theirs {
	case Rock:
		switch r.mine {
		case Rock:
			return u + 3
		case Paper:
			return u + 6
		case Scissors:
			return u
		}
	case Paper:
		switch r.mine {
		case Rock:
			return u
		case Paper:
			return u + 3
		case Scissors:
			return u + 6
		}
	case Scissors:
		switch r.mine {
		case Rock:
			return u + 6
		case Paper:
			return u
		case Scissors:
			return u + 3
		}
	}
	return 0
}

func (r *RPSRound) getMyChoiceScore() uint {
	switch r.mine {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic(fmt.Sprintf("unexpected choice r.mine=%v", r.mine))
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	scanner := bufio.NewScanner(f)

	RPSGame := part2(scanner)
	println(evaluateScore(RPSGame))
}

func part1(scanner *bufio.Scanner) []RPSRound {
	RPSGame := make([]RPSRound, 0)

	for scanner.Scan() {
		actions := scanner.Text()
		split := strings.Split(actions, " ")
		var theirs Choice
		var mine Choice
		switch split[0] {
		case "A":
			theirs = Rock
		case "B":
			theirs = Paper
		case "C":
			theirs = Scissors
		default:
			panic(fmt.Sprintf("Unexpected value %s", split[0]))
		}
		switch split[1] {
		case "X":
			mine = Rock
		case "Y":
			mine = Paper
		case "Z":
			mine = Scissors
		default:
			panic(fmt.Sprintf("Unexpected value %s", split[0]))
		}
		RPSGame = append(RPSGame, RPSRound{
			theirs: theirs,
			mine:   mine,
		})
	}
	return RPSGame
}

func part2(scanner *bufio.Scanner) []RPSRound {
	RPSGame := make([]RPSRound, 0)

	for scanner.Scan() {
		actions := scanner.Text()
		split := strings.Split(actions, " ")
		var theirs Choice
		var mine Choice
		switch split[0] {
		case "A":
			theirs = Rock
		case "B":
			theirs = Paper
		case "C":
			theirs = Scissors
		default:
			panic(fmt.Sprintf("Unexpected value %s", split[0]))
		}
		switch split[1] {
		case "X":
			switch theirs {
			case Rock:
				mine = Scissors
			case Paper:
				mine = Rock
			case Scissors:
				mine = Paper
			}
		case "Y":
			mine = theirs
		case "Z":
			switch theirs {
			case Rock:
				mine = Paper
			case Paper:
				mine = Scissors
			case Scissors:
				mine = Rock
			}
		default:
			panic(fmt.Sprintf("Unexpected value %s", split[0]))
		}
		RPSGame = append(RPSGame, RPSRound{
			theirs: theirs,
			mine:   mine,
		})
	}
	return RPSGame
}

func evaluateScore(games []RPSRound) uint {
	score := uint(0)
	for _, round := range games {
		score += round.getScore()
	}
	return score
}
