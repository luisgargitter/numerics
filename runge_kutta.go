package numerics

// Runge Kutta 4
type RK4Workspace[V any] struct {
	d  V
	k1 V
	k2 V
	k3 V
	k4 V
}

func RK4(w *RK4Workspace[Vector], f System[Vector], dt float64, y0 Vector, y *Vector) {
	f(y0, &w.k1)
	f(y0.Add(&w.d, w.k1.Mul(&w.d, dt/2.0)), &w.k2)
	f(y0.Add(&w.d, w.k2.Mul(&w.d, dt/2.0)), &w.k3)
	f(y0.Add(&w.d, w.k3.Mul(&w.d, dt)), &w.k4)

	// y = y0 + (h/6)(k1 + 2*k2 + 2*k3 + k4)
	*y = y0.Add(y, w.k2.Add(&w.d, w.k3).Mul(&w.d, 2.0).Add(&w.d, w.k1).Add(&w.d, w.k4).Mul(&w.d, dt/6.0))
}
