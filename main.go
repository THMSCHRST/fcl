package main

import (
	"fcl/mesh"
	"fcl/shader"
	"fcl/td"
	"fcl/w"
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	window, err := w.NewWindow(td.NewVec2(500, 500), "Window") //This creates a new window
	td.Check(err)
	defer window.Close() //Close the window after main func exit

	//A program is a bundle of vertex and fragment shader that is used to render a mesh
	program, err := shader.NewProgram(shader.VertexShader3D, shader.DefaultFragmentShader) //shader.VertexShader3D needs a 'model' uniform which is a transform, a 'view' uniform which should be cam.ViewMatrix and a 'projection' uniform which should be cam.ProjectionMatrix
	td.Check(err)

	defer program.Destroy() //Delete the program afterwards to prevent memory leaks

	builder := mesh.NewMeshBuilder() //A mesh builder can build a mesh from a number of triangles

	layout, err := mesh.NewLayout(2, []int{3, 3}) //A layout defines how a meshes uniforms work (this needs to match the reqs of the shaders)
	td.Check(err)

	//building a mesh using triangles
	apex := td.NewVec3(0, 1.5, 0)
	base1 := td.NewVec3(-1, -1, -1)
	base2 := td.NewVec3(1, -1, -1)
	base3 := td.NewVec3(1, -1, 1)
	base4 := td.NewVec3(-1, -1, 1)
	color := td.NewCol(33, 0, 78)

	builder.AddTriangle(mesh.NewTriangle(apex, base1, base2, color))
	builder.AddTriangle(mesh.NewTriangle(apex, base2, base3, color))
	builder.AddTriangle(mesh.NewTriangle(apex, base3, base4, color))
	builder.AddTriangle(mesh.NewTriangle(apex, base4, base1, color))

	mesh, err := builder.Build(layout) //build a mesh with a layout
	td.Check(err)

	defer mesh.Destroy() //prevent memory leaks

	cam := w.NewCamera(td.NewVec3(5, 0, 0), td.NewVec3(-1, 0, 0), *window, 45) //creates a new camera

	var angle float32
	pos := td.NewVec3(0, 0, 0)
	var speed float32 = 4

	for !window.ShouldClose() {
		deltaTime := window.GetDeltaTime()
		fmt.Println(window.GetFPS())

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

		cam.Update() //Only needed if the camera isnt static

		rotationMatrix := mgl32.HomogRotate3D(angle, mgl32.Vec3{0, 1, 0})

		translationMatrix := mgl32.Translate3D(pos.X, pos.Y, pos.Z)

		modelMatrix := translationMatrix.Mul4(rotationMatrix)

		window.StartFrame() //start drawing

		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()
		program.SetUniformMat4("model", modelMatrix)
		program.SetUniformMat4("view", cam.ViewMatrix)
		program.SetUniformMat4("projection", cam.ProjectionMatrix)
		mesh.Draw()

		window.EndFrame() //end drawing
	}
}
