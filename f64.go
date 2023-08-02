package fp

import (
    "reflect"

    "github.com/camry/fp/fix64"
)

var (
    Neg1        = F64FromRaw(fix64.Neg1)
    F64Zero     = F64FromRaw(fix64.Zero)
    F64Half     = F64FromRaw(fix64.Half)
    F64One      = F64FromRaw(fix64.One)
    F64Two      = F64FromRaw(fix64.Two)
    F64Pi       = F64FromRaw(fix64.Pi)
    F64Pi2      = F64FromRaw(fix64.Pi2)
    F64PiHalf   = F64FromRaw(fix64.PiHalf)
    F64E        = F64FromRaw(fix64.E)
    F64MinValue = F64FromRaw(fix64.MinValue)
    F64MaxValue = F64FromRaw(fix64.MaxValue)
)

// F64 Signed 32.32 fixed point value struct.
type F64 struct {
    Raw int64 // Raw fixed point value
}

/************************************/
/*********** Construction ***********/
/************************************/

func F64FromRaw(raw int64) F64 {
    var f F64
    f.Raw = raw
    return f
}

func F64FromInt32(v int32) F64 {
    return F64FromRaw(fix64.FromInt32(v))
}

func F64FromInt64(v int64) F64 {
    return F64FromRaw(fix64.FromInt64(v))
}

func F64FromFloat32(v float32) F64 {
    return F64FromRaw(fix64.FromFloat32(v))
}

func F64FromFloat64(v float64) F64 {
    return F64FromRaw(fix64.FromFloat64(v))
}

func F64FromF32(v F32) F64 {
    return F64FromRaw(int64(v.Raw) << 16)
}

// F64Ratio Creates the fixed point number that's a divided by b.
func F64Ratio(a, b int32) F64 {
    return F64FromRaw((int64(a) << 32) / int64(b))
}

// F64Ratio10 Creates the fixed point number that's a divided by 10.
func F64Ratio10(a int32) F64 {
    return F64FromRaw((int64(a) << 32) / 10)
}

// F64Ratio100 Creates the fixed point number that's a divided by 100.
func F64Ratio100(a int32) F64 {
    return F64FromRaw((int64(a) << 32) / 100)
}

// F64Ratio1000 Creates the fixed point number that's a divided by 1000.
func F64Ratio1000(a int32) F64 {
    return F64FromRaw((int64(a) << 32) / 1000)
}

// F64Min returns the smallest F64 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F64Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call F64Min with 0 arguments.
func F64Min(first F64, rest ...F64) F64 {
    ans := first
    for _, item := range rest {
        if item.Raw < ans.Raw {
            ans = item
        }
    }
    return ans
}

// F64Max returns the largest F64 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F64Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call F64Max with 0 arguments.
func F64Max(first F64, rest ...F64) F64 {
    ans := first
    for _, item := range rest {
        if item.Raw > ans.Raw {
            ans = item
        }
    }
    return ans
}

// F64Sum returns the combined total of the provided first and rest Decimals
func F64Sum(first F64, rest ...F64) F64 {
    total := first
    for _, item := range rest {
        total = total.Add(item)
    }

    return total
}

// F64Avg returns the average value of the provided first and rest Decimals
func F64Avg(first F64, rest ...F64) F64 {
    count := F64FromInt64(int64(len(rest) + 1))
    sum := F64Sum(first, rest...)
    return sum.Div(count)
}

/************************************/
/*********** Conversions ************/
/************************************/

func (f F64) FloorToInt() int32 {
    return fix64.FloorToInt(f.Raw)
}

func (f F64) CeilToInt() int32 {
    return fix64.CeilToInt(f.Raw)
}

func (f F64) RoundToInt() int32 {
    return fix64.RoundToInt(f.Raw)
}

func (f F64) Float32() float32 {
    return fix64.ToFloat32(f.Raw)
}

func (f F64) Float64() float64 {
    return fix64.ToFloat64(f.Raw)
}

func (f F64) F32() F32 {
    // TODO not implements
    panic("not implements")
}

/************************************/
/************ Operators *************/
/************************************/

// Negate -f
func (f F64) Negate() F64 {
    return F64FromRaw(-f.Raw)
}

// Add f + v2
func (f F64) Add(v2 F64) F64 {
    return F64FromRaw(fix64.Add(f.Raw, v2.Raw))
}

// Sub f - v2
func (f F64) Sub(v2 F64) F64 {
    return F64FromRaw(fix64.Sub(f.Raw, v2.Raw))
}

