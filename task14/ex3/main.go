package main

import "fmt"

func main() {
	var board [8][8]string

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = "."
			} else {
				board[i][j] = "*"
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}