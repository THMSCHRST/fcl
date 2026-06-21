package w

import (
	"fcl/td"
	"math"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.4/glfw"
)

type Window struct {
	GlfwWindow *glfw.Window
	lastFrame  time.Time
}

// NewWindow initializes glfw and opengl and returns a window with the given parameters.
func NewWindow(size td.Vec2, title string) (*Window, error) {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		glfw.Terminate()
		return nil, err
	}

	// req opengl 4.6
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(int(size.X), int(size.Y), title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, err
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		window.Destroy()
		glfw.Terminate()
		return nil, err
	}

	w, h := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(w), int32(h)) // set viewport
	gl.ClearColor(1, 1, 1, 1.0)

	gl.Enable(gl.DEPTH_TEST)

	return &Window{GlfwWindow: window, lastFrame: time.Now()}, nil
}

// for w.ShouldClose() {//code here}
func (w *Window) ShouldClose() bool {
	return w.GlfwWindow.ShouldClose()
}

// Close terminates glfw. You probably need to use this with 'defer w.Close()'.
func (w *Window) Close() {
	w.GlfwWindow.Destroy()
	glfw.Terminate()
}

// EndFrame should be called after completing drawing actions.
func (w *Window) EndFrame() {
	w.SwapBuffers()
	w.PollEvents()
}

func (w *Window) StartFrame() {
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	w.lastFrame = time.Now()
}

// swap to new frame
func (w *Window) SwapBuffers() {
	w.GlfwWindow.SwapBuffers()
}

// PollEvents processes window events (input, resize, etc).
func (w *Window) PollEvents() {
	glfw.PollEvents()
}

func (w *Window) GetDeltaTime() float32 {
	return float32(time.Since(w.lastFrame).Seconds())
}

func (w *Window) IsKeyDown(key td.Key) bool {
	return w.GlfwWindow.GetKey(glfw.Key(key)) == glfw.Press
}

func (w *Window) GetFPS() float32 { return float32(math.Round(float64(1 / w.GetDeltaTime()))) }
