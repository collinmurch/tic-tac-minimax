package main

import "fmt"

var fc int = 0

func minimax(b board, depth int) (int, int) {
	fc++
	if gameOver(b) {
		return score(b, depth), 0
	}
	depth++

	moves := getMoves(b)
	scores := make([]int, len(moves))
	
	for i := 0; i < len(moves); i++ {
		newGame := update(b, moves[i])
		scores[i], _ = minimax(newGame, depth)
	}

	fmt.Printf("%d, %d || %d\n", moves, scores, fc)
	
	var index int

	if b.turn == "X" {
		index = maxIndex(scores)
	} else {
		index = minIndex(scores)
	}

	return scores[index], moves[index]
}

func score(b board, depth int) int {
	if winning(b, "X") {
		return 10 - depth
	} else if winning(b, "O") {
		return depth - 10
	} else {
		return 0
	}
}

func maxIndex(slice []int) int {
	max := -1000
	var index int

	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
			index = i
		}
	}
	return index
}

func minIndex(slice []int) int {
	min := 1000
	var index int

	for i := 0; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
			index = i
		}
	}
	return index
}
