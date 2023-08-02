package fp

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix64"
)

var (
    Vec4Zero  = Vec4FromRaw(fix64.Zero, fix64.Zero, fix64.Zero, fix64.Zero)
    Vec4One   = Vec4FromRaw(fix64.One, fix64.One, fix64.One, fix64.One)
    Vec4AxisX = Vec4FromRaw(fix64.One, fix64.Zero, fix64.Zero, fix64.Zero)
    Vec4AxisY = Vec4FromRaw(fix64.Zero, fix64.One, fix64.Zero, fix64.Zero)
    Vec4AxisZ = Vec4FromRaw(fix64.Zero, fix64.Zero, fix64.One, fix64.Zero)
    Vec4AxisW = Vec4FromRaw(fix64.Zero, fix64.Zero, fix64.Zero, fix64.One)
)

// F64Vec4 struct with signed 16.16 fixed point components.
type F64Vec4 struct {
    RawX int64
    RawY int64
    RawZ int64
    RawW int64
}

func Vec4FromRaw(rawX, rawY, RawZ, RawW int64) F64Vec4 {
    return F64Vec4{
        RawX: rawX,
        RawY: rawY,
        RawZ: RawZ,
        RawW: RawW,
    }
}

func Vec4FromF64(x, y, z, w F64) F64Vec4 {
    return Vec4FromRaw(x.Raw, y.Raw, z.Raw, w.Raw)
}

func Vec4FromInt32(x, y, z, w int32) F64Vec4 {
    return Vec4FromRaw(fix64.FromInt32(x), fix64.FromInt32(y), fix64.FromInt32(z), fix64.FromInt32(w))
}

func Vec4FromInt64(x, y, z, w int64) F64Vec4 {
    return Vec4FromRaw(fix64.FromInt64(x), fix64.FromInt64(y), fix64.FromInt64(z), fix64.FromInt64(w))
}

func Vec4FromFloat32(x, y, z, w float32) F64Vec4 {
    return Vec4FromRaw(fix64.FromFloat32(x), fix64.FromFloat32(y), fix64.FromFloat32(z), fix64.FromFloat32(w))
}

func Vec4FromFloat64(x, y, z, w float64) F64Vec4 {
    return Vec4FromRaw(fix64.FromFloat64(x), fix64.FromFloat64(y), fix64.FromFloat64(z), fix64.FromFloat64(w))
}

// Vec4Min returns the smallest F64Vec4 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Vec4Min(v0, v1)
//
// This makes it harder to accidentally call Vec4Min with 0 arguments.
func Vec4Min(v0 F64Vec4, v1 F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Min(v0.RawX, v1.RawX), fix64.Min(v0.RawY, v1.RawY), fix64.Min(v0.RawZ, v1.RawZ), fix64.Min(v0.RawW, v1.RawW))
}

// Vec4Max returns the largest F64Vec4 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     Vec4Max(v0, v1)
//
// This makes it harder to accidentally call Vec4Max with 0 arguments.
func Vec4Max(v0 F64Vec4, v1 F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Max(v0.RawX, v1.RawX), fix64.Max(v0.RawY, v1.RawY), fix64.Max(v0.RawZ, v1.RawZ), fix64.Max(v0.RawW, v1.RawW))
}

func (v F64Vec4) X() F64 {
    return F64FromRaw(v.RawX)
}

func (v F64Vec4) Y() F64 {
    return F64FromRaw(v.RawY)
}

func (v F64Vec4) Z() F64 {
    return F64FromRaw(v.RawZ)
}

func (v F64Vec4) W() F64 {
    return F64FromRaw(v.RawW)
}

// Negate -v
func (v F64Vec4) Negate() F64Vec4 {
    return Vec4FromRaw(-v.RawX, -v.RawY, -v.RawZ, -v.RawW)
}

// Add v + b
func (v F64Vec4) Add(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(v.RawX+b.RawX, v.RawY+b.RawY, v.RawZ+b.RawZ, v.RawW+b.RawW)
}

// Sub v - b
func (v F64Vec4) Sub(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(v.RawX-b.RawX, v.RawY-b.RawY, v.RawZ-b.RawZ, v.RawW-b.RawW)
}

// Mul v * b
func (v F64Vec4) Mul(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Mul(v.RawX, b.RawX), fix64.Mul(v.RawY, b.RawY), fix64.Mul(v.RawZ, b.RawZ), fix64.Mul(v.RawW, b.RawW))
}

// DivPrecise v / b
func (v F64Vec4) DivPrecise(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.DivPrecise(v.RawX, b.RawX), fix64.DivPrecise(v.RawY, b.RawY), fix64.DivPrecise(v.RawZ, b.RawZ), fix64.DivPrecise(v.RawW, b.RawW))
}

