package routes

import (
	"fmt"
	"maze/quicksort"
)

//LeaderBoard show scores
func LeaderBoard() {
	arr := make([]int, Number)
	k := 0
	for _, j := range Players {
		arr[k] = j.Score
		k = k + 1
	}
	fmt.Println("*** Leader Board ***")
	arr = quicksort.QuickSort(arr)
	for i, j := range arr {
		name := findePerson(j)
		fmt.Printf("%d  %s==> %d\n", i, name, j)
	}
}
func findePerson(score int) string {
	for i, j := range Players {
		if j.Score == score {
			return i
		}
	}
	return ""
}
