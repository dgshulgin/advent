package game

import (
	"strconv"
	"strings"
)

type Set struct {
	red   int
	green int
	blue  int
}

// 3 green, 4 blue, 1 red
func NewSetFromString(input string) Set {
	set := Set{}

	fields := splitAtDelim(input, ',')
	for _, field := range fields {

		fdata := splitAtDelim(field, ' ')
		title := strings.TrimSpace(fdata[1])
		num, _ := strconv.Atoi(fdata[0])

		switch title {
		case "red":
			set.red = num
		case "green":
			set.green = num
		case "blue":
			set.blue = num
		}
	}
	return set
}

type Game struct {
	title  string
	party  []Set
	limits Set
}

// input form: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
func NewGameFromString(input string, limits Set) Game {

	gg := Game{}

	game := splitAtDelim(input, ':')
	gg.title = strings.TrimSpace(game[0])

	sets := splitAtDelim(game[1], ';')
	for _, set := range sets {
		gg.party = append(gg.party, NewSetFromString(set))
	}

	gg.limits = limits

	return gg
}

func splitAtDelim(input string, delim rune) []string {
	input = strings.TrimSpace(input)
	return strings.FieldsFunc(input, func(rr rune) bool {
		return rr == delim
	})
}

func Eval(game Game) bool {

	for _, pp := range game.party {
		if pp.red > game.limits.red || pp.green > game.limits.green || pp.blue > game.limits.blue {
			return false
		}
	}
	return true
}
