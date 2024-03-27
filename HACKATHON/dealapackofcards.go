package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	j := 0
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %v: %v, %v, %v\n", i+1, deck[j], deck[j+1], deck[j+2])
		j += 3
	}
}
