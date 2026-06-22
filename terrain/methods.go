package terrain

import "math"

func (h1 Heightmap) Multiply(h2 Heightmap) Heightmap {
	if len(h1) == len(h2) && len(h1[0]) == len(h2[0]) {
		for x := range h1 {
			for y := range h1[x] {
				h1[x][y] = h1[x][y] * h2[x][y]
			}
		}
	}
	return h1
}

func (h1 Heightmap) Round(s int) Heightmap {
	for x := range h1 {
		for y := range h1[x] {
			if s < 2 {
				h1[x][y] = 0.0
				continue
			}
			factor := float64(s - 1)
			rounded := math.Round(h1[x][y] * factor)
			result := rounded / factor

			if result < 0 {
				result = 0
			} else if result > 1 {
				result = 1
			}
			h1[x][y] = result
		}
	}
	return h1
}