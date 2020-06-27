package routes

import (
	"fmt"
	"maze/game"
	"maze/models"
	"maze/queueheap"
)

//Players of game
var Players = make(map[string]models.GameQueue)

//Number of family
var Number int

//Get details of family members
func Get() (*queueheap.MinHeap, int) {

	fmt.Println("Enter Number Of Family Members:")
	fmt.Scanln(&Number)
	gamequeue := make([]models.GameQueue, Number)
	fmt.Println("Enter ِNames and Ages:")
	for i := 0; i < Number; i++ {
		fmt.Scan(&gamequeue[i].Name, &gamequeue[i].Age)
		name := gamequeue[i].Name
		Players[name] = models.GameQueue{Name: gamequeue[i].Name, Age: gamequeue[i].Age, Score: game.Infinity}
	}
	// for i, j := range Players {
	// 	fmt.Println(i, j)
	// }
	minHeap := MakeQueue(gamequeue)
	return minHeap, len(gamequeue)
}

//MakeQueue of players' turn
func MakeQueue(gamequeue []models.GameQueue) *queueheap.MinHeap {
	minHeap := queueheap.NewMinHeap(len(gamequeue))
	for i := 0; i < len(gamequeue); i++ {
		minHeap.Insert(gamequeue[i])
	}
	minHeap.BuildMinHeap()
	return minHeap
	// for i := 0; i < len(gamequeue); i++ {
	// 	fmt.Println(minHeap.Remove())
	// }
	// Menue()
}

// func GameQueue(minHeap *queueheap.MinHeap) {
// 	for i := 0; i < minHeap.Size; i++ {
// 		fmt.Println(minHeap.Remove())
// 	}
// 	Menue()
// }
