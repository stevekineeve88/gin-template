package managers

import "shape/objects"

var shapeManager *ShapeManager

func GetShapeManager() ShapeManager {
	if shapeManager == nil {
		shapeManager = &ShapeManager{}
	}
	return *shapeManager
}

type ShapeManager struct{}

func (sh *ShapeManager) GetAreaRectangle(l uint64, w uint64) uint64 {
	r := objects.Rectangle{Width: w, Length: l}
	return r.GetArea()
}

func (sh *ShapeManager) GetPerimeterRectangle(l uint64, w uint64) uint64 {
	r := objects.Rectangle{Width: w, Length: l}
	return r.GetPerimeter()
}
