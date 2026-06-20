package td

type Vec2 struct {
	X float32
	Y float32
}

func NewVec2(x, y float32) Vec2 { return Vec2{X: x, Y: y} }

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func NewVec3(x, y, z float32) Vec3 { return Vec3{X: x, Y: y, Z: z} }
