package trebuchet

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

// Просматривает строку слева направо и возвращает первую найденную цифру num и ее позицию pos в строке line.
// Если в строке отсутствуют цифры, либо произошла ошибка в качетстве позиции возвращает pos=-1.
func number(line string) (num int, pos int) {
	pos = 0
	for {
		rr, w := utf8.DecodeRuneInString(line[pos:])
		// строка пуста, либо в ней присутствует ошибка кодировки
		// как бы то нибыло, цифра не найдена
		if rr == utf8.RuneError {
			return 0, -1
		}
		if unicode.IsDigit(rr) {
			num, _ := strconv.Atoi(string(rr))
			return num, pos
		}
		pos += w
	}
}

// Просматривает строку справа налево и возвращает первую найденную цифру num и ее позицию pos в строке line.
// Если в строке отсутствуют цифры, либо произошла ошибка в качетстве позиции возвращает pos=-1.
func numberLast(line string) (num int, pos int) {
	pos = len(line)
	for {
		rr, w := utf8.DecodeLastRuneInString(line[:pos])
		// строка пуста, либо в ней присутствует ошибка кодировки
		// как бы то нибыло, цифра не найдена
		if rr == utf8.RuneError {
			return 0, -1
		}
		if unicode.IsDigit(rr) {
			num, _ := strconv.Atoi(string(rr))
			return num, pos
		}
		pos -= w
	}
}
