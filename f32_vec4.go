package fp

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix32"
)

var (
    F32Vec4Zero  = F32Vec4FromRaw(fix32.Zero, fix32.Zero, fix32.Zero, fix32.Zero)
    F32Vec4One   = F32Vec4FromRaw(fix32.One, fix32.One, fix32.One, fix32.One)
    F32Vec4AxisX = F32Vec4FromRaw(fix32.One, fix32.Zero, fix32.Zero, fix32.Zero)
    F32Vec4AxisY = F32Vec4FromRaw(fix32.Zero, fix32.One, fix32.Zero, fix32.Zero)
    F32Vec4AxisZ = F32Vec4FromRaw(fix32.Zero, fix32.Zero, fix32.One, fix32.Zero)
    F32Vec4AxisW = F32Vec4FromRaw(fix32.Zero, fix32.Zero, fix32.Zero, fix32.One)
)

// F32Vec4 struct with signed 16.16 fixed point components.
type F32Vec4 struct {
    RawX int32
    RawY int32
    RawZ int32
    RawW int32
}

func F32Vec4FromRaw(rawX, rawY, RawZ, RawW int32) F32Vec4 {
    return F32Vec4{
        RawX: rawX,
        RawY: rawY,
        RawZ: RawZ,
        RawW: RawW,
    }
}

func F32Vec4FromF32(x, y, z, w F32) F32Vec4 {
    return F32Vec4FromRaw(x.Raw, y.Raw, z.Raw, w.Raw)
}

func F32Vec4FromInt32(x, y, z, w int32) F32Vec4 {
    return F32Vec4FromRaw(fix32.FromInt32(x), fix32.FromInt32(y), fix32.FromInt32(z), fix32.FromInt32(w))
}

func F32Vec4FromFloat32(x, y, z, w float32) F32Vec4 {
    return F32Vec4FromRaw(fix32.FromFloat32(x), fix32.FromFloat32(y), fix32.FromFloat32(z), fix32.FromFloat32(w))
}

func F32Vec4FromFloat64(x, y, z, w float64) F32Vec4 {
    return F32Vec4FromRaw(fix32.FromFloat64(x), fix32.FromFloat64(y), fix32.FromFloat64(z), fix32.FromFloat64(w))
}

// F32Vec4Min returns the smallest F32Vec4 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Vec4Min(v0, v1)
//
// This makes it harder to accidentally call F32Vec4Min with 0 arguments.
func F32Vec4Min(v0 F32Vec4, v1 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Min(v0.RawX, v1.RawX), fix32.Min(v0.RawY, v1.RawY), fix32.Min(v0.RawZ, v1.RawZ), fix32.Min(v0.RawW, v1.RawW))
}

// F32Vec4Max returns the largest F32Vec4 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Vec4Max(v0, v1)
//
// This makes it harder to accidentally call F32Vec4Max with 0 arguments.
func F32Vec4Max(v0 F32Vec4, v1 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Max(v0.RawX, v1.RawX), fix32.Max(v0.RawY, v1.RawY), fix32.Max(v0.RawZ, v1.RawZ), fix32.Max(v0.RawW, v1.RawW))
}

func (v F32Vec4) X() F32 {
    return F32FromRaw(v.RawX)
}

func (v F32Vec4) Y() F32 {
    return F32FromRaw(v.RawY)
}

func (v F32Vec4) Z() F32 {
    return F32FromRaw(v.RawZ)
}

func (v F32Vec4) W() F32 {
    return F32FromRaw(v.RawW)
}

// Negate -v
func (v F32Vec4) Negate() F32Vec4 {
    return F32Vec4FromRaw(-v.RawX, -v.RawY, -v.RawZ, -v.RawW)
}

// Add v + b
func (v F32Vec4) Add(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(v.RawX+b.RawX, v.RawY+b.RawY, v.RawZ+b.RawZ, v.RawW+b.RawW)
}

// Sub v - b
func (v F32Vec4) Sub(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(v.RawX-b.RawX, v.RawY-b.RawY, v.RawZ-b.RawZ, v.RawW-b.RawW)
}

// Mul v * b
func (v F32Vec4) Mul(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Mul(v.RawX, b.RawX), fix32.Mul(v.RawY, b.RawY), fix32.Mul(v.RawZ, b.RawZ), fix32.Mul(v.RawW, b.RawW))
}

// DivPrecise v / b
func (v F32Vec4) DivPrecise(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.DivPrecise(v.RawX, b.RawX), fix32.DivPrecise(v.RawY, b.RawY), fix32.DivPrecise(v.RawZ, b.RawZ), fix32.DivPrecise(v.RawW, b.RawW))
}

