package bloomclock

import "math"

func Probability(m, k, n float64) (p float64) {

	p = 1.0 - 1.0/m

	p = 1.0 - math.Pow(p, k*n)
	p = math.Pow(p, k)

	return
}
