package fp

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix32"
)

var (
    F32Vec3Zero    = F32Vec3FromRaw(fix32.Zero, fix32.Zero, fix32.Zero)
    F32Vec3One     = F32Vec3FromRaw(fix32.One, fix32.One, fix32.One)
    F32Vec3Down    = F32Vec3FromRaw(fix32.Zero, fix32.Neg1, fix32.Zero)
    F32Vec3Up      = F32Vec3FromRaw(fix32.Zero, fix32.One, fix32.Zero)
    F32Vec3Left    = F32Vec3FromRaw(fix32.Neg1, fix32.Zero, fix32.Zero)
    F32Vec3Right   = F32Vec3FromRaw(fix32.One, fix32.Zero, fix32.Zero)
    F32Vec3Forward = F32Vec3FromRaw(fix32.Zero, fix32.Zero, fix32.One)
    F32Vec3Back    = F32Vec3FromRaw(fix32.Zero, fix32.Zero, fix32.Neg1)
    F32Vec3AxisX   = F32Vec3FromRaw(fix32.One, fix32.Zero, fix32.Zero)
    F32Vec3AxisY   = F32Vec3FromRaw(fix32.Zero, fix32.One, fix32.Zero)
    F32Vec3AxisZ   = F32Vec3FromRaw(fix32.Zero, fix32.Zero, fix32.One)
)

// F32Vec3 struct with signed 16.16 fixed point components.
type F32Vec3 struct {
    RawX int32
    RawY int32
    RawZ int32
}

func F32Vec3FromRaw(rawX, rawY, RawZ int32) F32Vec3 {
    return F32Vec3{
        RawX: rawX,
        RawY: rawY,
        RawZ: RawZ,
    }
}

func F32Vec3FromF32(x, y, z F32) F32Vec3 {
    return F32Vec3FromRaw(x.Raw, y.Raw, z.Raw)
}

func F32Vec3FromInt32(x, y, z int32) F32Vec3 {
    return F32Vec3FromRaw(fix32.FromInt32(x), fix32.FromInt32(y), fix32.FromInt32(z))
}

func F32Vec3FromFloat32(x, y, z float32) F32Vec3 {
    return F32Vec3FromRaw(fix32.FromFloat32(x), fix32.FromFloat32(y), fix32.FromFloat32(z))
}

func F32Vec3FromFloat64(x, y, z float64) F32Vec3 {
    return F32Vec3FromRaw(fix32.FromFloat64(x), fix32.FromFloat64(y), fix32.FromFloat64(z))
}

// F32Vec3Min returns the smallest F32Vec3 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Vec3Min(v0, v1)
//
// This makes it harder to accidentally call F32Vec3Min with 0 arguments.
func F32Vec3Min(v0 F32Vec3, v1 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Min(v0.RawX, v1.RawX), fix32.Min(v0.RawY, v1.RawY), fix32.Min(v0.RawZ, v1.RawZ))
}

// F32Vec3Max returns the largest F32Vec3 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Vec3Max(v0, v1)
//
// This makes it harder to accidentally call F32Vec3Max with 0 arguments.
func F32Vec3Max(v0 F32Vec3, v1 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Max(v0.RawX, v1.RawX), fix32.Max(v0.RawY, v1.RawY), fix32.Min(v0.RawZ, v1.RawZ))
}

func (v F32Vec3) X() F32 {
    return F32FromRaw(v.RawX)
}

func (v F32Vec3) Y() F32 {
    return F32FromRaw(v.RawY)
}

func (v F32Vec3) Z() F32 {
    return F32FromRaw(v.RawZ)
}

// Negate -v
func (v F32Vec3) Negate() F32Vec3 {
    return F32Vec3FromRaw(-v.RawX, -v.RawY, -v.RawZ)
}

// Add v + b
func (v F32Vec3) Add(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(v.RawX+b.RawX, v.RawY+b.RawY, v.RawZ+b.RawZ)
}

// Sub v - b
func (v F32Vec3) Sub(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(v.RawX-b.RawX, v.RawY-b.RawY, v.RawZ-b.RawZ)
}