// Mod v % b
func (v F64Vec4) Mod(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(v.RawX%b.RawX, v.RawY%b.RawY, v.RawZ%b.RawZ, v.RawW%b.RawW)
}

// AddF64 v + b
func (v F64Vec4) AddF64(b F64) F64Vec4 {
    return Vec4FromRaw(v.RawX+b.Raw, v.RawY+b.Raw, v.RawZ+b.Raw, v.RawW+b.Raw)
}

// SubF64 v - b
func (v F64Vec4) SubF64(b F64) F64Vec4 {
    return Vec4FromRaw(v.RawX-b.Raw, v.RawY-b.Raw, v.RawZ-b.Raw, v.RawW-b.Raw)
}

// MulF64 v * b
func (v F64Vec4) MulF64(b F64) F64Vec4 {
    return Vec4FromRaw(fix64.Mul(v.RawX, b.Raw), fix64.Mul(v.RawY, b.Raw), fix64.Mul(v.RawZ, b.Raw), fix64.Mul(v.RawW, b.Raw))
}

// DivPreciseF64 v / b
func (v F64Vec4) DivPreciseF64(b F64) F64Vec4 {
    return Vec4FromRaw(fix64.DivPrecise(v.RawX, b.Raw), fix64.DivPrecise(v.RawY, b.Raw), fix64.DivPrecise(v.RawZ, b.Raw), fix64.DivPrecise(v.RawW, b.Raw))
}

// ModF64 v % b
func (v F64Vec4) ModF64(b F64) F64Vec4 {
    return Vec4FromRaw(v.RawX%b.Raw, v.RawY%b.Raw, v.RawZ%b.Raw, v.RawW%b.Raw)
}

// EQ v == b
func (v F64Vec4) EQ(b F64Vec4) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY && v.RawZ == b.RawZ && v.RawW == b.RawW
}

// NE v != b
func (v F64Vec4) NE(b F64Vec4) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY || v.RawZ != b.RawZ || v.RawW != b.RawW
}

func (v F64Vec4) Div(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Div(v.RawX, b.RawX), fix64.Div(v.RawY, b.RawY), fix64.Div(v.RawZ, b.RawZ), fix64.Div(v.RawW, b.RawW))
}

func (v F64Vec4) DivFast(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.DivFast(v.RawX, b.RawX), fix64.DivFast(v.RawY, b.RawY), fix64.DivFast(v.RawZ, b.RawZ), fix64.DivFast(v.RawW, b.RawW))
}

func (v F64Vec4) DivFastest(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.DivFastest(v.RawX, b.RawX), fix64.DivFastest(v.RawY, b.RawY), fix64.DivFastest(v.RawZ, b.RawZ), fix64.DivFastest(v.RawW, b.RawW))
}

func (v F64Vec4) SqrtPrecise() F64Vec4 {
    return Vec4FromRaw(fix64.SqrtPrecise(v.RawX), fix64.SqrtPrecise(v.RawY), fix64.SqrtPrecise(v.RawZ), fix64.SqrtPrecise(v.RawW))
}

func (v F64Vec4) Sqrt() F64Vec4 {
    return Vec4FromRaw(fix64.Sqrt(v.RawX), fix64.Sqrt(v.RawY), fix64.Sqrt(v.RawZ), fix64.Sqrt(v.RawW))
}

func (v F64Vec4) SqrtFast() F64Vec4 {
    return Vec4FromRaw(fix64.SqrtFast(v.RawX), fix64.SqrtFast(v.RawY), fix64.SqrtFast(v.RawZ), fix64.SqrtFast(v.RawW))
}

func (v F64Vec4) SqrtFastest() F64Vec4 {
    return Vec4FromRaw(fix64.SqrtFastest(v.RawX), fix64.SqrtFastest(v.RawY), fix64.SqrtFastest(v.RawZ), fix64.SqrtFastest(v.RawW))
}

func (v F64Vec4) RSqrt() F64Vec4 {
    return Vec4FromRaw(fix64.RSqrt(v.RawX), fix64.RSqrt(v.RawY), fix64.RSqrt(v.RawZ), fix64.RSqrt(v.RawW))
}

func (v F64Vec4) RSqrtFast() F64Vec4 {
    return Vec4FromRaw(fix64.RSqrtFast(v.RawX), fix64.RSqrtFast(v.RawY), fix64.RSqrtFast(v.RawZ), fix64.RSqrtFast(v.RawW))
}

func (v F64Vec4) RSqrtFastest() F64Vec4 {
    return Vec4FromRaw(fix64.RSqrtFastest(v.RawX), fix64.RSqrtFastest(v.RawY), fix64.RSqrtFastest(v.RawZ), fix64.RSqrtFastest(v.RawW))
}

