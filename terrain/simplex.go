package terrain

import (
	"math/rand"
)

type Noise struct {
	perm [512]uint8
	grad [512][2]float32
}

func New(seed int64) *Noise {
	n := &Noise{}
	n.init(seed)
	return n
}

const (
	f2 = 0.36602542
	g2 = 0.21132487
)

func (n *Noise) init(seed int64) {
	rng := rand.New(rand.NewSource(seed))

	perm := make([]uint8, 256)
	for i := 0; i < 256; i++ {
		perm[i] = uint8(i)
	}
	for i := 255; i > 0; i-- {
		j := int(rng.Int31n(int32(i + 1)))
		perm[i], perm[j] = perm[j], perm[i]
	}

	for i := 0; i < 512; i++ {
		n.perm[i] = perm[i&255]
	}

	var g2d = [12]uint16{
		0x0101, 0xff01, 0x01ff, 0xffff,
		0x0100, 0xff00, 0x0100, 0xff00,
		0x0001, 0x00ff, 0x0001, 0x00ff,
	}
	for i := 0; i < 512; i++ {
		idx := g2d[n.perm[i]%12]
		gx := int8(idx >> 8)
		gy := int8(idx)
		n.grad[i] = [2]float32{float32(gx), float32(gy)}
	}
}

// Noise2 computes a two dimensional simplex noise
// Public Domain: https://weber.itn.liu.se/~stegu/simplexnoise/simplexnoise.pdf
// Reference: https://mrl.cs.nyu.edu/~perlin/noise/
func (n *Noise) Noise2(x, y float32) float32 {
	s := (x + y) * f2
	i := floor(x + s)
	j := floor(y + s)
	t := float32(i+j) * g2
	x0 := x - (float32(i) - t)
	y0 := y - (float32(j) - t)

	var i1, j1 float32 = 0, 1
	if x0 > y0 {
		i1, j1 = 1, 0
	}
	x1 := x0 - i1 + g2
	y1 := y0 - j1 + g2
	x2 := x0 + 2*g2 - 1
	y2 := y0 + 2*g2 - 1

	// Use n.perm and n.grad
	pp := n.perm[j&255:]
	gg := n.grad[i&255:]
	p0 := int(pp[0])
	p1 := int(pp[int(j1)])
	p2 := int(pp[1])
	g0 := gg[p0]
	g1 := gg[int(i1)+p1]
	g2g := gg[1+p2]

	n0 := float32(0.0)
	if t := 0.5 - x0*x0 - y0*y0; t > 0 {
		n0 += pow4(t) * (g0[0]*x0 + g0[1]*y0)
	}
	if t := 0.5 - x1*x1 - y1*y1; t > 0 {
		n0 += pow4(t) * (g1[0]*x1 + g1[1]*y1)
	}
	if t := 0.5 - x2*x2 - y2*y2; t > 0 {
		n0 += pow4(t) * (g2g[0]*x2 + g2g[1]*y2)
	}
	return 70.0 * n0
}

// pow4 lifts the value to the power of 4
func pow4(v float32) float32 {
	v *= v
	return v * v
}

// floor floors the floating-point value to an integer
func floor(x float32) int {
	v := int(x)
	if x < float32(v) {
		return v - 1
	}
	return v
}