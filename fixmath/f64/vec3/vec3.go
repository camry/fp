package vec3

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix64"
    "github.com/camry/fp/fixmath/f64"
)

var (
    Zero    = FromRaw(fix64.Zero, fix64.Zero, fix64.Zero)
    One     = FromRaw(fix64.One, fix64.One, fix64.One)
    Down    = FromRaw(fix64.Zero, fix64.Neg1, fix64.Zero)
    Up      = FromRaw(fix64.Zero, fix64.One, fix64.Zero)
    Left    = FromRaw(fix64.Neg1, fix64.Zero, fix64.Zero)
    Right   = FromRaw(fix64.One, fix64.Zero, fix64.Zero)
    Forward = FromRaw(fix64.Zero, fix64.Zero, fix64.One)
    Back    = FromRaw(fix64.Zero, fix64.Zero, fix64.Neg1)
    AxisX   = FromRaw(fix64.One, fix64.Zero, fix64.Zero)
    AxisY   = FromRaw(fix64.Zero, fix64.One, fix64.Zero)
    AxisZ   = FromRaw(fix64.Zero, fix64.Zero, fix64.One)
)

// Vec3 struct with signed 32.32 fixed point components.
type Vec3 struct {
    RawX int64
    RawY int64
    RawZ int64
}

func FromRaw(rawX, rawY, RawZ int64) Vec3 {
    return Vec3{
        RawX: rawX,
        RawY: rawY,
        RawZ: RawZ,
    }
}

func FromF64(x, y, z f64.F64) Vec3 {
    return FromRaw(x.Raw, y.Raw, z.Raw)
}

func FromInt32(x, y, z int32) Vec3 {
    return FromRaw(fix64.FromInt32(x), fix64.FromInt32(y), fix64.FromInt32(z))
}

func FromInt64(x, y, z int64) Vec3 {
    return FromRaw(fix64.FromInt64(x), fix64.FromInt64(y), fix64.FromInt64(z))
}

func FromFloat32(x, y, z float32) Vec3 {
    return FromRaw(fix64.FromFloat32(x), fix64.FromFloat32(y), fix64.FromFloat32(z))
}

func FromFloat64(x, y, z float64) Vec3 {
    return FromRaw(fix64.FromFloat64(x), fix64.FromFloat64(y), fix64.FromFloat64(z))
}

// Min returns the smallest Vec3 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Min(v0, v1)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(v0 Vec3, v1 Vec3) Vec3 {
    return FromRaw(fix64.Min(v0.RawX, v1.RawX), fix64.Min(v0.RawY, v1.RawY), fix64.Min(v0.RawZ, v1.RawZ))
}

// Max returns the largest Vec3 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Max(v0, v1)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(v0 Vec3, v1 Vec3) Vec3 {
    return FromRaw(fix64.Max(v0.RawX, v1.RawX), fix64.Max(v0.RawY, v1.RawY), fix64.Min(v0.RawZ, v1.RawZ))
}

func (v Vec3) X() f64.F64 {
    return f64.FromRaw(v.RawX)
}

func (v Vec3) Y() f64.F64 {
    return f64.FromRaw(v.RawY)
}

func (v Vec3) Z() f64.F64 {
    return f64.FromRaw(v.RawZ)
}

// Negate -v
func (v Vec3) Negate() Vec3 {
    return FromRaw(-v.RawX, -v.RawY, -v.RawZ)
}

// Add v + b
func (v Vec3) Add(b Vec3) Vec3 {
    return FromRaw(v.RawX+b.RawX, v.RawY+b.RawY, v.RawZ+b.RawZ)
}

// Sub v - b
func (v Vec3) Sub(b Vec3) Vec3 {
    return FromRaw(v.RawX-b.RawX, v.RawY-b.RawY, v.RawZ-b.RawZ)
}

// Mul v * b
func (v Vec3) Mul(b Vec3) Vec3 {
    return FromRaw(fix64.Mul(v.RawX, b.RawX), fix64.Mul(v.RawY, b.RawY), fix64.Mul(v.RawZ, b.RawZ))
}

// DivPrecise v / b
func (v Vec3) DivPrecise(b Vec3) Vec3 {
    return FromRaw(fix64.DivPrecise(v.RawX, b.RawX), fix64.DivPrecise(v.RawY, b.RawY), fix64.DivPrecise(v.RawZ, b.RawZ))
}

