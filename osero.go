package main

import "fmt"

func print_board(a [8][8]string) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}

func input_bord() [2]int {
	var n [2]int
	fmt.Scan(&n[0], &n[1])
	return n
}

func check_number(a [8][8]string) {
	n := [2]int{0, 0}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if a[i][j] == "O" {
				n[0]++
			} else if a[i][j] == "X" {
				n[1]++
			}
		}
	}
	fmt.Println("O:%v, X:%v.", n[0], n[1])
}

func main() {
	a := [8][8]string{{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "O", "X", ".", ".", "."},
		{".", ".", ".", "X", "O", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."}}
	print_board(a)

	for i := 0; i < 60; i++ {
		fmt.Printf("Player %v 's turn. Enter position (x,y)\n", i%2+1)
		var n [2]int
		n = input_bord()
		if (a[n[0]-1][n[1]-1] == "O") || (a[n[0]-1][n[1]-1] == "X") {
			fmt.Println("Cheat!!!!!")
			break
		}
		if i%2 == 0 {
			a[n[0]-1][n[1]-1] = "O"

			// win := check_winner(a)

			print_board(a)
			check_number(a)
			/*
				if win == "O" {
					fmt.Printf("Player %v win!\n", i%2+1)
					break
				}
			*/
		} else {
			a[n[0]-1][n[1]-1] = "X"

			// win := check_winner(a)

			print_board(a)
			check_number(a)
			/*
				if win == "X" {
					fmt.Printf("Player %v win!\n", i%2+1)
					break
				}
			*/
		}
		fmt.Println("Draw.")
	}
}
