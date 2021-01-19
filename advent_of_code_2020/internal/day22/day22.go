package day22

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type playerDeckHistory struct {
	p1Deck string
	p2Deck string
}

// GetWinningScore returns the winners score after simulating the combat game
func GetWinningScore(input []string, problem2 bool) int {
	playerDecks := parseInput(input)
	if len(playerDecks) != 2 {
		log.Fatalf("Sanity check failed! Expected 2 decks, but there were %v", len(playerDecks))
	}

	var winnerDeck []int
	if problem2 {
		_, winnerDeck = simulateRecursiveCombatToCompletion(playerDecks[0], playerDecks[1])
	} else {
		winnerDeck = simulateToCompletion(playerDecks[0], playerDecks[1])
	}

	score := 0
	for i, card := range winnerDeck {
		score += card * (len(winnerDeck) - i)
	}

	return score
}

func simulateRecursiveCombatToCompletion(p1Deck, p2Deck []int) (int, []int) {
	deckHistory := map[string]bool{}

	// count := 0
	for len(p1Deck) > 0 && len(p2Deck) > 0 {
		// count++
	  	// if count >= 1000 {
		// 	fmt.Printf("IN\n")
		// }
		historyString := "1:" + deckToString(p1Deck) + "+2:" + deckToString(p2Deck)
		if _, isPresent := deckHistory[historyString]; isPresent {
			return 1, p1Deck
		}
		deckHistory[historyString] = true

		p1Card := p1Deck[0]
		p1Deck = p1Deck[1:]
		p2Card := p2Deck[0]
		p2Deck = p2Deck[1:]

		var winner int
		if p1Card <= len(p1Deck) && p2Card <= len(p2Deck) {
			p1DeckCopy := make([]int, p1Card)
			copy(p1DeckCopy, p1Deck[:p1Card])
			p2DeckCopy := make([]int, p2Card)
			copy(p2DeckCopy, p2Deck[:p2Card])


			winner, _ = simulateRecursiveCombatToCompletion(p1DeckCopy, p2DeckCopy)
		} else if p1Card > p2Card {
			winner = 1
		} else {
			winner = 2
		}

		if winner == 1 {
			p1Deck = append(p1Deck, p1Card)
			p1Deck = append(p1Deck, p2Card)
		} else {
			p2Deck = append(p2Deck, p2Card)
			p2Deck = append(p2Deck, p1Card)
		}
	}

	if len(p1Deck) == 0 {
		return 2, p2Deck
	}

	return 1, p1Deck
}

func deckToString(deck []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(deck)), ","), "[]")
}

func parseInput(input []string) ([][]int) {
	playerDecks := [][]int{}
	if input[len(input)-1] != "" {
		input = append(input, "")
	}

	playerDeck := []int{}
	for _, line := range input {
		if line == "" {
			playerDecks = append(playerDecks, playerDeck)
			playerDeck = []int{}
			continue
		}

		intValue, e := strconv.Atoi(line)
		if e != nil {
			continue
		}

		playerDeck = append(playerDeck, intValue)
	}

	return playerDecks
}

func simulateToCompletion(p1Deck, p2Deck []int) []int {
	for len(p1Deck) > 0 && len(p2Deck) > 0 {
		p1Card := p1Deck[0]
		p1Deck = p1Deck[1:]
		p2Card := p2Deck[0]
		p2Deck = p2Deck[1:]

		if p1Card > p2Card {
			p1Deck = append(p1Deck, p1Card)
			p1Deck = append(p1Deck, p2Card)
		} else {
			p2Deck = append(p2Deck, p2Card)
			p2Deck = append(p2Deck, p1Card)
		}
	}

	if len(p1Deck) == 0 {
		return p2Deck
	}

	return p1Deck
}
