package render

import (
	"fmt" // <-- Added for proper error handling

	"github.com/go-gl/gl/v4.6-core/gl"
)

// Attribute describes a single component of a vertex (e.g., Position, Color).
// Index: The 'location' in the shader (layout (location = X)).
// Size: Number of float32 components (1, 2, 3, or 4).
// Offset: Byte offset from the start of the vertex.
type Attribute struct {
    Index  uint32
    Size   int32
    Offset int
}

// Mesh holds the geometry data on the GPU.
type Mesh struct {
    VAO         uint32
    VBO         uint32
    EBO         uint32
    vertexCount int32
    indexCount  int32
    isIndexed   bool
}

func NewMesh(vertices []float32, indices []uint32, attribs []Attribute) (*Mesh, error) {
    if len(vertices) == 0 {
        return nil, fmt.Errorf("vertices slice cannot be empty")
    }
    if len(attribs) == 0 {
        return nil, fmt.Errorf("attribute layout cannot be empty")
    }

    lastAttr := attribs[len(attribs)-1]
    stride := int32(lastAttr.Offset) + lastAttr.Size*4 // 4 bytes per float32

    if stride == 0 {
        return nil, fmt.Errorf("calculated stride is 0, check your attribute offsets")
    }

    floatsPerVertex := stride / 4
    vertexCount := int32(len(vertices)) / floatsPerVertex

    if int32(len(vertices))%floatsPerVertex != 0 {
        return nil, fmt.Errorf("vertex data length (%d) is not a multiple of floats per vertex (%d)", len(vertices), floatsPerVertex)
    }

    var vao, vbo, ebo uint32

    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)

    // upload vbo
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

    // use layout
    for _, attr := range attribs {
        gl.VertexAttribPointerWithOffset(attr.Index, attr.Size, gl.FLOAT, false, stride, uintptr(attr.Offset))
        gl.EnableVertexAttribArray(attr.Index)
    }

    // handle idx buffer
    isIndexed := len(indices) > 0
    var indexCount int32
    if isIndexed {
        gl.GenBuffers(1, &ebo)
        gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
        gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
        indexCount = int32(len(indices))
    }

    // unbind vao-
    gl.BindVertexArray(0)

    return &Mesh{
        VAO:         vao,
        VBO:         vbo,
        EBO:         ebo,
        vertexCount: vertexCount,
        indexCount:  indexCount,
        isIndexed:   isIndexed,
    }, nil
}

// Draw renders the mesh using the currently active program.
func (m *Mesh) Draw() {
    gl.BindVertexArray(m.VAO)
    if m.isIndexed {
        gl.DrawElements(gl.TRIANGLES, m.indexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
    } else {
        gl.DrawArrays(gl.TRIANGLES, 0, m.vertexCount)
    }
    gl.BindVertexArray(0)
}

// Destroy frees all GPU resources with the mesh.
func (m *Mesh) Destroy() {
    gl.DeleteVertexArrays(1, &m.VAO)
    gl.DeleteBuffers(1, &m.VBO)
    if m.isIndexed {
        gl.DeleteBuffers(1, &m.EBO)
    }
}