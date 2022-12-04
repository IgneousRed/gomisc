package gomisc

type Vec2F [2]float64
type Vec2I [2]int

func (v Vec2F) Eq(other Vec2F) bool {
	return v[0] == other[0] && v[1] == other[1]
}
func (v Vec2F) Add(other Vec2F) Vec2F {
	return Vec2F{v[0] + other[0], v[1] + other[1]}
}
func (v Vec2F) Add1(other float64) Vec2F {
	return Vec2F{v[0] + other, v[1] + other}
}
func (v Vec2F) Sub(other Vec2F) Vec2F {
	return Vec2F{v[0] - other[0], v[1] - other[1]}
}
func (v Vec2F) Sub1(other float64) Vec2F {
	return Vec2F{v[0] - other, v[1] - other}
}
func (v Vec2F) Mul(other Vec2F) Vec2F {
	return Vec2F{v[0] * other[0], v[1] * other[1]}
}
func (v Vec2F) Mul1(other float64) Vec2F {
	return Vec2F{v[0] * other, v[1] * other}
}
func (v Vec2F) Div(other Vec2F) Vec2F {
	return Vec2F{v[0] / other[0], v[1] / other[1]}
}
func (v Vec2F) Div1(other float64) Vec2F {
	return Vec2F{v[0] / other, v[1] / other}
}
func (v Vec2F) Abs() Vec2F {
	return Vec2F{Abs(v[0]), Abs(v[1])}
}
func (v Vec2F) Min() float64 {
	return Min(v[:]...)
}
func (v Vec2F) Max() float64 {
	return Max(v[:]...)
}
func (v Vec2F) Floor() Vec2F {
	return Vec2F{Floor(v[0]), Floor(v[1])}
}
func (v Vec2F) FloorI() Vec2I {
	return Vec2I{FloorI(v[0]), FloorI(v[1])}
}
func (v Vec2F) Round() Vec2F {
	return Vec2F{Round(v[0]), Round(v[1])}
}
func (v Vec2F) RoundI() Vec2I {
	return Vec2I{RoundI(v[0]), RoundI(v[1])}
}
func (v Vec2F) Fade() Vec2F {
	return Vec2F{Fade(v[0]), Fade(v[1])}
}
func (v Vec2F) Wrap(lens Vec2F) Vec2F {
	return Vec2F{Wrap(v[0], lens[0]), Wrap(v[1], lens[1])}
}
func (v Vec2F) Wrap1(len float64) Vec2F {
	return Vec2F{Wrap(v[0], len), Wrap(v[1], len)}
}
func (v Vec2F) Dot(other Vec2F) float64 {
	temp := v.Mul(other)
	return Sum(temp[:]...)
}
func (v Vec2F) Mag() float64 {
	return Sqrt(v[0]*v[0] + v[1]*v[1])
}
func (v Vec2F) MagSet(value float64) Vec2F {
	if mag := v.Mag(); mag != 0. {
		fix := value / mag
		return Vec2F{v[0] * fix, v[1] * fix}
	}
	return v
}
func (v Vec2F) Norm() Vec2F {
	return v.MagSet(1.)
}
func (v Vec2F) Project(other Vec2F) Vec2F {
	return v.Mul1(v.Dot(other))
}
func (v Vec2F) Rot90() Vec2F {
	return Vec2F{-v[1], v[0]}
}
func RadToVec2F(ang float64) Vec2F {
	return Vec2F{Cos(ang), Sin(ang)}
}
func Vec2FToRad(v Vec2F) float64 {
	return Atan2(v[1], v[0])
}
func TranslateVec2F(points []Vec2F, amount Vec2F) []Vec2F {
	return MapF(points, func(p Vec2F) Vec2F { return p.Add(amount) })
}
func RotateVec2F(points []Vec2F, amount float64) []Vec2F {
	newX := RadToVec2F(amount)
	newY := newX.Rot90()
	return MapF(points, func(p Vec2F) Vec2F {
		return newX.Mul1(p[0]).Add(newY.Mul1(p[1]))
	})
}
func ScaleVec2F(points []Vec2F, amount Vec2F) []Vec2F {
	return MapF(points, func(p Vec2F) Vec2F { return p.Mul(amount) })
}
