package fp

import (
    "reflect"

    "github.com/camry/fp/fix32"
)

var (
    F32Neg1     = F32FromRaw(fix32.Neg1)
    F32Zero     = F32FromRaw(fix32.Zero)
    F32Half     = F32FromRaw(fix32.Half)
    F32One      = F32FromRaw(fix32.One)
    F32Two      = F32FromRaw(fix32.Two)
    F32Pi       = F32FromRaw(fix32.Pi)
    F32Pi2      = F32FromRaw(fix32.Pi2)
    F32PiHalf   = F32FromRaw(fix32.PiHalf)
    F32E        = F32FromRaw(fix32.E)
    F32MinValue = F32FromRaw(fix32.MinValue)
    F32MaxValue = F32FromRaw(fix32.MaxValue)
)

// F32 Signed 16.16 fixed point value struct.
type F32 struct {
    Raw int32
}

/************************************/
/*********** Construction ***********/
/************************************/

func F32FromRaw(raw int32) F32 {
    var f F32
    f.Raw = raw
    return f
}

func F32FromInt32(v int32) F32 {
    return F32FromRaw(fix32.FromInt32(v))
}

func F32FromFloat32(v float32) F32 {
    return F32FromRaw(fix32.FromFloat32(v))
}

func F32FromFloat64(v float64) F32 {
    return F32FromRaw(fix32.FromFloat64(v))
}

func F32FromF64(v F64) F32 {
    return F32FromRaw(int32(v.Raw >> 16))
}

// F32Ratio Creates the fixed point number that's a divided by b.
func F32Ratio(a, b int32) F32 {
    return F32FromRaw(int32((int64(a) << 16) / int64(b)))
}

// F32Ratio10 Creates the fixed point number that's a divided by 10.
func F32Ratio10(a int32) F32 {
    return F32FromRaw(int32((int64(a) << 16) / 10))
}

// F32Ratio100 Creates the fixed point number that's a divided by 100.
func F32Ratio100(a int32) F32 {
    return F32FromRaw(int32((int64(a) << 16) / 100))
}

// F32Ratio1000 Creates the fixed point number that's a divided by 1000.
func F32Ratio1000(a int32) F32 {
    return F32FromRaw(int32((int64(a) << 16) / 1000))
}

// F32Min returns the smallest F64 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call F64Min with 0 arguments.
func F32Min(first F32, rest ...F32) F32 {
    ans := first
    for _, item := range rest {
        if item.Raw < ans.Raw {
            ans = item
        }
    }
    return ans
}

// F32Max returns the largest F64 that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//     F32Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call F64Max with 0 arguments.
func F32Max(first F32, rest ...F32) F32 {
    ans := first
    for _, item := range rest {
        if item.Raw > ans.Raw {
            ans = item
        }
    }
    return ans
}

// F32Sum returns the combined total of the provided first and rest Decimals
func F32Sum(first F32, rest ...F32) F32 {
    total := first
    for _, item := range rest {
        total = total.Add(item)
    }

    return total
}

// F32Avg returns the average value of the provided first and rest Decimals
func F32Avg(first F32, rest ...F32) F32 {
    count := F32FromInt32(int32(len(rest) + 1))
    sum := F32Sum(first, rest...)
    return sum.Div(count)
}

/************************************/
/*********** Conversions ************/
/************************************/

func (f F32) FloorToInt() int32 {
    return fix32.FloorToInt(f.Raw)
}

func (f F32) CeilToInt() int32 {
    return fix32.CeilToInt(f.Raw)
}

func (f F32) RoundToInt() int32 {
    return fix32.RoundToInt(f.Raw)
}

func (f F32) Float32() float32 {
    return fix32.ToFloat32(f.Raw)
}

func (f F32) Float64() float64 {
    return fix32.ToFloat64(f.Raw)
}

/************************************/
/************ Operators *************/
/************************************/

// Negate -f
func (f F32) Negate() F32 {
    return F32FromRaw(-f.Raw)
}

// Add f + v2
func (f F32) Add(v2 F32) F32 {
    return F32FromRaw(fix32.Add(f.Raw, v2.Raw))
}

// Sub f - v2
func (f F32) Sub(v2 F32) F32 {
    return F32FromRaw(fix32.Sub(f.Raw, v2.Raw))
}

