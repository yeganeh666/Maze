package routes

import (
	"bufio"
	"fmt"
	"log"
	"maze/game"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//ReadFile of path
func ReadFile() ([9]string, int, [9][9]int) {
	lines, err := readLines("maze.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var track [9]string
	for i, line := range lines {
		fmt.Println(i, line)
		track[i] = line
	}
	//fmt.Println(track)
	matrix := game.AdjacencyMatrix(track)
	//fmt.Println(matrix[0][8])
	dijkstra := game.MakeGraph(matrix)
	return track, dijkstra, matrix
}