// Mod v % b
func (v Vec3) Mod(b Vec3) Vec3 {
    return FromRaw(v.RawX%b.RawX, v.RawY%b.RawY, v.RawZ%b.RawZ)
}

// AddF64 v + b
func (v Vec3) AddF64(b f64.F64) Vec3 {
    return FromRaw(v.RawX+b.Raw, v.RawY+b.Raw, v.RawZ+b.Raw)
}

// SubF64 v - b
func (v Vec3) SubF64(b f64.F64) Vec3 {
    return FromRaw(v.RawX-b.Raw, v.RawY-b.Raw, v.RawZ-b.Raw)
}

// MulF64 v * b
func (v Vec3) MulF64(b f64.F64) Vec3 {
    return FromRaw(fix64.Mul(v.RawX, b.Raw), fix64.Mul(v.RawY, b.Raw), fix64.Mul(v.RawZ, b.Raw))
}

// DivPreciseF64 v / b
func (v Vec3) DivPreciseF64(b f64.F64) Vec3 {
    return FromRaw(fix64.DivPrecise(v.RawX, b.Raw), fix64.DivPrecise(v.RawY, b.Raw), fix64.DivPrecise(v.RawZ, b.Raw))
}

// ModF64 v % b
func (v Vec3) ModF64(b f64.F64) Vec3 {
    return FromRaw(v.RawX%b.Raw, v.RawY%b.Raw, v.RawZ%b.Raw)
}

// EQ v == b
func (v Vec3) EQ(b Vec3) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY && v.RawZ == b.RawZ
}

// NE v != b
func (v Vec3) NE(b Vec3) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY || v.RawZ != b.RawZ
}

func (v Vec3) DivF64(b f64.F64) Vec3 {
    return FromRaw(fix64.Div(v.RawX, b.Raw), fix64.Div(v.RawY, b.Raw), fix64.Div(v.RawZ, b.Raw))
}

func (v Vec3) DivFastF64(b f64.F64) Vec3 {
    return FromRaw(fix64.DivFast(v.RawX, b.Raw), fix64.DivFast(v.RawY, b.Raw), fix64.DivFast(v.RawZ, b.Raw))
}

func (v Vec3) DivFastestF64(b f64.F64) Vec3 {
    return FromRaw(fix64.DivFastest(v.RawX, b.Raw), fix64.DivFastest(v.RawY, b.Raw), fix64.DivFastest(v.RawZ, b.Raw))
}

func (v Vec3) Div(b Vec3) Vec3 {
    return FromRaw(fix64.Div(v.RawX, b.RawX), fix64.Div(v.RawY, b.RawY), fix64.Div(v.RawZ, b.RawZ))
}

func (v Vec3) DivFast(b Vec3) Vec3 {
    return FromRaw(fix64.DivFast(v.RawX, b.RawX), fix64.DivFast(v.RawY, b.RawY), fix64.DivFast(v.RawZ, b.RawZ))
}

func (v Vec3) DivFastest(b Vec3) Vec3 {
    return FromRaw(fix64.DivFastest(v.RawX, b.RawX), fix64.DivFastest(v.RawY, b.RawY), fix64.DivFastest(v.RawZ, b.RawZ))
}

func (v Vec3) SqrtPrecise() Vec3 {
    return FromRaw(fix64.SqrtPrecise(v.RawX), fix64.SqrtPrecise(v.RawY), fix64.SqrtPrecise(v.RawZ))
}

func (v Vec3) Sqrt() Vec3 {
    return FromRaw(fix64.Sqrt(v.RawX), fix64.Sqrt(v.RawY), fix64.Sqrt(v.RawZ))
}

func (v Vec3) SqrtFast() Vec3 {
    return FromRaw(fix64.SqrtFast(v.RawX), fix64.SqrtFast(v.RawY), fix64.SqrtFast(v.RawZ))
}

func (v Vec3) SqrtFastest() Vec3 {
    return FromRaw(fix64.SqrtFastest(v.RawX), fix64.SqrtFastest(v.RawY), fix64.SqrtFastest(v.RawZ))
}

func (v Vec3) RSqrt() Vec3 {
    return FromRaw(fix64.RSqrt(v.RawX), fix64.RSqrt(v.RawY), fix64.RSqrt(v.RawZ))
}

