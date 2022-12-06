package gomisc

type Vec2F [2]float64
type Vec2I [2]int

// Are `v` and `other` identical.
func (v Vec2F) Eq(other Vec2F) bool {
	return v[0] == other[0] && v[1] == other[1]
}

// Changes sign of each `v` element.
func (v Vec2F) Neg() Vec2F {
	return Vec2F{-v[0], -v[1]}
}

// Reciprocates each `v` element.
func (v Vec2F) Rcp() Vec2F {
	return Vec2F{1 / v[0], 1 / v[1]}
}

// `v` and `other` pairwise add.
func (v Vec2F) Add(other Vec2F) Vec2F {
	return Vec2F{v[0] + other[0], v[1] + other[1]}
}

// Add `other` to each `v` element.
func (v Vec2F) Add1(other float64) Vec2F {
	return Vec2F{v[0] + other, v[1] + other}
}

// `v` and `other` pairwise subtract.
func (v Vec2F) Sub(other Vec2F) Vec2F {
	return Vec2F{v[0] - other[0], v[1] - other[1]}
}

// Subtract `other` from each `v` element.
func (v Vec2F) Sub1(other float64) Vec2F {
	return Vec2F{v[0] - other, v[1] - other}
}

// `v` and `other` pairwise multiply.
func (v Vec2F) Mul(other Vec2F) Vec2F {
	return Vec2F{v[0] * other[0], v[1] * other[1]}
}

// Multiply `other` with each `v` element.
func (v Vec2F) Mul1(other float64) Vec2F {
	return Vec2F{v[0] * other, v[1] * other}
}

// `v` and `other` pairwise divide.
func (v Vec2F) Div(other Vec2F) Vec2F {
	return Vec2F{v[0] / other[0], v[1] / other[1]}
}

// Divide `other` from each `v` element.
func (v Vec2F) Div1(other float64) Vec2F {
	return Vec2F{v[0] / other, v[1] / other}
}

// `v` and `other` pairwise wrap.
func (v Vec2F) Wrap(lens Vec2F) Vec2F {
	return Vec2F{Wrap(v[0], lens[0]), Wrap(v[1], lens[1])}
}

// Wrap `len` to each `v` element.
func (v Vec2F) Wrap1(len float64) Vec2F {
	return Vec2F{Wrap(v[0], len), Wrap(v[1], len)}
}

// Make `v` elements absolute.
func (v Vec2F) Abs() Vec2F {
	return Vec2F{Abs(v[0]), Abs(v[1])}
}

// Lowest `v` element.
func (v Vec2F) Min() float64 {
	if v[1] < v[0] {
		return v[1]
	}
	return v[0]
}

// Highest `v` element.
func (v Vec2F) Max() float64 {
	if v[1] > v[0] {
		return v[1]
	}
	return v[0]
}

// `v` element Sum.
func (v Vec2F) Sum() float64 {
	return v[0] + v[1]
}

// Angle to direction.
func (a Rad) Vec2F() Vec2F {
	return Vec2F{a.Cos(), a.Sin()}
}

// Angle to direction.
func (a Deg) Vec2F() Vec2F {
	return Vec2F{a.Cos(), a.Sin()}
}

// Direction to angle.
func (v Vec2F) Rad() Rad {
	return Atan2(v[1], v[0])
}

// Direction to angle.
func (v Vec2F) Deg() Deg {
	return Atan2(v[1], v[0]).Deg()
}

// Floor `v` elements.
func (v Vec2F) Floor() Vec2F {
	return Vec2F{Floor(v[0]), Floor(v[1])}
}

// FloorI `v` elements.
func (v Vec2F) FloorI() Vec2I {
	return Vec2I{FloorI(v[0]), FloorI(v[1])}
}

// Round `v` elements.
func (v Vec2F) Round() Vec2F {
	return Vec2F{Round(v[0]), Round(v[1])}
}

// RoundI `v` elements.
func (v Vec2F) RoundI() Vec2I {
	return Vec2I{RoundI(v[0]), RoundI(v[1])}
}

