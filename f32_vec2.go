package fp

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix32"
)

var (
    F32Vec2Zero  = F32Vec2FromRaw(fix32.Zero, fix32.Zero)
    F32Vec2One   = F32Vec2FromRaw(fix32.One, fix32.One)
    F32Vec2Down  = F32Vec2FromRaw(fix32.Zero, fix32.Neg1)
    F32Vec2Up    = F32Vec2FromRaw(fix32.Zero, fix32.One)
    F32Vec2Left  = F32Vec2FromRaw(fix32.Neg1, fix32.Zero)
    F32Vec2Right = F32Vec2FromRaw(fix32.One, fix32.Zero)
    F32Vec2AxisX = F32Vec2FromRaw(fix32.One, fix32.Zero)
    F32Vec2AxisY = F32Vec2FromRaw(fix32.Zero, fix32.One)
)

// F32Vec2 struct with signed 16.16 fixed point components.
type F32Vec2 struct {
    RawX int32
    RawY int32
}

func F32Vec2FromRaw(rawX, rawY int32) F32Vec2 {
    return F32Vec2{
        RawX: rawX,
        RawY: rawY,
    }
}

func F32Vec2FromF32(x, y F32) F32Vec2 {
    return F32Vec2FromRaw(x.Raw, y.Raw)
}

func F32Vec2FromInt32(x, y int32) F32Vec2 {
    return F32Vec2FromRaw(fix32.FromInt32(x), fix32.FromInt32(y))
}

func F32Vec2FromFloat32(x, y float32) F32Vec2 {
    return F32Vec2FromRaw(fix32.FromFloat32(x), fix32.FromFloat32(y))
}

func F32Vec2FromFloat64(x, y float64) F32Vec2 {
    return F32Vec2FromRaw(fix32.FromFloat64(x), fix32.FromFloat64(y))
}

// F32Vec2Min returns the smallest F32Vec2 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Vec2Min(v0, v1)
//
// This makes it harder to accidentally call F32Min with 0 arguments.
func F32Vec2Min(v0 F32Vec2, v1 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Min(v0.RawX, v1.RawX), fix32.Min(v0.RawY, v1.RawY))
}

// F32Vec2Max returns the largest F32Vec2 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Vec2Max(v0, v1)
//
// This makes it harder to accidentally call F32Max with 0 arguments.
func F32Vec2Max(v0 F32Vec2, v1 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Max(v0.RawX, v1.RawX), fix32.Max(v0.RawY, v1.RawY))
}

func (v F32Vec2) X() F32 {
    return F32FromRaw(v.RawX)
}

func (v F32Vec2) Y() F32 {
    return F32FromRaw(v.RawY)
}

// Negate -v
func (v F32Vec2) Negate() F32Vec2 {
    return F32Vec2FromRaw(-v.RawX, -v.RawY)
}

// Add v + b
func (v F32Vec2) Add(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(v.RawX+b.RawX, v.RawY+b.RawY)
}

// Sub v - b
func (v F32Vec2) Sub(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(v.RawX-b.RawX, v.RawY-b.RawY)
}

// Mul v * b
func (v F32Vec2) Mul(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Mul(v.RawX, b.RawX), fix32.Mul(v.RawY, b.RawY))
}

// DivPrecise v / b
func (v F32Vec2) DivPrecise(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.DivPrecise(v.RawX, b.RawX), fix32.DivPrecise(v.RawY, b.RawY))
}

// Mod v % b
func (v F32Vec2) Mod(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(v.RawX%b.RawX, v.RawY%b.RawY)
}

// AddF32 v + b
func (v F32Vec2) AddF32(b F32) F32Vec2 {
    return F32Vec2FromRaw(v.RawX+b.Raw, v.RawY+b.Raw)
}

// SubF32 v - b
func (v F32Vec2) SubF32(b F32) F32Vec2 {
    return F32Vec2FromRaw(v.RawX-b.Raw, v.RawY-b.Raw)
}

// MulF32 v * b
func (v F32Vec2) MulF32(b F32) F32Vec2 {
    return F32Vec2FromRaw(fix32.Mul(v.RawX, b.Raw), fix32.Mul(v.RawY, b.Raw))
}