func (v F64Vec4) Rcp() F64Vec4 {
    return Vec4FromRaw(fix64.Rcp(v.RawX), fix64.Rcp(v.RawY), fix64.Rcp(v.RawZ), fix64.Rcp(v.RawW))
}

func (v F64Vec4) RcpFast() F64Vec4 {
    return Vec4FromRaw(fix64.RcpFast(v.RawX), fix64.RcpFast(v.RawY), fix64.RcpFast(v.RawZ), fix64.RcpFast(v.RawW))
}

func (v F64Vec4) RcpFastest() F64Vec4 {
    return Vec4FromRaw(fix64.RcpFastest(v.RawX), fix64.RcpFastest(v.RawY), fix64.RcpFastest(v.RawZ), fix64.RcpFastest(v.RawW))
}

func (v F64Vec4) Exp() F64Vec4 {
    return Vec4FromRaw(fix64.Exp(v.RawX), fix64.Exp(v.RawY), fix64.Exp(v.RawZ), fix64.Exp(v.RawW))
}

func (v F64Vec4) ExpFast() F64Vec4 {
    return Vec4FromRaw(fix64.ExpFast(v.RawX), fix64.ExpFast(v.RawY), fix64.ExpFast(v.RawZ), fix64.ExpFast(v.RawW))
}

func (v F64Vec4) ExpFastest() F64Vec4 {
    return Vec4FromRaw(fix64.ExpFastest(v.RawX), fix64.ExpFastest(v.RawY), fix64.ExpFastest(v.RawZ), fix64.ExpFastest(v.RawW))
}

func (v F64Vec4) Exp2() F64Vec4 {
    return Vec4FromRaw(fix64.Exp2(v.RawX), fix64.Exp2(v.RawY), fix64.Exp2(v.RawZ), fix64.Exp2(v.RawW))
}

func (v F64Vec4) Exp2Fast() F64Vec4 {
    return Vec4FromRaw(fix64.Exp2Fast(v.RawX), fix64.Exp2Fast(v.RawY), fix64.Exp2Fast(v.RawZ), fix64.Exp2Fast(v.RawW))
}

func (v F64Vec4) Exp2Fastest() F64Vec4 {
    return Vec4FromRaw(fix64.Exp2Fastest(v.RawX), fix64.Exp2Fastest(v.RawY), fix64.Exp2Fastest(v.RawZ), fix64.Exp2Fastest(v.RawW))
}

func (v F64Vec4) Log() F64Vec4 {
    return Vec4FromRaw(fix64.Log(v.RawX), fix64.Log(v.RawY), fix64.Log(v.RawZ), fix64.Log(v.RawW))
}

func (v F64Vec4) LogFast() F64Vec4 {
    return Vec4FromRaw(fix64.LogFast(v.RawX), fix64.LogFast(v.RawY), fix64.LogFast(v.RawZ), fix64.LogFast(v.RawW))
}

func (v F64Vec4) LogFastest() F64Vec4 {
    return Vec4FromRaw(fix64.LogFastest(v.RawX), fix64.LogFastest(v.RawY), fix64.LogFastest(v.RawZ), fix64.LogFastest(v.RawW))
}

func (v F64Vec4) Log2() F64Vec4 {
    return Vec4FromRaw(fix64.Log2(v.RawX), fix64.Log2(v.RawY), fix64.Log2(v.RawZ), fix64.Log2(v.RawW))
}

func (v F64Vec4) Log2Fast() F64Vec4 {
    return Vec4FromRaw(fix64.Log2Fast(v.RawX), fix64.Log2Fast(v.RawY), fix64.Log2Fast(v.RawZ), fix64.Log2Fast(v.RawW))
}

func (v F64Vec4) Log2Fastest() F64Vec4 {
    return Vec4FromRaw(fix64.Log2Fastest(v.RawX), fix64.Log2Fastest(v.RawY), fix64.Log2Fastest(v.RawZ), fix64.Log2Fastest(v.RawW))
}

func (v F64Vec4) Sin() F64Vec4 {
    return Vec4FromRaw(fix64.Sin(v.RawX), fix64.Sin(v.RawY), fix64.Sin(v.RawZ), fix64.Sin(v.RawW))
}

func (v F64Vec4) SinFast() F64Vec4 {
    return Vec4FromRaw(fix64.SinFast(v.RawX), fix64.SinFast(v.RawY), fix64.SinFast(v.RawZ), fix64.SinFast(v.RawW))
}

func (v F64Vec4) SinFastest() F64Vec4 {
    return Vec4FromRaw(fix64.SinFastest(v.RawX), fix64.SinFastest(v.RawY), fix64.SinFastest(v.RawZ), fix64.SinFastest(v.RawW))
}

