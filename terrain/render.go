package terrain

import (
	"fcl/mesh"
	"fcl/td"
)

func (hm Heightmap) GenTriangles(scale td.Vec3, offset td.Vec3, lod int) []mesh.Triangle {
    if lod < 1 {
        lod = 1
    }

    tri := []mesh.Triangle{}
    w := len(hm)
    if w == 0 {
        return nil
    }
    h := len(hm[0])
    if h == 0 {
        return nil
    }

    // Build world‑space vertex grid
    verts := make([][]td.Vec3, w)
    for x := 0; x < w; x++ {
        col := make([]td.Vec3, h)
        row := hm[x]
        worldX := float32(x)*scale.X + offset.X
        for y := 0; y < h; y++ {
            worldY := float32(row[y])*scale.Y + offset.Y
            worldZ := float32(y)*scale.Z + offset.Z
            col[y] = td.NewVec3(worldX, worldY, worldZ)
        }
        verts[x] = col
    }

    rangeH := 1
    invRange := float32(1.0 / float64(rangeH))

    // Terrain triangles (unchanged)
    for x := 0; x < w-1; {
        nextX := x + lod
        if nextX > w-1 {
            nextX = w - 1
        }
        for y := 0; y < h-1; {
            nextY := y + lod
            if nextY > h-1 {
                nextY = h - 1
            }
            h00 := float32(hm[x][y])
            h10 := float32(hm[nextX][y])
            h01 := float32(hm[x][nextY])
            h11 := float32(hm[nextX][nextY])
            avg := (h00 + h10 + h01 + h11) * 0.25
            norm := avg * invRange
            color := heightColor(norm)
            p00 := verts[x][y]
            p01 := verts[x][nextY]
            p10 := verts[nextX][y]
            p11 := verts[nextX][nextY]
            tri = append(tri, mesh.NewTriangle(p00, p01, p10, color))
            tri = append(tri, mesh.NewTriangle(p10, p01, p11, color))
            y = nextY
        }
        x = nextX
    }

    skirtColor := heightColor(0.05)
    skirtY := offset.Y - 1000.0

    for x := 0; x < w-1; {
        nextX := x + lod
        if nextX > w-1 {
            nextX = w - 1
        }
        p1 := verts[x][0]
        p2 := verts[nextX][0]
        b1 := td.NewVec3(p1.X, skirtY, p1.Z)
        b2 := td.NewVec3(p2.X, skirtY, p2.Z)
        tri = append(tri, mesh.NewTriangle(p1, p2, b1, skirtColor))
        tri = append(tri, mesh.NewTriangle(p2, b2, b1, skirtColor))
        x = nextX
    }
    for x := 0; x < w-1; {
        nextX := x + lod
        if nextX > w-1 {
            nextX = w - 1
        }
        p1 := verts[x][h-1]
        p2 := verts[nextX][h-1]
        b1 := td.NewVec3(p1.X, skirtY, p1.Z)
        b2 := td.NewVec3(p2.X, skirtY, p2.Z)
        tri = append(tri, mesh.NewTriangle(p1, b1, p2, skirtColor))
        tri = append(tri, mesh.NewTriangle(p2, b1, b2, skirtColor))
        x = nextX
    }
    for y := 0; y < h-1; {
        nextY := y + lod
        if nextY > h-1 {
            nextY = h - 1
        }
        p1 := verts[0][y]
        p2 := verts[0][nextY]
        b1 := td.NewVec3(p1.X, skirtY, p1.Z)
        b2 := td.NewVec3(p2.X, skirtY, p2.Z)
        tri = append(tri, mesh.NewTriangle(p1, b1, p2, skirtColor))
        tri = append(tri, mesh.NewTriangle(p2, b1, b2, skirtColor))
        y = nextY
    }
    for y := 0; y < h-1; {
        nextY := y + lod
        if nextY > h-1 {
            nextY = h - 1
        }
        p1 := verts[w-1][y]
        p2 := verts[w-1][nextY]
        b1 := td.NewVec3(p1.X, skirtY, p1.Z)
        b2 := td.NewVec3(p2.X, skirtY, p2.Z)
        tri = append(tri, mesh.NewTriangle(p1, p2, b1, skirtColor))
        tri = append(tri, mesh.NewTriangle(p2, b2, b1, skirtColor))
        y = nextY
    }

    return tri
}

func heightColor(h float32) td.Col {
	if h < -1 {
		h = -1
	}
	if h > 1 {
		h = 1
	}
	t := (h + 1) / 2
	r := float32(255 * t)
	g := float32(255 * (1 - t*t))
	b := float32(255 * (1 - t))
	return td.NewCol(r, g, b)
}