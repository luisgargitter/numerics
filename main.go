package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
	p1 := Particle{mgl64.Vec3{0, 0, 0}, mgl64.Vec3{-0.001, 0, 0}, 1000000.0, 0.0}
	p2 := Particle{mgl64.Vec3{1, 0, 0}, mgl64.Vec3{0, 0.01, 0}, 10.0, 0.0}

	y := mgl64.NewVecN(7 * 2)

	VecNFromVec3(y, p1.Position, 0)
	VecNFromVec3(y, p1.Velocity, 3)
	y.Set(6, p1.Mass)

	VecNFromVec3(y, p2.Position, 7+0)
	VecNFromVec3(y, p2.Velocity, 7+3)
	y.Set(7+6, p2.Mass)

	rk4w := NewRK4Workspace(y.Size())

	for i := 0; i < 200; i += 1 {
		RK4(rk4w, Gravitation, 50, y, y)
		fmt.Printf("%f, %f, %f, %f, %f, %f\n",
			y.Get(0), y.Get(1), y.Get(2), y.Get(7), y.Get(8), y.Get(9))
	}
}