func (v F64Vec4) Cos() F64Vec4 {
    return Vec4FromRaw(fix64.Cos(v.RawX), fix64.Cos(v.RawY), fix64.Cos(v.RawZ), fix64.Cos(v.RawW))
}

func (v F64Vec4) CosFast() F64Vec4 {
    return Vec4FromRaw(fix64.CosFast(v.RawX), fix64.CosFast(v.RawY), fix64.CosFast(v.RawZ), fix64.CosFast(v.RawW))
}

func (v F64Vec4) CosFastest() F64Vec4 {
    return Vec4FromRaw(fix64.CosFastest(v.RawX), fix64.CosFastest(v.RawY), fix64.CosFastest(v.RawZ), fix64.CosFastest(v.RawW))
}

func (v F64Vec4) Pow(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Pow(v.RawX, b.RawX), fix64.Pow(v.RawY, b.RawY), fix64.Pow(v.RawZ, b.RawZ), fix64.Pow(v.RawW, b.RawW))
}

func (v F64Vec4) PowFast(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.PowFast(v.RawX, b.RawX), fix64.PowFast(v.RawY, b.RawY), fix64.PowFast(v.RawZ, b.RawZ), fix64.PowFast(v.RawW, b.RawW))
}

func (v F64Vec4) PowFastest(b F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.PowFastest(v.RawX, b.RawX), fix64.PowFastest(v.RawY, b.RawY), fix64.PowFastest(v.RawZ, b.RawZ), fix64.PowFastest(v.RawW, b.RawW))
}

func (v F64Vec4) Length() F64 {
    return F64FromRaw(fix64.Sqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW)))
}

func (v F64Vec4) LengthFast() F64 {
    return F64FromRaw(fix64.SqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW)))
}

func (v F64Vec4) LengthFastest() F64 {
    return F64FromRaw(fix64.SqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW)))
}

func (v F64Vec4) LengthSqr() F64 {
    return F64FromRaw(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW))
}

func (v F64Vec4) Normalize() F64Vec4 {
    ooLen := F64FromRaw(fix64.RSqrt(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW)))
    return Vec4FromF64(ooLen, ooLen, ooLen, ooLen).Mul(v)
}

func (v F64Vec4) NormalizeFast() F64Vec4 {
    ooLen := F64FromRaw(fix64.RSqrtFast(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW)))
    return Vec4FromF64(ooLen, ooLen, ooLen, ooLen).Mul(v)
}

func (v F64Vec4) NormalizeFastest() F64Vec4 {
    ooLen := F64FromRaw(fix64.RSqrtFastest(fix64.Mul(v.RawX, v.RawX) + fix64.Mul(v.RawY, v.RawY) + fix64.Mul(v.RawZ, v.RawZ) + fix64.Mul(v.RawW, v.RawW)))
    return Vec4FromF64(ooLen, ooLen, ooLen, ooLen).Mul(v)
}

func (v F64Vec4) Dot(b F64Vec4) F64 {
    return F64FromRaw(fix64.Mul(v.RawX, b.RawX) + fix64.Mul(v.RawY, b.RawY) + fix64.Mul(v.RawZ, b.RawZ) + fix64.Mul(v.RawW, b.RawW))
}

func (v F64Vec4) Distance(b F64Vec4) F64 {
    return v.Sub(b).Length()
}

func (v F64Vec4) DistanceFast(b F64Vec4) F64 {
    return v.Sub(b).LengthFast()
}

func (v F64Vec4) DistanceFastest(b F64Vec4) F64 {
    return v.Sub(b).LengthFastest()
}

func (v F64Vec4) Clamp(min, max F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Clamp(v.RawX, min.RawX, max.RawX), fix64.Clamp(v.RawY, min.RawY, max.RawY), fix64.Clamp(v.RawZ, min.RawZ, max.RawZ), fix64.Clamp(v.RawW, min.RawW, max.RawW))
}

func (v F64Vec4) Lerp(b F64Vec4, t F64) F64Vec4 {
    tb := t.Raw
    ta := fix64.One - tb
    return Vec4FromRaw(fix64.Mul(v.RawX, ta)+fix64.Mul(b.RawX, tb), fix64.Mul(v.RawY, ta)+fix64.Mul(b.RawY, tb), fix64.Mul(v.RawZ, ta)+fix64.Mul(b.RawZ, tb), fix64.Mul(v.RawW, ta)+fix64.Mul(b.RawW, tb))
}

func (v F64Vec4) Equals(obj F64Vec4) bool {
    return reflect.DeepEqual(v, obj)
}

func (v F64Vec4) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s, %s)`, fix64.ToString(v.RawX), fix64.ToString(v.RawY), fix64.ToString(v.RawZ), fix64.ToString(v.RawW))
}