// Mul f * v2
func (f F64) Mul(v2 F64) F64 {
    return F64FromRaw(fix64.Mul(f.Raw, v2.Raw))
}

// DivPrecise f / v2
func (f F64) DivPrecise(v2 F64) F64 {
    return F64FromRaw(fix64.DivPrecise(f.Raw, v2.Raw))
}

// Mod f % v2
func (f F64) Mod(v2 F64) F64 {
    return F64FromRaw(fix64.Mod(f.Raw, v2.Raw))
}

// AddVec2 f + v2
func (f F64) AddVec2(v2 F64Vec2) F64Vec2 {
    return Vec2FromRaw(f.Raw+v2.RawX, f.Raw+v2.RawY)
}

// SubVec2 f - v2
func (f F64) SubVec2(v2 F64Vec2) F64Vec2 {
    return Vec2FromRaw(f.Raw-v2.RawX, f.Raw-v2.RawY)
}

// MulVec2 f * v2
func (f F64) MulVec2(v2 F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.Mul(f.Raw, v2.RawX), fix64.Mul(f.Raw, v2.RawY))
}

// DivPreciseVec2 f / v2
func (f F64) DivPreciseVec2(v2 F64Vec2) F64Vec2 {
    return Vec2FromRaw(fix64.DivPrecise(f.Raw, v2.RawX), fix64.DivPrecise(f.Raw, v2.RawY))
}

// ModVec2 f % v2
func (f F64) ModVec2(v2 F64Vec2) F64Vec2 {
    return Vec2FromRaw(f.Raw%v2.RawX, f.Raw%v2.RawY)
}

// AddVec3 f + v2
func (f F64) AddVec3(v2 F64Vec3) F64Vec3 {
    return Vec3FromRaw(f.Raw+v2.RawX, f.Raw+v2.RawY, f.Raw+v2.RawZ)
}

// SubVec3 f - v2
func (f F64) SubVec3(v2 F64Vec3) F64Vec3 {
    return Vec3FromRaw(f.Raw-v2.RawX, f.Raw-v2.RawY, f.Raw-v2.RawZ)
}

// MulVec3 f * v2
func (f F64) MulVec3(v2 F64Vec3) F64Vec3 {
    return Vec3FromRaw(fix64.Mul(f.Raw, v2.RawX), fix64.Mul(f.Raw, v2.RawY), fix64.Mul(f.Raw, v2.RawZ))
}

// DivPreciseVec3 f / v2
func (f F64) DivPreciseVec3(v2 F64Vec3) F64Vec3 {
    return Vec3FromRaw(fix64.DivPrecise(f.Raw, v2.RawX), fix64.DivPrecise(f.Raw, v2.RawY), fix64.DivPrecise(f.Raw, v2.RawZ))
}

// ModVec3 f % v2
func (f F64) ModVec3(v2 F64Vec3) F64Vec3 {
    return Vec3FromRaw(f.Raw%v2.RawX, f.Raw%v2.RawY, f.Raw%v2.RawZ)
}

// AddVec4 f + v2
func (f F64) AddVec4(v2 F64Vec4) F64Vec4 {
    return Vec4FromRaw(f.Raw+v2.RawX, f.Raw+v2.RawY, f.Raw+v2.RawZ, f.Raw+v2.RawW)
}

// SubVec4 f - v2
func (f F64) SubVec4(v2 F64Vec4) F64Vec4 {
    return Vec4FromRaw(f.Raw-v2.RawX, f.Raw-v2.RawY, f.Raw-v2.RawZ, f.Raw-v2.RawW)
}

// MulVec4 f * v2
func (f F64) MulVec4(v2 F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.Mul(f.Raw, v2.RawX), fix64.Mul(f.Raw, v2.RawY), fix64.Mul(f.Raw, v2.RawZ), fix64.Mul(f.Raw, v2.RawW))
}

// DivPreciseVec4 f / v2
func (f F64) DivPreciseVec4(v2 F64Vec4) F64Vec4 {
    return Vec4FromRaw(fix64.DivPrecise(f.Raw, v2.RawX), fix64.DivPrecise(f.Raw, v2.RawY), fix64.DivPrecise(f.Raw, v2.RawZ), fix64.DivPrecise(f.Raw, v2.RawW))
}

// ModVec4 f % v2
func (f F64) ModVec4(v2 F64Vec4) F64Vec4 {
    return Vec4FromRaw(f.Raw%v2.RawX, f.Raw%v2.RawY, f.Raw%v2.RawZ, f.Raw%v2.RawW)
}

