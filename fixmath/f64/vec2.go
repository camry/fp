package f64

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix64"
)

var (
    Vec2Zero  = Vec2FromRaw(fix64.Zero, fix64.Zero)
    Vec2One   = Vec2FromRaw(fix64.One, fix64.One)
    Vec2Down  = Vec2FromRaw(fix64.Zero, fix64.Neg1)
    Vec2Up    = Vec2FromRaw(fix64.Zero, fix64.One)
    Vec2Left  = Vec2FromRaw(fix64.Neg1, fix64.Zero)
    Vec2Right = Vec2FromRaw(fix64.One, fix64.Zero)
    Vec2AxisX = Vec2FromRaw(fix64.One, fix64.Zero)
    Vec2AxisY = Vec2FromRaw(fix64.Zero, fix64.One)
)

// Vec2 struct with signed 32.32 fixed point components.
type Vec2 struct {
    RawX int64
    RawY int64
}

func Vec2FromRaw(rawX, rawY int64) Vec2 {
    return Vec2{
        RawX: rawX,
        RawY: rawY,
    }
}

func Vec2FromF64(x, y F64) Vec2 {
    return Vec2FromRaw(x.Raw, y.Raw)
}

func Vec2FromInt32(x, y int32) Vec2 {
    return Vec2FromRaw(fix64.FromInt32(x), fix64.FromInt32(y))
}

func Vec2FromInt64(x, y int64) Vec2 {
    return Vec2FromRaw(fix64.FromInt64(x), fix64.FromInt64(y))
}

func Vec2FromFloat32(x, y float32) Vec2 {
    return Vec2FromRaw(fix64.FromFloat32(x), fix64.FromFloat32(y))
}

func Vec2FromFloat64(x, y float64) Vec2 {
    return Vec2FromRaw(fix64.FromFloat64(x), fix64.FromFloat64(y))
}

// Vec2Min returns the smallest Vec2 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Vec2Min(v0, v1)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Vec2Min(v0 Vec2, v1 Vec2) Vec2 {
    return Vec2FromRaw(fix64.Min(v0.RawX, v1.RawX), fix64.Min(v0.RawY, v1.RawY))
}

// Vec2Max returns the largest Vec2 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Vec2Max(v0, v1)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Vec2Max(v0 Vec2, v1 Vec2) Vec2 {
    return Vec2FromRaw(fix64.Max(v0.RawX, v1.RawX), fix64.Max(v0.RawY, v1.RawY))
}

func (v Vec2) X() F64 {
    return FromRaw(v.RawX)
}

func (v Vec2) Y() F64 {
    return FromRaw(v.RawY)
}

// Negate -v
func (v Vec2) Negate() Vec2 {
    return Vec2FromRaw(-v.RawX, -v.RawY)
}

// Add v + b
func (v Vec2) Add(b Vec2) Vec2 {
    return Vec2FromRaw(v.RawX+b.RawX, v.RawY+b.RawY)
}

// Sub v - b
func (v Vec2) Sub(b Vec2) Vec2 {
    return Vec2FromRaw(v.RawX-b.RawX, v.RawY-b.RawY)
}

// Mul v * b
func (v Vec2) Mul(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.Mul(v.RawX, b.RawX), fix64.Mul(v.RawY, b.RawY))
}

// DivPrecise v / b
func (v Vec2) DivPrecise(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.DivPrecise(v.RawX, b.RawX), fix64.DivPrecise(v.RawY, b.RawY))
}

// Mod v % b
func (v Vec2) Mod(b Vec2) Vec2 {
    return Vec2FromRaw(v.RawX%b.RawX, v.RawY%b.RawY)
}

// AddF64 v + b
func (v Vec2) AddF64(b F64) Vec2 {
    return Vec2FromRaw(v.RawX+b.Raw, v.RawY+b.Raw)
}

// SubF64 v - b
func (v Vec2) SubF64(b F64) Vec2 {
    return Vec2FromRaw(v.RawX-b.Raw, v.RawY-b.Raw)
}

// MulF64 v * b
func (v Vec2) MulF64(b F64) Vec2 {
    return Vec2FromRaw(fix64.Mul(v.RawX, b.Raw), fix64.Mul(v.RawY, b.Raw))
}

