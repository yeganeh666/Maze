package routes

import (
	"fmt"
	"maze/models"
	"maze/queueheap"
)

var Players models.Members

//Get details of family members
func Get() (*queueheap.MinHeap, int) {
	var number int

	fmt.Println("Enter Number Of Family Members:")
	fmt.Scanln(&number)
	Players.Size = number
	gamequeue := make([]models.GameQueue, number)
	fmt.Println("Enter ِNames and Ages:")
	for i := 0; i < number; i++ {
		fmt.Scan(&gamequeue[i].Name, &gamequeue[i].Age)
		Players.Member[i].Name = gamequeue[i].Name
		Players.Member[i].Age = gamequeue[i].Age
	}
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
