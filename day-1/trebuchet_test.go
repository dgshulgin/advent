package trebuchet

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleStroke(t *testing.T) {
	stroke := "1abc2"
	caliber, err := calibrate(stroke)
	assert.NoError(t, err, "calibration error")

	assert.Equal(t, 12, caliber, "calibration failed")
}

func TestCalibrateTrebuchet(t *testing.T) {

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
}

func TestPuzzleInput(t *testing.T) {

	path, ok := os.LookupEnv("TR_INPUT")
	if !ok {
		t.Fatalf("input data path undefined, declare TR_INPUT")
	}
	file, _ := os.Open(path)

	var strokes []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strokes = append(strokes, scanner.Text())
	}

	file.Close()

	var sum int
	for _, stroke := range strokes {
		actCaliber, err := calibrate(stroke)
		assert.NoError(t, err, fmt.Sprintf("error for stoke '%s'", stroke))
		sum += actCaliber
	}
	fmt.Printf("final caliber is %d\n", sum)
}