// DivPreciseF32 v / b
func (v F32Vec2) DivPreciseF32(b F32) F32Vec2 {
    return F32Vec2FromRaw(fix32.DivPrecise(v.RawX, b.Raw), fix32.DivPrecise(v.RawY, b.Raw))
}

// ModF32 v % b
func (v F32Vec2) ModF32(b F32) F32Vec2 {
    return F32Vec2FromRaw(v.RawX%b.Raw, v.RawY%b.Raw)
}

// EQ v == b
func (v F32Vec2) EQ(b F32Vec2) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY
}

// NE v != b
func (v F32Vec2) NE(b F32Vec2) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY
}

func (v F32Vec2) Div(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Div(v.RawX, b.RawX), fix32.Div(v.RawY, b.RawY))
}

func (v F32Vec2) DivFast(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.DivFast(v.RawX, b.RawX), fix32.DivFast(v.RawY, b.RawY))
}

func (v F32Vec2) DivFastest(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.DivFastest(v.RawX, b.RawX), fix32.DivFastest(v.RawY, b.RawY))
}

func (v F32Vec2) SqrtPrecise() F32Vec2 {
    return F32Vec2FromRaw(fix32.SqrtPrecise(v.RawX), fix32.SqrtPrecise(v.RawY))
}

func (v F32Vec2) Sqrt() F32Vec2 {
    return F32Vec2FromRaw(fix32.Sqrt(v.RawX), fix32.Sqrt(v.RawY))
}

func (v F32Vec2) SqrtFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.SqrtFast(v.RawX), fix32.SqrtFast(v.RawY))
}

func (v F32Vec2) SqrtFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.SqrtFastest(v.RawX), fix32.SqrtFastest(v.RawY))
}

func (v F32Vec2) RSqrt() F32Vec2 {
    return F32Vec2FromRaw(fix32.RSqrt(v.RawX), fix32.RSqrt(v.RawY))
}

func (v F32Vec2) RSqrtFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.RSqrtFast(v.RawX), fix32.RSqrtFast(v.RawY))
}

func (v F32Vec2) RSqrtFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.RSqrtFastest(v.RawX), fix32.RSqrtFastest(v.RawY))
}

func (v F32Vec2) Rcp() F32Vec2 {
    return F32Vec2FromRaw(fix32.Rcp(v.RawX), fix32.Rcp(v.RawY))
}

func (v F32Vec2) RcpFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.RcpFast(v.RawX), fix32.RcpFast(v.RawY))
}

func (v F32Vec2) RcpFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.RcpFastest(v.RawX), fix32.RcpFastest(v.RawY))
}

func (v F32Vec2) Exp() F32Vec2 {
    return F32Vec2FromRaw(fix32.Exp(v.RawX), fix32.Exp(v.RawY))
}

func (v F32Vec2) ExpFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.ExpFast(v.RawX), fix32.ExpFast(v.RawY))
}

func (v F32Vec2) ExpFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.ExpFastest(v.RawX), fix32.ExpFastest(v.RawY))
}

func (v F32Vec2) Exp2() F32Vec2 {
    return F32Vec2FromRaw(fix32.Exp2(v.RawX), fix32.Exp2(v.RawY))
}

func (v F32Vec2) Exp2Fast() F32Vec2 {
    return F32Vec2FromRaw(fix32.Exp2Fast(v.RawX), fix32.Exp2Fast(v.RawY))
}

func (v F32Vec2) Exp2Fastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.Exp2Fastest(v.RawX), fix32.Exp2Fastest(v.RawY))
}

func (v F32Vec2) Log() F32Vec2 {
    return F32Vec2FromRaw(fix32.Log(v.RawX), fix32.Log(v.RawY))
}

func (v F32Vec2) LogFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.LogFast(v.RawX), fix32.LogFast(v.RawY))
}

func (v F32Vec2) LogFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.LogFastest(v.RawX), fix32.LogFastest(v.RawY))
}

func (v F32Vec2) Log2() F32Vec2 {
    return F32Vec2FromRaw(fix32.Log2(v.RawX), fix32.Log2(v.RawY))
}

