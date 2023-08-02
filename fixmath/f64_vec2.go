package fixmath

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

// F64Vec2 struct with signed 32.32 fixed point components.
type F64Vec2 struct {
    RawX int64
    RawY int64
}

func Vec2FromRaw(rawX, rawY int64) F64Vec2 {
    return F64Vec2{
        RawX: rawX,
        RawY: rawY,
    }
}

func Vec2FromF64(x, y F64) F64Vec2 {
    return Vec2FromRaw(x.Raw, y.Raw)
}

func Vec2FromInt32(x, y int32) F64Vec2 {
    return Vec2FromRaw(fix64.FromInt32(x), fix64.FromInt32(y))
}

func Vec2FromInt64(x, y int64) F64Vec2 {
    return Vec2FromRaw(fix64.FromInt64(x), fix64.FromInt64(y))
}

func Vec2FromFloat32(x, y float32) F64Vec2 {
    return Vec2FromRaw(fix64.FromFloat32(x), fix64.FromFloat32(y))
}

func Vec2FromFloat64(x, y float64) F64Vec2 {
    return Vec2FromRaw(fix64.FromFloat64(x), fix64.FromFloat64(y))
}

// Vec2Min returns the smallest F64Vec2 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Vec2Min(v0, v1)
//
// This makes it harder to accidentally call F64Min with 0 arguments.
func Vec2Min(v0 F64Vec2, v1 F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Min(v0.RawX, v1.RawX), fix64.Min(v0.RawY, v1.RawY))
}

// Vec2Max returns the largest F64Vec2 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Vec2Max(v0, v1)
//
// This makes it harder to accidentally call F64Max with 0 arguments.
func Vec2Max(v0 F64Vec2, v1 F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Max(v0.RawX, v1.RawX), fix64.Max(v0.RawY, v1.RawY))
}

func (v F64Vec2) X() F64 {
    return F64FromRaw(v.RawX)
}

func (v F64Vec2) Y() F64 {
    return F64FromRaw(v.RawY)
}

// Negate -v
func (v F64Vec2) Negate() F64Vec2 {
    return Vec2FromRaw(-v.RawX, -v.RawY)
}

// Add v + b
func (v F64Vec2) Add(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(v.RawX+b.RawX, v.RawY+b.RawY)
}

// Sub v - b
func (v F64Vec2) Sub(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(v.RawX-b.RawX, v.RawY-b.RawY)
}

// Mul v * b
func (v F64Vec2) Mul(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Mul(v.RawX, b.RawX), fix64.Mul(v.RawY, b.RawY))
}

// DivPrecise v / b
func (v F64Vec2) DivPrecise(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.DivPrecise(v.RawX, b.RawX), fix64.DivPrecise(v.RawY, b.RawY))
}

// Mod v % b
func (v F64Vec2) Mod(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(v.RawX%b.RawX, v.RawY%b.RawY)
}

// AddF64 v + b
func (v F64Vec2) AddF64(b F64) F64Vec2 {
    return Vec2FromRaw(v.RawX+b.Raw, v.RawY+b.Raw)
}

// SubF64 v - b
func (v F64Vec2) SubF64(b F64) F64Vec2 {
    return Vec2FromRaw(v.RawX-b.Raw, v.RawY-b.Raw)
}

// MulF64 v * b
func (v F64Vec2) MulF64(b F64) F64Vec2 {
    return Vec2FromRaw(fix64.Mul(v.RawX, b.Raw), fix64.Mul(v.RawY, b.Raw))
}

// DivPreciseF64 v / b
func (v F64Vec2) DivPreciseF64(b F64) F64Vec2 {
    return Vec2FromRaw(fix64.DivPrecise(v.RawX, b.Raw), fix64.DivPrecise(v.RawY, b.Raw))
}

// ModF64 v % b
func (v F64Vec2) ModF64(b F64) F64Vec2 {
    return Vec2FromRaw(v.RawX%b.Raw, v.RawY%b.Raw)
}

// EQ v == b
func (v F64Vec2) EQ(b F64Vec2) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY
}

// NE v != b
func (v F64Vec2) NE(b F64Vec2) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY
}

