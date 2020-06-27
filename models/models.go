package models

//GameQueue consists of the names and ages of the family members.
type GameQueue struct {
	Name  string
	Age   int
	Score int
}

//NodeStack uses for find sloution
type NodeStack struct {
	X   int
	Y   int
	Dir int
}
