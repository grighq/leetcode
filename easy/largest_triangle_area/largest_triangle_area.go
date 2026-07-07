package largesttrianglearea

import "math"

func largestTriangleArea(points [][]int) float64 {
	res := 0.0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				p1, p2, p3 := points[i], points[j], points[k]

				x1, x2, x3 := float64(p1[0]), float64(p2[0]), float64(p3[0])
				y1, y2, y3 := float64(p1[1]), float64(p2[1]), float64(p3[1])

				area := 0.5 * math.Abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))
				res = math.Max(res, area)
			}
		}
	}

	return res
}