// Add2 f++
func (f F64) Add2() F64 {
    return F64FromRaw(f.Raw + fix64.One)
}

// Sub2 f--
func (f F64) Sub2() F64 {
    return F64FromRaw(f.Raw - fix64.One)
}

// EQ f == v2
func (f F64) EQ(v2 F64) bool {
    return f.Raw == v2.Raw
}

// NE f != v2
func (f F64) NE(v2 F64) bool {
    return f.Raw != v2.Raw
}

// LT f < v2
func (f F64) LT(v2 F64) bool {
    return f.Raw < v2.Raw
}

// LE f <= v2
func (f F64) LE(v2 F64) bool {
    return f.Raw <= v2.Raw
}

// GT f > v2
func (f F64) GT(v2 F64) bool {
    return f.Raw > v2.Raw
}

// GE f >= v2
func (f F64) GE(v2 F64) bool {
    return f.Raw >= v2.Raw
}

// RadToDeg 180 / F64.Pi
func (f F64) RadToDeg() F64 {
    return F64FromRaw(fix64.Mul(f.Raw, 246083499198))
}

// DegToRad F64.Pi / 180
func (f F64) DegToRad() F64 {
    return F64FromRaw(fix64.Mul(f.Raw, 74961320))
}

func (f F64) Div2() F64 {
    return F64FromRaw(f.Raw >> 1)
}

func (f F64) Abs() F64 {
    return F64FromRaw(fix64.Abs(f.Raw))
}

func (f F64) Nabs() F64 {
    return F64FromRaw(fix64.Nabs(f.Raw))
}

func (f F64) Sign() int32 {
    return fix64.Sign(f.Raw)
}

func (f F64) Ceil() F64 {
    return F64FromRaw(fix64.Ceil(f.Raw))
}

func (f F64) Floor() F64 {
    return F64FromRaw(fix64.Floor(f.Raw))
}

func (f F64) Round() F64 {
    return F64FromRaw(fix64.Round(f.Raw))
}

func (f F64) Fract() F64 {
    return F64FromRaw(fix64.Fract(f.Raw))
}

func (f F64) Div(b F64) F64 {
    return F64FromRaw(fix64.Div(f.Raw, b.Raw))
}

func (f F64) DivFast(b F64) F64 {
    return F64FromRaw(fix64.DivFast(f.Raw, b.Raw))
}

func (f F64) DivFastest(b F64) F64 {
    return F64FromRaw(fix64.DivFastest(f.Raw, b.Raw))
}

func (f F64) SqrtPrecise() F64 {
    return F64FromRaw(fix64.SqrtPrecise(f.Raw))
}

func (f F64) Sqrt() F64 {
    return F64FromRaw(fix64.Sqrt(f.Raw))
}

func (f F64) SqrtFast() F64 {
    return F64FromRaw(fix64.SqrtFast(f.Raw))
}

func (f F64) SqrtFastest() F64 {
    return F64FromRaw(fix64.SqrtFastest(f.Raw))
}

func (f F64) RSqrt() F64 {
    return F64FromRaw(fix64.RSqrt(f.Raw))
}

func (f F64) RSqrtFast() F64 {
    return F64FromRaw(fix64.RSqrtFast(f.Raw))
}

func (f F64) RSqrtFastest() F64 {
    return F64FromRaw(fix64.RSqrtFastest(f.Raw))
}

func (f F64) Rcp() F64 {
    return F64FromRaw(fix64.Rcp(f.Raw))
}

func (f F64) RcpFast() F64 {
    return F64FromRaw(fix64.RcpFast(f.Raw))
}

func (f F64) RcpFastest() F64 {
    return F64FromRaw(fix64.RcpFastest(f.Raw))
}

func (f F64) Exp() F64 {
    return F64FromRaw(fix64.Exp(f.Raw))
}

func (f F64) ExpFast() F64 {
    return F64FromRaw(fix64.ExpFast(f.Raw))
}

func (f F64) ExpFastest() F64 {
    return F64FromRaw(fix64.ExpFastest(f.Raw))
}

func (f F64) Exp2() F64 {
    return F64FromRaw(fix64.Exp2(f.Raw))
}

func (f F64) Exp2Fast() F64 {
    return F64FromRaw(fix64.Exp2Fast(f.Raw))
}

func (f F64) Exp2Fastest() F64 {
    return F64FromRaw(fix64.Exp2Fastest(f.Raw))
}

