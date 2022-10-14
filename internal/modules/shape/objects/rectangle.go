package objects

type Rectangle struct {
	Width  uint64
	Length uint64
}

func (r *Rectangle) GetArea() uint64 {
	return r.Length * r.Width
}

func (r *Rectangle) GetPerimeter() uint64 {
	return (2 * r.Length) + (2 * r.Width)
}
