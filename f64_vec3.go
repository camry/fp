package fp

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix64"
)

var (
    F64Vec3Zero    = F64Vec3FromRaw(fix64.Zero, fix64.Zero, fix64.Zero)
    F64Vec3One     = F64Vec3FromRaw(fix64.One, fix64.One, fix64.One)
    F64Vec3Down    = F64Vec3FromRaw(fix64.Zero, fix64.Neg1, fix64.Zero)
    F64Vec3Up      = F64Vec3FromRaw(fix64.Zero, fix64.One, fix64.Zero)
    F64Vec3Left    = F64Vec3FromRaw(fix64.Neg1, fix64.Zero, fix64.Zero)
    F64Vec3Right   = F64Vec3FromRaw(fix64.One, fix64.Zero, fix64.Zero)
    F64Vec3Forward = F64Vec3FromRaw(fix64.Zero, fix64.Zero, fix64.One)
    F64Vec3Back    = F64Vec3FromRaw(fix64.Zero, fix64.Zero, fix64.Neg1)
    F64Vec3AxisX   = F64Vec3FromRaw(fix64.One, fix64.Zero, fix64.Zero)
    F64Vec3AxisY   = F64Vec3FromRaw(fix64.Zero, fix64.One, fix64.Zero)
    F64Vec3AxisZ   = F64Vec3FromRaw(fix64.Zero, fix64.Zero, fix64.One)
)

// F64Vec3 struct with signed 32.32 fixed point components.
type F64Vec3 struct {
    RawX int64
    RawY int64
    RawZ int64
}

func F64Vec3FromRaw(rawX, rawY, RawZ int64) F64Vec3 {
    return F64Vec3{
        RawX: rawX,
        RawY: rawY,
        RawZ: RawZ,
    }
}

func F64Vec3FromF64(x, y, z F64) F64Vec3 {
    return F64Vec3FromRaw(x.Raw, y.Raw, z.Raw)
}

func F64Vec3FromInt32(x, y, z int32) F64Vec3 {
    return F64Vec3FromRaw(fix64.FromInt32(x), fix64.FromInt32(y), fix64.FromInt32(z))
}

func F64Vec3FromInt64(x, y, z int64) F64Vec3 {
    return F64Vec3FromRaw(fix64.FromInt64(x), fix64.FromInt64(y), fix64.FromInt64(z))
}

func F64Vec3FromFloat32(x, y, z float32) F64Vec3 {
    return F64Vec3FromRaw(fix64.FromFloat32(x), fix64.FromFloat32(y), fix64.FromFloat32(z))
}

func F64Vec3FromFloat64(x, y, z float64) F64Vec3 {
    return F64Vec3FromRaw(fix64.FromFloat64(x), fix64.FromFloat64(y), fix64.FromFloat64(z))
}

// F64Vec3Min returns the smallest F64Vec3 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F64Vec3Min(v0, v1)
//
// This makes it harder to accidentally call F64Vec3Min with 0 arguments.
func F64Vec3Min(v0 F64Vec3, v1 F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Min(v0.RawX, v1.RawX), fix64.Min(v0.RawY, v1.RawY), fix64.Min(v0.RawZ, v1.RawZ))
}

// F64Vec3Max returns the largest F64Vec3 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F64Vec3Max(v0, v1)
//
// This makes it harder to accidentally call F64Vec3Max with 0 arguments.
func F64Vec3Max(v0 F64Vec3, v1 F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Max(v0.RawX, v1.RawX), fix64.Max(v0.RawY, v1.RawY), fix64.Min(v0.RawZ, v1.RawZ))
}

func (v F64Vec3) X() F64 {
    return F64FromRaw(v.RawX)
}

func (v F64Vec3) Y() F64 {
    return F64FromRaw(v.RawY)
}

func (v F64Vec3) Z() F64 {
    return F64FromRaw(v.RawZ)
}

// Negate -v
func (v F64Vec3) Negate() F64Vec3 {
    return F64Vec3FromRaw(-v.RawX, -v.RawY, -v.RawZ)
}

// Add v + b
func (v F64Vec3) Add(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(v.RawX+b.RawX, v.RawY+b.RawY, v.RawZ+b.RawZ)
}

