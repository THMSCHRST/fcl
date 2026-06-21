package td

import "github.com/go-gl/mathgl/mgl32"

type Col struct {
	R float32
	G float32
	B float32
}

func NewCol(r, g, b float32) Col {
	return Col{r, g, b}
}

func ColToMGL32(c Col) mgl32.Vec3 { return mgl32.Vec3{c.R / 255, c.G / 255, c.B / 255} }
