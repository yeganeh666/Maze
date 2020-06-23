package routes

import (
	"fmt"
	"maze/game"
)

//Start the game
func Start() {
	g := [9]string{"##+###++C", "++++++++#", "###++++##", "##+++#++#", "###++#+##", "#####++#+", "##+++++++", "##++++#+#", "M++####+#"}
	fmt.Println(g)
	matrix := game.AdjacencyMatrix(g)
	//fmt.Println(matrix[0][8])
	game.MakeGraph(matrix)
}
