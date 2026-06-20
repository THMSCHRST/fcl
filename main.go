package main

import (
	"fcl/helper"
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

    for !window.ShouldClose() {
        gl.Clear(gl.COLOR_BUFFER_BIT)
        // draw here
		window.EndFrame()
    }
}