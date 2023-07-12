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

func check_winner(b [3][3]string) string {
	a := [9]string{b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2], b[2][0], b[2][1], b[2][2]}

	if (a[0]+a[1]+a[2] == "XXX") || (a[0]+a[1]+a[2] == "OOO") {
		if a[0]+a[1]+a[2] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[3]+a[4]+a[5] == "XXX" || a[3]+a[4]+a[5] == "OOO" {
		if a[3]+a[4]+a[5] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[6]+a[7]+a[8] == "XXX" || a[6]+a[7]+a[8] == "OOO" {
		if a[6]+a[7]+a[8] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[0]+a[3]+a[6] == "XXX" || a[0]+a[3]+a[6] == "OOO" {
		if a[0]+a[3]+a[6] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[1]+a[4]+a[7] == "XXX" || a[1]+a[4]+a[7] == "OOO" {
		if a[1]+a[4]+a[7] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[2]+a[5]+a[8] == "XXX" || a[2]+a[5]+a[8] == "OOO" {
		if a[2]+a[5]+a[8] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[0]+a[4]+a[8] == "XXX" || a[0]+a[4]+a[8] == "OOO" {
		if a[0]+a[4]+a[8] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	} else if a[2]+a[4]+a[6] == "XXX" || a[2]+a[4]+a[6] == "OOO" {
		if a[2]+a[4]+a[6] == "XXX" {
			return "X"
		} else {
			return "O"
		}
	}

	return "NO"
}

func main() {
	a := [3][3]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}

	for i := 0; i < 9; i++ {
		fmt.Printf("Player %v 's turn. Enter position (x,y)\n", i%2+1)
		var n [2]int
		n = input_bord()
		if (a[n[0]-1][n[1]-1] == "O") || (a[n[0]-1][n[1]-1] == "X") {
			fmt.Println("Cheat!!!!!")
			break
		}
		if i%2 == 0 {
			a[n[0]-1][n[1]-1] = "O"

			win := check_winner(a)

			print_board(a)
			if win == "O" {
				fmt.Printf("Player %v win!\n", i%2+1)
				break
			}
		} else {
			a[n[0]-1][n[1]-1] = "X"

			win := check_winner(a)

			print_board(a)
			if win == "X" {
				fmt.Printf("Player %v win!\n", i%2+1)
				break
			}
		}

	}

	fmt.Printf("draw")
}