// Mod v % b
func (v F32Vec4) Mod(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(v.RawX%b.RawX, v.RawY%b.RawY, v.RawZ%b.RawZ, v.RawW%b.RawW)
}

// AddF32 v + b
func (v F32Vec4) AddF32(b F32) F32Vec4 {
    return F32Vec4FromRaw(v.RawX+b.Raw, v.RawY+b.Raw, v.RawZ+b.Raw, v.RawW+b.Raw)
}

// SubF32 v - b
func (v F32Vec4) SubF32(b F32) F32Vec4 {
    return F32Vec4FromRaw(v.RawX-b.Raw, v.RawY-b.Raw, v.RawZ-b.Raw, v.RawW-b.Raw)
}

// MulF32 v * b
func (v F32Vec4) MulF32(b F32) F32Vec4 {
    return F32Vec4FromRaw(fix32.Mul(v.RawX, b.Raw), fix32.Mul(v.RawY, b.Raw), fix32.Mul(v.RawZ, b.Raw), fix32.Mul(v.RawW, b.Raw))
}

// DivPreciseF32 v / b
func (v F32Vec4) DivPreciseF32(b F32) F32Vec4 {
    return F32Vec4FromRaw(fix32.DivPrecise(v.RawX, b.Raw), fix32.DivPrecise(v.RawY, b.Raw), fix32.DivPrecise(v.RawZ, b.Raw), fix32.DivPrecise(v.RawW, b.Raw))
}

// ModF32 v % b
func (v F32Vec4) ModF32(b F32) F32Vec4 {
    return F32Vec4FromRaw(v.RawX%b.Raw, v.RawY%b.Raw, v.RawZ%b.Raw, v.RawW%b.Raw)
}

// EQ v == b
func (v F32Vec4) EQ(b F32Vec4) bool {
    return v.RawX == b.RawX && v.RawY == b.RawY && v.RawZ == b.RawZ && v.RawW == b.RawW
}

// NE v != b
func (v F32Vec4) NE(b F32Vec4) bool {
    return v.RawX != b.RawX || v.RawY != b.RawY || v.RawZ != b.RawZ || v.RawW != b.RawW
}

func (v F32Vec4) Div(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Div(v.RawX, b.RawX), fix32.Div(v.RawY, b.RawY), fix32.Div(v.RawZ, b.RawZ), fix32.Div(v.RawW, b.RawW))
}

func (v F32Vec4) DivFast(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.DivFast(v.RawX, b.RawX), fix32.DivFast(v.RawY, b.RawY), fix32.DivFast(v.RawZ, b.RawZ), fix32.DivFast(v.RawW, b.RawW))
}

func (v F32Vec4) DivFastest(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.DivFastest(v.RawX, b.RawX), fix32.DivFastest(v.RawY, b.RawY), fix32.DivFastest(v.RawZ, b.RawZ), fix32.DivFastest(v.RawW, b.RawW))
}

func (v F32Vec4) SqrtPrecise() F32Vec4 {
    return F32Vec4FromRaw(fix32.SqrtPrecise(v.RawX), fix32.SqrtPrecise(v.RawY), fix32.SqrtPrecise(v.RawZ), fix32.SqrtPrecise(v.RawW))
}

func (v F32Vec4) Sqrt() F32Vec4 {
    return F32Vec4FromRaw(fix32.Sqrt(v.RawX), fix32.Sqrt(v.RawY), fix32.Sqrt(v.RawZ), fix32.Sqrt(v.RawW))
}

func (v F32Vec4) SqrtFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.SqrtFast(v.RawX), fix32.SqrtFast(v.RawY), fix32.SqrtFast(v.RawZ), fix32.SqrtFast(v.RawW))
}

func (v F32Vec4) SqrtFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.SqrtFastest(v.RawX), fix32.SqrtFastest(v.RawY), fix32.SqrtFastest(v.RawZ), fix32.SqrtFastest(v.RawW))
}

func (v F32Vec4) RSqrt() F32Vec4 {
    return F32Vec4FromRaw(fix32.RSqrt(v.RawX), fix32.RSqrt(v.RawY), fix32.RSqrt(v.RawZ), fix32.RSqrt(v.RawW))
}

func (v F32Vec4) RSqrtFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.RSqrtFast(v.RawX), fix32.RSqrtFast(v.RawY), fix32.RSqrtFast(v.RawZ), fix32.RSqrtFast(v.RawW))
}