// Mul v * b
func (v F32Vec3) Mul(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Mul(v.RawX, b.RawX), fix32.Mul(v.RawY, b.RawY), fix32.Mul(v.RawZ, b.RawZ))
}

// DivPrecise v / b
func (v F32Vec3) DivPrecise(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivPrecise(v.RawX, b.RawX), fix32.DivPrecise(v.RawY, b.RawY), fix32.DivPrecise(v.RawZ, b.RawZ))
}

// Mod v % b
func (v F32Vec3) Mod(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(v.RawX%b.RawX, v.RawY%b.RawY, v.RawZ%b.RawZ)
}

// AddF32 v + b
func (v F32Vec3) AddF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(v.RawX+b.Raw, v.RawY+b.Raw, v.RawZ+b.Raw)
}

// SubF32 v - b
func (v F32Vec3) SubF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(v.RawX-b.Raw, v.RawY-b.Raw, v.RawZ-b.Raw)
}

// MulF32 v * b
func (v F32Vec3) MulF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(fix32.Mul(v.RawX, b.Raw), fix32.Mul(v.RawY, b.Raw), fix32.Mul(v.RawZ, b.Raw))
}

// DivPreciseF32 v / b
func (v F32Vec3) DivPreciseF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivPrecise(v.RawX, b.Raw), fix32.DivPrecise(v.RawY, b.Raw), fix32.DivPrecise(v.RawZ, b.Raw))
}

// ModF32 v % b
func (v F32Vec3) ModF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(v.RawX%b.Raw, v.RawY%b.Raw, v.RawZ%b.Raw)
}

// EQ v == b
func (v F32Vec3) EQ(b F32Vec3) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY && v.RawZ == b.RawZ
}

// NE v != b
func (v F32Vec3) NE(b F32Vec3) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY || v.RawZ != b.RawZ
}

func (v F32Vec3) DivF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(fix32.Div(v.RawX, b.Raw), fix32.Div(v.RawY, b.Raw), fix32.Div(v.RawZ, b.Raw))
}

func (v F32Vec3) DivFastF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivFast(v.RawX, b.Raw), fix32.DivFast(v.RawY, b.Raw), fix32.DivFast(v.RawZ, b.Raw))
}

func (v F32Vec3) DivFastestF32(b F32) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivFastest(v.RawX, b.Raw), fix32.DivFastest(v.RawY, b.Raw), fix32.DivFastest(v.RawZ, b.Raw))
}

func (v F32Vec3) Div(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Div(v.RawX, b.RawX), fix32.Div(v.RawY, b.RawY), fix32.Div(v.RawZ, b.RawZ))
}

func (v F32Vec3) DivFast(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivFast(v.RawX, b.RawX), fix32.DivFast(v.RawY, b.RawY), fix32.DivFast(v.RawZ, b.RawZ))
}

func (v F32Vec3) DivFastest(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivFastest(v.RawX, b.RawX), fix32.DivFastest(v.RawY, b.RawY), fix32.DivFastest(v.RawZ, b.RawZ))
}

func (v F32Vec3) SqrtPrecise() F32Vec3 {
    return F32Vec3FromRaw(fix32.SqrtPrecise(v.RawX), fix32.SqrtPrecise(v.RawY), fix32.SqrtPrecise(v.RawZ))
}

func (v F32Vec3) Sqrt() F32Vec3 {
    return F32Vec3FromRaw(fix32.Sqrt(v.RawX), fix32.Sqrt(v.RawY), fix32.Sqrt(v.RawZ))
}

func (v F32Vec3) SqrtFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.SqrtFast(v.RawX), fix32.SqrtFast(v.RawY), fix32.SqrtFast(v.RawZ))
}

func (v F32Vec3) SqrtFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.SqrtFastest(v.RawX), fix32.SqrtFastest(v.RawY), fix32.SqrtFastest(v.RawZ))
}

func (v F32Vec3) RSqrt() F32Vec3 {
    return F32Vec3FromRaw(fix32.RSqrt(v.RawX), fix32.RSqrt(v.RawY), fix32.RSqrt(v.RawZ))
}

