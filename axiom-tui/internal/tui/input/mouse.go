package input

type MouseRegion struct {
	ID     string
	X      int
	Y      int
	Width  int
	Height int
}

func (r MouseRegion) Contains(x int, y int) bool {
	return x >= r.X && x < r.X+r.Width && y >= r.Y && y < r.Y+r.Height
}
