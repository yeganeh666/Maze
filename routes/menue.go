package routes

import (
	"fmt"
	"maze/models"
)

//Menue of game
func Menue(name string, path [9]string, dijkstra int, matrix [9][9]int) {

	fmt.Printf("1)Start the game as %s\n2)Leaderboard\n3)Solve\n", name)
	var n int
	for {
		fmt.Scan(&n)
		switch n {
		case 1:
			score := Start(path, dijkstra, matrix)
			Players[name] = models.GameQueue{Score: score}
		case 2:
			LeaderBoard()
			fmt.Print("\n")
		case 4:
			return
		}
	}

}