// DivPreciseF64 v / b
func (v Vec2) DivPreciseF64(b F64) Vec2 {
    return Vec2FromRaw(fix64.DivPrecise(v.RawX, b.Raw), fix64.DivPrecise(v.RawY, b.Raw))
}

// ModF64 v % b
func (v Vec2) ModF64(b F64) Vec2 {
    return Vec2FromRaw(v.RawX%b.Raw, v.RawY%b.Raw)
}

// EQ v == b
func (v Vec2) EQ(b Vec2) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY
}

// NE v != b
func (v Vec2) NE(b Vec2) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY
}

func (v Vec2) Div(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.Div(v.RawX, b.RawX), fix64.Div(v.RawY, b.RawY))
}

func (v Vec2) DivFast(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.DivFast(v.RawX, b.RawX), fix64.DivFast(v.RawY, b.RawY))
}

func (v Vec2) DivFastest(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.DivFastest(v.RawX, b.RawX), fix64.DivFastest(v.RawY, b.RawY))
}

func (v Vec2) SqrtPrecise() Vec2 {
    return Vec2FromRaw(fix64.SqrtPrecise(v.RawX), fix64.SqrtPrecise(v.RawY))
}

func (v Vec2) Sqrt() Vec2 {
    return Vec2FromRaw(fix64.Sqrt(v.RawX), fix64.Sqrt(v.RawY))
}

func (v Vec2) SqrtFast() Vec2 {
    return Vec2FromRaw(fix64.SqrtFast(v.RawX), fix64.SqrtFast(v.RawY))
}

func (v Vec2) SqrtFastest() Vec2 {
    return Vec2FromRaw(fix64.SqrtFastest(v.RawX), fix64.SqrtFastest(v.RawY))
}

func (v Vec2) RSqrt() Vec2 {
    return Vec2FromRaw(fix64.RSqrt(v.RawX), fix64.RSqrt(v.RawY))
}

func (v Vec2) RSqrtFast() Vec2 {
    return Vec2FromRaw(fix64.RSqrtFast(v.RawX), fix64.RSqrtFast(v.RawY))
}

func (v Vec2) RSqrtFastest() Vec2 {
    return Vec2FromRaw(fix64.RSqrtFastest(v.RawX), fix64.RSqrtFastest(v.RawY))
}

func (v Vec2) Rcp() Vec2 {
    return Vec2FromRaw(fix64.Rcp(v.RawX), fix64.Rcp(v.RawY))
}

func (v Vec2) RcpFast() Vec2 {
    return Vec2FromRaw(fix64.RcpFast(v.RawX), fix64.RcpFast(v.RawY))
}

func (v Vec2) RcpFastest() Vec2 {
    return Vec2FromRaw(fix64.RcpFastest(v.RawX), fix64.RcpFastest(v.RawY))
}

func (v Vec2) Exp() Vec2 {
    return Vec2FromRaw(fix64.Exp(v.RawX), fix64.Exp(v.RawY))
}

func (v Vec2) ExpFast() Vec2 {
    return Vec2FromRaw(fix64.ExpFast(v.RawX), fix64.ExpFast(v.RawY))
}

func (v Vec2) ExpFastest() Vec2 {
    return Vec2FromRaw(fix64.ExpFastest(v.RawX), fix64.ExpFastest(v.RawY))
}

func (v Vec2) Exp2() Vec2 {
    return Vec2FromRaw(fix64.Exp2(v.RawX), fix64.Exp2(v.RawY))
}

func (v Vec2) Exp2Fast() Vec2 {
    return Vec2FromRaw(fix64.Exp2Fast(v.RawX), fix64.Exp2Fast(v.RawY))
}

func (v Vec2) Exp2Fastest() Vec2 {
    return Vec2FromRaw(fix64.Exp2Fastest(v.RawX), fix64.Exp2Fastest(v.RawY))
}

func (v Vec2) Log() Vec2 {
    return Vec2FromRaw(fix64.Log(v.RawX), fix64.Log(v.RawY))
}

func (v Vec2) LogFast() Vec2 {
    return Vec2FromRaw(fix64.LogFast(v.RawX), fix64.LogFast(v.RawY))
}

func (v Vec2) LogFastest() Vec2 {
    return Vec2FromRaw(fix64.LogFastest(v.RawX), fix64.LogFastest(v.RawY))
}

