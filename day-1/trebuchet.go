package trebuchet

import (
	"errors"
)

// 1abc2
func calibrate(stroke string) (int, error) {

	if len(stroke) == 0 {
		return 0, errors.New("калибровочная строка пуста")
	}

	snum := startNumber(stroke)
	if snum == -1 {
		return 0, errors.New("калибровочная строка не содержит цифр")
	}
	lnum := endNumber(stroke)

	return snum*10 + lnum, nil
}

// возвращает первую цифру последовательности, либо -1 если последовательность не содержит цифры
func startNumber(line string) int {
	nnum, npos := number(line)     // цифра числом
	wnum, wpos := NumberWord(line) // цифра прописью

	if wpos == -1 && npos == -1 {
		return -1 // цифры не найдены в строке
	}
	if wpos == -1 {
		return nnum
	}
	if npos == -1 {
		return wnum
	}
	if wpos < npos {
		return wnum
	}

	return nnum
}

// возвращает последнюю (крайнюю справа) цифру последовательности
func endNumber(line string) int {
	nnum, npos := numberLast(line)     // цифра числом
	wnum, wpos := LastNumberWord(line) // цифра прописью

	if wpos == -1 {
		return nnum
	}
	if npos == -1 {
		return wnum
	}
	if wpos > npos {
		return wnum
	}

	return nnum
}
