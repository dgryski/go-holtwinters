package holtwinters

import (
	"testing"
)

func TestHoltWinters(t *testing.T) {

	// test data from the blog post

	var series = []float64{
		30, 21, 29, 31, 40, 48, 53, 47, 37, 39, 31, 29, 17, 9, 20, 24, 27, 35, 41, 38,
		27, 31, 27, 26, 21, 13, 21, 18, 33, 35, 40, 36, 22, 24, 21, 20, 17, 14, 17, 19,
		26, 29, 40, 31, 20, 24, 18, 26, 17, 9, 17, 21, 28, 32, 46, 33, 23, 28, 22, 27,
		18, 8, 17, 21, 31, 34, 44, 38, 31, 30, 26, 32,
	}

	t.Log(initialTrend(series, 12))
	t.Log(initialSeasonalComponents(series, 12))
	forecast := TripleExponentialSmoothing(series, 12, 0.716, 0.029, 0.993, 24)
	t.Log(forecast[len(forecast)-24:])
}
