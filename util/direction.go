package util

type Direction struct {
	Y    int
	X    int
	Name string
}

var DirectionMap = map[string]Direction{
	"up":    {-1, 0, "up"},
	"down":  {1, 0, "down"},
	"left":  {0, -1, "left"},
	"right": {0, 1, "right"},
}
