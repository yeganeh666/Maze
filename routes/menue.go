package routes

import "fmt"

//Menue of game
func Menue(name string, path [9]string, dijkstra int, matrix [9][9]int) {

	fmt.Printf("1)Start the game as %s\n2)Leaderboard\n3)Solve\n", name)
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		Start(path, dijkstra, matrix)
	}

}