func (v F32Vec4) RSqrtFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.RSqrtFastest(v.RawX), fix32.RSqrtFastest(v.RawY), fix32.RSqrtFastest(v.RawZ), fix32.RSqrtFastest(v.RawW))
}

func (v F32Vec4) Rcp() F32Vec4 {
    return F32Vec4FromRaw(fix32.Rcp(v.RawX), fix32.Rcp(v.RawY), fix32.Rcp(v.RawZ), fix32.Rcp(v.RawW))
}

func (v F32Vec4) RcpFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.RcpFast(v.RawX), fix32.RcpFast(v.RawY), fix32.RcpFast(v.RawZ), fix32.RcpFast(v.RawW))
}

func (v F32Vec4) RcpFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.RcpFastest(v.RawX), fix32.RcpFastest(v.RawY), fix32.RcpFastest(v.RawZ), fix32.RcpFastest(v.RawW))
}

func (v F32Vec4) Exp() F32Vec4 {
    return F32Vec4FromRaw(fix32.Exp(v.RawX), fix32.Exp(v.RawY), fix32.Exp(v.RawZ), fix32.Exp(v.RawW))
}

func (v F32Vec4) ExpFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.ExpFast(v.RawX), fix32.ExpFast(v.RawY), fix32.ExpFast(v.RawZ), fix32.ExpFast(v.RawW))
}

func (v F32Vec4) ExpFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.ExpFastest(v.RawX), fix32.ExpFastest(v.RawY), fix32.ExpFastest(v.RawZ), fix32.ExpFastest(v.RawW))
}

func (v F32Vec4) Exp2() F32Vec4 {
    return F32Vec4FromRaw(fix32.Exp2(v.RawX), fix32.Exp2(v.RawY), fix32.Exp2(v.RawZ), fix32.Exp2(v.RawW))
}

func (v F32Vec4) Exp2Fast() F32Vec4 {
    return F32Vec4FromRaw(fix32.Exp2Fast(v.RawX), fix32.Exp2Fast(v.RawY), fix32.Exp2Fast(v.RawZ), fix32.Exp2Fast(v.RawW))
}

func (v F32Vec4) Exp2Fastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.Exp2Fastest(v.RawX), fix32.Exp2Fastest(v.RawY), fix32.Exp2Fastest(v.RawZ), fix32.Exp2Fastest(v.RawW))
}

func (v F32Vec4) Log() F32Vec4 {
    return F32Vec4FromRaw(fix32.Log(v.RawX), fix32.Log(v.RawY), fix32.Log(v.RawZ), fix32.Log(v.RawW))
}

func (v F32Vec4) LogFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.LogFast(v.RawX), fix32.LogFast(v.RawY), fix32.LogFast(v.RawZ), fix32.LogFast(v.RawW))
}

func (v F32Vec4) LogFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.LogFastest(v.RawX), fix32.LogFastest(v.RawY), fix32.LogFastest(v.RawZ), fix32.LogFastest(v.RawW))
}

func (v F32Vec4) Log2() F32Vec4 {
    return F32Vec4FromRaw(fix32.Log2(v.RawX), fix32.Log2(v.RawY), fix32.Log2(v.RawZ), fix32.Log2(v.RawW))
}

func (v F32Vec4) Log2Fast() F32Vec4 {
    return F32Vec4FromRaw(fix32.Log2Fast(v.RawX), fix32.Log2Fast(v.RawY), fix32.Log2Fast(v.RawZ), fix32.Log2Fast(v.RawW))
}

func (v F32Vec4) Log2Fastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.Log2Fastest(v.RawX), fix32.Log2Fastest(v.RawY), fix32.Log2Fastest(v.RawZ), fix32.Log2Fastest(v.RawW))
}

func (v F32Vec4) Sin() F32Vec4 {
    return F32Vec4FromRaw(fix32.Sin(v.RawX), fix32.Sin(v.RawY), fix32.Sin(v.RawZ), fix32.Sin(v.RawW))
}

func (v F32Vec4) SinFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.SinFast(v.RawX), fix32.SinFast(v.RawY), fix32.SinFast(v.RawZ), fix32.SinFast(v.RawW))
}

func (v F32Vec4) SinFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.SinFastest(v.RawX), fix32.SinFastest(v.RawY), fix32.SinFastest(v.RawZ), fix32.SinFastest(v.RawW))
}

func (v F32Vec4) Cos() F32Vec4 {
    return F32Vec4FromRaw(fix32.Cos(v.RawX), fix32.Cos(v.RawY), fix32.Cos(v.RawZ), fix32.Cos(v.RawW))
}