// Sub v - b
func (v F64Vec3) Sub(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(v.RawX-b.RawX, v.RawY-b.RawY, v.RawZ-b.RawZ)
}

// Mul v * b
func (v F64Vec3) Mul(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Mul(v.RawX, b.RawX), fix64.Mul(v.RawY, b.RawY), fix64.Mul(v.RawZ, b.RawZ))
}

// DivPrecise v / b
func (v F64Vec3) DivPrecise(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.DivPrecise(v.RawX, b.RawX), fix64.DivPrecise(v.RawY, b.RawY), fix64.DivPrecise(v.RawZ, b.RawZ))
}

// Mod v % b
func (v F64Vec3) Mod(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(v.RawX%b.RawX, v.RawY%b.RawY, v.RawZ%b.RawZ)
}

// AddF64 v + b
func (v F64Vec3) AddF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(v.RawX+b.Raw, v.RawY+b.Raw, v.RawZ+b.Raw)
}

// SubF64 v - b
func (v F64Vec3) SubF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(v.RawX-b.Raw, v.RawY-b.Raw, v.RawZ-b.Raw)
}

// MulF64 v * b
func (v F64Vec3) MulF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(fix64.Mul(v.RawX, b.Raw), fix64.Mul(v.RawY, b.Raw), fix64.Mul(v.RawZ, b.Raw))
}

// DivPreciseF64 v / b
func (v F64Vec3) DivPreciseF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(fix64.DivPrecise(v.RawX, b.Raw), fix64.DivPrecise(v.RawY, b.Raw), fix64.DivPrecise(v.RawZ, b.Raw))
}

// ModF64 v % b
func (v F64Vec3) ModF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(v.RawX%b.Raw, v.RawY%b.Raw, v.RawZ%b.Raw)
}

// EQ v == b
func (v F64Vec3) EQ(b F64Vec3) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY && v.RawZ == b.RawZ
}

// NE v != b
func (v F64Vec3) NE(b F64Vec3) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY || v.RawZ != b.RawZ
}

func (v F64Vec3) DivF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(fix64.Div(v.RawX, b.Raw), fix64.Div(v.RawY, b.Raw), fix64.Div(v.RawZ, b.Raw))
}

func (v F64Vec3) DivFastF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(fix64.DivFast(v.RawX, b.Raw), fix64.DivFast(v.RawY, b.Raw), fix64.DivFast(v.RawZ, b.Raw))
}

func (v F64Vec3) DivFastestF64(b F64) F64Vec3 {
    return F64Vec3FromRaw(fix64.DivFastest(v.RawX, b.Raw), fix64.DivFastest(v.RawY, b.Raw), fix64.DivFastest(v.RawZ, b.Raw))
}

func (v F64Vec3) Div(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Div(v.RawX, b.RawX), fix64.Div(v.RawY, b.RawY), fix64.Div(v.RawZ, b.RawZ))
}

func (v F64Vec3) DivFast(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.DivFast(v.RawX, b.RawX), fix64.DivFast(v.RawY, b.RawY), fix64.DivFast(v.RawZ, b.RawZ))
}

func (v F64Vec3) DivFastest(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.DivFastest(v.RawX, b.RawX), fix64.DivFastest(v.RawY, b.RawY), fix64.DivFastest(v.RawZ, b.RawZ))
}

func (v F64Vec3) SqrtPrecise() F64Vec3 {
    return F64Vec3FromRaw(fix64.SqrtPrecise(v.RawX), fix64.SqrtPrecise(v.RawY), fix64.SqrtPrecise(v.RawZ))
}

func (v F64Vec3) Sqrt() F64Vec3 {
    return F64Vec3FromRaw(fix64.Sqrt(v.RawX), fix64.Sqrt(v.RawY), fix64.Sqrt(v.RawZ))
}

func (v F64Vec3) SqrtFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.SqrtFast(v.RawX), fix64.SqrtFast(v.RawY), fix64.SqrtFast(v.RawZ))
}

func (v F64Vec3) SqrtFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.SqrtFastest(v.RawX), fix64.SqrtFastest(v.RawY), fix64.SqrtFastest(v.RawZ))
}

