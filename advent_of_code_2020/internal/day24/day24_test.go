package day24

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBlackTileLayoutCount(t *testing.T) {
	testCases := []struct {
		input           []string
		expectedOutput  int
		expectedOutput2 int
	}{
		{
			input: []string{
				"sesenwnenenewseeswwswswwnenewsewsw",
				"neeenesenwnwwswnenewnwwsewnenwseswesw",
				"seswneswswsenwwnwse",
				"nwnwneseeswswnenewneswwnewseswneseene",
				"swweswneswnenwsewnwneneseenw",
				"eesenwseswswnenwswnwnwsewwnwsene",
				"sewnenenenesenwsewnenwwwse",
				"wenwwweseeeweswwwnwwe",
				"wsweesenenewnwwnwsenewsenwwsesesenwne",
				"neeswseenwwswnwswswnw",
				"nenwswwsewswnenenewsenwsenwnesesenew",
				"enewnwewneswsewnwswenweswnenwsenwsw",
				"sweneswneswneneenwnewenewwneswswnese",
				"swwesenesewenwneswnwwneseswwne",
				"enesenwswwswneneswsenwnewswseenwsese",
				"wnwnesenesenenwwnenwsewesewsesesew",
				"nenewswnwewswnenesenwnesewesw",
				"eneswnwswnwsenenwnwnwwseeswneewsenese",
				"neswnwewnwnwseenwseesewsenwsweewe",
				"wseweeenwnesenwwwswnew",
			},
			expectedOutput: 10,
			expectedOutput2: 2208,
		},
	}

	for _, testCase := range testCases {
		actualOutput := BlackTileLayoutCount(testCase.input, 0)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}

		actualOutput2 := BlackTileLayoutCount(testCase.input, 100)

		if testCase.expectedOutput2 != actualOutput2 {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput2, actualOutput2)
		}
	}
}

func TestFlipByRules(t *testing.T) {
	testCases := []struct {
		tileLocSet map[tileLoc]bool
		rounds int
		expectedOutput int
	}{
		// {
		// 	tileLocSet: map[tileLoc]bool{
		// 		{row: 1, col: 0}: true,
		// 		{row: 1, col: 1}: true,
		// 		{row: 0, col: 0}: true,
		// 	},
		// 	rounds: 1,
		// 	expectedOutput: 6,
		// },
		// {
		// 	tileLocSet: map[tileLoc]bool{
		// 		{row: 1, col: 0}: true,
		// 		{row: 1, col: 1}: true,
		// 		{row: 0, col: 0}: true,
		// },
		// 	rounds: 2,
		// 	expectedOutput: 9,
		// },
		{
			tileLocSet: map[tileLoc]bool{
				{row: 0, col: 0}: true,
			},
			rounds: 1,
			expectedOutput: 0,
		},
	}

	for _, testCase := range testCases {
		actualTileLocSet := flipByRules(testCase.tileLocSet, testCase.rounds)
		actualOutput := countBlackTiles(actualTileLocSet)

		if testCase.expectedOutput != actualOutput {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}

func TestParseInput(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput [][]move
	}{
		{
			input: []string{
				"seswneswswsenwwnwse",
				"wseweeenwnesenwwwswnew",
			},
			expectedOutput: [][]move{
				{
					seMove, swMove, neMove, swMove, swMove,
					seMove, nwMove, wMove, nwMove, seMove,
				},
				{
					wMove, seMove, wMove, eMove, eMove, eMove, nwMove, neMove,
					seMove, nwMove, wMove, wMove, swMove, neMove, wMove,
				},
			},
		},
	}

	for _, testCase := range testCases {
		actualOutput := parseInput(testCase.input)

		if !cmp.Equal(testCase.expectedOutput, actualOutput) {
			t.Errorf("expected %v, but got %v\n", testCase.expectedOutput, actualOutput)
		}
	}
}
