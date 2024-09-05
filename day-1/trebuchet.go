package trebuchet

import (
	"errors"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// 1abc2
func calibrate(stroke string) (int, error) {

	if len(stroke) == 0 {
		return 0, errors.New("calibration string is empty")
	}

	var firstDigit, lastDigit rune

	start := 0
	for {
		rr, w := utf8.DecodeRuneInString(stroke[start:])
		if rr == utf8.RuneError && w == 0 {
			return 0, errors.New("no digits in calibration string")
		}
		if rr == utf8.RuneError && w == 1 {
			return 0, errors.New("calibration string is invalid")
		}
		if unicode.IsDigit(rr) {
			firstDigit = rr
			break
		}
		start += w
	}

	last := len(stroke)
	for {
		rr, w := utf8.DecodeLastRuneInString(stroke[:last])
		if rr == utf8.RuneError && w == 0 {
			return 0, errors.New("no digits in calibration string")
		}
		if rr == utf8.RuneError && w == 1 {
			return 0, errors.New("calibration string is invalid")
		}
		if unicode.IsDigit(rr) {
			lastDigit = rr
			break
		}
		last -= w
	}

	val, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
	if err != nil {
		return 0, errors.New("error while converting digits in calibration string")
	}

	return val, nil
}
