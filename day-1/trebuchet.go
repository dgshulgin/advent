package trebuchet

import (
	"errors"
)

// 1abc2
func calibrate(stroke string) (int, error) {

	if len(stroke) == 0 {
		return 0, errors.New("калибровочная строка пуста")
	}

	first, err := number(stroke)
	if err != nil {
		return 0, err
	}

	last, err := numberLast(stroke)
	if err != nil {
		return 0, err
	}

	return first*10 + last, nil
}
