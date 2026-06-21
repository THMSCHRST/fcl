package main

import (
	"fcl/helper"
	"fcl/meshUtil"
	"fcl/render"
	"fcl/td"
	"fcl/w"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window, err := w.NewWindow(td.NewVec2(500, 500), "Window")
	helper.Check(err)
	defer window.Close()

	program, err := render.NewProgram(render.VertexShader3D, render.DefaultFragmentShader)
	helper.Check(err)

	defer program.Destroy()

    builder := meshUtil.NewMeshBuilder()

    layout := []render.Attribute{
        {Index: 0, Size: 3, Offset: 0},
        {Index: 1, Size: 3, Offset: 12},
    }

    apex := mgl32.Vec3{0, 1.5, 0}
    base1 := mgl32.Vec3{-1, -1, -1}
    base2 := mgl32.Vec3{1, -1, -1}
    base3 := mgl32.Vec3{1, -1, 1}
    base4 := mgl32.Vec3{-1, -1, 1}

    builder.AddTriangle(apex, base1, base2, mgl32.Vec3{1, 0, 0})
    builder.AddTriangle(apex, base2, base3, mgl32.Vec3{0, 1, 0})
    builder.AddTriangle(apex, base3, base4, mgl32.Vec3{0, 0, 1})
    builder.AddTriangle(apex, base4, base1, mgl32.Vec3{1, 1, 0})

	mesh, err := builder.Build(layout)
	helper.Check(err)

	defer mesh.Destroy()

	cam := render.NewCamera(td.NewVec3(5, 0, 0), td.NewVec3(-1, 0, 0), *window, 45)

	var angle float32
	pos := td.NewVec3(0, 0, 0)
	var speed float32 = 4

	render.EnableDepth()

	for !window.ShouldClose() {
		deltaTime := window.GetDeltaTime()

		if window.IsKeyDown(td.KeyW) {
			cam.Pos.X -= speed * deltaTime
		}
		if window.IsKeyDown(td.KeyS) {
			cam.Pos.X += speed * deltaTime
		}
		if window.IsKeyDown(td.KeyA) {
			cam.Pos.Z += speed * deltaTime
		}
		if window.IsKeyDown(td.KeyD) {
			cam.Pos.Z -= speed * deltaTime
		}
		if window.IsKeyDown(td.KeySpace) {
			cam.Pos.Y += speed * deltaTime
		}
		if window.IsKeyDown(td.KeyLeftShift) {
			cam.Pos.Y -= speed * deltaTime
		}

		angle += deltaTime

		cam.Update()

		render.UpdateDepth()

		rotationMatrix := mgl32.HomogRotate3D(angle, mgl32.Vec3{0, 1, 0})

		translationMatrix := mgl32.Translate3D(pos.X, pos.Y, pos.Z)

		modelMatrix := translationMatrix.Mul4(rotationMatrix)

		window.StartFrame()

		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()
		program.SetUniformMat4("model", modelMatrix)
		program.SetUniformMat4("view", cam.ViewMatrix)
		program.SetUniformMat4("projection", cam.ProjectionMatrix)
		mesh.Draw()

		window.EndFrame()
	}
}
