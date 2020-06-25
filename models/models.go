package models

//GameQueue consists of the names and ages of the family members.
type GameQueue struct {
	Name  string
	Age   int
	Score int
}

//Members of family
type Members struct {
	Size   int
	Member []GameQueue
}