// Mul f * v2
func (f F32) Mul(v2 F32) F32 {
    return F32FromRaw(fix32.Mul(f.Raw, v2.Raw))
}

// DivPrecise f / v2
func (f F32) DivPrecise(v2 F32) F32 {
    return F32FromRaw(fix32.DivPrecise(f.Raw, v2.Raw))
}

// Mod f % v2
func (f F32) Mod(v2 F32) F32 {
    return F32FromRaw(fix32.Mod(f.Raw, v2.Raw))
}

// AddVec2 f + v2
func (f F32) AddVec2(v2 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(f.Raw+v2.RawX, f.Raw+v2.RawY)
}

// SubVec2 f - v2
func (f F32) SubVec2(v2 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(f.Raw-v2.RawX, f.Raw-v2.RawY)
}

// MulVec2 f * v2
func (f F32) MulVec2(v2 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.Mul(f.Raw, v2.RawX), fix32.Mul(f.Raw, v2.RawY))
}

// DivPreciseVec2 f / v2
func (f F32) DivPreciseVec2(v2 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(fix32.DivPrecise(f.Raw, v2.RawX), fix32.DivPrecise(f.Raw, v2.RawY))
}

// ModVec2 f % v2
func (f F32) ModVec2(v2 F32Vec2) F32Vec2 {
    return F32Vec2FromRaw(f.Raw%v2.RawX, f.Raw%v2.RawY)
}

// AddVec3 f + v2
func (f F32) AddVec3(v2 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(f.Raw+v2.RawX, f.Raw+v2.RawY, f.Raw+v2.RawZ)
}

// SubVec3 f - v2
func (f F32) SubVec3(v2 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(f.Raw-v2.RawX, f.Raw-v2.RawY, f.Raw-v2.RawZ)
}

// MulVec3 f * v2
func (f F32) MulVec3(v2 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.Mul(f.Raw, v2.RawX), fix32.Mul(f.Raw, v2.RawY), fix32.Mul(f.Raw, v2.RawZ))
}

// DivPreciseVec3 f / v2
func (f F32) DivPreciseVec3(v2 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(fix32.DivPrecise(f.Raw, v2.RawX), fix32.DivPrecise(f.Raw, v2.RawY), fix32.DivPrecise(f.Raw, v2.RawZ))
}

// ModVec3 f % v2
func (f F32) ModVec3(v2 F32Vec3) F32Vec3 {
    return F32Vec3FromRaw(f.Raw%v2.RawX, f.Raw%v2.RawY, f.Raw%v2.RawZ)
}

// AddVec4 f + v2
func (f F32) AddVec4(v2 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(f.Raw+v2.RawX, f.Raw+v2.RawY, f.Raw+v2.RawZ, f.Raw+v2.RawW)
}

// SubVec4 f - v2
func (f F32) SubVec4(v2 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(f.Raw-v2.RawX, f.Raw-v2.RawY, f.Raw-v2.RawZ, f.Raw-v2.RawW)
}

// MulVec4 f * v2
func (f F32) MulVec4(v2 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.Mul(f.Raw, v2.RawX), fix32.Mul(f.Raw, v2.RawY), fix32.Mul(f.Raw, v2.RawZ), fix32.Mul(f.Raw, v2.RawW))
}

// DivPreciseVec4 f / v2
func (f F32) DivPreciseVec4(v2 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(fix32.DivPrecise(f.Raw, v2.RawX), fix32.DivPrecise(f.Raw, v2.RawY), fix32.DivPrecise(f.Raw, v2.RawZ), fix32.DivPrecise(f.Raw, v2.RawW))
}

// ModVec4 f % v2
func (f F32) ModVec4(v2 F32Vec4) F32Vec4 {
    return F32Vec4FromRaw(f.Raw%v2.RawX, f.Raw%v2.RawY, f.Raw%v2.RawZ, f.Raw%v2.RawW)
}

// Add2 f++
func (f F32) Add2() F32 {
    return F32FromRaw(f.Raw + fix32.One)
}

// Sub2 f--
func (f F32) Sub2() F32 {
    return F32FromRaw(f.Raw - fix32.One)
}

// EQ f == v2
func (f F32) EQ(v2 F32) bool {
    return f.Raw == v2.Raw
}

