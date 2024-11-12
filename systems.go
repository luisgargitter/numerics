package main

import (
	"github.com/go-gl/mathgl/mgl64"
)

func Vec3FromVecN(vn *mgl64.VecN, i int) mgl64.Vec3 {
	raw := vn.Raw()
	return mgl64.Vec3{raw[i+0], raw[i+1], raw[i+2]}
}

func VecNFromVec3(vn *mgl64.VecN, v mgl64.Vec3, i int) {
	raw := vn.Raw()
	raw[i+0], raw[i+1], raw[i+2] = v[0], v[1], v[2]
}

func Gravitation(y *mgl64.VecN, dy *mgl64.VecN) {
	const offsetPos = 0
	const offsetVel = 3
	const offsetMass = 6
	const stride = offsetMass + 1

	dy.Zero(dy.Size())
	for i := 0; i < y.Size(); i += stride {
		p1 := Particle{Vec3FromVecN(y, i+offsetPos), Vec3FromVecN(y, i+offsetVel), y.Get(i + offsetMass), 0.0}

		for j := i + stride; j < y.Size(); j += stride {
			p2 := Particle{Vec3FromVecN(y, j+offsetPos), Vec3FromVecN(y, j+offsetVel), y.Get(j + offsetMass), 0.0}

			f := p1.ForceV(&p2)
			fp1 := f.Mul(1.0 / p1.Mass)
			fp2 := f.Mul(-1.0 / p2.Mass)

			VecNFromVec3(dy, fp1, i+offsetVel)
			VecNFromVec3(dy, fp2, j+offsetVel)
		}
		// change in Position
		VecNFromVec3(dy, Vec3FromVecN(y, i+offsetVel), i+offsetPos)
	}
}