func (v F32Vec3) RSqrtFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.RSqrtFast(v.RawX), fix32.RSqrtFast(v.RawY), fix32.RSqrtFast(v.RawZ))
}

func (v F32Vec3) RSqrtFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.RSqrtFastest(v.RawX), fix32.RSqrtFastest(v.RawY), fix32.RSqrtFastest(v.RawZ))
}

func (v F32Vec3) Rcp() F32Vec3 {
    return F32Vec3FromRaw(fix32.Rcp(v.RawX), fix32.Rcp(v.RawY), fix32.Rcp(v.RawZ))
}

func (v F32Vec3) RcpFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.RcpFast(v.RawX), fix32.RcpFast(v.RawY), fix32.RcpFast(v.RawZ))
}

func (v F32Vec3) RcpFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.RcpFastest(v.RawX), fix32.RcpFastest(v.RawY), fix32.RcpFastest(v.RawZ))
}

func (v F32Vec3) Exp() F32Vec3 {
    return F32Vec3FromRaw(fix32.Exp(v.RawX), fix32.Exp(v.RawY), fix32.Exp(v.RawZ))
}

func (v F32Vec3) ExpFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.ExpFast(v.RawX), fix32.ExpFast(v.RawY), fix32.ExpFast(v.RawZ))
}

func (v F32Vec3) ExpFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.ExpFastest(v.RawX), fix32.ExpFastest(v.RawY), fix32.ExpFastest(v.RawZ))
}

func (v F32Vec3) Exp2() F32Vec3 {
    return F32Vec3FromRaw(fix32.Exp2(v.RawX), fix32.Exp2(v.RawY), fix32.Exp2(v.RawZ))
}

func (v F32Vec3) Exp2Fast() F32Vec3 {
    return F32Vec3FromRaw(fix32.Exp2Fast(v.RawX), fix32.Exp2Fast(v.RawY), fix32.Exp2Fast(v.RawZ))
}

func (v F32Vec3) Exp2Fastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.Exp2Fastest(v.RawX), fix32.Exp2Fastest(v.RawY), fix32.Exp2Fastest(v.RawZ))
}

func (v F32Vec3) Log() F32Vec3 {
    return F32Vec3FromRaw(fix32.Log(v.RawX), fix32.Log(v.RawY), fix32.Log(v.RawZ))
}

func (v F32Vec3) LogFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.LogFast(v.RawX), fix32.LogFast(v.RawY), fix32.LogFast(v.RawZ))
}

func (v F32Vec3) LogFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.LogFastest(v.RawX), fix32.LogFastest(v.RawY), fix32.LogFastest(v.RawZ))
}

func (v F32Vec3) Log2() F32Vec3 {
    return F32Vec3FromRaw(fix32.Log2(v.RawX), fix32.Log2(v.RawY), fix32.Log2(v.RawZ))
}

func (v F32Vec3) Log2Fast() F32Vec3 {
    return F32Vec3FromRaw(fix32.Log2Fast(v.RawX), fix32.Log2Fast(v.RawY), fix32.Log2Fast(v.RawZ))
}

func (v F32Vec3) Log2Fastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.Log2Fastest(v.RawX), fix32.Log2Fastest(v.RawY), fix32.Log2Fastest(v.RawZ))
}

func (v F32Vec3) Sin() F32Vec3 {
    return F32Vec3FromRaw(fix32.Sin(v.RawX), fix32.Sin(v.RawY), fix32.Sin(v.RawZ))
}

func (v F32Vec3) SinFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.SinFast(v.RawX), fix32.SinFast(v.RawY), fix32.SinFast(v.RawZ))
}

func (v F32Vec3) SinFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.SinFastest(v.RawX), fix32.SinFastest(v.RawY), fix32.SinFastest(v.RawZ))
}

func (v F32Vec3) Cos() F32Vec3 {
    return F32Vec3FromRaw(fix32.Cos(v.RawX), fix32.Cos(v.RawY), fix32.Cos(v.RawZ))
}