func (v F64Vec3) RSqrt() F64Vec3 {
    return F64Vec3FromRaw(fix64.RSqrt(v.RawX), fix64.RSqrt(v.RawY), fix64.RSqrt(v.RawZ))
}

func (v F64Vec3) RSqrtFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.RSqrtFast(v.RawX), fix64.RSqrtFast(v.RawY), fix64.RSqrtFast(v.RawZ))
}

func (v F64Vec3) RSqrtFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.RSqrtFastest(v.RawX), fix64.RSqrtFastest(v.RawY), fix64.RSqrtFastest(v.RawZ))
}

func (v F64Vec3) Rcp() F64Vec3 {
    return F64Vec3FromRaw(fix64.Rcp(v.RawX), fix64.Rcp(v.RawY), fix64.Rcp(v.RawZ))
}

func (v F64Vec3) RcpFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.RcpFast(v.RawX), fix64.RcpFast(v.RawY), fix64.RcpFast(v.RawZ))
}

func (v F64Vec3) RcpFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.RcpFastest(v.RawX), fix64.RcpFastest(v.RawY), fix64.RcpFastest(v.RawZ))
}

func (v F64Vec3) Exp() F64Vec3 {
    return F64Vec3FromRaw(fix64.Exp(v.RawX), fix64.Exp(v.RawY), fix64.Exp(v.RawZ))
}

func (v F64Vec3) ExpFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.ExpFast(v.RawX), fix64.ExpFast(v.RawY), fix64.ExpFast(v.RawZ))
}

func (v F64Vec3) ExpFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.ExpFastest(v.RawX), fix64.ExpFastest(v.RawY), fix64.ExpFastest(v.RawZ))
}

func (v F64Vec3) Exp2() F64Vec3 {
    return F64Vec3FromRaw(fix64.Exp2(v.RawX), fix64.Exp2(v.RawY), fix64.Exp2(v.RawZ))
}

func (v F64Vec3) Exp2Fast() F64Vec3 {
    return F64Vec3FromRaw(fix64.Exp2Fast(v.RawX), fix64.Exp2Fast(v.RawY), fix64.Exp2Fast(v.RawZ))
}

func (v F64Vec3) Exp2Fastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.Exp2Fastest(v.RawX), fix64.Exp2Fastest(v.RawY), fix64.Exp2Fastest(v.RawZ))
}

func (v F64Vec3) Log() F64Vec3 {
    return F64Vec3FromRaw(fix64.Log(v.RawX), fix64.Log(v.RawY), fix64.Log(v.RawZ))
}

func (v F64Vec3) LogFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.LogFast(v.RawX), fix64.LogFast(v.RawY), fix64.LogFast(v.RawZ))
}

func (v F64Vec3) LogFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.LogFastest(v.RawX), fix64.LogFastest(v.RawY), fix64.LogFastest(v.RawZ))
}

func (v F64Vec3) Log2() F64Vec3 {
    return F64Vec3FromRaw(fix64.Log2(v.RawX), fix64.Log2(v.RawY), fix64.Log2(v.RawZ))
}

func (v F64Vec3) Log2Fast() F64Vec3 {
    return F64Vec3FromRaw(fix64.Log2Fast(v.RawX), fix64.Log2Fast(v.RawY), fix64.Log2Fast(v.RawZ))
}

func (v F64Vec3) Log2Fastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.Log2Fastest(v.RawX), fix64.Log2Fastest(v.RawY), fix64.Log2Fastest(v.RawZ))
}

func (v F64Vec3) Sin() F64Vec3 {
    return F64Vec3FromRaw(fix64.Sin(v.RawX), fix64.Sin(v.RawY), fix64.Sin(v.RawZ))
}

func (v F64Vec3) SinFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.SinFast(v.RawX), fix64.SinFast(v.RawY), fix64.SinFast(v.RawZ))
}

func (v F64Vec3) SinFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.SinFastest(v.RawX), fix64.SinFastest(v.RawY), fix64.SinFastest(v.RawZ))
}

func (v F64Vec3) Cos() F64Vec3 {
    return F64Vec3FromRaw(fix64.Cos(v.RawX), fix64.Cos(v.RawY), fix64.Cos(v.RawZ))
}

