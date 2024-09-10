package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	partnumber "github.com/dshulgin/advent/day-3"
)

func main() {

	path, _ := os.LookupEnv("TR_INPUT")
	file, _ := os.Open(path)

	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	file.Close()

	numbers := partnumber.ParseNumbers(data)
	symbols := partnumber.ParseSymbols(data)

	partnumber.Match(numbers, symbols)

	var sum int
	for _, symbol := range symbols {
		for _, num := range symbol.Numbers {
			n := numbers[num]
			nn, _ := strconv.Atoi(n.Raw)
			sum += nn
		}
	}

	fmt.Printf("Anwer: %d\n", sum)
}
