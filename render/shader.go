package render

/*
#include <stdlib.h>
*/
import "C"

import (
	"fcl/helper"
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Program
type Program struct {
	ID uint32
}

func compileShader(src string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	// allocate c mem directly
	cSrc := (*uint8)(unsafe.Pointer(C.CString(src + "\x00")))
	defer C.free(unsafe.Pointer(cSrc)) // finally

	gl.ShaderSource(shader, 1, &cSrc, nil)
	gl.CompileShader(shader)

	// check compile err and handle them
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		// allocate c buffer for log too
		cLog := (*uint8)(C.malloc(C.size_t(logLength + 1)))
		defer C.free(unsafe.Pointer(cLog))

		gl.GetShaderInfoLog(shader, logLength, nil, cLog)

		// c string -> go string
		logStr := C.GoString((*C.char)(unsafe.Pointer(cLog)))
		return 0, fmt.Errorf("shader compilation error: %s", logStr)
	}

	return shader, nil
}

// NewProgram returns a new program.
func NewProgram(vertexSrc, fragmentSrc string) (*Program, error) {
	// compile vertex shader
	vertexShader, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	helper.Check(err)

	defer gl.DeleteShader(vertexShader)

	// compile fragment
	fragmentShader, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	helper.Check(err)

	defer gl.DeleteShader(fragmentShader)

	// create a program and attach shaders
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	// link program
	gl.LinkProgram(program)

	// check and handle link errors
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status) // get the link status
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log+"\x00"))
		return nil, fmt.Errorf("program linking error: %s", log)
	}

	return &Program{ID: program}, nil
}

// Use inits the program for rendering.
func (p *Program) Use() {
	gl.UseProgram(p.ID)
}

// Destroy frees the program.
func (p *Program) Destroy() {
	gl.DeleteProgram(p.ID)
}

// SetUniformMat4 sets a 4x4 matrix uniform by name.
func (p *Program) SetUniformMat4(name string, mat mgl32.Mat4) {
	cName := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(p.ID, cName)
	gl.UniformMatrix4fv(location, 1, false, &mat[0])
}
