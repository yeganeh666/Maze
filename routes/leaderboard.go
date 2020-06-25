package routes

import (
	"fmt"
	"maze/quicksort"
)

//LeaderBoard show scores
func LeaderBoard() {
	var arr []int
	for i, j := range Players.Member {
		arr[i] = j.Score
	}
	fmt.Println("*** Leader Board ***")
	quicksort.QuickSort(arr)
	for i, j := range arr {
		fmt.Printf("%d ==> %d\n", i, j)
	}
}