func (v F64Vec2) Div(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Div(v.RawX, b.RawX), fix64.Div(v.RawY, b.RawY))
}

func (v F64Vec2) DivFast(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.DivFast(v.RawX, b.RawX), fix64.DivFast(v.RawY, b.RawY))
}

func (v F64Vec2) DivFastest(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.DivFastest(v.RawX, b.RawX), fix64.DivFastest(v.RawY, b.RawY))
}

func (v F64Vec2) SqrtPrecise() F64Vec2 {
    return Vec2FromRaw(fix64.SqrtPrecise(v.RawX), fix64.SqrtPrecise(v.RawY))
}

func (v F64Vec2) Sqrt() F64Vec2 {
    return Vec2FromRaw(fix64.Sqrt(v.RawX), fix64.Sqrt(v.RawY))
}

func (v F64Vec2) SqrtFast() F64Vec2 {
    return Vec2FromRaw(fix64.SqrtFast(v.RawX), fix64.SqrtFast(v.RawY))
}

func (v F64Vec2) SqrtFastest() F64Vec2 {
    return Vec2FromRaw(fix64.SqrtFastest(v.RawX), fix64.SqrtFastest(v.RawY))
}

func (v F64Vec2) RSqrt() F64Vec2 {
    return Vec2FromRaw(fix64.RSqrt(v.RawX), fix64.RSqrt(v.RawY))
}

func (v F64Vec2) RSqrtFast() F64Vec2 {
    return Vec2FromRaw(fix64.RSqrtFast(v.RawX), fix64.RSqrtFast(v.RawY))
}

func (v F64Vec2) RSqrtFastest() F64Vec2 {
    return Vec2FromRaw(fix64.RSqrtFastest(v.RawX), fix64.RSqrtFastest(v.RawY))
}

func (v F64Vec2) Rcp() F64Vec2 {
    return Vec2FromRaw(fix64.Rcp(v.RawX), fix64.Rcp(v.RawY))
}

func (v F64Vec2) RcpFast() F64Vec2 {
    return Vec2FromRaw(fix64.RcpFast(v.RawX), fix64.RcpFast(v.RawY))
}

func (v F64Vec2) RcpFastest() F64Vec2 {
    return Vec2FromRaw(fix64.RcpFastest(v.RawX), fix64.RcpFastest(v.RawY))
}

func (v F64Vec2) Exp() F64Vec2 {
    return Vec2FromRaw(fix64.Exp(v.RawX), fix64.Exp(v.RawY))
}

func (v F64Vec2) ExpFast() F64Vec2 {
    return Vec2FromRaw(fix64.ExpFast(v.RawX), fix64.ExpFast(v.RawY))
}

func (v F64Vec2) ExpFastest() F64Vec2 {
    return Vec2FromRaw(fix64.ExpFastest(v.RawX), fix64.ExpFastest(v.RawY))
}

func (v F64Vec2) Exp2() F64Vec2 {
    return Vec2FromRaw(fix64.Exp2(v.RawX), fix64.Exp2(v.RawY))
}

func (v F64Vec2) Exp2Fast() F64Vec2 {
    return Vec2FromRaw(fix64.Exp2Fast(v.RawX), fix64.Exp2Fast(v.RawY))
}

func (v F64Vec2) Exp2Fastest() F64Vec2 {
    return Vec2FromRaw(fix64.Exp2Fastest(v.RawX), fix64.Exp2Fastest(v.RawY))
}

func (v F64Vec2) Log() F64Vec2 {
    return Vec2FromRaw(fix64.Log(v.RawX), fix64.Log(v.RawY))
}

func (v F64Vec2) LogFast() F64Vec2 {
    return Vec2FromRaw(fix64.LogFast(v.RawX), fix64.LogFast(v.RawY))
}

func (v F64Vec2) LogFastest() F64Vec2 {
    return Vec2FromRaw(fix64.LogFastest(v.RawX), fix64.LogFastest(v.RawY))
}

func (v F64Vec2) Log2() F64Vec2 {
    return Vec2FromRaw(fix64.Log2(v.RawX), fix64.Log2(v.RawY))
}

func (v F64Vec2) Log2Fast() F64Vec2 {
    return Vec2FromRaw(fix64.Log2Fast(v.RawX), fix64.Log2Fast(v.RawY))
}

