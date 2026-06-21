package mesh

var (
	TriangleVertices = []float32{
		-0.5, -0.5, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0,
		0.0, 0.5, 0.0, 0.0, 0.0, 1.0,
	}
	CubeVertices = []float32{
		-1, -1, -1, 1, 0, 0,
		1, -1, -1, 1, 0, 0,
		1, 1, -1, 1, 0, 0,
		-1, 1, -1, 1, 0, 0,
		-1, -1, 1, 0, 1, 0,
		1, -1, 1, 0, 1, 0,
		1, 1, 1, 0, 1, 0,
		-1, 1, 1, 0, 1, 0,
	}
	CubeIndices = []uint32{
		0, 1, 2, 0, 2, 3,
		4, 5, 6, 4, 6, 7,
		0, 3, 7, 0, 7, 4,
		1, 5, 6, 1, 6, 2,
		0, 4, 5, 0, 5, 1,
		3, 2, 6, 3, 6, 7,
	}
	DefaultLayout = []Attribute{
		{Index: 0, Size: 3, Offset: 0},
		{Index: 1, Size: 3, Offset: 12},
	}
)

func GenRec(width, height, depth float32) ([]float32, []uint32) {
	hw, hh, hd := width/2, height/2, depth/2

	v := []float32{
		-hw, -hh, -hd, 1, 0, 0,
		hw, -hh, -hd, 1, 0, 0,
		hw, hh, -hd, 1, 0, 0,
		-hw, hh, -hd, 1, 0, 0,
		-hw, -hh, hd, 0, 1, 0,
		hw, -hh, hd, 0, 1, 0,
		hw, hh, hd, 0, 1, 0,
		-hw, hh, hd, 0, 1, 0,
	}

	i := []uint32{
		0, 1, 2, 0, 2, 3,
		4, 5, 6, 4, 6, 7,
		0, 3, 7, 0, 7, 4,
		1, 5, 6, 1, 6, 2,
		0, 4, 5, 0, 5, 1,
		3, 2, 6, 3, 6, 7,
	}
	return v, i
}

func GenPlane(width, depth float32, segments int) ([]float32, []uint32) {
	hw, hd := width/2, depth/2
	stepX := width / float32(segments)
	stepZ := depth / float32(segments)

	var verts []float32
	var indices []uint32

	for z := 0; z <= segments; z++ {
		for x := 0; x <= segments; x++ {
			px := -hw + float32(x)*stepX
			pz := -hd + float32(z)*stepZ
			height := float32(0.0)
			r := (px + hw) / width
			g := float32(0.5)
			b := (pz + hd) / depth

			verts = append(verts, px, height, pz, r, g, b)
		}
	}

	for z := 0; z < segments; z++ {
		for x := 0; x < segments; x++ {
			tl := uint32(z*(segments+1) + x)
			tr := tl + 1
			bl := tl + uint32(segments+1)
			br := bl + 1

			indices = append(indices, tl, bl, tr)
			indices = append(indices, tr, bl, br)
		}
	}

	return verts, indices
}