// NE f != v2
func (f F32) NE(v2 F32) bool {
    return f.Raw != v2.Raw
}

// LT f < v2
func (f F32) LT(v2 F32) bool {
    return f.Raw < v2.Raw
}

// LE f <= v2
func (f F32) LE(v2 F32) bool {
    return f.Raw <= v2.Raw
}

// GT f > v2
func (f F32) GT(v2 F32) bool {
    return f.Raw > v2.Raw
}

// GE f >= v2
func (f F32) GE(v2 F32) bool {
    return f.Raw >= v2.Raw
}

// RadToDeg 180 / F32.Pi
func (f F32) RadToDeg() F32 {
    return F32FromRaw(fix32.Mul(f.Raw, 3754943))
}

// DegToRad F32.Pi / 180
func (f F32) DegToRad() F32 {
    return F32FromRaw(fix32.Mul(f.Raw, 1143))
}

func (f F32) Div2() F32 {
    return F32FromRaw(f.Raw >> 1)
}

func (f F32) Abs() F32 {
    return F32FromRaw(fix32.Abs(f.Raw))
}

func (f F32) Nabs() F32 {
    return F32FromRaw(fix32.Nabs(f.Raw))
}

func (f F32) Sign() int32 {
    return fix32.Sign(f.Raw)
}

func (f F32) Ceil() F32 {
    return F32FromRaw(fix32.Ceil(f.Raw))
}

func (f F32) Floor() F32 {
    return F32FromRaw(fix32.Floor(f.Raw))
}

func (f F32) Round() F32 {
    return F32FromRaw(fix32.Round(f.Raw))
}

func (f F32) Fract() F32 {
    return F32FromRaw(fix32.Fract(f.Raw))
}

func (f F32) Div(b F32) F32 {
    return F32FromRaw(fix32.Div(f.Raw, b.Raw))
}

func (f F32) DivFast(b F32) F32 {
    return F32FromRaw(fix32.DivFast(f.Raw, b.Raw))
}

func (f F32) DivFastest(b F32) F32 {
    return F32FromRaw(fix32.DivFastest(f.Raw, b.Raw))
}

func (f F32) SqrtPrecise() F32 {
    return F32FromRaw(fix32.SqrtPrecise(f.Raw))
}

func (f F32) Sqrt() F32 {
    return F32FromRaw(fix32.Sqrt(f.Raw))
}

func (f F32) SqrtFast() F32 {
    return F32FromRaw(fix32.SqrtFast(f.Raw))
}

func (f F32) SqrtFastest() F32 {
    return F32FromRaw(fix32.SqrtFastest(f.Raw))
}

func (f F32) RSqrt() F32 {
    return F32FromRaw(fix32.RSqrt(f.Raw))
}

func (f F32) RSqrtFast() F32 {
    return F32FromRaw(fix32.RSqrtFast(f.Raw))
}

func (f F32) RSqrtFastest() F32 {
    return F32FromRaw(fix32.RSqrtFastest(f.Raw))
}

func (f F32) Rcp() F32 {
    return F32FromRaw(fix32.Rcp(f.Raw))
}

func (f F32) RcpFast() F32 {
    return F32FromRaw(fix32.RcpFast(f.Raw))
}

func (f F32) RcpFastest() F32 {
    return F32FromRaw(fix32.RcpFastest(f.Raw))
}

func (f F32) Exp() F32 {
    return F32FromRaw(fix32.Exp(f.Raw))
}

func (f F32) ExpFast() F32 {
    return F32FromRaw(fix32.ExpFast(f.Raw))
}

func (f F32) ExpFastest() F32 {
    return F32FromRaw(fix32.ExpFastest(f.Raw))
}

func (f F32) Exp2() F32 {
    return F32FromRaw(fix32.Exp2(f.Raw))
}

func (f F32) Exp2Fast() F32 {
    return F32FromRaw(fix32.Exp2Fast(f.Raw))
}

func (f F32) Exp2Fastest() F32 {
    return F32FromRaw(fix32.Exp2Fastest(f.Raw))
}

func (f F32) Log() F32 {
    return F32FromRaw(fix32.Log(f.Raw))
}

func (f F32) LogFast() F32 {
    return F32FromRaw(fix32.LogFast(f.Raw))
}

