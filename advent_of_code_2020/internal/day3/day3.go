package day3

// Slope the slope to toboggan through the trees with
type Slope struct {
	Right int
	Down  int
}

// TreesHit returns and integer describing trees hit when going on the inputSlope on the inputLandscapes
func TreesHit(inputLandscapes []string, inputSlope Slope) (int, error) {
	treesHit := 0
	RightwardMovement := 0
	DownwardMovement := 0

	inputLandscapes = inputLandscapes[1:]

	for _, landscape := range inputLandscapes {
		DownwardMovement++
		if DownwardMovement%inputSlope.Down != 0 {
			continue
		}

		RightwardMovement += inputSlope.Right
		index := RightwardMovement % len(landscape)
		if landscape[index] == '#' {
			// fmt.Printf("%v - index: %v\n", landscape, index)
			treesHit++
		}
	}
	return treesHit, nil
}
