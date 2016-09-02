// Package holtwinters implements simple Holt-Winters forecasting
/*
Translated from the python code at

        http://grisha.org/blog/2016/02/17/triple-exponential-smoothing-forecasting-part-iii/
*/
package holtwinters

func initialTrend(series []float64, slen int) float64 {
	var sum float64
	slenf := float64(slen)
	sum = 0.0
	for i := 0; i < slen; i++ {
		sum += (series[i+slen] - series[i]) / slenf
	}

	return sum / slenf
}

func initialSeasonalComponents(series []float64, slen int) []float64 {
	seasonals := make([]float64, slen)
	seasonAverages := make([]float64, slen)

	nSeasons := int(len(series) / slen)
	// compute season averages
	for j := 0; j < nSeasons; j++ {
		seasonAverages[j] = fsum(series[slen*j:slen*j+slen]) / float64(slen)
	}
	// compute initial values
	for i := 0; i < slen; i++ {
		var sum float64
		for j := 0; j < nSeasons; j++ {
			sum += series[slen*j+i] - seasonAverages[j]
			seasonals[i] = sum / float64(nSeasons)
		}
	}
	return seasonals
}

func fsum(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	return sum
}

func TripleExponentialSmoothing(series []float64, slen int, alpha, beta, gamma float64, nPredictions int) []float64 {
	var result []float64
	seasonals := initialSeasonalComponents(series, slen)

	var trend float64
	var smooth, lastSmooth float64

	for i := 0; i < len(series)+nPredictions; i++ {
		if i == 0 {
			// initial values
			smooth = series[0]
			trend = initialTrend(series, slen)
			result = append(result, series[0])
			continue
		}

		if i >= len(series) {
			// we are forecasting
			m := float64(i - len(series) + 1)
			result = append(result, (smooth+m*trend)+seasonals[i%slen])
			continue
		}

		val := series[i]
		lastSmooth, smooth = smooth, alpha*(val-seasonals[i%slen])+(1-alpha)*(smooth+trend)
		trend = beta*(smooth-lastSmooth) + (1-beta)*trend
		seasonals[i%slen] = gamma*(val-smooth) + (1-gamma)*seasonals[i%slen]
		result = append(result, smooth+trend+seasonals[i%slen])
	}
	return result
}
