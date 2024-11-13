package numerics

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Runge Kutta 4
type RK4Workspace struct {
	d  *mgl64.VecN
	k1 *mgl64.VecN
	k2 *mgl64.VecN
	k3 *mgl64.VecN
	k4 *mgl64.VecN
}

func NewRK4Workspace(n int) *RK4Workspace {
	rk4w := RK4Workspace{
		mgl64.NewVecN(n),
		mgl64.NewVecN(n),
		mgl64.NewVecN(n),
		mgl64.NewVecN(n),
		mgl64.NewVecN(n),
	}
	return &rk4w
}

func RK4(w *RK4Workspace, f System, dt float64, y0 *mgl64.VecN, y *mgl64.VecN) {
	d, k1, k2, k3, k4 := w.d, w.k1, w.k2, w.k3, w.k4
	f(y0, k1)

	k1.Mul(d, dt/2.0)
	y0.Add(d, d)
	f(d, k2)

	k2.Mul(d, dt/2.0)
	y0.Add(d, d)
	f(d, k3)

	k3.Mul(d, dt)
	y0.Add(d, d)
	f(d, k4)

	// y = y0 + (h/6)(k1 + 2*k2 + 2*k3 + k4)
	k2.Add(d, k3)
	d.Mul(d, 2.0)
	d.Add(d, k1)
	d.Add(d, k4)
	d.Mul(d, dt/6.0)
	y0.Add(y, d)
}
