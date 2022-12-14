package gomisc

type Vector2 [2]f64

// New Vector2.
func Vec2(x, y f64) Vector2 {
	return Vector2{x, y}
}

// Are `v` and `other` identical.
func (v Vector2) Eq(other Vector2) bool {
	return v[0] == other[0] && v[1] == other[1]
}

// Changes sign of each `v` element.
func (v Vector2) Neg() Vector2 {
	return Vec2(-v[0], -v[1])
}

// Reciprocates each `v` element.
func (v Vector2) Rcp() Vector2 {
	return Vec2(1/v[0], 1/v[1])
}

// `v` and `other` pairwise add.
func (v Vector2) Add(other Vector2) Vector2 {
	return Vec2(v[0]+other[0], v[1]+other[1])
}

// Add `other` to each `v` element.
func (v Vector2) Add1(other f64) Vector2 {
	return Vec2(v[0]+other, v[1]+other)
}

// `v` and `other` pairwise subtract.
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vec2(v[0]-other[0], v[1]-other[1])
}

// Subtract `other` from each `v` element.
func (v Vector2) Sub1(other f64) Vector2 {
	return Vec2(v[0]-other, v[1]-other)
}

// `v` and `other` pairwise multiply.
func (v Vector2) Mul(other Vector2) Vector2 {
	return Vec2(v[0]*other[0], v[1]*other[1])
}

// Multiply `other` with each `v` element.
func (v Vector2) Mul1(other f64) Vector2 {
	return Vec2(v[0]*other, v[1]*other)
}

// `v` and `other` pairwise divide.
func (v Vector2) Div(other Vector2) Vector2 {
	return Vec2(v[0]/other[0], v[1]/other[1])
}

// Divide `other` from each `v` element.
func (v Vector2) Div1(other f64) Vector2 {
	return Vec2(v[0]/other, v[1]/other)
}

// `v` and `other` pairwise wrap.
func (v Vector2) Wrap(lens Vector2) Vector2 {
	return Vec2(Wrap(v[0], lens[0]), Wrap(v[1], lens[1]))
}

// Wrap `len` to each `v` element.
func (v Vector2) Wrap1(len f64) Vector2 {
	return Vec2(Wrap(v[0], len), Wrap(v[1], len))
}

// Make `v` elements absolute.
func (v Vector2) Abs() Vector2 {
	return Vec2(Abs(v[0]), Abs(v[1]))
}

// Lowest `v` element.
func (v Vector2) Min() f64 {
	if v[1] < v[0] {
		return v[1]
	}
	return v[0]
}

// Highest `v` element.
func (v Vector2) Max() f64 {
	if v[1] > v[0] {
		return v[1]
	}
	return v[0]
}

// `v` element Su
func (v Vector2) Sum() f64 {
	return v[0] + v[1]
}

// Sign `v` elements.
func (v Vector2) Sign() Vector2 {
	return Vec2(Sign(v[0]), Sign(v[1]))
}

// Floor `v` elements.
func (v Vector2) Floor() Vector2 {
	return Vec2(Floor(v[0]), Floor(v[1]))
}

// Round `v` elements.
func (v Vector2) Round() Vector2 {
	return Vec2(Round(v[0]), Round(v[1]))
}

// Swaps x and y
func (v Vector2) ReverseOrder() Vector2 {
	return Vec2(v[1], v[0])
}

// `v` and `other` linear interpolation.
func (v Vector2) Lerp(other Vector2, t f64) Vector2 {
	return other.Sub(v).Mul1(t).Add(v)
}

// // `v` and `other` inverse linear interpolation.
// func (v Vec2) LerpInv(other, value Vec2) f64 {
// 	// return value.Sub(v).Div(other.Sub(v))

// 	return Vec2{
// 		(value - v) / (other - v),
// 	}
// }

// `v` and `other` dot product.
func (v Vector2) Dot(other Vector2) f64 {
	return v.Mul(other).Sum()
}

// Magnitude squared.
func (v Vector2) MagSq() f64 {
	return v.Dot(v)
}

// Magnitude.
func (v Vector2) Mag() f64 {
	return Sqrt(v.MagSq())
}

// `v` direction with `value` magnitude.
func (v Vector2) MagSet(value f64) Vector2 {
	if mag := v.Mag(); mag != 0 {
		return v.Mul1(value / mag)
	}
	return Vector2{}
}

// `v` direction with 1 magnitude.
func (v Vector2) Norm() Vector2 {
	return v.MagSet(1)
}

// Clamps `v` magnitude.
func (v Vector2) ClampMag(max f64) Vector2 {
	if v.Mag() > max {
		return v.MagSet(max)
	}
	return v
}

// Distance between `v` and `other`.
func (v Vector2) Dst(other Vector2) f64 {
	return v.Sub(other).Mag()
}

// Move `v` towards `other` by `dlt`.
func (v Vector2) MoveTowards(other Vector2, dlt f64) Vector2 {
	return other.Sub(v).MagSet(Min(dlt, v.Dst(other))).Add(v)
}

// Project `other` onto `v`, changing magnitude of `v`.
// Both magnitudes affects result magnitude.
func (v Vector2) Project(other Vector2) Vector2 {
	return v.MagSet(v.Dot(other))
}

// Reflect `v` on normal `norm`.
// `norm` should be normalized.
func (v Vector2) Reflect(norm Vector2) Vector2 {
	return v.Sub(norm.Mul1(v.Dot(norm) * 2))
}

// Rotate `v` 90 degrees.
func (v Vector2) Rot90() Vector2 {
	return Vec2(-v[1], v[0])
}

// Angle to direction.
func (a Rad) Vec2() Vector2 {
	return Vec2(a.Cos(), a.Sin())
}

// Direction to angle.
func (v Vector2) Rad() Rad {
	return Rad(Atan2(v[1], v[0]))
}

// Angle from `v` to `other`.
func (v Vector2) AngTo(other Vector2) Rad {
	return other.Sub(v).Rad()
}

// Rotate `v` with angle `amount`.
func (v Vector2) Rot(amount Rad) Vector2 {
	newX := amount.Vec2()
	return newX.Rot90().Mul1(v[1]).Add(newX.Mul1(v[0]))
}
