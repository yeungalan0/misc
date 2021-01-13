package day20

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yeungalan0/misc/advent_of_code_2020/internal/utils"
)

type testCase struct {
	input           []string
	expectedOutput  int
	expectedOutput2 int
}

func TestFitTiles(t *testing.T) {
	testCases := []testCase{
		{
			input: []string{
				"Tile 2311:",
				"..##.#..#.",
				"##..#.....",
				"#...##..#.",
				"####.#...#",
				"##.##.###.",
				"##...#.###",
				".#.#.#..##",
				"..#....#..",
				"###...#.#.",
				"..###..###",
				"",
				"Tile 1951:",
				"#.##...##.",
				"#.####...#",
				".....#..##",
				"#...######",
				".##.#....#",
				".###.#####",
				"###.##.##.",
				".###....#.",
				"..#.#..#.#",
				"#...##.#..",
				"",
				"Tile 1171:",
				"####...##.",
				"#..##.#..#",
				"##.#..#.#.",
				".###.####.",
				"..###.####",
				".##....##.",
				".#...####.",
				"#.##.####.",
				"####..#...",
				".....##...",
				"",
				"Tile 1427:",
				"###.##.#..",
				".#..#.##..",
				".#.##.#..#",
				"#.#.#.##.#",
				"....#...##",
				"...##..##.",
				"...#.#####",
				".#.####.#.",
				"..#..###.#",
				"..##.#..#.",
				"",
				"Tile 1489:",
				"##.#.#....",
				"..##...#..",
				".##..##...",
				"..#...#...",
				"#####...#.",
				"#..#.#.#.#",
				"...#.#.#..",
				"##.#...##.",
				"..##.##.##",
				"###.##.#..",
				"",
				"Tile 2473:",
				"#....####.",
				"#..#.##...",
				"#.##..#...",
				"######.#.#",
				".#...#.#.#",
				".#########",
				".###.#..#.",
				"########.#",
				"##...##.#.",
				"..###.#.#.",
				"",
				"Tile 2971:",
				"..#.#....#",
				"#...###...",
				"#.#.###...",
				"##.##..#..",
				".#####..##",
				".#..####.#",
				"#..#.#..#.",
				"..####.###",
				"..#.#.###.",
				"...#.#.#.#",
				"",
				"Tile 2729:",
				"...#.#.#.#",
				"####.#....",
				"..#.#.....",
				"....#..#.#",
				".##..##.#.",
				".#.####...",
				"####.#.#..",
				"##.####...",
				"##..#.##..",
				"#.##...##.",
				"",
				"Tile 3079:",
				"#.#.#####.",
				".#..######",
				"..#.......",
				"######....",
				"####.#..#.",
				".#...#.##.",
				"#.#####.##",
				"..#.###...",
				"..#.......",
				"..#.###...",
			},
			expectedOutput:  20899048083289,
			expectedOutput2: 273,
		},
	}

	for _, testCase := range testCases {
		actualOutput := FitTiles(testCase.input)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		// actualOutput2 := CalculateRoughWaters(testCase.input)

		// if testCase.expectedOutput2 != actualOutput2 {
		// 	t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		// }
	}
}

func TestOrientTile(t *testing.T) {
	origTileEdges := tileComponents{
		top:    "..##.#..#.",
		right:  "...#.##..#",
		bottom: "..###..###",
		left:   ".#####..#.",
	}

	flippedTileEdges := tileComponents{
		top:    utils.ReverseString(origTileEdges.top),
		right:  origTileEdges.left,
		bottom: utils.ReverseString(origTileEdges.bottom),
		left:   origTileEdges.right,
	}

	testCases := []struct {
		inputOrientation int
		inputTileEdges   tileComponents
		expectedOutput   tileComponents
	}{
		{
			inputOrientation: 0,
			inputTileEdges:   origTileEdges,
			expectedOutput:   origTileEdges,
		},
		{
			inputOrientation: 1,
			inputTileEdges:   origTileEdges,
			expectedOutput: tileComponents{
				top:    utils.ReverseString(origTileEdges.left),
				right:  origTileEdges.top,
				bottom: utils.ReverseString(origTileEdges.right),
				left:   origTileEdges.bottom,
			},
		},
		{
			inputOrientation: 2,
			inputTileEdges:   origTileEdges,
			expectedOutput: tileComponents{
				top:    utils.ReverseString(origTileEdges.bottom),
				right:  utils.ReverseString(origTileEdges.left),
				bottom: utils.ReverseString(origTileEdges.top),
				left:   utils.ReverseString(origTileEdges.right),
			},
		},
		{
			inputOrientation: 3,
			inputTileEdges:   origTileEdges,
			expectedOutput: tileComponents{
				top:    origTileEdges.right,
				right:  utils.ReverseString(origTileEdges.bottom),
				bottom: origTileEdges.left,
				left:   utils.ReverseString(origTileEdges.top),
			},
		},
		{
			inputOrientation: 4,
			inputTileEdges:   origTileEdges,
			expectedOutput:   flippedTileEdges,
		},
		{
			inputOrientation: 5,
			inputTileEdges:   origTileEdges,
			expectedOutput: tileComponents{
				top:    utils.ReverseString(flippedTileEdges.left),
				right:  flippedTileEdges.top,
				bottom: utils.ReverseString(flippedTileEdges.right),
				left:   flippedTileEdges.bottom,
			},
		},
		{
			inputOrientation: 6,
			inputTileEdges:   origTileEdges,
			expectedOutput: tileComponents{
				top:    utils.ReverseString(flippedTileEdges.bottom),
				right:  utils.ReverseString(flippedTileEdges.left),
				bottom: utils.ReverseString(flippedTileEdges.top),
				left:   utils.ReverseString(flippedTileEdges.right),
			},
		},
		{
			inputOrientation: 7,
			inputTileEdges:   origTileEdges,
			expectedOutput: tileComponents{
				top:    flippedTileEdges.right,
				right:  utils.ReverseString(flippedTileEdges.bottom),
				bottom: flippedTileEdges.left,
				left:   utils.ReverseString(flippedTileEdges.top),
			},
		},
	}

	for _, testCase := range testCases {
		actualOutput := orientTileEdges(testCase.inputOrientation, testCase.inputTileEdges)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}