package mesh

import (
	"fcl/td"
)

type MeshBuilder struct {
	vertices []float32
	indices  []uint32
	offset   uint32
}

type Triangle struct {
	P1  td.Vec3
	P2  td.Vec3
	P3  td.Vec3
	Col td.Col
}

func NewTriangle(p1, p2, p3 td.Vec3, col td.Col) Triangle {
	return Triangle{p1, p2, p3, col}
}

type Quad struct {
	P1  td.Vec3
	P2  td.Vec3
	P3  td.Vec3
	P4  td.Vec3
	Col td.Col
}

func NewQuad(p1, p2, p3, p4 td.Vec3, col td.Col) Quad {
	return Quad{p1, p2, p3, p4, col}
}

func NewMeshBuilder() *MeshBuilder {
	return &MeshBuilder{
		vertices: []float32{},
		indices:  []uint32{},
		offset:   0,
	}
}

func (mb *MeshBuilder) AddTriangle(t Triangle) {
	col := td.ColToMGL32(t.Col)
	mb.vertices = append(mb.vertices,
		t.P1.X, t.P1.Y, t.P1.Z,
		col[0], col[1], col[2],
	)
	mb.vertices = append(mb.vertices,
		t.P2.X, t.P2.Y, t.P2.Z,
		col[0], col[1], col[2],
	)
	mb.vertices = append(mb.vertices,
		t.P3.X, t.P3.Y, t.P3.Z,
		col[0], col[1], col[2],
	)

	mb.indices = append(mb.indices, mb.offset, mb.offset+1, mb.offset+2)

	mb.offset += 3
}

func (mb *MeshBuilder) AddQuad(q Quad) {
	mb.AddTriangle(NewTriangle(q.P1, q.P3, q.P2, q.Col))
	mb.AddTriangle(NewTriangle(q.P2, q.P3, q.P4, q.Col))
}

func (mb *MeshBuilder) Build(layout []Attribute) (*Mesh, error) {
	return NewMesh(mb.vertices, mb.indices, layout)
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
