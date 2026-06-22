package main

import (
	"fcl/chunk"
	"fcl/mesh"
	"fcl/shader"
	"fcl/td"
	"fcl/terrain"
	"fcl/window"
	"math"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {

	win, err := window.NewWindow(td.NewVec2(800, 800), "Window")
	td.Check(err)
	defer win.Close()

	program, err := shader.NewProgram(shader.VertexShader3D, shader.DefaultFragmentShader)
	td.Check(err)

	defer program.Destroy()

	ch := chunk.NewChunkManager(
		func(chunkSize, offsetX, offsetZ, lod int, seed int64) *mesh.Mesh {
			seedVal := seed
			hm := terrain.GenHeightmap(chunkSize, chunkSize, offsetX, offsetZ, 25, seedVal).
				Multiply(terrain.GenHeightmap(chunkSize, chunkSize, offsetX, offsetZ, 25, seedVal)).Round(5).
				Multiply(terrain.GenHeightmap(chunkSize, chunkSize, offsetX, offsetZ, 100, seedVal))
			scale := td.NewVec3(1,15,1)
			builder := mesh.NewMeshBuilder()

			layout, err := mesh.NewLayout(2, []int{3, 3})
			td.Check(err)

			for _, t := range hm.GenTriangles(scale,td.NewVec3(float32(offsetX), 0, float32(offsetZ)),lod) {
				builder.AddTriangle(t)
			}

			m, err := builder.Build(layout)
			td.Check(err)
			return m
		},
		7,
		func(i int) int {return int(math.Max(1,math.Min(50,float64(i-1))))},
		200,
		terrain.GetSeed(),
	)

	cam := window.NewCamera(td.NewVec3(0, 0, 0), td.NewVec3(-1, 0, 0), *win, 75, float32(ch.ViewDist*ch.ChunkSize)*1.41)

	var angle float32
	pos := td.NewVec3(0, -1, 0)
	var speed float32 = 32
	var lookSpeed float32 = 2

	for !win.ShouldClose() {
		ch.Update(cam.Pos)

		deltaTime := win.GetDeltaTime()

		if win.IsKeyDown(td.KeyW) {
			cam.Pos.X -= speed * deltaTime
		}
		if win.IsKeyDown(td.KeyS) {
			cam.Pos.X += speed * deltaTime
		}
		if win.IsKeyDown(td.KeyA) {
			cam.Pos.Z += speed * deltaTime
		}
		if win.IsKeyDown(td.KeyD) {
			cam.Pos.Z -= speed * deltaTime
		}

		if win.IsKeyDown(td.KeySpace) {
			cam.Pos.Y += speed * deltaTime
		}
		if win.IsKeyDown(td.KeyLeftShift) {
			cam.Pos.Y -= speed * deltaTime
		}

		if win.IsKeyDown(td.KeyUp) {
			cam.Pitch(lookSpeed * deltaTime)
		}
		if win.IsKeyDown(td.KeyDown) {
			cam.Pitch(-lookSpeed * deltaTime)
		}
		if win.IsKeyDown(td.KeyLeft) {
			cam.Yaw(lookSpeed * deltaTime)
		}
		if win.IsKeyDown(td.KeyRight) {
			cam.Yaw(-lookSpeed * deltaTime)
		}

		angle += deltaTime

		cam.Update() //Only needed if the camera isnt static

		translationMatrix := mgl32.Translate3D(pos.X, pos.Y, pos.Z)

		modelMatrix := translationMatrix

		win.StartFrame() //start drawing

		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		program.Use()
		program.SetUniformMat4("model", modelMatrix)
		program.SetUniformMat4("view", cam.ViewMatrix)
		program.SetUniformMat4("projection", cam.ProjectionMatrix)
		ch.Draw()

		win.EndFrame() //end drawing
	}
}
