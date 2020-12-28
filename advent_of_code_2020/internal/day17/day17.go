package day17

type cubeLocation struct {
	x int
	y int
	z int
	w int
}

// GetActiveCubes returns the number of active cubes after simulating number of rounds
func GetActiveCubes(input []string, rounds int, activate4D bool) int {
	worldModel := parseInput(input, activate4D)

	for i := 0; i < rounds; i++ {
		worldModel = simulateRound(worldModel, activate4D)
	}

	return countActive(worldModel)
}

func simulateRound(worldModel map[cubeLocation]bool, activate4D bool) map[cubeLocation]bool {
	changeStateQueue := []cubeLocation{}

	for cube, isActive := range worldModel {
		activeNeighbors := countActiveNeighbors(cube, worldModel, activate4D)
		if isActive && !(activeNeighbors == 2 || activeNeighbors == 3) {
			changeStateQueue = append(changeStateQueue, cube)
		} else if !isActive && activeNeighbors == 3 {
			changeStateQueue = append(changeStateQueue, cube)
		}
	}

	for _, cube := range changeStateQueue {
		if worldModel[cube] {
			worldModel[cube] = false
		} else {
			worldModel[cube] = true
			worldModel = addInactiveNeighbors(cube, worldModel, activate4D)
		}
	}

	return worldModel
}

func countActive(worldModel map[cubeLocation]bool) int {
	activecount := 0

	for _, isActive := range worldModel {
		if isActive {
			activecount++
		}
	}

	return activecount
}

func parseInput(input []string, activate4D bool) map[cubeLocation]bool {
	worldModel := map[cubeLocation]bool{}

	// Assume all input is in the z=0 plane
	// Add all active cubes
	for y, line := range input {
		for x, isActive := range line {
			if isActive == '#' {
				worldModel[cubeLocation{x: x, y: y, z: 0, w: 0}] = true
			}
		}
	}

	// Add all neighboring inactive cubes to keep an eye on
	for cubeLoc := range worldModel {
		worldModel = addInactiveNeighbors(cubeLoc, worldModel, activate4D)
	}

	return worldModel
}

func addInactiveNeighbors(cubeLoc cubeLocation, worldModel map[cubeLocation]bool, activate4D bool) map[cubeLocation]bool {
	getNeighbors := getNeighbors3D
	if activate4D {
		getNeighbors = getNeighbors4D
	}

	for _, neighbor := range getNeighbors(cubeLoc) {
		if _, isPresent := worldModel[neighbor]; !isPresent {
			worldModel[neighbor] = false
		}
	}

	return worldModel
}

func getNeighbors3D(cubeLoc cubeLocation) []cubeLocation {
	neighbors := []cubeLocation{}
	for x := cubeLoc.x - 1; x <= cubeLoc.x+1; x++ {
		for y := cubeLoc.y - 1; y <= cubeLoc.y+1; y++ {
			for z := cubeLoc.z - 1; z <= cubeLoc.z+1; z++ {
				possibleNeighbor := cubeLocation{x: x, y: y, z: z}
				if possibleNeighbor != cubeLoc {
					neighbors = append(neighbors, possibleNeighbor)
				}
			}
		}
	}

	return neighbors
}

func getNeighbors4D(cubeLoc cubeLocation) []cubeLocation {
	neighbors := []cubeLocation{}
	for x := cubeLoc.x - 1; x <= cubeLoc.x+1; x++ {
		for y := cubeLoc.y - 1; y <= cubeLoc.y+1; y++ {
			for z := cubeLoc.z - 1; z <= cubeLoc.z+1; z++ {
				for w := cubeLoc.w - 1; w <= cubeLoc.w+1; w++ {
					possibleNeighbor := cubeLocation{x: x, y: y, z: z, w: w}
					if possibleNeighbor != cubeLoc {
						neighbors = append(neighbors, possibleNeighbor)
					}
				}
			}
		}
	}

	return neighbors
}

func countActiveNeighbors(cubeLoc cubeLocation, worldModel map[cubeLocation]bool, activate4D bool) int {
	activeCount := 0
	getNeighbors := getNeighbors3D
	if activate4D {
		getNeighbors = getNeighbors4D
	}

	for _, neighbor := range getNeighbors(cubeLoc) {
		if worldModel[neighbor] {
			activeCount++
		}
	}

	return activeCount
}
