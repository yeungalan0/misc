package day12

import (
	"log"
	"math"
	"strconv"
)

const east = "E"
const south = "S"
const west = "W"
const north = "N"

type action struct {
	action string
	value  int
}

type waypoint struct {
	northSouthDistance int
	eastWestDistance   int
}

// TODO: Need to refactor this, there's gotta be an easier way to determine the direction...

// DetermineManhattanDistance determines the Manhattan distance after following a given instruction set
func DetermineManhattanDistance(input []string) int {
	northSouthDistance := 0
	eastWestDistance := 0
	direction := east

	actions := convertToActionsSlice(input)

	for _, ferryAction := range actions {
		if ferryAction.action == "F" {
			ferryAction.action = direction
		}
		if ferryAction.action == "N" {
			northSouthDistance += ferryAction.value
		} else if ferryAction.action == "E" {
			eastWestDistance += ferryAction.value
		} else if ferryAction.action == "S" {
			northSouthDistance -= ferryAction.value
		} else if ferryAction.action == "W" {
			eastWestDistance -= ferryAction.value
		} else if ferryAction.action == "R" {
			direction = determineDirection(direction, ferryAction.value)
		} else if ferryAction.action == "L" {
			direction = determineDirection(direction, -ferryAction.value)
		}
	}

	return int(math.Abs(float64(northSouthDistance)) + math.Abs(float64(eastWestDistance)))
}

// DetermineManhattanDistance2 determines the Manhattan distance after following a given instruction set with a waypoint
func DetermineManhattanDistance2(input []string) int {
	currentWaypoint := waypoint{northSouthDistance: 1, eastWestDistance: 10}
	northSouthDistance := 0
	eastWestDistance := 0

	actions := convertToActionsSlice(input)

	for _, ferryAction := range actions {
		if ferryAction.action == "F" {
			northSouthDistance += (currentWaypoint.northSouthDistance * ferryAction.value)
			eastWestDistance += (currentWaypoint.eastWestDistance * ferryAction.value)
		}
		if ferryAction.action == "N" {
			currentWaypoint.northSouthDistance += ferryAction.value
		} else if ferryAction.action == "E" {
			currentWaypoint.eastWestDistance += ferryAction.value
		} else if ferryAction.action == "S" {
			currentWaypoint.northSouthDistance -= ferryAction.value
		} else if ferryAction.action == "W" {
			currentWaypoint.eastWestDistance -= ferryAction.value
		} else if ferryAction.action == "R" {
			currentWaypoint = determineDirection2(currentWaypoint, ferryAction.value)
		} else if ferryAction.action == "L" {
			currentWaypoint = determineDirection2(currentWaypoint, -ferryAction.value)
		}
	}

	return int(math.Abs(float64(northSouthDistance)) + math.Abs(float64(eastWestDistance)))
}

func convertToActionsSlice(input []string) []action {
	actions := []action{}
	for _, line := range input {
		value, error := strconv.Atoi(line[1:])
		if error != nil {
			log.Fatalf("Couldn't process %v as int!", line[1:])
		}

		actions = append(actions, action{
			action: string(line[0]),
			value:  value,
		})
	}

	return actions
}

func determineDirection(currentDirection string, rotation int) string {
	directionMap := map[string]int{east: 0, south: 1, west: 2, north: 3}
	directionSlice := []string{east, south, west, north}

	currentDirectionIndex := directionMap[currentDirection]
	movement := rotation / 90
	newDirectionIndex := (currentDirectionIndex + movement) % len(directionSlice)
	if newDirectionIndex < 0 {
		newDirectionIndex = len(directionSlice) + newDirectionIndex
	}

	return directionSlice[newDirectionIndex]
}

func determineDirection2(currentWaypoint waypoint, rotation int) waypoint {
	newWaypoint := waypoint{
		northSouthDistance: currentWaypoint.northSouthDistance,
		eastWestDistance:   currentWaypoint.eastWestDistance,
	}

	var newDirectionFromNorthSouth string
	var newDirectionFromEastWest string

	if currentWaypoint.northSouthDistance > 0 {
		newDirectionFromNorthSouth = determineDirection(north, rotation)
	} else {
		newDirectionFromNorthSouth = determineDirection(south, rotation)
	}

	if currentWaypoint.eastWestDistance > 0 {
		newDirectionFromEastWest = determineDirection(east, rotation)
	} else {
		newDirectionFromEastWest = determineDirection(west, rotation)
	}

	newWaypoint = setWaypointValue(newWaypoint, newDirectionFromNorthSouth, currentWaypoint.northSouthDistance)
	newWaypoint = setWaypointValue(newWaypoint, newDirectionFromEastWest, currentWaypoint.eastWestDistance)

	return newWaypoint
}

func setWaypointValue(newWaypoint waypoint, newDirection string, value int) waypoint {
	if newDirection == north {
		newWaypoint.northSouthDistance = int(math.Abs(float64(value)))
	} else if newDirection == south {
		newWaypoint.northSouthDistance = -int(math.Abs(float64(value)))
	} else if newDirection == east {
		newWaypoint.eastWestDistance = int(math.Abs(float64(value)))
	} else {
		newWaypoint.eastWestDistance = -int(math.Abs(float64(value)))
	}

	return newWaypoint
}
