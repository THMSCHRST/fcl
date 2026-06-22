package mesh

import (
	"fcl/td"
)

type MeshBuilder struct {
	Vertices []float32
	Indices  []uint32
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
		Vertices: []float32{},
		Indices:  []uint32{},
		offset:   0,
	}
}

func (mb *MeshBuilder) AddTriangle(t Triangle) {
	col := td.ColToMGL32(t.Col)
	mb.Vertices = append(mb.Vertices,
		t.P1.X, t.P1.Y, t.P1.Z,
		col[0], col[1], col[2],
	)
	mb.Vertices = append(mb.Vertices,
		t.P2.X, t.P2.Y, t.P2.Z,
		col[0], col[1], col[2],
	)
	mb.Vertices = append(mb.Vertices,
		t.P3.X, t.P3.Y, t.P3.Z,
		col[0], col[1], col[2],
	)

	mb.Indices = append(mb.Indices, mb.offset, mb.offset+1, mb.offset+2)

	mb.offset += 3
}

func (mb *MeshBuilder) AddQuad(q Quad) {
	mb.AddTriangle(NewTriangle(q.P1, q.P3, q.P2, q.Col))
	mb.AddTriangle(NewTriangle(q.P2, q.P3, q.P4, q.Col))
}

func (mb *MeshBuilder) Build(layout []Attribute) (*Mesh, error) {
	/*
	newTriangles := []Triangle{}

    for i := 0; i < len(mb.Vertices); i += 18 {
        v1 := td.NewVec3(mb.Vertices[i], mb.Vertices[i+1], mb.Vertices[i+2])
        v2 := td.NewVec3(mb.Vertices[i+6], mb.Vertices[i+7], mb.Vertices[i+8])
        v3 := td.NewVec3(mb.Vertices[i+12], mb.Vertices[i+13], mb.Vertices[i+14])

		color := td.NewCol(
			mb.Vertices[i+3]*255,
			mb.Vertices[i+4]*255,
			mb.Vertices[i+5]*255,
		)

        newTriangles = append(newTriangles, NewTriangle(v1, v2, v3, color))
    }

    mb.Vertices = []float32{}
    mb.Indices = []uint32{}
    mb.offset = 0

    for _, t := range newTriangles {
        mb.AddTriangle(t)
    }*/

    return NewMesh(mb.Vertices, mb.Indices, layout)
}

func (mb *MeshBuilder) Clear() {
	mb.Vertices = mb.Vertices[:0]
	mb.Indices = mb.Indices[:0]
	mb.offset = 0
}

func (mb *MeshBuilder) VertexCount() int {
	return int(mb.offset)
}

func (mb *MeshBuilder) TriangleCount() int {
	return len(mb.Indices) / 3
}
