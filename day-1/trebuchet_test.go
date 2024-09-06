package trebuchet

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Test utilities */
func TestReverse(t *testing.T) {
	rev := reverse("nine")
	assert.EqualValues(t, "enin", rev)
}

func TestCalibrateTrebuchet(t *testing.T) {
	t.Run("Числа только цифрами", func(t *testing.T) {
		strokes := map[string]int{"1abc2": 12, "pqr3stu8vwx": 38, "a1b2c3d4e5f": 15, "treb7uchet": 77}

		var sum int
		for stroke, expectCaliber := range strokes {
			actCaliber, err := calibrate(stroke)
			assert.NoError(t, err, fmt.Sprintf("error for stoke '%s'", stroke))
			if assert.Equal(t, expectCaliber, actCaliber, fmt.Sprintf("calibration failed for stoke '%s'", stroke)) {
				sum += actCaliber
			}
		}
		assert.Equal(t, 142, sum, "trebuchet calibration failed")
	})

	t.Run("Числа цифрами и прописью", func(t *testing.T) {
		strokes := map[string]int{"two1nine": 29, "eightwothree": 83, "abcone2threexyz": 13, "xtwone3four": 24, "4nineeightseven2": 42, "zoneight234": 14, "7pqrstsixteen": 76}

		var sum int
		for stroke, expectCaliber := range strokes {
			actCaliber, err := calibrate(stroke)
			assert.NoError(t, err, fmt.Sprintf("error for stoke '%s'", stroke))
			if assert.Equal(t, expectCaliber, actCaliber, fmt.Sprintf("calibration failed for stoke '%s'", stroke)) {
				sum += actCaliber
			}
		}
		assert.Equal(t, 281, sum, "trebuchet calibration failed")
	})

	t.Run("Большой тест, данные из input.txt", func(t *testing.T) {
		path, ok := os.LookupEnv("TR_INPUT")
		if !ok {
			t.Fatalf("input data path undefined, declare TR_INPUT")
		}
		file, _ := os.Open(path)
		defer file.Close()

		var strokes []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			strokes = append(strokes, scanner.Text())
		}

		var sum int
		for _, stroke := range strokes {
			actCaliber, err := calibrate(stroke)
			assert.NoError(t, err, fmt.Sprintf("error for stoke '%s'", stroke))
			sum += actCaliber
		}
		fmt.Printf("--- final caliber is %d ----\n", sum)
	})

}
