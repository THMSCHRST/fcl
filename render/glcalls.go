package render

import "github.com/go-gl/gl/v4.6-core/gl"

func EnableDepth() {
	gl.Enable(gl.DEPTH_TEST)
}

func UpdateDepth() {
	gl.Clear(gl.DEPTH_BUFFER_BIT)
}
