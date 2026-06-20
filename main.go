package main

import (
	"fcl/td"
	"fcl/w"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
)

func init() {
    runtime.LockOSThread()
}

func main() {
	window := w.InitWindow(td.NewVec2(500,500),"Test")

    for !window.ShouldClose() {
        gl.Clear(gl.COLOR_BUFFER_BIT)
        // draw here
        window.SwapBuffers()
        glfw.PollEvents()
    }
}