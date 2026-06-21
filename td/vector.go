package td

var (
	Vec2Zero = NewVec2(0, 0)
	Vec3Zero = NewVec3(0, 0, 0)
)

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

func (v1 Vec3) Add(v2 Vec3) Vec3 { return NewVec3(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z) }
