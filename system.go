package numerics

import (
	"github.com/go-gl/mathgl/mgl64"
)

type System func(x *mgl64.VecN, y *mgl64.VecN)
