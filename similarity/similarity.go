package similarity

import (
	"math"
)

func Cosine(v1, v2 map[string]float64) float64 {
	var top, b1, b2 float64 = 0.0, 0.0, 0.0
	for key, val1 := range v1 {
		if val2, ok := (v2)[key]; ok {
			top += val1 * val2
		}
		b1 += val1 * val1
	}

	for _, val2 := range v2 {
		b2 += val2 * val2
	}

	if b1 == 0 || b2 == 0 {
		return 0.0
	}

	return top / (math.Sqrt(b1) * math.Sqrt(b2))
}

func IntCosine(v1, v2 map[string]int) float64 {
	var top, b1, b2 float64 = 0.0, 0.0, 0.0
	for key, val1 := range v1 {
		if val2, ok := (v2)[key]; ok {
			top += float64(val1) * float64(val2)
		}
		b1 += float64(val1) * float64(val1)
	}

	for _, val2 := range v2 {
		b2 += float64(val2) * float64(val2)
	}

	if b1 == 0 || b2 == 0 {
		return 0.0
	}

	return top / (math.Sqrt(b1) * math.Sqrt(b2))
}
