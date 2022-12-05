package gomisc

type Vec2F struct {
	x, y float64
}
type Vec2I struct {
	x, y int
}

// Are `v` and `other` identical
func (v Vec2F) Eq(other Vec2F) bool {
	return v.x == other.x && v.y == other.y
}

// `v` and `other` pairwise add
func (v Vec2F) Add(other Vec2F) Vec2F {
	return Vec2F{v.x + other.x, v.y + other.y}
}

// Add `other` to each `v` element
func (v Vec2F) Add1(other float64) Vec2F {
	return Vec2F{v.x + other, v.y + other}
}

// `v` and `other` pairwise subtract
func (v Vec2F) Sub(other Vec2F) Vec2F {
	return Vec2F{v.x - other.x, v.y - other.y}
}

// Subtract `other` from each `v` element
func (v Vec2F) Sub1(other float64) Vec2F {
	return Vec2F{v.x - other, v.y - other}
}

// `v` and `other` pairwise multiply
func (v Vec2F) Mul(other Vec2F) Vec2F {
	return Vec2F{v.x * other.x, v.y * other.y}
}

// Multiply `other` with each `v` element
func (v Vec2F) Mul1(other float64) Vec2F {
	return Vec2F{v.x * other, v.y * other}
}

// `v` and `other` pairwise divide
func (v Vec2F) Div(other Vec2F) Vec2F {
	return Vec2F{v.x / other.x, v.y / other.y}
}

// Divide `other` from each `v` element
func (v Vec2F) Div1(other float64) Vec2F {
	return Vec2F{v.x / other, v.y / other}
}

// `v` and `other` pairwise wrap
func (v Vec2F) Wrap(lens Vec2F) Vec2F {
	return Vec2F{Wrap(v.x, lens.x), Wrap(v.y, lens.y)}
}

// Wrap `len` to each `v` element
func (v Vec2F) Wrap1(len float64) Vec2F {
	return Vec2F{Wrap(v.x, len), Wrap(v.y, len)}
}

// Make `v` elements absolute
func (v Vec2F) Abs() Vec2F {
	return Vec2F{Abs(v.x), Abs(v.y)}
}

// Lowest `v` element
func (v Vec2F) Min() float64 {
	if v.y < v.x {
		return v.y
	}
	return v.x
}

// Highest `v` element
func (v Vec2F) Max() float64 {
	if v.y > v.x {
		return v.y
	}
	return v.x
}

// `v` element Sum
func (v Vec2F) Sum() float64 {
	return v.x + v.y
}

// Angle to direction
func (a Rad64) Vec2F() Vec2F {
	return Vec2F{a.Cos(), a.Sin()}
}

// Angle to direction
func (a Deg64) Vec2F() Vec2F {
	return Vec2F{a.Cos(), a.Sin()}
}

// Direction to angle
func (v Vec2F) Rad() Rad64 {
	return Atan2(v.y, v.x)
}

// Direction to angle
func (v Vec2F) Deg() Deg64 {
	return Atan2(v.y, v.x).Deg()
}

// Floor `v` elements
func (v Vec2F) Floor() Vec2F {
	return Vec2F{Floor(v.x), Floor(v.y)}
}

// FloorI `v` elements
func (v Vec2F) FloorI() Vec2I {
	return Vec2I{FloorI(v.x), FloorI(v.y)}
}

// Round `v` elements
func (v Vec2F) Round() Vec2F {
	return Vec2F{Round(v.x), Round(v.y)}
}

// RoundI `v` elements
func (v Vec2F) RoundI() Vec2I {
	return Vec2I{RoundI(v.x), RoundI(v.y)}
}

// `v` and `other` linear interpolation
func (v Vec2F) Lerp(other Vec2F, t float64) Vec2F {
	return other.Sub(v).Mul1(t).Add(v)
}

// Magnitude
func (v Vec2F) Mag() float64 {
	return Sqrt(v.x*v.x + v.y*v.y)
}

// `v` direction with `value` magnitude
func (v Vec2F) MagSet(value float64) Vec2F {
	if mag := v.Mag(); mag != 0 {
		fix := value / mag
		return Vec2F{v.x * fix, v.y * fix}
	}
	return v
}

// `v` direction with 1 magnitude
func (v Vec2F) Norm() Vec2F {
	return v.MagSet(1)
}

// `v` and `other` dot product
func (v Vec2F) Dot(other Vec2F) float64 {
	return v.Mul(other).Sum()
}

// Angle from `v` to `other`
func (v Vec2F) AngTo(other Vec2F) Rad64 {
	return other.Sub(v).Rad()
}

// Clamps `v` magnitude
func (v Vec2F) ClampMag(max float64) Vec2F {
	if v.Mag() > max {
		return v.MagSet(max)
	}
	return v
}

// Distance between `v` and `other`
func (v Vec2F) Dst(other Vec2F) float64 {
	return v.Sub(other).Mag()
}

// Move `v` towards `other` by `dlt`
func (v Vec2F) MoveTowards(other Vec2F, dlt float64) Vec2F {
	if v.Dst(other) <= dlt {
		return other
	}
	return other.Sub(v).MagSet(dlt).Add(v)
}

// Project `other` onto `v`, changing magnitude of `v`
func (v Vec2F) Project(other Vec2F) Vec2F {
	return v.MagSet(v.Dot(other))
}

// Rotate `v` 90 degrees
func (v Vec2F) Rot90() Vec2F {
	return Vec2F{-v.y, v.x}
}

// Reflect `v` on `norm`
func (v Vec2F) Reflect(norm Vec2F) Vec2F {
	return norm.Rot90().Norm().Project(v).Add(norm.Project(v))
}

// Add `amount` to each point
func TranslateVec2F(points []Vec2F, amount Vec2F) []Vec2F {
	return MapF(points, func(p Vec2F) Vec2F { return p.Add(amount) })
}

// Rotate every point around origin by `amount`
func RotateVec2FRad(points []Vec2F, amount Rad64) []Vec2F {
	newX := amount.Vec2F()
	newY := newX.Rot90()
	return MapF(points, func(p Vec2F) Vec2F {
		return newX.Mul1(p.x).Add(newY.Mul1(p.y))
	})
}

// Rotate every point around origin by `amount`
func RotateVec2FDeg(points []Vec2F, amount Deg64) []Vec2F {
	newX := amount.Vec2F()
	newY := newX.Rot90()
	return MapF(points, func(p Vec2F) Vec2F {
		return newX.Mul1(p.x).Add(newY.Mul1(p.y))
	})
}

// Multiply `amount` with each point
func ScaleVec2F(points []Vec2F, amount Vec2F) []Vec2F {
	return MapF(points, func(p Vec2F) Vec2F { return p.Mul(amount) })
}
