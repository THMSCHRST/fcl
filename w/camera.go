package w

import (
	"fcl/td"

	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Pos              td.Vec3
	Dir              td.Vec3
	ViewMatrix       mgl32.Mat4
	ProjectionMatrix mgl32.Mat4
	FOV              float32
}

func NewCamera(pos, dir td.Vec3, window Window, fov float32) Camera {
	w, h := window.GlfwWindow.GetFramebufferSize()
	return Camera{
		Pos: pos,
		Dir: dir,
		ViewMatrix: mgl32.LookAtV(
			mgl32.Vec3{pos.X, pos.Y, pos.Z},
			mgl32.Vec3{dir.X + pos.X, dir.Y + pos.Y, dir.Z + pos.Z},
			mgl32.Vec3{0, 1, 0},
		),
		ProjectionMatrix: mgl32.Perspective(
			mgl32.DegToRad(fov),
			float32(w)/float32(h),
			0.1,
			100.0,
		),
		FOV: fov,
	}
}

func (c *Camera) Update() {
	c.ViewMatrix = mgl32.LookAtV(
		mgl32.Vec3{c.Pos.X, c.Pos.Y, c.Pos.Z},
		mgl32.Vec3{c.Dir.X + c.Pos.X, c.Dir.Y + c.Pos.Y, c.Dir.Z + c.Pos.Z},
		mgl32.Vec3{0, 1, 0},
	)
}