// `v` and `other` linear interpolation.
func (v Vec2F) Lerp(other Vec2F, t float64) Vec2F {
	return other.Sub(v).Mul1(t).Add(v)
}

// Magnitude.
func (v Vec2F) Mag() float64 {
	return Sqrt(v[0]*v[0] + v[1]*v[1])
}

// `v` direction with `value` magnitude.
func (v Vec2F) MagSet(value float64) Vec2F {
	if mag := v.Mag(); mag != 0 {
		return v.Mul1(value / mag)
	}
	return v
}

// `v` direction with 1 magnitude.
func (v Vec2F) Norm() Vec2F {
	return v.MagSet(1)
}

// `v` and `other` dot product.
func (v Vec2F) Dot(other Vec2F) float64 {
	return v.Mul(other).Sum()
}

// Angle from `v` to `other`.
func (v Vec2F) AngTo(other Vec2F) Rad {
	return other.Sub(v).Rad()
}

// Clamps `v` magnitude.
func (v Vec2F) ClampMag(max float64) Vec2F {
	if v.Mag() > max {
		return v.MagSet(max)
	}
	return v
}

// Distance between `v` and `other`.
func (v Vec2F) Dst(other Vec2F) float64 {
	return other.Sub(v).Mag()
}

// Move `v` towards `other` by `dlt`.
// `dlt` must be >= 0.
func (v Vec2F) MoveTowards(other Vec2F, dlt float64) Vec2F {
	offset := other.Sub(v)
	dst := offset.Mag()
	if dst <= dlt {
		return other
	}
	return offset.Mul1(dlt / dst).Add(v)
}

// Project `other` onto `v`, changing magnitude of `v`.
func (v Vec2F) Project(other Vec2F) Vec2F {
	return v.MagSet(v.Dot(other))
}

// Rotate `v` 90 degrees.
func (v Vec2F) Rot90() Vec2F {
	return Vec2F{-v[1], v[0]}
}

// Rotate `v` with angle `amount`.
func (v Vec2F) RotRad(amount Rad) Vec2F {
	newX := amount.Vec2F()
	return newX.Rot90().Mul1(v[1]).Add(newX.Mul1(v[0]))
}

// Rotate `v` with angle `amount`.
func (v Vec2F) RotDeg(amount Deg) Vec2F {
	return v.RotRad(amount.Rad())
}

// Reflect `v` on `norm`.
// `norm` magnitude determines bounce magnitude.
func (v Vec2F) Reflect(norm Vec2F) Vec2F {
	return norm.Rot90().Norm().Project(v).Sub(norm.Project(v))
}

// Add `amount` to each point.
func TranslateVec2F(points []Vec2F, amount Vec2F) []Vec2F {
	result := make([]Vec2F, len(points))
	for i, p := range points {
		result[i] = p.Add(amount)
	}
	return result
}

// Rotate every point around origin by `amount`.
func RotateVec2FRad(points []Vec2F, amount Rad) []Vec2F {
	newX := amount.Vec2F()
	newY := newX.Rot90()
	result := make([]Vec2F, len(points))
	for i, p := range points {
		result[i] = newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
	}
	return result
}

// Rotate every point around origin by `amount`.
func RotateVec2FDeg(points []Vec2F, amount Deg) []Vec2F {
	newX := amount.Vec2F()
	newY := newX.Rot90()
	result := make([]Vec2F, len(points))
	for i, p := range points {
		result[i] = newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
	}
	return result
}

// Multiply `amount` with each point.
func ScaleVec2F(points []Vec2F, amount Vec2F) []Vec2F {
	result := make([]Vec2F, len(points))
	for i, p := range points {
		result[i] = p.Mul(amount)
	}
	return result
}

// Multiply `amount` with each point.
func Scale1Vec2F(points []Vec2F, amount float64) []Vec2F {
	result := make([]Vec2F, len(points))
	for i, p := range points {
		result[i] = p.Mul1(amount)
	}
	return result
}
