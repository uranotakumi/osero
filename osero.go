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

func input_board() [2]int {
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

func check_change(a [8][8]string, n [2]int, c int) [8][8]string {
	if c%2 == 0 {
		// yoko
		for i := 0; i < 8; i++ {
			if i < n[1]-1 {
				if a[n[0]-1][i] == "O" {
					for j := i + 1; j <= n[1]-2; j++ {
						a[n[0]-1][j] = "O"
					}
				}
			}
			if i > n[1]-1 {
				if a[n[0]-1][i] == "O" {
					for j := n[1] - 2; j <= i-1; j++ {
						a[n[0]-1][j] = "O"
					}
				}
			}
		}
		// tate
		for i := 0; i < 8; i++ {
			if i < n[0]-1 {
				if a[i][n[1]-1] == "O" {
					for j := i + 1; j <= n[0]-2; j++ {
						a[j][n[1]-1] = "O"
					}
				}
			}
			if i > n[0]-1 {
				if a[i][n[1]-1] == "O" {
					for j := n[0] - 2; j <= i-1; j++ {
						a[j][n[1]-1] = "O"
					}
				}
			}
		}
	} else {
		// yoko
		for i := 0; i < 8; i++ {
			if i < n[1]-1 {
				if a[n[0]-1][i] == "X" {
					for j := i + 1; j <= n[1]-2; j++ {
						a[n[0]-1][j] = "X"
					}
				}
			}
			// koko ga hen
			/*
				if i > n[1]-1 {
					if a[n[0]-1][i] == "X" {
						for j := n[1] - 2; j <= i-1; j++ {
							a[n[0]-1][j] = "X"
						}
					}
				}
			*/
		}
		// tate
		for i := 0; i < 8; i++ {
			if i < n[0]-1 {
				if a[i][n[1]-1] == "X" {
					for j := i + 1; j <= n[0]-2; j++ {
						a[j][n[1]-1] = "X"
					}
				}
			}
			// kokoga hen
			/*
				if i > n[0]-1 {
					if a[i][n[1]-1] == "X" {
						for j := n[0] - 2; j <= i-1; j++ {
							a[j][n[1]-1] = "X"
						}
					}
				}
			*/
		}
	}

	return a
}

func main() {
	a := [8][8]string{
		{".", ".", ".", ".", ".", ".", ".", "."},
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
		n = input_board()
		if (a[n[0]-1][n[1]-1] == "O") || (a[n[0]-1][n[1]-1] == "X") {
			fmt.Println("Cheat!!!!!")
			break
		}
		if i%2 == 0 {
			a[n[0]-1][n[1]-1] = "O"

			// win := check_winner(a)

			a = check_change(a, n, i)
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

			a = check_change(a, n, i)
			print_board(a)
			check_number(a)
			/*
				if win == "X" {
					fmt.Printf("Player %v win!\n", i%2+1)
					break
				}
			*/
		}
	}
	fmt.Println("Draw.")
}
