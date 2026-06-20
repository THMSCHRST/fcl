package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
)

func init() {
    runtime.LockOSThread()
}

func main() {
    if err := glfw.Init(); err != nil {
        panic(err)
    }
    defer glfw.Terminate()

    // req opengl 4.6
    glfw.WindowHint(glfw.ContextVersionMajor, 4)
    glfw.WindowHint(glfw.ContextVersionMinor, 6)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(800, 600, "Hello OpenGL", nil, nil)
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()

    // init opengl
    if err := gl.Init(); err != nil {
        panic(err)
    }

    // set viewport
    width, height := window.GetFramebufferSize()
    gl.Viewport(0, 0, int32(width), int32(height))
    gl.ClearColor(1, 1, 1, 1.0)

    for !window.ShouldClose() {
        gl.Clear(gl.COLOR_BUFFER_BIT)
        // draw here
        window.SwapBuffers()
        glfw.PollEvents()
    }
}