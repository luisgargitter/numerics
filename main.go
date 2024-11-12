package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
	p1 := Particle{mgl64.Vec3{0, 0, 0}, mgl64.Vec3{-0.0001, 0, 0}, 10000.0, 0.0}
	p2 := Particle{mgl64.Vec3{1, 0, 0}, mgl64.Vec3{0, 0.001, 0}, 10.0, 0.0}

	y := ParticlesToVecN([]Particle{p1, p2})

	rk4w := NewRK4Workspace(y.Size())

	for i := 0; i < 200; i += 1 {
		RK4(rk4w, dParticleSystem, 50, y, y)
		objs := VecNToParticles(y)
		p1 := objs[0].Position
		p2 := objs[1].Position
		fmt.Printf("%f, %f, %f, %f, %f, %f\n", p1[0], p1[1], p1[2], p2[0], p2[1], p2[2])
	}
}
