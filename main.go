package main

import (
	"fcl/helper"
	"fcl/render"
	"fcl/td"
	"fcl/w"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
)

func init() {
    runtime.LockOSThread()
}

func main() {
	window, err := w.NewWindow(td.NewVec2(500,500),"Hello"); helper.Check(err)
	defer window.Close()

    const (
        vertexShaderSrc = `#version 460 core
        layout (location = 0) in vec3 aPos;
        layout (location = 1) in vec3 aColor;
        out vec3 color;
        void main() {
            gl_Position = vec4(aPos, 1.0);
            color = aColor;
        }`

        fragmentShaderSrc = `#version 460 core
        in vec3 color;
        out vec4 fragColor;
        void main() {
            fragColor = vec4(color, 1.0);
        }`
    )

    program, err := render.NewProgram(vertexShaderSrc, fragmentShaderSrc)
	helper.Check(err)

    defer program.Destroy()

    // triangle
    vertices := []float32{
        // pos             // col
        -0.5, -0.5, 0.0,    1.0, 0.0, 0.0, // r
         0.5, -0.5, 0.0,    0.0, 1.0, 0.0, // g
         0.0,  0.5, 0.0,    0.0, 0.0, 1.0, // b
    }
    mesh, err := render.NewMesh(vertices, nil)
	helper.Check(err)

    defer mesh.Destroy()

    for !window.ShouldClose() {
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
        gl.Clear(gl.COLOR_BUFFER_BIT)

        program.Use()
        mesh.Draw()

		window.EndFrame()
    }
}