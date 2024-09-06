package game

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {

	/*
		Game 1: Set{blue:3, red: 4}, Set{red:1, green:2, blue:6}, Set{ green:2}
		Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

		В приведенном выше примере игры 1, 2 и 5 были бы возможны, если бы сумка была загружена
		 с такой конфигурацией. Тем не менее, игра 3 была бы невозможна, потому что в какой-то
		 момент эльф показал вам сразу 20 красных кубиков; точно так же игра 4 также была бы
		 невозможна, потому что эльф показал вам сразу 15 синих кубиков. Если сложить
		 идентификаторы игр, которые были бы возможны, то получится .8
	*/

	t.Run("пробная игра", func(t *testing.T) {
		limits := Set{red: 12, green: 13, blue: 14}

		party1 := []Set{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}
		party2 := []Set{{blue: 1, green: 2}, {green: 3, blue: 4, red: 1}, {green: 1, blue: 1}}
		party3 := []Set{{green: 8, blue: 6, red: 20}, {blue: 5, red: 4, green: 13}, {green: 5, red: 1}}
		party4 := []Set{{green: 1, red: 3, blue: 6}, {green: 3, red: 6}, {green: 3, blue: 15, red: 14}}
		party5 := []Set{{red: 6, blue: 1, green: 3}, {blue: 2, red: 1, green: 2}}

		games := []Game{
			{title: "Game 1", party: party1, limits: limits},
			{title: "Game 2", party: party2, limits: limits},
			{title: "Game 3", party: party3, limits: limits},
			{title: "Game 4", party: party4, limits: limits},
			{title: "Game 5", party: party5, limits: limits},
		}

		var sum int
		for idx, gg := range games {
			if Eval(gg) {
				sum += (idx + 1)
			}
		}
		assert.EqualValues(t, 8, sum)
	})

	t.Run("Большой тест, данные из input.txt", func(t *testing.T) {
		path, ok := os.LookupEnv("TR_INPUT")
		if !ok {
			t.Fatalf("input data path undefined, declare TR_INPUT")
		}
		file, _ := os.Open(path)
		defer file.Close()

		var gamesInput []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			gamesInput = append(gamesInput, scanner.Text())
		}

		limits := Set{red: 12, green: 13, blue: 14}

		var sum int
		for idx, gi := range gamesInput {
			game := NewGameFromString(gi, limits)
			if Eval(game) {
				sum += (idx + 1)
			}
		}
		fmt.Printf("--- Answer is %d ---\n", sum)
	})
}
