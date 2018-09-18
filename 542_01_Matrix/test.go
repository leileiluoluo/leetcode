func updateMatrix(m [][]int) [][]int {
	const MaxInt = 999999999

	max := func(n1, n2 int) int {
		if n1 >= n2 {
			return n1
		}
		return n2
	}

	steps := func(x1, y1, x2, y2 int) int {
		abs := func(a, b int) int {
			c := a - b
			if c < 0 {
				return -c
			}
			return c
		}
		return abs(x1, x2) + abs(y1, y2)
	}

	containsZero := func(sumGraph [][]int, x, y, radius int) bool {
		xMin, yMin, xMax, yMax := x-radius-1, y-radius-1, x+radius, y+radius
		if xMax > len(sumGraph)-1 {
			xMax = len(sumGraph) - 1
		}
		if yMax > len(sumGraph[0])-1 {
			yMax = len(sumGraph[0]) - 1
		}

		sumAllNotZeros := MaxInt
		sum := sumAllNotZeros
		if xMin < 0 && yMin < 0 {
			sumAllNotZeros = (xMax + 1) * (yMax + 1)
			sum = sumGraph[xMax][yMax]
		} else if xMin < 0 && yMin >= 0 {
			sumAllNotZeros = (xMax + 1) * (yMax - yMin)
			sum = sumGraph[xMax][yMax] - sumGraph[xMax][yMin]
		} else if xMin >= 0 && yMin < 0 {
			sumAllNotZeros = (xMax - xMin) * (yMax + 1)
			sum = sumGraph[xMax][yMax] - sumGraph[xMin][yMax]
		} else {
			sumAllNotZeros = (xMax - xMin) * (yMax - yMin)
			sum = sumGraph[xMax][yMax] - sumGraph[xMin][yMax] - sumGraph[xMax][yMin] + sumGraph[xMin][yMin]
		}
		return sum < sumAllNotZeros
	}

	findMinStepsAround := func(m [][]int, x, y, radius int) int {
		xMin, yMin, xMax, yMax := x-radius, y-radius, x+radius, y+radius
		if xMin < 0 {
			xMin = 0
		}
		if yMin < 0 {
			yMin = 0
		}
		if xMax > len(m)-1 {
			xMax = len(m) - 1
		}
		if yMax > len(m[0])-1 {
			yMax = len(m[0]) - 1
		}
		minSteps := MaxInt
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				if 0 == m[i][j] {
					steps := steps(x, y, i, j)
					if steps < minSteps {
						minSteps = steps
					}
				}
			}
		}
		return minSteps
	}

	findMinStepsWithScalingRadius := func(m, sumGraph [][]int, x, y int) (int, int) {
		minRadius, maxRadius := 1, max(max(x, y), max(len(m)-1-x, len(m)-1-y))
		minSteps := MaxInt
		r := minRadius
		for r = minRadius; r <= maxRadius; r *= 2 {
			containsZero := containsZero(sumGraph, x, y, r)
			if containsZero {
				minSteps = findMinStepsAround(m, x, y, r)
				if minSteps > r {
					continue
				}
				return minSteps, r
			}
		}
		if maxRadius != r {
			minSteps = findMinStepsAround(m, x, y, maxRadius)
		}
		return minSteps, r
	}

	sumGraph := func(m [][]int) [][]int {
		sumGraph := make([][]int, len(m))
		for i := range m {
			sumGraph[i] = make([]int, len(m[i]))
			for j := range m[i] {
				sumGraph[i][j] = 0
				if 0 == j && 0 == i {
					sumGraph[i][j] = m[i][j]
				} else if 0 == j && i > 0 {
					sumGraph[i][j] = sumGraph[i-1][j] + m[i][j]
				} else if j > 0 && 0 == i {
					sumGraph[i][j] = sumGraph[i][j-1] + m[i][j]
				} else {
					sumGraph[i][j] = sumGraph[i][j-1] + sumGraph[i-1][j] - sumGraph[i-1][j-1] + m[i][j]
				}
			}
		}
		return sumGraph
	}

	sumG := sumGraph(m)
	updated := make([][]int, len(m))
	for i := range m {
		updated[i] = make([]int, len(m[i]))
		for j := range m[i] {
			minSteps, _ := findMinStepsWithScalingRadius(m, sumG, i, j)
			updated[i][j] = minSteps
		}
	}
	return updated
}