package routes

import (
	"fmt"
	"time"
)

//Start the game
func Start(path [9]string, dijkstra int, matrix [9][9]int) int {
	choose := ""
	score := 0
	x, y, end := 8, 0, 0
	track := to2D(path)
	showPath(track)
	timer := time.Now()
	for {
		fmt.Scan(&choose)
		switch choose {
		case "W":
			if x-1 >= 0 && matrix[x-1][y] != 0 {
				temp := track[x][y]
				track[x][y] = track[x-1][y]
				track[x-1][y] = temp
				x = x - 1
				showPath(track)

			} else {
				fmt.Println("Invalid Move!")
				showPath(track)
			}
		case "S":
			if x+1 < 9 && matrix[x+1][y] != 0 {
				temp := track[x][y]
				track[x][y] = track[x+1][y]
				track[x+1][y] = temp
				x = x + 1
				showPath(track)

			} else {
				fmt.Println("Invalid Move!")
				showPath(track)
			}
		case "D":
			if y+1 < 9 && matrix[x][y+1] != 0 {
				temp := track[x][y]
				track[x][y] = track[x][y+1]
				track[x][y+1] = temp
				y = y + 1
				showPath(track)

			} else {
				fmt.Println("Invalid Move!")
				showPath(track)
			}
		case "A":
			if y-1 >= 0 && matrix[x][y-1] != 0 {
				temp := track[x][y]
				track[x][y] = track[x][y-1]
				track[x][y-1] = temp
				y = y - 1
				showPath(track)

			} else {
				fmt.Println("Invalid Move!")
				showPath(track)
			}

		}
		score = score + 1
		if (x == 0) && (y == 8) {
			end = int(time.Since(timer).Seconds())
			score = calulateScore(dijkstra, score, end)
			break
		}
	}

	fmt.Print("\n~~~~ *_* YOU WIN *_* ~~~~\n\n")
	fmt.Printf("TIME: %d    SCORE: %d\n", end, score)
	return score

}
func calulateScore(dijkstra, score, time int) int {
	return (score - dijkstra) * time
}
func to2D(path [9]string) [9][9]string {
	var track [9][9]string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			track[i][j] = string(path[i][j])
		}
	}
	return track
}
func showPath(path [9][9]string) {
	CallClear()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(path[i][j])
		}
		fmt.Print("\n")
	}
}