func (v Vec2) Log2() Vec2 {
    return Vec2FromRaw(fix64.Log2(v.RawX), fix64.Log2(v.RawY))
}

func (v Vec2) Log2Fast() Vec2 {
    return Vec2FromRaw(fix64.Log2Fast(v.RawX), fix64.Log2Fast(v.RawY))
}

func (v Vec2) Log2Fastest() Vec2 {
    return Vec2FromRaw(fix64.Log2Fastest(v.RawX), fix64.Log2Fastest(v.RawY))
}

func (v Vec2) Sin() Vec2 {
    return Vec2FromRaw(fix64.Sin(v.RawX), fix64.Sin(v.RawY))
}

func (v Vec2) SinFast() Vec2 {
    return Vec2FromRaw(fix64.SinFast(v.RawX), fix64.SinFast(v.RawY))
}

func (v Vec2) SinFastest() Vec2 {
    return Vec2FromRaw(fix64.SinFastest(v.RawX), fix64.SinFastest(v.RawY))
}

func (v Vec2) Cos() Vec2 {
    return Vec2FromRaw(fix64.Cos(v.RawX), fix64.Cos(v.RawY))
}

func (v Vec2) CosFast() Vec2 {
    return Vec2FromRaw(fix64.CosFast(v.RawX), fix64.CosFast(v.RawY))
}

func (v Vec2) CosFastest() Vec2 {
    return Vec2FromRaw(fix64.CosFastest(v.RawX), fix64.CosFastest(v.RawY))
}

func (v Vec2) Pow(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.Pow(v.RawX, b.RawX), fix64.Pow(v.RawY, b.RawY))
}

func (v Vec2) PowFast(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.PowFast(v.RawX, b.RawX), fix64.PowFast(v.RawY, b.RawY))
}

func (v Vec2) PowFastest(b Vec2) Vec2 {
    return Vec2FromRaw(fix64.PowFastest(v.RawX, b.RawX), fix64.PowFastest(v.RawY, b.RawY))
}

func (v Vec2) Length() F64 {
    return FromRaw(fix64.Sqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
}

func (v Vec2) LengthFast() F64 {
    return FromRaw(fix64.SqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
}

func (v Vec2) LengthFastest() F64 {
    return FromRaw(fix64.SqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
}

func (v Vec2) LengthSqr() F64 {
    return FromRaw(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY))
}

func (v Vec2) Normalize() Vec2 {
    ooLen := FromRaw(fix64.RSqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
    return Vec2FromF64(ooLen, ooLen).Mul(v)
}

func (v Vec2) NormalizeFast() Vec2 {
    ooLen := FromRaw(fix64.RSqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
    return Vec2FromF64(ooLen, ooLen).Mul(v)
}

func (v Vec2) NormalizeFastest() Vec2 {
    ooLen := FromRaw(fix64.RSqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
    return Vec2FromF64(ooLen, ooLen).Mul(v)
}

func (v Vec2) Dot(b Vec2) F64 {
    return FromRaw(fix64.Mul(v.RawX, b.RawX) + fix64.Mul(v.RawY, b.RawY))
}

func (v Vec2) Distance(b Vec2) F64 {
    return v.Sub(b).Length()
}

func (v Vec2) DistanceFast(b Vec2) F64 {
    return v.Sub(b).LengthFast()
}

func (v Vec2) DistanceFastest(b Vec2) F64 {
    return v.Sub(b).LengthFastest()
}

func (v Vec2) Clamp(min, max Vec2) Vec2 {
    return Vec2FromRaw(fix64.Clamp(v.RawX, min.RawX, max.RawX), fix64.Clamp(v.RawY, min.RawY, max.RawY))
}

func (v Vec2) Lerp(b Vec2, t F64) Vec2 {
    tb := t.Raw
    ta := fix64.One - tb
    return Vec2FromRaw(fix64.Mul(v.RawX, ta)+fix64.Mul(b.RawX, tb), fix64.Mul(v.RawY, ta)+fix64.Mul(b.RawY, tb))
}

func (v Vec2) Equals(obj Vec2) bool {
    return reflect.DeepEqual(v, obj)
}

func (v Vec2) ToString() string {
    return fmt.Sprintf(`(%s, %s)`, fix64.ToString(v.RawX), fix64.ToString(v.RawY))
}
