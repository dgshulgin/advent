package trebuchet

import (
	"strings"
)

const (
	longestPattern int = 5 // кол-во символов самой длинной цифры - three, seven, eight
)

// возвращает первое (слева) число представленное цифрой прописью и его позицию во входной строке
func NumberWord(line string) (num int, pos int) {

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var start int
	for {

		r := min(len(line)-start, start+longestPattern)
		if r <= 0 {
			break
		}
		chunk := line[start : start+r]
		for idx, num := range numbers {
			if strings.HasPrefix(chunk, num) {
				return idx + 1, start
			}
		}
		start += 1
	}
	return 0, -1
}

// возвращает последнее (крайнее справа) число представленное цифрой прописью и его позицию во входной строке
// инвертирует сроку и ищет вхождение паттерна с левого края, затем пересчитывает позицию
func LastNumberWord(line string) (num int, pos int) {
	numbers := []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

	rev := reverse(line)
	var start int
	for {

		r := min(len(rev)-start, start+longestPattern)
		if r <= 0 {
			break
		}
		chunk := rev[start : start+r]
		for idx, num := range numbers {
			if strings.HasPrefix(chunk, num) {
				return idx + 1, len(rev) - start - len(num) + 1
			}
		}
		start += 1
	}
	return 0, -1
}

// возвращает инвертированную строку
func reverse(line string) string {
	var rev string
	for _, ch := range line {
		rev = string(ch) + rev
	}
	return rev
}
