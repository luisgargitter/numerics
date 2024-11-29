package numerics

type Vector interface {
	Add(*Vector, Vector) Vector
	Mul(*Vector, float64) Vector
}
