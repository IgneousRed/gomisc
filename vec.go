package gomisc

type Vector2 struct {
	X, Y float64
}
type Rad float64

// New Vector2.
func Vec2(x, y float64) Vector2 {
	return Vector2{x, y}
}

// Are `v` and `other` identical.
func (v Vector2) Eq(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

// Changes sign of each `v` element.
func (v Vector2) Neg() Vector2 {
	return Vec2(-v.X, -v.Y)
}

// Reciprocates each `v` element.
func (v Vector2) Rcp() Vector2 {
	return Vec2(1/v.X, 1/v.Y)
}

// `v` and `other` pairwise add.
func (v Vector2) Add(other Vector2) Vector2 {
	return Vec2(v.X+other.X, v.Y+other.Y)
}

// Add `other` to each `v` element.
func (v Vector2) Add1(other float64) Vector2 {
	return Vec2(v.X+other, v.Y+other)
}

// `v` and `other` pairwise subtract.
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vec2(v.X-other.X, v.Y-other.Y)
}

// Subtract `other` from each `v` element.
func (v Vector2) Sub1(other float64) Vector2 {
	return Vec2(v.X-other, v.Y-other)
}

// `v` and `other` pairwise multiply.
func (v Vector2) Mul(other Vector2) Vector2 {
	return Vec2(v.X*other.X, v.Y*other.Y)
}

// Multiply `other` with each `v` element.
func (v Vector2) Mul1(other float64) Vector2 {
	return Vec2(v.X*other, v.Y*other)
}

// `v` and `other` pairwise divide.
func (v Vector2) Div(other Vector2) Vector2 {
	return Vec2(v.X/other.X, v.Y/other.Y)
}

// Divide `other` from each `v` element.
func (v Vector2) Div1(other float64) Vector2 {
	return Vec2(v.X/other, v.Y/other)
}

// `v` and `other` pairwise wrap.
func (v Vector2) Wrap(lens Vector2) Vector2 {
	return Vec2(Wrap(v.X, lens.X), Wrap(v.Y, lens.Y))
}

// Wrap `len` to each `v` element.
func (v Vector2) Wrap1(len float64) Vector2 {
	return Vec2(Wrap(v.X, len), Wrap(v.Y, len))
}

// Make `v` elements absolute.
func (v Vector2) Abs() Vector2 {
	return Vec2(Abs(v.X), Abs(v.Y))
}

// Lowest `v` element.
func (v Vector2) Min() float64 {
	if v.Y < v.X {
		return v.Y
	}
	return v.X
}

// Highest `v` element.
func (v Vector2) Max() float64 {
	if v.Y > v.X {
		return v.Y
	}
	return v.X
}

// `v` element Su
func (v Vector2) Sum() float64 {
	return v.X + v.Y
}

// Floor `v` elements.
func (v Vector2) Floor() Vector2 {
	return Vec2(Floor(v.X), Floor(v.Y))
}

// Round `v` elements.
func (v Vector2) Round() Vector2 {
	return Vec2(Round(v.X), Round(v.Y))
}

// `v` and `other` linear interpolation.
func (v Vector2) Lerp(other Vector2, t float64) Vector2 {
	return other.Sub(v).Mul1(t).Add(v)
}

// // `v` and `other` inverse linear interpolation.
// func (v Vec2) LerpInv(other, value Vec2) float64 {
// 	// return value.Sub(v).Div(other.Sub(v))

// 	return Vec2{
// 		(value - v) / (other - v),
// 	}
// }

// `v` and `other` dot product.
func (v Vector2) Dot(other Vector2) float64 {
	return v.Mul(other).Sum()
}

// Magnitude squared.
func (v Vector2) MagSq() float64 {
	return v.Dot(v)
}

// Magnitude.
func (v Vector2) Mag() float64 {
	return Sqrt(v.MagSq())
}

// `v` direction with `value` magnitude.
func (v Vector2) MagSet(value float64) Vector2 {
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
func (v Vector2) ClampMag(max float64) Vector2 {
	if v.Mag() > max {
		return v.MagSet(max)
	}
	return v
}

// Distance between `v` and `other`.
func (v Vector2) Dst(other Vector2) float64 {
	return v.Sub(other).Mag()
}

// Move `v` towards `other` by `dlt`.
func (v Vector2) MoveTowards(other Vector2, dlt float64) Vector2 {
	return other.Sub(v).MagSet(Min(dlt, v.Dst(other))).Add(v)
}

// Project `other` onto `v`, changing magnitude of `v`.
// Both magnitudes affects result magnitude.
func (v Vector2) Project(other Vector2) Vector2 {
	return v.MagSet(v.Dot(other))
}

// Rotate `v` 90 degrees.
func (v Vector2) Rot90() Vector2 {
	return Vec2(-v.Y, v.X)
}

// Reflect `v` on normal `norm`.
// `norm` should be normalized.
func (v Vector2) Reflect(norm Vector2) Vector2 {
	return v.Sub(norm.Mul1(v.Dot(norm) * 2))
}

// Angle in degrees.
func Deg(value float64) Rad {
	return Rad(value * Deg2Rad)
}

// Radian cosine.
func (r Rad) Cos() float64 {
	return Cos(float64(r))
}

// Radian sine.
func (r Rad) Sin() float64 {
	return Sin(float64(r))
}

// Angle to direction.
func (a Rad) Vec2() Vector2 {
	return Vec2(a.Cos(), a.Sin())
}

// Direction to angle.
func (v Vector2) Rad() Rad {
	return Rad(Atan2(v.Y, v.X))
}

// Angle from `v` to `other`.
func (v Vector2) AngTo(other Vector2) Rad {
	return other.Sub(v).Rad()
}

// Rotate `v` with angle `amount`.
func (v Vector2) Rot(amount Rad) Vector2 {
	newX := amount.Vec2()
	return newX.Rot90().Mul1(v.Y).Add(newX.Mul1(v.X))
}
