package main

import (
	"fcl/helper"
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

	program, err := render.NewProgram(render.TransformVertexShader, render.DefaultFragmentShader)
	helper.Check(err)

	defer program.Destroy()

	// triangle
	vertices := []float32{
		// pos             // col
		-0.5, -0.5, 0.0, 1.0, 0.0, 0.0, // r
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, // g
		0.0, 0.5, 0.0, 0.0, 0.0, 1.0, // b
	}
	mesh, err := render.NewMesh(vertices, nil)
	helper.Check(err)

	defer mesh.Destroy()

	var angle float32
	pos := td.NewVec3(0, 0, 0)
	var speed float32 = 4

	for !window.ShouldClose() {
		deltaTime := window.GetDeltaTime()

		if window.IsKeyDown(td.KeyW) {
			pos.Y += speed * deltaTime
		}
		if window.IsKeyDown(td.KeyS) {
			pos.Y -= speed * deltaTime
		}
		if window.IsKeyDown(td.KeyA) {
			pos.X -= speed * deltaTime
		}
		if window.IsKeyDown(td.KeyD) {
			pos.X += speed * deltaTime
		}

		angle += deltaTime

		rotationMatrix := mgl32.HomogRotate3D(angle, mgl32.Vec3{0, 0, 1})

		translationMatrix := mgl32.Translate3D(pos.X, pos.Y, pos.Z)

		modelMatrix := translationMatrix.Mul4(rotationMatrix)

		window.StartFrame()

		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()
		program.SetUniformMat4("model", modelMatrix)
		mesh.Draw()

		window.EndFrame()
	}
}
