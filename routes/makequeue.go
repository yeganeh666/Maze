package routes

import (
	"fmt"
	"maze/models"
	"maze/queueheap"
)

//Get details of family members
func Get() {
	var number int

	fmt.Println("Enter Number Of Family Members:")
	fmt.Scanln(&number)

	gamequeue := make([]models.GameQueue, number)

	for i := 0; i < number; i++ {
		fmt.Scan(&gamequeue[i].Name, &gamequeue[i].Age)
	}
	MakeQueue(gamequeue)
}

//MakeQueue of players' turn
func MakeQueue(gamequeue []models.GameQueue) {
	minHeap := queueheap.NewMinHeap(len(gamequeue))
	for i := 0; i < len(gamequeue); i++ {
		minHeap.Insert(gamequeue[i])
	}
	minHeap.BuildMinHeap()
	for i := 0; i < len(gamequeue); i++ {
		fmt.Println(minHeap.Remove())
	}
	Menue()
}
