package move

type Ant struct {
	Name string
	CurrentRoom string
}


type Ants struct {
	Ants map[int]Ant
}