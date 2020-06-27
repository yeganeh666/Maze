package routes

import (
	"fmt"
	"maze/implementDS"
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
	arr = implementDS.QuickSort(arr)
	var names []string
	for i, j := range arr {
		name := findePerson(j, names)
		names = append(names, name)
		fmt.Printf("%d  %s==> %d\n", i, name, j)
	}
}
func findePerson(score int, names []string) string {
	for i, j := range Players {
		flag := false
		for k := 0; k < len(names); k++ {
			if j.Name == names[k] {
				flag = true
			}
		}
		if j.Score == score && flag == false {
			return i
		}
	}
	return ""
}
