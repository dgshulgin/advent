package trebuchet

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func number(line string) (int, error) {
	start := 0
	for {
		rr, w := utf8.DecodeRuneInString(line[start:])
		if rr == utf8.RuneError && w == 0 {
			return 0, errors.New("в калибровочной строке отсутствуют цифры")
		}
		if rr == utf8.RuneError && w == 1 {
			return 0, errors.New("ошибка кодировки в калибровочной строке")
		}
		if unicode.IsDigit(rr) {
			num, _ := strconv.Atoi(string(rr))
			return num, nil
		}
		start += w
	}
}

func numberLast(line string) (int, error) {
	last := len(line)
	for {
		rr, w := utf8.DecodeLastRuneInString(line[:last])
		if rr == utf8.RuneError && w == 0 {
			return 0, errors.New("в калибровочной строке отсутствуют цифры")
		}
		if rr == utf8.RuneError && w == 1 {
			return 0, errors.New("ошибка кодировки в калибровочной строке")
		}
		if unicode.IsDigit(rr) {
			num, _ := strconv.Atoi(string(rr))
			return num, nil
		}
		last -= w
	}
}