func (f F32) LogFastest() F32 {
    return F32FromRaw(fix32.LogFastest(f.Raw))
}

func (f F32) Log2() F32 {
    return F32FromRaw(fix32.Log2(f.Raw))
}

func (f F32) Log2Fast() F32 {
    return F32FromRaw(fix32.Log2Fast(f.Raw))
}

func (f F32) Log2Fastest() F32 {
    return F32FromRaw(fix32.Log2Fastest(f.Raw))
}

func (f F32) Sin() F32 {
    return F32FromRaw(fix32.Sin(f.Raw))
}

func (f F32) SinFast() F32 {
    return F32FromRaw(fix32.SinFast(f.Raw))
}

func (f F32) SinFastest() F32 {
    return F32FromRaw(fix32.SinFastest(f.Raw))
}

func (f F32) Cos() F32 {
    return F32FromRaw(fix32.Cos(f.Raw))
}

func (f F32) CosFast() F32 {
    return F32FromRaw(fix32.CosFast(f.Raw))
}

func (f F32) CosFastest() F32 {
    return F32FromRaw(fix32.CosFastest(f.Raw))
}

func (f F32) Tan() F32 {
    return F32FromRaw(fix32.Tan(f.Raw))
}

func (f F32) TanFast() F32 {
    return F32FromRaw(fix32.TanFast(f.Raw))
}

func (f F32) TanFastest() F32 {
    return F32FromRaw(fix32.TanFastest(f.Raw))
}

func (f F32) Asin() F32 {
    return F32FromRaw(fix32.Asin(f.Raw))
}

func (f F32) AsinFast() F32 {
    return F32FromRaw(fix32.AsinFast(f.Raw))
}

func (f F32) AsinFastest() F32 {
    return F32FromRaw(fix32.AsinFastest(f.Raw))
}

func (f F32) Acos() F32 {
    return F32FromRaw(fix32.Acos(f.Raw))
}

func (f F32) AcosFast() F32 {
    return F32FromRaw(fix32.AcosFast(f.Raw))
}

func (f F32) AcosFastest() F32 {
    return F32FromRaw(fix32.AcosFastest(f.Raw))
}

func (f F32) Atan() F32 {
    return F32FromRaw(fix32.Atan(f.Raw))
}

func (f F32) AtanFast() F32 {
    return F32FromRaw(fix32.AtanFast(f.Raw))
}

func (f F32) AtanFastest() F32 {
    return F32FromRaw(fix32.AtanFastest(f.Raw))
}

func (f F32) Atan2(x F32) F32 {
    return F32FromRaw(fix32.Atan2(f.Raw, x.Raw))
}

func (f F32) Atan2Fast(x F32) F32 {
    return F32FromRaw(fix32.Atan2Fast(f.Raw, x.Raw))
}

func (f F32) Atan2Fastest(x F32) F32 {
    return F32FromRaw(fix32.Atan2Fastest(f.Raw, x.Raw))
}

func (f F32) Pow(b F32) F32 {
    return F32FromRaw(fix32.Pow(f.Raw, b.Raw))
}

func (f F32) PowFast(b F32) F32 {
    return F32FromRaw(fix32.PowFast(f.Raw, b.Raw))
}

func (f F32) PowFastest(b F32) F32 {
    return F32FromRaw(fix32.PowFastest(f.Raw, b.Raw))
}

func (f F32) Clamp(min, max F32) F32 {
    return F32FromRaw(fix32.Clamp(f.Raw, min.Raw, max.Raw))
}

func (f F32) Clamp01() F32 {
    return F32FromRaw(fix32.Clamp(f.Raw, fix32.Zero, fix32.One))
}

func (f F32) Lerp(b, t F32) F32 {
    tb := t.Raw
    ta := fix32.One - tb
    return F32FromRaw(fix32.Mul(f.Raw, ta) + fix32.Mul(b.Raw, tb))
}

func (f F32) Equals(obj F32) bool {
    return reflect.DeepEqual(f, obj)
}

func (f F32) CompareTo(other F32) int32 {
    if f.Raw < other.Raw {
        return -1
    }
    if f.Raw > other.Raw {
        return +1
    }
    return 0
}

func (f F32) ToString() string {
    return fix32.ToString(f.Raw)
}
