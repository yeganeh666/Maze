package main

import (
	"maze/routes"
)

func main() {
	path, dijkstra, matrix := routes.ReadFile()
	GameQueue, len := routes.Get()
	for i := 0; i < len; i++ {
		person := GameQueue.Remove()
		//fmt.Println(person)
		routes.Menue(person.Name, path, dijkstra, matrix)
	}

}
