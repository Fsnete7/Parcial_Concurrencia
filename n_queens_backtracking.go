package main

import (
	"fmt"
)

var avail = 0
var noAvail = 1

type Solution struct {
	n            int
	placedQueens []int
	step         int
}

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}

	var finalAns = 0
	for i := 0; i < n; i++ {
		var placedQueens = make([]int, n)
		placedQueens[0] = i
		var solution = Solution{n, placedQueens, 1}
		var ans = resolvedSolution(solution)
		finalAns = finalAns + len(ans)
	}

	return finalAns
}

func resolvedSolution(solution Solution) [][]int {
	var step = solution.step + 1
	var rowNumber = solution.step
	var availablePosition = calculateAvailablePosition(solution.placedQueens, rowNumber, solution.n)

	var finalAnswer = [][]int{}

	if step == solution.n {
		for i := 0; i < solution.n; i++ {
			if availablePosition[i] == avail {
				solution.placedQueens[rowNumber] = i
				var successAns = make([]int, solution.n)
				for j := 0; j < solution.n; j++ {
					successAns[j] = solution.placedQueens[j]
				}
				return append(finalAnswer, successAns)
			}
		}
		return finalAnswer
	} else {
		for i := 0; i < solution.n; i++ {
			if availablePosition[i] == avail {
				solution.placedQueens[rowNumber] = i
				solution.step = step
				var ans = resolvedSolution(solution)
				for j := 0; j < len(ans); j++ {
					finalAnswer = append(finalAnswer, ans[j])
				}
			}
		}
		return finalAnswer
	}
}

func calculateAvailablePosition(placedQueens []int, rowNumber int, n int) []int {
	var row = make([]int, n)
	if rowNumber == 0 {
		return row
	}

	for i := 0; i < rowNumber; i++ {
		var prevQueen = placedQueens[i]
		row[prevQueen] = noAvail
		var diff = rowNumber - i
		if prevQueen-diff >= 0 {
			row[prevQueen-diff] = noAvail
		}
		if prevQueen+diff <= n-1 {
			row[prevQueen+diff] = noAvail
		}
	}
	return row
}

func main() {
	fmt.Println(totalNQueens(5))
}