func (v F64Vec2) Log2Fastest() F64Vec2 {
    return Vec2FromRaw(fix64.Log2Fastest(v.RawX), fix64.Log2Fastest(v.RawY))
}

func (v F64Vec2) Sin() F64Vec2 {
    return Vec2FromRaw(fix64.Sin(v.RawX), fix64.Sin(v.RawY))
}

func (v F64Vec2) SinFast() F64Vec2 {
    return Vec2FromRaw(fix64.SinFast(v.RawX), fix64.SinFast(v.RawY))
}

func (v F64Vec2) SinFastest() F64Vec2 {
    return Vec2FromRaw(fix64.SinFastest(v.RawX), fix64.SinFastest(v.RawY))
}

func (v F64Vec2) Cos() F64Vec2 {
    return Vec2FromRaw(fix64.Cos(v.RawX), fix64.Cos(v.RawY))
}

func (v F64Vec2) CosFast() F64Vec2 {
    return Vec2FromRaw(fix64.CosFast(v.RawX), fix64.CosFast(v.RawY))
}

func (v F64Vec2) CosFastest() F64Vec2 {
    return Vec2FromRaw(fix64.CosFastest(v.RawX), fix64.CosFastest(v.RawY))
}

func (v F64Vec2) Pow(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Pow(v.RawX, b.RawX), fix64.Pow(v.RawY, b.RawY))
}

func (v F64Vec2) PowFast(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.PowFast(v.RawX, b.RawX), fix64.PowFast(v.RawY, b.RawY))
}

func (v F64Vec2) PowFastest(b F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.PowFastest(v.RawX, b.RawX), fix64.PowFastest(v.RawY, b.RawY))
}

func (v F64Vec2) Length() F64 {
    return F64FromRaw(fix64.Sqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
}

func (v F64Vec2) LengthFast() F64 {
    return F64FromRaw(fix64.SqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
}

func (v F64Vec2) LengthFastest() F64 {
    return F64FromRaw(fix64.SqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
}

func (v F64Vec2) LengthSqr() F64 {
    return F64FromRaw(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY))
}

func (v F64Vec2) Normalize() F64Vec2 {
    ooLen := F64FromRaw(fix64.RSqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
    return Vec2FromF64(ooLen, ooLen).Mul(v)
}

func (v F64Vec2) NormalizeFast() F64Vec2 {
    ooLen := F64FromRaw(fix64.RSqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
    return Vec2FromF64(ooLen, ooLen).Mul(v)
}

func (v F64Vec2) NormalizeFastest() F64Vec2 {
    ooLen := F64FromRaw(fix64.RSqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY)))
    return Vec2FromF64(ooLen, ooLen).Mul(v)
}

func (v F64Vec2) Dot(b F64Vec2) F64 {
    return F64FromRaw(fix64.Mul(v.RawX, b.RawX) + fix64.Mul(v.RawY, b.RawY))
}

func (v F64Vec2) Distance(b F64Vec2) F64 {
    return v.Sub(b).Length()
}

func (v F64Vec2) DistanceFast(b F64Vec2) F64 {
    return v.Sub(b).LengthFast()
}

func (v F64Vec2) DistanceFastest(b F64Vec2) F64 {
    return v.Sub(b).LengthFastest()
}

func (v F64Vec2) Clamp(min, max F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Clamp(v.RawX, min.RawX, max.RawX), fix64.Clamp(v.RawY, min.RawY, max.RawY))
}

func (v F64Vec2) Lerp(b F64Vec2, t F64) F64Vec2 {
    tb := t.Raw
    ta := fix64.One - tb
    return Vec2FromRaw(fix64.Mul(v.RawX, ta)+fix64.Mul(b.RawX, tb), fix64.Mul(v.RawY, ta)+fix64.Mul(b.RawY, tb))
}

func (v F64Vec2) Equals(obj F64Vec2) bool {
    return reflect.DeepEqual(v, obj)
}

func (v F64Vec2) ToString() string {
    return fmt.Sprintf(`(%s, %s)`, fix64.ToString(v.RawX), fix64.ToString(v.RawY))
}