func (v Vec3) RSqrtFast() Vec3 {
    return FromRaw(fix64.RSqrtFast(v.RawX), fix64.RSqrtFast(v.RawY), fix64.RSqrtFast(v.RawZ))
}

func (v Vec3) RSqrtFastest() Vec3 {
    return FromRaw(fix64.RSqrtFastest(v.RawX), fix64.RSqrtFastest(v.RawY), fix64.RSqrtFastest(v.RawZ))
}

func (v Vec3) Rcp() Vec3 {
    return FromRaw(fix64.Rcp(v.RawX), fix64.Rcp(v.RawY), fix64.Rcp(v.RawZ))
}

func (v Vec3) RcpFast() Vec3 {
    return FromRaw(fix64.RcpFast(v.RawX), fix64.RcpFast(v.RawY), fix64.RcpFast(v.RawZ))
}

func (v Vec3) RcpFastest() Vec3 {
    return FromRaw(fix64.RcpFastest(v.RawX), fix64.RcpFastest(v.RawY), fix64.RcpFastest(v.RawZ))
}

func (v Vec3) Exp() Vec3 {
    return FromRaw(fix64.Exp(v.RawX), fix64.Exp(v.RawY), fix64.Exp(v.RawZ))
}

func (v Vec3) ExpFast() Vec3 {
    return FromRaw(fix64.ExpFast(v.RawX), fix64.ExpFast(v.RawY), fix64.ExpFast(v.RawZ))
}

func (v Vec3) ExpFastest() Vec3 {
    return FromRaw(fix64.ExpFastest(v.RawX), fix64.ExpFastest(v.RawY), fix64.ExpFastest(v.RawZ))
}

func (v Vec3) Exp2() Vec3 {
    return FromRaw(fix64.Exp2(v.RawX), fix64.Exp2(v.RawY), fix64.Exp2(v.RawZ))
}

func (v Vec3) Exp2Fast() Vec3 {
    return FromRaw(fix64.Exp2Fast(v.RawX), fix64.Exp2Fast(v.RawY), fix64.Exp2Fast(v.RawZ))
}

func (v Vec3) Exp2Fastest() Vec3 {
    return FromRaw(fix64.Exp2Fastest(v.RawX), fix64.Exp2Fastest(v.RawY), fix64.Exp2Fastest(v.RawZ))
}

func (v Vec3) Log() Vec3 {
    return FromRaw(fix64.Log(v.RawX), fix64.Log(v.RawY), fix64.Log(v.RawZ))
}

func (v Vec3) LogFast() Vec3 {
    return FromRaw(fix64.LogFast(v.RawX), fix64.LogFast(v.RawY), fix64.LogFast(v.RawZ))
}

func (v Vec3) LogFastest() Vec3 {
    return FromRaw(fix64.LogFastest(v.RawX), fix64.LogFastest(v.RawY), fix64.LogFastest(v.RawZ))
}

func (v Vec3) Log2() Vec3 {
    return FromRaw(fix64.Log2(v.RawX), fix64.Log2(v.RawY), fix64.Log2(v.RawZ))
}

func (v Vec3) Log2Fast() Vec3 {
    return FromRaw(fix64.Log2Fast(v.RawX), fix64.Log2Fast(v.RawY), fix64.Log2Fast(v.RawZ))
}

func (v Vec3) Log2Fastest() Vec3 {
    return FromRaw(fix64.Log2Fastest(v.RawX), fix64.Log2Fastest(v.RawY), fix64.Log2Fastest(v.RawZ))
}

func (v Vec3) Sin() Vec3 {
    return FromRaw(fix64.Sin(v.RawX), fix64.Sin(v.RawY), fix64.Sin(v.RawZ))
}

func (v Vec3) SinFast() Vec3 {
    return FromRaw(fix64.SinFast(v.RawX), fix64.SinFast(v.RawY), fix64.SinFast(v.RawZ))
}

func (v Vec3) SinFastest() Vec3 {
    return FromRaw(fix64.SinFastest(v.RawX), fix64.SinFastest(v.RawY), fix64.SinFastest(v.RawZ))
}

func (v Vec3) Cos() Vec3 {
    return FromRaw(fix64.Cos(v.RawX), fix64.Cos(v.RawY), fix64.Cos(v.RawZ))
}