func (v F64Vec3) CosFast() F64Vec3 {
    return F64Vec3FromRaw(fix64.CosFast(v.RawX), fix64.CosFast(v.RawY), fix64.CosFast(v.RawZ))
}

func (v F64Vec3) CosFastest() F64Vec3 {
    return F64Vec3FromRaw(fix64.CosFastest(v.RawX), fix64.CosFastest(v.RawY), fix64.CosFastest(v.RawZ))
}

func (v F64Vec3) Pow(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Pow(v.RawX, b.RawX), fix64.Pow(v.RawY, b.RawY), fix64.Pow(v.RawZ, b.RawZ))
}

func (v F64Vec3) PowFast(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.PowFast(v.RawX, b.RawX), fix64.PowFast(v.RawY, b.RawY), fix64.PowFast(v.RawZ, b.RawZ))
}

func (v F64Vec3) PowFastest(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.PowFastest(v.RawX, b.RawX), fix64.PowFastest(v.RawY, b.RawY), fix64.PowFastest(v.RawZ, b.RawZ))
}

func (v F64Vec3) Length() F64 {
    return F64FromRaw(fix64.Sqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
}

func (v F64Vec3) LengthFast() F64 {
    return F64FromRaw(fix64.SqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
}

func (v F64Vec3) LengthFastest() F64 {
    return F64FromRaw(fix64.SqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
}

func (v F64Vec3) LengthSqr() F64 {
    return F64FromRaw(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ))
}

func (v F64Vec3) Normalize() F64Vec3 {
    ooLen := F64FromRaw(fix64.RSqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
    return F64Vec3FromF64(ooLen, ooLen, ooLen).Mul(v)
}

func (v F64Vec3) NormalizeFast() F64Vec3 {
    ooLen := F64FromRaw(fix64.RSqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
    return F64Vec3FromF64(ooLen, ooLen, ooLen).Mul(v)
}

func (v F64Vec3) NormalizeFastest() F64Vec3 {
    ooLen := F64FromRaw(fix64.RSqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
    return F64Vec3FromF64(ooLen, ooLen, ooLen).Mul(v)
}

func (v F64Vec3) Dot(b F64Vec3) F64 {
    return F64FromRaw(fix64.Mul(v.RawX, b.RawX) + fix64.Mul(v.RawY, b.RawY) + fix64.Mul(v.RawZ, b.RawZ))
}

func (v F64Vec3) Distance(b F64Vec3) F64 {
    return v.Sub(b).Length()
}

func (v F64Vec3) DistanceFast(b F64Vec3) F64 {
    return v.Sub(b).LengthFast()
}

func (v F64Vec3) DistanceFastest(b F64Vec3) F64 {
    return v.Sub(b).LengthFastest()
}

func (v F64Vec3) Clamp(min, max F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Clamp(v.RawX, min.RawX, max.RawX), fix64.Clamp(v.RawY, min.RawY, max.RawY), fix64.Clamp(v.RawZ, min.RawZ, max.RawZ))
}

func (v F64Vec3) Lerp(b F64Vec3, t F64) F64Vec3 {
    tb := t.Raw
    ta := fix64.One - tb
    return F64Vec3FromRaw(fix64.Mul(v.RawX, ta)+fix64.Mul(b.RawX, tb), fix64.Mul(v.RawY, ta)+fix64.Mul(b.RawY, tb), fix64.Mul(v.RawZ, ta)+fix64.Mul(b.RawZ, tb))
}

func (v F64Vec3) Cross(b F64Vec3) F64Vec3 {
    return F64Vec3FromRaw(fix64.Mul(v.RawY, b.RawZ)-fix64.Mul(v.RawZ, b.RawY), fix64.Mul(v.RawZ, b.RawX)-fix64.Mul(v.RawX, b.RawZ), fix64.Mul(v.RawX, b.RawY)-fix64.Mul(v.RawY, b.RawX))
}

func (v F64Vec3) Equals(obj F64Vec3) bool {
    return reflect.DeepEqual(v, obj)
}

func (v F64Vec3) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s)`, fix64.ToString(v.RawX), fix64.ToString(v.RawY), fix64.ToString(v.RawZ))
}
