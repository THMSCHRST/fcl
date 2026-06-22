package terrain

func GenHeightmap(w,h,oX,oY,d int, s int64) Heightmap {
	n := New(s)

	hm := make(Heightmap,w)
	for x := range w {
		hm[x] = make([]float64, h)
		for y := range h {
			hm[x][y] = float64(n.Noise2(float32(x+oX)/float32(d),float32(y+oY)/float32(d)))
		}
	}
	return hm
}

func StaticHeightmap(w,h,v int) Heightmap {
	hm := make(Heightmap,w)
	for x := range w {
		hm[x] = make([]float64, h)
		for y := range h {
			hm[x][y] = float64(v)
		}
	}
	return hm
}