func (f F64) Log() F64 {
    return F64FromRaw(fix64.Log(f.Raw))
}

func (f F64) LogFast() F64 {
    return F64FromRaw(fix64.LogFast(f.Raw))
}

func (f F64) LogFastest() F64 {
    return F64FromRaw(fix64.LogFastest(f.Raw))
}

func (f F64) Log2() F64 {
    return F64FromRaw(fix64.Log2(f.Raw))
}

func (f F64) Log2Fast() F64 {
    return F64FromRaw(fix64.Log2Fast(f.Raw))
}

func (f F64) Log2Fastest() F64 {
    return F64FromRaw(fix64.Log2Fastest(f.Raw))
}

func (f F64) Sin() F64 {
    return F64FromRaw(fix64.Sin(f.Raw))
}

func (f F64) SinFast() F64 {
    return F64FromRaw(fix64.SinFast(f.Raw))
}

func (f F64) SinFastest() F64 {
    return F64FromRaw(fix64.SinFastest(f.Raw))
}

func (f F64) Cos() F64 {
    return F64FromRaw(fix64.Cos(f.Raw))
}

func (f F64) CosFast() F64 {
    return F64FromRaw(fix64.CosFast(f.Raw))
}

func (f F64) CosFastest() F64 {
    return F64FromRaw(fix64.CosFastest(f.Raw))
}

func (f F64) Tan() F64 {
    return F64FromRaw(fix64.Tan(f.Raw))
}

func (f F64) TanFast() F64 {
    return F64FromRaw(fix64.TanFast(f.Raw))
}

func (f F64) TanFastest() F64 {
    return F64FromRaw(fix64.TanFastest(f.Raw))
}

func (f F64) Asin() F64 {
    return F64FromRaw(fix64.Asin(f.Raw))
}

func (f F64) AsinFast() F64 {
    return F64FromRaw(fix64.AsinFast(f.Raw))
}

func (f F64) AsinFastest() F64 {
    return F64FromRaw(fix64.AsinFastest(f.Raw))
}

func (f F64) Acos() F64 {
    return F64FromRaw(fix64.Acos(f.Raw))
}

func (f F64) AcosFast() F64 {
    return F64FromRaw(fix64.AcosFast(f.Raw))
}

func (f F64) AcosFastest() F64 {
    return F64FromRaw(fix64.AcosFastest(f.Raw))
}

func (f F64) Atan() F64 {
    return F64FromRaw(fix64.Atan(f.Raw))
}

func (f F64) AtanFast() F64 {
    return F64FromRaw(fix64.AtanFast(f.Raw))
}

func (f F64) AtanFastest() F64 {
    return F64FromRaw(fix64.AtanFastest(f.Raw))
}

func (f F64) Atan2(x F64) F64 {
    return F64FromRaw(fix64.Atan2(f.Raw, x.Raw))
}

func (f F64) Atan2Fast(x F64) F64 {
    return F64FromRaw(fix64.Atan2Fast(f.Raw, x.Raw))
}

func (f F64) Atan2Fastest(x F64) F64 {
    return F64FromRaw(fix64.Atan2Fastest(f.Raw, x.Raw))
}

func (f F64) Pow(b F64) F64 {
    return F64FromRaw(fix64.Pow(f.Raw, b.Raw))
}

func (f F64) PowFast(b F64) F64 {
    return F64FromRaw(fix64.PowFast(f.Raw, b.Raw))
}

func (f F64) PowFastest(b F64) F64 {
    return F64FromRaw(fix64.PowFastest(f.Raw, b.Raw))
}

func (f F64) Clamp(min, max F64) F64 {
    return F64FromRaw(fix64.Clamp(f.Raw, min.Raw, max.Raw))
}

func (f F64) Clamp01() F64 {
    return F64FromRaw(fix64.Clamp(f.Raw, fix64.Zero, fix64.One))
}

func (f F64) Lerp(b, t F64) F64 {
    tb := t.Raw
    ta := fix64.One - tb
    return F64FromRaw(fix64.Mul(f.Raw, ta) + fix64.Mul(b.Raw, tb))
}

func (f F64) Equals(obj F64) bool {
    return reflect.DeepEqual(f, obj)
}

func (f F64) CompareTo(other F64) int32 {
    if f.Raw < other.Raw {
        return -1
    }
    if f.Raw > other.Raw {
        return +1
    }
    return 0
}

func (f F64) ToString() string {
    return fix64.ToString(f.Raw)
}
