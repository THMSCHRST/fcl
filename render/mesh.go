package render

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

// Mesh holds the geometry data on the GPU.
type Mesh struct {
	VAO         uint32
	VBO         uint32
	EBO         uint32
	vertexCount int32
	indexCount  int32
	isIndexed   bool
}

// NewMesh creates a mesh from a slice of vertices and optional slice of indices.
// layout: position vec3 + color vec3 = 6 floats in 1 vertex.
func NewMesh(vertices []float32, indices []uint32) (*Mesh, error) {
	var vao, vbo, ebo uint32

	// create vao
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// create vbo -> upload verticeS
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// set vertex attr ptr
	// position (location = 0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 6*4, 0)
	gl.EnableVertexAttribArray(0)
	// color (location = 1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)
	gl.EnableVertexAttribArray(1)

	isIndexed := len(indices) > 0
	var indexCount int32
	if isIndexed {
		gl.GenBuffers(1, &ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
		indexCount = int32(len(indices))
	}

	// unbind vao
	gl.BindVertexArray(0)

	return &Mesh{
		VAO:         vao,
		VBO:         vbo,
		EBO:         ebo,
		vertexCount: int32(len(vertices) / 6), // 6 float per vertex
		indexCount:  indexCount,
		isIndexed:   isIndexed,
	}, nil
}

// Draw renders the mesh using the active program.
func (m *Mesh) Draw() {
	gl.BindVertexArray(m.VAO)
	if m.isIndexed {
		gl.DrawElementsWithOffset(gl.TRIANGLES, m.indexCount, gl.UNSIGNED_INT, 0)
	} else {
		// fallback no index
		gl.DrawArrays(gl.TRIANGLES, 0, m.vertexCount)
	}
	gl.BindVertexArray(0)
}

// Destroy frees GPU resources.
func (m *Mesh) Destroy() {
	gl.DeleteVertexArrays(1, &m.VAO)
	gl.DeleteBuffers(1, &m.VBO)
	if m.isIndexed {
		gl.DeleteBuffers(1, &m.EBO)
	}
}
