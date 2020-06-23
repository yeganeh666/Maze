package game

import (
	"fmt"
	"strconv"
)

var plus int

//AdjacencyMatrix of our maze
func AdjacencyMatrix(path [9]string) [9][9]int {
	var matrix [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if path[i][j] == '+' || (i == 0 && j == 8) || (i == 8 && j == 0) {
				matrix[i][j] = 1
				plus = plus + 1
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
func MakeGraph(matrix [9][9]int) {
	graph := Graph{}
	nodes := make(map[string]*Node)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if matrix[i][j] == 1 {
				name := strconv.Itoa(i) + strconv.Itoa(j)
				nodes[name] = &Node{Name: name}
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if matrix[i][j] == 1 {
				if i+1 < 9 && matrix[i+1][j] == 1 {
					name := strconv.Itoa(i) + strconv.Itoa(j)
					name2 := strconv.Itoa(i+1) + strconv.Itoa(j)
					graph.AddEdge(nodes[name], nodes[name2], 1)
					// fmt.Println(nodes[name].Name)
					// fmt.Println(nodes[name2].Name)
				} else if i-1 >= 0 && matrix[i-1][j] == 1 {
					name := strconv.Itoa(i) + strconv.Itoa(j)
					name2 := strconv.Itoa(i-1) + strconv.Itoa(j)
					graph.AddEdge(nodes[name], nodes[name2], 1)
					// fmt.Println(nodes[name].Name)
					// fmt.Println(nodes[name2].Name)
				} else if j-1 >= 0 && matrix[i][j-1] == 1 {
					name := strconv.Itoa(i) + strconv.Itoa(j)
					name2 := strconv.Itoa(i) + strconv.Itoa(j-1)
					graph.AddEdge(nodes[name], nodes[name2], 1)
					// fmt.Println(nodes[name].Name)
					// fmt.Println(nodes[name2].Name)
				} else if j+1 < 9 && matrix[i][j+1] == 1 {
					name := strconv.Itoa(i) + strconv.Itoa(j)
					name2 := strconv.Itoa(i) + strconv.Itoa(j+1)
					graph.AddEdge(nodes[name], nodes[name2], 1)
					// fmt.Println(nodes[name].Name)
					// fmt.Println(nodes[name2].Name)
				}
			}
		}

	}
	//fmt.Println(nodes["08"].Name)

	fmt.Println(graph.Dijkstra(nodes["80"], nodes["08"]))
}