func (v F32Vec4) CosFast() F32Vec4 {
    return F32Vec4FromRaw(fix32.CosFast(v.RawX), fix32.CosFast(v.RawY), fix32.CosFast(v.RawZ), fix32.CosFast(v.RawW))
}

func (v F32Vec4) CosFastest() F32Vec4 {
    return F32Vec4FromRaw(fix32.CosFastest(v.RawX), fix32.CosFastest(v.RawY), fix32.CosFastest(v.RawZ), fix32.CosFastest(v.RawW))
}

func (v F32Vec4) Pow(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Pow(v.RawX, b.RawX), fix32.Pow(v.RawY, b.RawY), fix32.Pow(v.RawZ, b.RawZ), fix32.Pow(v.RawW, b.RawW))
}

func (v F32Vec4) PowFast(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.PowFast(v.RawX, b.RawX), fix32.PowFast(v.RawY, b.RawY), fix32.PowFast(v.RawZ, b.RawZ), fix32.PowFast(v.RawW, b.RawW))
}

func (v F32Vec4) PowFastest(b F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.PowFastest(v.RawX, b.RawX), fix32.PowFastest(v.RawY, b.RawY), fix32.PowFastest(v.RawZ, b.RawZ), fix32.PowFastest(v.RawW, b.RawW))
}

func (v F32Vec4) Length() F32 {
    return F32FromRaw(fix32.Sqrt(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW)))
}

func (v F32Vec4) LengthFast() F32 {
    return F32FromRaw(fix32.SqrtFast(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW)))
}

func (v F32Vec4) LengthFastest() F32 {
    return F32FromRaw(fix32.SqrtFastest(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW)))
}

func (v F32Vec4) LengthSqr() F32 {
    return F32FromRaw(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW))
}

func (v F32Vec4) Normalize() F32Vec4 {
    ooLen := F32FromRaw(fix32.RSqrt(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW)))
    return F32Vec4FromF32(ooLen, ooLen, ooLen, ooLen).Mul(v)
}

func (v F32Vec4) NormalizeFast() F32Vec4 {
    ooLen := F32FromRaw(fix32.RSqrtFast(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW)))
    return F32Vec4FromF32(ooLen, ooLen, ooLen, ooLen).Mul(v)
}

func (v F32Vec4) NormalizeFastest() F32Vec4 {
    ooLen := F32FromRaw(fix32.RSqrtFastest(fix32.Mul(v.RawX, v.RawX) + fix32.Mul(v.RawY, v.RawY) + fix32.Mul(v.RawZ, v.RawZ) + fix32.Mul(v.RawW, v.RawW)))
    return F32Vec4FromF32(ooLen, ooLen, ooLen, ooLen).Mul(v)
}

func (v F32Vec4) Dot(b F32Vec4) F32 {
    return F32FromRaw(fix32.Mul(v.RawX, b.RawX) + fix32.Mul(v.RawY, b.RawY) + fix32.Mul(v.RawZ, b.RawZ) + fix32.Mul(v.RawW, b.RawW))
}

func (v F32Vec4) Distance(b F32Vec4) F32 {
    return v.Sub(b).Length()
}

func (v F32Vec4) DistanceFast(b F32Vec4) F32 {
    return v.Sub(b).LengthFast()
}

func (v F32Vec4) DistanceFastest(b F32Vec4) F32 {
    return v.Sub(b).LengthFastest()
}

func (v F32Vec4) Clamp(min, max F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Clamp(v.RawX, min.RawX, max.RawX), fix32.Clamp(v.RawY, min.RawY, max.RawY), fix32.Clamp(v.RawZ, min.RawZ, max.RawZ), fix32.Clamp(v.RawW, min.RawW, max.RawW))
}

func (v F32Vec4) Lerp(b F32Vec4, t F32) F32Vec4 {
    tb := t.Raw
    ta := fix32.One - tb
    return F32Vec4FromRaw(fix32.Mul(v.RawX, ta)+fix32.Mul(b.RawX, tb), fix32.Mul(v.RawY, ta)+fix32.Mul(b.RawY, tb), fix32.Mul(v.RawZ, ta)+fix32.Mul(b.RawZ, tb), fix32.Mul(v.RawW, ta)+fix32.Mul(b.RawW, tb))
}

func (v F32Vec4) Equals(obj F32Vec4) bool {
    return reflect.DeepEqual(v, obj)
}

func (v F32Vec4) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s, %s)`, fix32.ToString(v.RawX), fix32.ToString(v.RawY), fix32.ToString(v.RawZ), fix32.ToString(v.RawW))
}