func (v F32Vec3) CosFast() F32Vec3 {
    return F32Vec3FromRaw(fix32.CosFast(v.RawX), fix32.CosFast(v.RawY), fix32.CosFast(v.RawZ))
}

func (v F32Vec3) CosFastest() F32Vec3 {
    return F32Vec3FromRaw(fix32.CosFastest(v.RawX), fix32.CosFastest(v.RawY), fix32.CosFastest(v.RawZ))
}

func (v F32Vec3) Pow(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Pow(v.RawX, b.RawX), fix32.Pow(v.RawY, b.RawY), fix32.Pow(v.RawZ, b.RawZ))
}

func (v F32Vec3) PowFast(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.PowFast(v.RawX, b.RawX), fix32.PowFast(v.RawY, b.RawY), fix32.PowFast(v.RawZ, b.RawZ))
}

func (v F32Vec3) PowFastest(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.PowFastest(v.RawX, b.RawX), fix32.PowFastest(v.RawY, b.RawY), fix32.PowFastest(v.RawZ, b.RawZ))
}

func (v F32Vec3) Length() F32 {
    return F32FromRaw(fix32.Sqrt(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ)))
}

func (v F32Vec3) LengthFast() F32 {
    return F32FromRaw(fix32.SqrtFast(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ)))
}

func (v F32Vec3) LengthFastest() F32 {
    return F32FromRaw(fix32.SqrtFastest(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ)))
}

func (v F32Vec3) LengthSqr() F32 {
    return F32FromRaw(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ))
}

func (v F32Vec3) Normalize() F32Vec3 {
    ooLen := F32FromRaw(fix32.RSqrt(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ)))
    return F32Vec3FromF32(ooLen, ooLen, ooLen).Mul(v)
}

func (v F32Vec3) NormalizeFast() F32Vec3 {
    ooLen := F32FromRaw(fix32.RSqrtFast(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ)))
    return F32Vec3FromF32(ooLen, ooLen, ooLen).Mul(v)
}

func (v F32Vec3) NormalizeFastest() F32Vec3 {
    ooLen := F32FromRaw(fix32.RSqrtFastest(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ)))
    return F32Vec3FromF32(ooLen, ooLen, ooLen).Mul(v)
}

func (v F32Vec3) Dot(b F32Vec3) F32 {
    return F32FromRaw(fix32.Mul(v.RawX, b.RawX) + fix32.Mul(v.RawY, b.RawY) + fix32.Mul(v.RawZ, b.RawZ))
}

func (v F32Vec3) Distance(b F32Vec3) F32 {
    return v.Sub(b).Length()
}

func (v F32Vec3) DistanceFast(b F32Vec3) F32 {
    return v.Sub(b).LengthFast()
}

func (v F32Vec3) DistanceFastest(b F32Vec3) F32 {
    return v.Sub(b).LengthFastest()
}

func (v F32Vec3) Clamp(min, max F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Clamp(v.RawX, min.RawX, max.RawX), fix32.Clamp(v.RawY, min.RawY, max.RawY), fix32.Clamp(v.RawZ, min.RawZ, max.RawZ))
}

func (v F32Vec3) Lerp(b F32Vec3, t F32) F32Vec3 {
    tb := t.Raw
    ta := fix32.One - tb
    return F32Vec3FromRaw(fix32.Mul(v.RawX, ta)+fix32.Mul(b.RawX, tb), fix32.Mul(v.RawY, ta)+fix32.Mul(b.RawY, tb), fix32.Mul(v.RawZ, ta)+fix32.Mul(b.RawZ, tb))
}

func (v F32Vec3) Cross(b F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Mul(v.RawY, b.RawZ)-fix32.Mul(v.RawZ, b.RawY), fix32.Mul(v.RawZ, b.RawX)-fix32.Mul(v.RawX, b.RawZ), fix32.Mul(v.RawX, b.RawY)-fix32.Mul(v.RawY, b.RawX))
}

func (v F32Vec3) Equals(obj F32Vec3) bool {
    return reflect.DeepEqual(v, obj)
}

func (v F32Vec3) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s)`, fix32.ToString(v.RawX), fix32.ToString(v.RawY), fix32.ToString(v.RawZ))
}
