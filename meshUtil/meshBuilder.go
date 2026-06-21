package meshUtil

import (
	"fcl/render"

	"github.com/go-gl/mathgl/mgl32"
)

type MeshBuilder struct {
    vertices []float32
    indices  []uint32
    offset   uint32
}

func NewMeshBuilder() *MeshBuilder {
    return &MeshBuilder{
        vertices: []float32{},
        indices:  []uint32{},
        offset:   0,
    }
}

func (mb *MeshBuilder) AddTriangle(v1, v2, v3, color mgl32.Vec3) {
    mb.vertices = append(mb.vertices,
        v1[0], v1[1], v1[2],
        color[0], color[1], color[2],
    )
    mb.vertices = append(mb.vertices,
        v2[0], v2[1], v2[2],
        color[0], color[1], color[2],
    )
    mb.vertices = append(mb.vertices,
        v3[0], v3[1], v3[2],
        color[0], color[1], color[2],
    )

    mb.indices = append(mb.indices, mb.offset, mb.offset+1, mb.offset+2)

    mb.offset += 3
}

func (mb *MeshBuilder) AddQuad(v1, v2, v3, v4 mgl32.Vec3, color mgl32.Vec3) {
    // Triangle 1: v1, v3, v2
    mb.AddTriangle(v1, v3, v2, color)
    // Triangle 2: v2, v3, v4
    mb.AddTriangle(v2, v3, v4, color)
}

func (mb *MeshBuilder) Build(layout []render.Attribute) (*render.Mesh, error) {
    return render.NewMesh(mb.vertices, mb.indices, layout)
}

func (mb *MeshBuilder) Clear() {
    mb.vertices = mb.vertices[:0]
    mb.indices = mb.indices[:0]
    mb.offset = 0
}

func (mb *MeshBuilder) VertexCount() int {
    return int(mb.offset)
}

func (mb *MeshBuilder) TriangleCount() int {
    return len(mb.indices) / 3
}