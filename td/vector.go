package td

import "github.com/go-gl/mathgl/mgl32"

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

func Vec3Mgl32(v Vec3) mgl32.Vec3 {
    return mgl32.Vec3{v.X, v.Y, v.Z}
}

func Mgl32Vec3(v mgl32.Vec3) Vec3 {
    return Vec3{v.X(), v.Y(), v.Z()}
}