func (v Vec3) CosFast() Vec3 {
    return FromRaw(fix64.CosFast(v.RawX), fix64.CosFast(v.RawY), fix64.CosFast(v.RawZ))
}

func (v Vec3) CosFastest() Vec3 {
    return FromRaw(fix64.CosFastest(v.RawX), fix64.CosFastest(v.RawY), fix64.CosFastest(v.RawZ))
}

func (v Vec3) Pow(b Vec3) Vec3 {
    return FromRaw(fix64.Pow(v.RawX, b.RawX), fix64.Pow(v.RawY, b.RawY), fix64.Pow(v.RawZ, b.RawZ))
}

func (v Vec3) PowFast(b Vec3) Vec3 {
    return FromRaw(fix64.PowFast(v.RawX, b.RawX), fix64.PowFast(v.RawY, b.RawY), fix64.PowFast(v.RawZ, b.RawZ))
}

func (v Vec3) PowFastest(b Vec3) Vec3 {
    return FromRaw(fix64.PowFastest(v.RawX, b.RawX), fix64.PowFastest(v.RawY, b.RawY), fix64.PowFastest(v.RawZ, b.RawZ))
}

func (v Vec3) Length() f64.F64 {
    return f64.FromRaw(fix64.Sqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
}

func (v Vec3) LengthFast() f64.F64 {
    return f64.FromRaw(fix64.SqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
}

func (v Vec3) LengthFastest() f64.F64 {
    return f64.FromRaw(fix64.SqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
}

func (v Vec3) LengthSqr() f64.F64 {
    return f64.FromRaw(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ))
}

func (v Vec3) Normalize() Vec3 {
    ooLen := f64.FromRaw(fix64.RSqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
    return FromF64(ooLen, ooLen, ooLen).Mul(v)
}

func (v Vec3) NormalizeFast() Vec3 {
    ooLen := f64.FromRaw(fix64.RSqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
    return FromF64(ooLen, ooLen, ooLen).Mul(v)
}

func (v Vec3) NormalizeFastest() Vec3 {
    ooLen := f64.FromRaw(fix64.RSqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ)))
    return FromF64(ooLen, ooLen, ooLen).Mul(v)
}

func (v Vec3) Dot(b Vec3) f64.F64 {
    return f64.FromRaw(fix64.Mul(v.RawX, b.RawX) + fix64.Mul(v.RawY, b.RawY) + fix64.Mul(v.RawZ, b.RawZ))
}

func (v Vec3) Distance(b Vec3) f64.F64 {
    return v.Sub(b).Length()
}

func (v Vec3) DistanceFast(b Vec3) f64.F64 {
    return v.Sub(b).LengthFast()
}

func (v Vec3) DistanceFastest(b Vec3) f64.F64 {
    return v.Sub(b).LengthFastest()
}

func (v Vec3) Clamp(min, max Vec3) Vec3 {
    return FromRaw(fix64.Clamp(v.RawX, min.RawX, max.RawX), fix64.Clamp(v.RawY, min.RawY, max.RawY), fix64.Clamp(v.RawZ, min.RawZ, max.RawZ))
}

func (v Vec3) Lerp(b Vec3, t f64.F64) Vec3 {
    tb := t.Raw
    ta := fix64.One - tb
    return FromRaw(fix64.Mul(v.RawX, ta)+fix64.Mul(b.RawX, tb), fix64.Mul(v.RawY, ta)+fix64.Mul(b.RawY, tb), fix64.Mul(v.RawZ, ta)+fix64.Mul(b.RawZ, tb))
}

func (v Vec3) Cross(b Vec3) Vec3 {
    return FromRaw(fix64.Mul(v.RawY, b.RawZ)-fix64.Mul(v.RawZ, b.RawY), fix64.Mul(v.RawZ, b.RawX)-fix64.Mul(v.RawX, b.RawZ), fix64.Mul(v.RawX, b.RawY)-fix64.Mul(v.RawY, b.RawX))
}

func (v Vec3) Equals(obj Vec3) bool {
    return reflect.DeepEqual(v, obj)
}

func (v Vec3) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s)`, fix64.ToString(v.RawX), fix64.ToString(v.RawY), fix64.ToString(v.RawZ))
}
