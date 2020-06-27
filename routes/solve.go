package routes

import (
	"fmt"
	"maze/implementDS"
	"maze/models"
)

var (
	n       int
	m       int
	fx      int
	fy      int
	Visited [9][9]int
)

//IsReachable for find solution
func IsReachable(maze [9][9]int) bool {
	n = 9
	m = 9
	i, j := 0, 8
	s := &implementDS.CustomQueue{
		Stack: make([]models.NodeStack, 0),
	}
	temp := models.NodeStack{i, j, 0}
	s.Push(temp)
	for !s.Empty() {
		// Pop the top node and move to the
		// left, right, top, down or retract
		// back according the value of node's
		// Dir variable.
		temp = s.Top()
		d := temp.Dir
		i = temp.X
		j = temp.Y
		// Increment the Direction and
		// push the node in the stack again.
		temp.Dir = temp.Dir + 1
		s.Pop()
		s.Push(temp)
		// If we reach the Food coordinates
		// return true
		if i == fx && j == fy {
			// length := s.Size()
			// for i := 0; i < length; i++ {
			// 	fmt.Printf("%d: %d %d \n", i, s.Top().X, s.Top().Y)
			// 	s.Pop()
			// }
			return true
		}
		// Checking the Up Direction.
		if d == 0 {
			if i-1 >= 0 && maze[i-1][j] == 1 && Visited[i-1][j] == 1 {
				temp1 := models.NodeStack{X: i - 1, Y: j}
				Visited[i-1][j] = 0
				s.Push(temp1)
			}
		} else if d == 1 { // Checking the left Direction
			if j-1 >= 0 && maze[i][j-1] == 1 && Visited[i][j-1] == 1 {
				temp1 := models.NodeStack{X: i, Y: j - 1}
				Visited[i][j-1] = 0
				s.Push(temp1)
			}
		} else if d == 2 { // Checking the down Direction
			if i+1 < n && maze[i+1][j] == 1 && Visited[i+1][j] == 1 {
				temp1 := models.NodeStack{X: i + 1, Y: j}
				Visited[i+1][j] = 0
				s.Push(temp1)
			}
		} else if d == 3 { // Checking the right Direction
			if j+1 < m && maze[i][j+1] == 1 && Visited[i][j+1] == 1 {
				temp1 := models.NodeStack{X: i, Y: j + 1}
				Visited[i][j+1] = 0
				s.Push(temp1)
			}
		} else {
			// If none of the Direction can take
			// the rat to the Food, retract back
			// to the path where the rat came from.
			Visited[temp.X][temp.Y] = 1
			s.Pop()
		}
	}
	// If the stack is empty and
	// no path is found return false.
	return false

}
func Solve(maze [9][9]int) {
	Visited = [9][9]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1}}

	// Food coordinates
	fx = 8
	fy = 0

	if IsReachable(maze) == true {
		fmt.Println("Path Found!")
		for i, j := range Visited {
			fmt.Println(i, j)

		}
	} else {
		fmt.Println("No Path Found!")
	}
	return
}
