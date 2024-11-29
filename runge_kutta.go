package numerics

type VecAdd[T any] func(d *T, a *T, b *T) *T
type VecMul[T any] func(d *T, a *T, c float64) *T

// Runge Kutta 4
type RK4Workspace[T any] struct {
	add VecAdd[T]
	mul VecMul[T]
	d   T
	k1  T
	k2  T
	k3  T
	k4  T
}

func RK4[T any](w *RK4Workspace[T], f System[T], dt float64, y0 *T, y *T) {
	f(y0, &w.k1)
	f(w.add(&w.d, y0, w.mul(&w.d, &w.k1, dt/2.0)), &w.k2)
	f(w.add(&w.d, y0, w.mul(&w.d, &w.k2, dt/2.0)), &w.k3)
	f(w.add(&w.d, y0, w.mul(&w.d, &w.k3, dt)), &w.k4)

	// y = y0 + (h/6)(k1 + 2*k2 + 2*k3 + k4)
	y = w.add(y, y0, w.mul(&w.d, w.add(&w.d, w.mul(&w.d, w.add(&w.d, &w.k2, &w.k3), 2.0), w.add(&w.d, &w.k1, &w.k4)), dt/6.0))
}
