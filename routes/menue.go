package routes

import "fmt"

//Menue of game
func Menue() {

	fmt.Println("1)Start the game as NAME\n2)Leaderboard\n3)Solve")
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		Start()
	}

}
