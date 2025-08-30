package solutions

import "fmt"

// https://leetcode.com/problems/word-search/description/
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range len(board) {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == word[0] {
				if recurse(board, visited, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func recurse(board [][]byte, visited [][]bool, word string, start, i, j int) bool {
	if len(word) == start {
		return true
	}
	if i >= len(board) || j >= len(board[0]) {
		return false
	}
	if i < 0 || j < 0 {
		return false
	}
	if visited[i][j] {
		return false
	}
	if word[start] != board[i][j] {
		return false
	}
	visited[i][j] = true
	valid := recurse(board, visited, word, start+1, i+1, j) ||
		recurse(board, visited, word, start+1, i-1, j) ||
		recurse(board, visited, word, start+1, i, j+1) ||
		recurse(board, visited, word, start+1, i, j-1)
	visited[i][j] = false
	return valid
}

func WordSearch() {
	fmt.Print(exist([][]byte{}, ""))
}