func (v F32Vec2) Log2Fast() F32Vec2 {
    return F32Vec2FromRaw(fix32.Log2Fast(v.RawX), fix32.Log2Fast(v.RawY))
}

func (v F32Vec2) Log2Fastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.Log2Fastest(v.RawX), fix32.Log2Fastest(v.RawY))
}

func (v F32Vec2) Sin() F32Vec2 {
    return F32Vec2FromRaw(fix32.Sin(v.RawX), fix32.Sin(v.RawY))
}

func (v F32Vec2) SinFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.SinFast(v.RawX), fix32.SinFast(v.RawY))
}

func (v F32Vec2) SinFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.SinFastest(v.RawX), fix32.SinFastest(v.RawY))
}

func (v F32Vec2) Cos() F32Vec2 {
    return F32Vec2FromRaw(fix32.Cos(v.RawX), fix32.Cos(v.RawY))
}

func (v F32Vec2) CosFast() F32Vec2 {
    return F32Vec2FromRaw(fix32.CosFast(v.RawX), fix32.CosFast(v.RawY))
}

func (v F32Vec2) CosFastest() F32Vec2 {
    return F32Vec2FromRaw(fix32.CosFastest(v.RawX), fix32.CosFastest(v.RawY))
}

func (v F32Vec2) Pow(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Pow(v.RawX, b.RawX), fix32.Pow(v.RawY, b.RawY))
}

func (v F32Vec2) PowFast(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.PowFast(v.RawX, b.RawX), fix32.PowFast(v.RawY, b.RawY))
}

func (v F32Vec2) PowFastest(b F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.PowFastest(v.RawX, b.RawX), fix32.PowFastest(v.RawY, b.RawY))
}

func (v F32Vec2) Length() F32 {
    return F32FromRaw(fix32.Sqrt(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY)))
}

func (v F32Vec2) LengthFast() F32 {
    return F32FromRaw(fix32.SqrtFast(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY)))
}

func (v F32Vec2) LengthFastest() F32 {
    return F32FromRaw(fix32.SqrtFastest(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY)))
}

func (v F32Vec2) LengthSqr() F32 {
    return F32FromRaw(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY))
}

func (v F32Vec2) Normalize() F32Vec2 {
    ooLen := F32FromRaw(fix32.RSqrt(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY)))
    return F32Vec2FromF32(ooLen, ooLen).Mul(v)
}

func (v F32Vec2) NormalizeFast() F32Vec2 {
    ooLen := F32FromRaw(fix32.RSqrtFast(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY)))
    return F32Vec2FromF32(ooLen, ooLen).Mul(v)
}

func (v F32Vec2) NormalizeFastest() F32Vec2 {
    ooLen := F32FromRaw(fix32.RSqrtFastest(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY)))
    return F32Vec2FromF32(ooLen, ooLen).Mul(v)
}

func (v F32Vec2) Dot(b F32Vec2) F32 {
    return F32FromRaw(fix32.Mul(v.RawX, b.RawX) + fix32.Mul(v.RawY, b.RawY))
}

func (v F32Vec2) Distance(b F32Vec2) F32 {
    return v.Sub(b).Length()
}

func (v F32Vec2) DistanceFast(b F32Vec2) F32 {
    return v.Sub(b).LengthFast()
}

func (v F32Vec2) DistanceFastest(b F32Vec2) F32 {
    return v.Sub(b).LengthFastest()
}

func (v F32Vec2) Clamp(min, max F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Clamp(v.RawX, min.RawX, max.RawX), fix32.Clamp(v.RawY, min.RawY, max.RawY))
}

func (v F32Vec2) Lerp(b F32Vec2, t F32) F32Vec2 {
    tb := t.Raw
    ta := fix32.One - tb
    return F32Vec2FromRaw(fix32.Mul(v.RawX, ta)+fix32.Mul(b.RawX, tb), fix32.Mul(v.RawY, ta)+fix32.Mul(b.RawY, tb))
}

func (v F32Vec2) Equals(obj F32Vec2) bool {
    return reflect.DeepEqual(v, obj)
}

func (v F32Vec2) ToString() string {
    return fmt.Sprintf(`(%s, %s)`, fix32.ToString(v.RawX), fix32.ToString(v.RawY))
}
