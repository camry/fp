package fix32

import (
    "strconv"

    "github.com/camry/fp/fix64"
    "github.com/camry/fp/fixutil"
)

const (
    Shift        int32 = 16
    FractionMask int32 = (1 << Shift) - 1
    IntegerMask        = ^FractionMask
)

// Constants
const (
    Zero   int32 = 0
    Neg1   int32 = -1 << Shift
    One    int32 = 1 << Shift
    Two    int32 = 2 << Shift
    Three  int32 = 3 << Shift
    Four   int32 = 4 << Shift
    Half         = One >> 1
    Pi           = int32(13493037705 >> 16) // (int32)(Math.PI * 65536.0) << 16
    Pi2          = int32(26986075409 >> 16)
    PiHalf       = int32(6746518852 >> 16)
    E            = int32(11674931555 >> 16)
)

const (
    MinValue int32 = -2147483648
    MaxValue int32 = 2147483647
)

// Private constants
const (
    RcpLn2         = int32(0x171547652 >> 16) // 1.0 / log(2.0) ~= 1.4426950408889634
    RcpLog2E       = int32(2977044471 >> 16)  // 1.0 / log2(e) ~= 0.6931471805599453
    RcpTwoPi int32 = 683565276                // 1.0 / (4.0 * 0.5 * pi);  -- the 4.0 factor converts directly to s2.30
)

// FromInt32 Converts an integer to a fp-point value.
func FromInt32(v int32) int32 {
    return v << Shift
}

// FromFloat32 Converts a float32 to a fp-point value.
func FromFloat32(v float32) int32 {
    return int32(v * 65536.0)
}

// FromFloat64 Converts a float64 to a fp-point value.
func FromFloat64(v float64) int32 {
    return int32(v * 65536.0)
}

// CeilToInt Converts a fixed-point value into an integer by rounding it up to nearest integer.
func CeilToInt(v int32) int32 {
    return int32((v + (One - 1)) >> Shift)
}

// FloorToInt Converts a fixed-point value into an integer by rounding it down to nearest integer.
func FloorToInt(v int32) int32 {
    return int32(v >> Shift)
}

// RoundToInt Converts a fixed-point value into an integer by rounding it to nearest integer.
func RoundToInt(v int32) int32 {
    return int32((v + Half) >> Shift)
}

// ToFloat64 Converts a fixed-point value into a double.
func ToFloat64(v int32) float64 {
    return float64(v) * (1.0 / 65536.0)
}

// ToFloat32 Converts a FP value into a float.
func ToFloat32(v int32) float32 {
    return float32(v) * (1.0 / 65536.0)
}

// ToString Converts the value to a human readable string.
func ToString(v int32) string {
    return strconv.FormatFloat(ToFloat64(v), 'f', 16, 32)
}

// Abs Returns the absolute (positive) value of v.
func Abs(v int32) int32 {
    // note fails with MinValue
    mask := v >> 31
    return (v + mask) ^ mask
}

// Nabs Negative absolute value (returns -abs(x)).
func Nabs(v int32) int32 {
    return -Abs(v)
}

// Ceil Round up to nearest integer.
func Ceil(v int32) int32 {
    return (v + FractionMask) & IntegerMask
}

// Floor Round down to nearest integer.
func Floor(v int32) int32 {
    return v & IntegerMask
}

// Round to nearest integer.
func Round(v int32) int32 {
    return (v + Half) & IntegerMask
}

// Fract Returns the fractional part of x. Equal to 'x - floor(x)'.
func Fract(v int32) int32 {
    return v & FractionMask
}

// Min Returns the minimum of the two values.
func Min(a, b int32) int32 {
    if a < b {
        return a
    } else {
        return b
    }
}

// Max Returns the maximum of the two values.
func Max(a, b int32) int32 {
    if a > b {
        return a
    } else {
        return b
    }
}

// Clamp Returns the value clamped between min and max.
func Clamp(a, min, max int32) int32 {
    if a > max {
        return max
    } else {
        if a < min {
            return min
        } else {
            return a
        }
    }
}

// Sign Returns the sign of the value (-1 if negative, 0 if zero, 1 if positive).
func Sign(v int32) int32 {
    return (v >> 31) | int32((uint64(-v))>>31)
}

// Add Adds the two FP numbers together.
func Add(a, b int32) int32 {
    return a + b
}

// Sub Subtracts the two FP numbers from each other.
func Sub(a, b int32) int32 {
    return a - b
}

// Mul Multiplies two FP values together.
func Mul(a, b int32) int32 {
    return int32((int64(a) * int64(b)) >> Shift)
}

// Lerp Linearly interpolate from a to b by t.
func Lerp(a, b, t int32) int32 {
    ta := int64(a) * (int64(One) - int64(t))
    tb := int64(b) * int64(t)
    return int32((ta + tb) >> Shift)
}

func Nlz(v uint32) int32 {
    var n int32 = 0
    if v <= 0x0000FFFF {
        n = n + 16
        v = v << 16
    }
    if v <= 0x00FFFFFF {
        n = n + 8
        v = v << 8
    }
    if v <= 0x0FFFFFFF {
        n = n + 4
        v = v << 4
    }
    if v <= 0x3FFFFFFF {
        n = n + 2
        v = v << 2
    }
    if v <= 0x7FFFFFFF {
        n = n + 1
    }
    if v == 0 {
        return 32
    }
    return n
}

// DivPrecise Divides two FP values.
func DivPrecise(a, b int32) int32 {
    if b == MinValue || b == 0 {
        return 0
    }
    res := int32((int64(a) << Shift) / int64(b))
    return res
}

// Div Calculates division approximation.
func Div(a, b int32) int32 {
    if b == MinValue || b == 0 {
        return 0
    }
    return int32((int64(a) << 16) / int64(b))
}

// DivFast Calculates division approximation.
func DivFast(a, b int32) int32 {
    if b == MinValue || b == 0 {
        return 0
    }

    // Handle negative values.
    var sign int32
    if b < 0 {
        sign = -1
    } else {
        sign = 1
    }
    b *= sign

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 29 - Nlz(uint32(b))
    n := fixutil.ShiftRight(b, offset-28)
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly6(n - ONE)

    // Multiply by reciprocal, apply exponent, convert back to s16.16.
    y := fixutil.Qmul30(res, a)
    return fixutil.ShiftRight(sign*y, offset-14)
}

// DivFastest Calculates division approximation.
func DivFastest(a, b int32) int32 {
    if b == MinValue || b == 0 {
        return 0
    }

    // Handle negative values.
    var sign int32
    if b < 0 {
        sign = -1
    } else {
        sign = 1
    }
    b *= sign

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 29 - Nlz(uint32(b))
    n := fixutil.ShiftRight(b, offset-28)
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4(n - ONE)

    // Multiply by reciprocal, apply exponent, convert back to s16.16.
    y := fixutil.Qmul30(res, a)
    return fixutil.ShiftRight(sign*y, offset-14)
}

// Mod Divides two FP values and returns the modulus.
func Mod(a, b int32) int32 {
    if b == 0 {
        return 0
    }
    return a % b
}

// SqrtPrecise Calculates the square root of the given number.
func SqrtPrecise(a int32) int32 {
    // Adapted from https://github.com/chmike/fpsqrt
    if a <= 0 {
        return 0
    }
    var r = uint32(a)
    var b uint32 = 0x40000000
    var q uint32 = 0
    for b > 0x40 {
        t := q + b
        if r >= t {
            r -= t
            q = t + b
        }
        r <<= 1
        b >>= 1
    }
    q >>= 8
    return int32(q)
}

func Sqrt(x int32) int32 {
    // Return 0 for all non-positive values.
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const SQRT2 int32 = 1518500249 // sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)
    y := fixutil.SqrtPoly3Lut8(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = SQRT2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s16.16.
    yr := fixutil.Qmul30(adjust, y)
    return fixutil.ShiftRight(yr, 14-offset)
}

func SqrtFast(x int32) int32 {
    // Return 0 for all non-positive values.
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const SQRT2 int32 = 1518500249 // sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)
    y := fixutil.SqrtPoly4(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = SQRT2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s16.16.
    yr := fixutil.Qmul30(adjust, y)
    return fixutil.ShiftRight(yr, 14-offset)
}

func SqrtFastest(x int32) int32 {
    // Return 0 for all non-positive values.
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const SQRT2 int32 = 1518500249 // sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)
    y := fixutil.SqrtPoly3(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = SQRT2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s16.16.
    yr := fixutil.Qmul30(adjust, y)
    return fixutil.ShiftRight(yr, 14-offset)
}

// RSqrt Calculates the reciprocal square root.
func RSqrt(x int32) int32 {
    // Return 0 for invalid values
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const HalfSqrt2 int32 = 759250125 // 0.5 * sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 1 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset)
    y := fixutil.RSqrtPoly3Lut16(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = HalfSqrt2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s16.16.
    yr := fixutil.Qmul30(adjust, y)
    return fixutil.ShiftRight(yr, offset+21)
}

// RSqrtFast Calculates the reciprocal square root.
func RSqrtFast(x int32) int32 {
    // Return 0 for invalid values
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const HalfSqrt2 int32 = 759250125 // 0.5 * sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 1 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset)
    y := fixutil.RSqrtPoly5(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = HalfSqrt2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s16.16.
    yr := fixutil.Qmul30(adjust, y)
    return fixutil.ShiftRight(yr, offset+21)
}

// RSqrtFastest Calculates the reciprocal square root.
func RSqrtFastest(x int32) int32 {
    // Return 0 for invalid values
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const HalfSqrt2 int32 = 759250125 // 0.5 * sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 1 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset)
    y := fixutil.RSqrtPoly3(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = HalfSqrt2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s16.16.
    yr := fixutil.Qmul30(adjust, y)
    return fixutil.ShiftRight(yr, offset+21)
}

// Rcp Calculates reciprocal approximation.
func Rcp(x int32) int32 {
    if x == MinValue || x == 0 {
        return 0
    }

    // Handle negative values.
    var sign int32
    if x < 0 {
        sign = -1
    } else {
        sign = 1
    }
    x *= sign

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 29 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-28)
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4Lut8(n - ONE)

    // Apply exponent, convert back to s16.16.
    return fixutil.ShiftRight(sign*res, offset)
}

// RcpFast Calculates reciprocal approximation.
func RcpFast(x int32) int32 {
    if x == MinValue || x == 0 {
        return 0
    }

    // Handle negative values.
    var sign int32
    if x < 0 {
        sign = -1
    } else {
        sign = 1
    }
    x *= sign

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 29 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-28)
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly6(n - ONE)

    // Apply exponent, convert back to s16.16.
    return fixutil.ShiftRight(sign*res, offset)
}

// RcpFastest Calculates reciprocal approximation.
func RcpFastest(x int32) int32 {
    if x == MinValue || x == 0 {
        return 0
    }

    // Handle negative values.
    var sign int32
    if x < 0 {
        sign = -1
    } else {
        sign = 1
    }
    x *= sign

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 29 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-28)
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4(n - ONE)

    // Apply exponent, convert back to s16.16.
    return fixutil.ShiftRight(sign*res, offset)
}

// Exp2 Calculates the base 2 exponent.
func Exp2(x int32) int32 {
    // Handle values that would under or overflow.
    if x >= 15*One {
        return MaxValue
    }
    if x <= -16*One {
        return 0
    }

    // Compute exp2 for fractional part.
    k := (x & FractionMask) << 14
    y := fixutil.Exp2Poly5(k)

    // Combine integer and fractional result, and convert back to s16.16.
    intPart := x >> Shift
    return fixutil.ShiftRight(y, 14-intPart)
}

// Exp2Fast Calculates the base 2 exponent.
func Exp2Fast(x int32) int32 {
    // Handle values that would under or overflow.
    if x >= 15*One {
        return MaxValue
    }
    if x <= -16*One {
        return 0
    }

    // Compute exp2 for fractional part.
    k := (x & FractionMask) << 14
    y := fixutil.Exp2Poly4(k)

    // Combine integer and fractional result, and convert back to s16.16.
    intPart := x >> Shift
    return fixutil.ShiftRight(y, 14-intPart)
}

// Exp2Fastest Calculates the base 2 exponent.
func Exp2Fastest(x int32) int32 {
    // Handle values that would under or overflow.
    if x >= 15*One {
        return MaxValue
    }
    if x <= -16*One {
        return 0
    }

    // Compute exp2 for fractional part.
    k := (x & FractionMask) << 14
    y := fixutil.Exp2Poly3(k)

    // Combine integer and fractional result, and convert back to s16.16.
    intPart := x >> Shift
    return fixutil.ShiftRight(y, 14-intPart)
}

func Exp(x int32) int32 {
    // e^x == 2^(x / ln(2))
    return Exp2(Mul(x, RcpLn2))
}

func ExpFast(x int32) int32 {
    // e^x == 2^(x / ln(2))
    return Exp2Fast(Mul(x, RcpLn2))
}

func ExpFastest(x int32) int32 {
    // e^x == 2^(x / ln(2))
    return Exp2Fastest(Mul(x, RcpLn2))
}

func Log(x int32) int32 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)

    // Polynomial approximation.
    const ONE int32 = 1 << 30
    y := fixutil.LogPoly5Lut8(n - ONE)

    // Combine integer and fractional parts (into s16.16).
    return offset*RcpLog2E + (y >> 14)
}

func LogFast(x int32) int32 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)

    // Polynomial approximation.
    const ONE int32 = 1 << 30
    y := fixutil.LogPoly3Lut8(n - ONE)

    // Combine integer and fractional parts (into s16.16).
    return offset*RcpLog2E + (y >> 14)
}

func LogFastest(x int32) int32 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)

    // Polynomial approximation.
    const ONE int32 = 1 << 30
    y := fixutil.LogPoly5(n - ONE)

    // Combine integer and fractional parts (into s16.16).
    return offset*RcpLog2E + (y >> 14)
}

func Log2(x int32) int32 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)

    // Polynomial approximation of mantissa.
    const ONE int32 = 1 << 30
    y := fixutil.Log2Poly4Lut16(n - ONE)

    // Combine integer and fractional parts (into s16.16).
    return (offset << Shift) + (y >> 14)
}

func Log2Fast(x int32) int32 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)

    // Polynomial approximation of mantissa.
    const ONE int32 = 1 << 30
    y := fixutil.Log2Poly3Lut16(n - ONE)

    // Combine integer and fractional parts (into s16.16).
    return (offset << Shift) + (y >> 14)
}

func Log2Fastest(x int32) int32 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 15 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset-14)

    // Polynomial approximation of mantissa.
    const ONE int32 = 1 << 30
    y := fixutil.Log2Poly5(n - ONE)

    // Combine integer and fractional parts (into s16.16).
    return (offset << Shift) + (y >> 14)
}

// Pow Calculates x to the power of the exponent.
func Pow(x, exponent int32) int32 {
    // n^0 == 1
    if exponent == 0 {
        return One
    }
    // Return 0 for invalid values
    if x <= 0 {
        return 0
    }
    return Exp(Mul(exponent, Log(x)))
}

// PowFast Calculates x to the power of the exponent.
func PowFast(x, exponent int32) int32 {
    // n^0 == 1
    if exponent == 0 {
        return One
    }
    // Return 0 for invalid values
    if x <= 0 {
        return 0
    }
    return ExpFast(Mul(exponent, LogFast(x)))
}

// PowFastest Calculates x to the power of the exponent.
func PowFastest(x, exponent int32) int32 {
    // n^0 == 1
    if exponent == 0 {
        return One
    }
    // Return 0 for invalid values
    if x <= 0 {
        return 0
    }
    return ExpFastest(Mul(exponent, LogFastest(x)))
}

func UnitSin(z int32) int32 {
    // See: http://www.coranac.com/2009/07/sines/

    // Handle quadrants 1 and 2 by mirroring the [1, 3] range to [-1, 1] (by calculating 2 - z).
    // The if condition uses the fact that for the quadrants of interest are 0b01 and 0b10 (top two bits are different).
    if (z ^ (z << 1)) < 0 {
        z = int32((1 << 31) - int64(z))
    }

    // Now z is in range [-1, 1].
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    zz := fixutil.Qmul30(z, z)
    res := fixutil.Qmul30(fixutil.SinPoly4(zz), z)

    // Return as s2.30.
    return res
}

func UnitSinFast(z int32) int32 {
    // See: http://www.coranac.com/2009/07/sines/

    // Handle quadrants 1 and 2 by mirroring the [1, 3] range to [-1, 1] (by calculating 2 - z).
    // The if condition uses the fact that for the quadrants of interest are 0b01 and 0b10 (top two bits are different).
    if (z ^ (z << 1)) < 0 {
        z = int32((1 << 31) - int64(z))
    }

    // Now z is in range [-1, 1].
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    zz := fixutil.Qmul30(z, z)
    res := fixutil.Qmul30(fixutil.SinPoly3(zz), z)

    // Return as s2.30.
    return res
}

func UnitSinFastest(z int32) int32 {
    // See: http://www.coranac.com/2009/07/sines/

    // Handle quadrants 1 and 2 by mirroring the [1, 3] range to [-1, 1] (by calculating 2 - z).
    // The if condition uses the fact that for the quadrants of interest are 0b01 and 0b10 (top two bits are different).
    if (z ^ (z << 1)) < 0 {
        z = int32((1 << 31) - int64(z))
    }

    // Now z is in range [-1, 1].
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    zz := fixutil.Qmul30(z, z)
    res := fixutil.Qmul30(fixutil.SinPoly2(zz), z)

    // Return as s2.30.
    return res
}

func Sin(x int32) int32 {
    // Map [0, 2pi] to [0, 4] (as s2.30).
    // This also wraps the values into one period.
    z := Mul(RcpTwoPi, x)

    // Compute sin from s2.30 and convert back to s16.16.
    return UnitSin(z) >> 14
}

func SinFast(x int32) int32 {
    // Map [0, 2pi] to [0, 4] (as s2.30).
    // This also wraps the values into one period.
    z := Mul(RcpTwoPi, x)

    // Compute sin from s2.30 and convert back to s16.16.
    return UnitSinFast(z) >> 14
}

func SinFastest(x int32) int32 {
    // Map [0, 2pi] to [0, 4] (as s2.30).
    // This also wraps the values into one period.
    z := Mul(RcpTwoPi, x)

    // Compute sin from s2.30 and convert back to s16.16.
    return UnitSinFastest(z) >> 14
}

func Cos(x int32) int32 {
    return Sin(x + PiHalf)
}

func CosFast(x int32) int32 {
    return SinFast(x + PiHalf)
}

func CosFastest(x int32) int32 {
    return SinFastest(x + PiHalf)
}

func Tan(x int32) int32 {
    z := Mul(RcpTwoPi, x)
    sinX := UnitSin(z)
    cosX := UnitSin(z + (1 << 30))
    return Div(sinX, cosX)
}

func TanFast(x int32) int32 {
    z := Mul(RcpTwoPi, x)
    sinX := UnitSinFast(z)
    cosX := UnitSinFast(z + (1 << 30))
    return DivFast(sinX, cosX)
}

func TanFastest(x int32) int32 {
    z := Mul(RcpTwoPi, x)
    sinX := UnitSinFastest(z)
    cosX := UnitSinFastest(z + (1 << 30))
    return DivFastest(sinX, cosX)
}

func Atan2Div(y, x int32) int32 {
    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    const ONE int32 = 1 << 30
    const HALF int32 = 1 << 29
    offset := 1 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset)

    // Polynomial approximation of reciprocal.
    oox := fixutil.RcpPoly4Lut8(n - ONE)

    // Apply exponent and multiply.
    yr := fixutil.ShiftRight(y, offset)
    return fixutil.Qmul30(yr, oox)
}

func Atan2DivFast(y, x int32) int32 {
    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    const ONE int32 = 1 << 30
    const HALF int32 = 1 << 29
    offset := 1 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset)

    // Polynomial approximation of reciprocal.
    oox := fixutil.RcpPoly6(n - ONE)

    // Apply exponent and multiply.
    yr := fixutil.ShiftRight(y, offset)
    return fixutil.Qmul30(yr, oox)
}

func Atan2DivFastest(y, x int32) int32 {
    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    const ONE int32 = 1 << 30
    const HALF int32 = 1 << 29
    offset := 1 - Nlz(uint32(x))
    n := fixutil.ShiftRight(x, offset)

    // Polynomial approximation of reciprocal.
    oox := fixutil.RcpPoly4(n - ONE)

    // Apply exponent and multiply.
    yr := fixutil.ShiftRight(y, offset)
    return fixutil.Qmul30(yr, oox)
}

func Atan2(y, x int32) int32 {
    // See: https://www.dsprelated.com/showarticle/1052.php
    if x == 0 {
        if y > 0 {
            return PiHalf
        }
        if y < 0 {
            return -PiHalf
        }

        return 0
    }

    nx := Abs(x)
    ny := Abs(y)
    negMask := (x ^ y) >> 31

    if nx >= ny {
        k := Atan2Div(ny, nx)
        z := fixutil.AtanPoly5Lut8(k)
        angle := (negMask ^ (z >> 14)) - negMask
        if x > 0 {
            return angle
        }
        if y >= 0 {
            return angle + Pi
        }
        return angle - Pi
    } else {
        k := Atan2Div(nx, ny)
        z := fixutil.AtanPoly5Lut8(k)
        angle := negMask ^ (z >> 14)

        if y > 0 {
            return PiHalf - angle
        } else {
            return -PiHalf - angle
        }
    }
}

func Atan2Fast(y, x int32) int32 {
    // See: https://www.dsprelated.com/showarticle/1052.php
    if x == 0 {
        if y > 0 {
            return PiHalf
        }
        if y < 0 {
            return -PiHalf
        }
        return 0
    }

    nx := Abs(x)
    ny := Abs(y)
    negMask := (x ^ y) >> 31

    if nx >= ny {
        k := Atan2DivFast(ny, nx)
        z := fixutil.AtanPoly3Lut8(k)
        angle := negMask ^ (z >> 14)
        if x > 0 {
            return angle
        }
        if y >= 0 {
            return angle + Pi
        }
        return angle - Pi
    } else {
        k := Atan2DivFast(nx, ny)
        z := fixutil.AtanPoly3Lut8(k)
        angle := negMask ^ (z >> 14)

        if y > 0 {
            return PiHalf - angle
        } else {
            return -PiHalf - angle
        }
    }
}

func Atan2Fastest(y, x int32) int32 {
    // See: https://www.dsprelated.com/showarticle/1052.php
    if x == 0 {
        if y > 0 {
            return PiHalf
        }
        if y < 0 {
            return -PiHalf
        }
        return 0
    }

    nx := Abs(x)
    ny := Abs(y)
    negMask := (x ^ y) >> 31

    if nx >= ny {
        k := Atan2DivFastest(ny, nx)
        z := fixutil.AtanPoly4(k)
        angle := negMask ^ (z >> 14)
        if x > 0 {
            return angle
        }
        if y >= 0 {
            return angle + Pi
        }
        return angle - Pi
    } else {
        k := Atan2DivFastest(nx, ny)
        z := fixutil.AtanPoly4(k)
        angle := negMask ^ (z >> 14)

        if y > 0 {
            return PiHalf - angle
        } else {
            return -PiHalf - angle
        }
    }
}

func Asin(x int32) int32 {
    if x < -One || x > One {
        return 0
    }
    // Compute Atan2(x, Sqrt((1+x) * (1-x))), using s32.32.
    xx := int64(One+x) * int64(One-x)
    y := fix64.Sqrt(xx)
    return int32(fix64.Atan2(int64(x)<<16, y) >> 16)
}

func AsinFast(x int32) int32 {
    if x < -One || x > One {
        return 0
    }
    // Compute Atan2(x, Sqrt((1+x) * (1-x))), using s32.32.
    xx := int64(One+x) * int64(One-x)
    y := fix64.SqrtFast(xx)
    return int32(fix64.Atan2Fast(int64(x)<<16, y) >> 16)
}

func AsinFastest(x int32) int32 {
    if x < -One || x > One {
        return 0
    }
    // Compute Atan2(x, Sqrt((1+x) * (1-x))), using s32.32.
    xx := int64(One+x) * int64(One-x)
    y := fix64.SqrtFastest(xx)
    return int32(fix64.Atan2Fastest(int64(x)<<16, y) >> 16)
}

func Acos(x int32) int32 {
    if x < -One || x > One {
        return 0
    }

    // Compute Atan2(Sqrt((1+x) * (1-x)), x), using s32.32.
    xx := int64(One+x) * int64(One-x)
    y := fix64.Sqrt(xx)
    return int32(fix64.Atan2(y, int64(x)<<16) >> 16)
}

func AcosFast(x int32) int32 {
    if x < -One || x > One {
        return 0
    }

    // Compute Atan2(Sqrt((1+x) * (1-x)), x), using s32.32.
    xx := int64(One+x) * int64(One-x)
    y := fix64.SqrtFast(xx)
    return int32(fix64.Atan2Fast(y, int64(x)<<16) >> 16)
}

func AcosFastest(x int32) int32 {
    if x < -One || x > One {
        return 0
    }

    // Compute Atan2(Sqrt((1+x) * (1-x)), x), using s32.32.
    xx := int64(One+x) * int64(One-x)
    y := fix64.SqrtFastest(xx)
    return int32(fix64.Atan2Fastest(y, int64(x)<<16) >> 16)
}

func Atan(x int32) int32 {
    return Atan2(x, One)
}

func AtanFast(x int32) int32 {
    return Atan2Fast(x, One)
}

func AtanFastest(x int32) int32 {
    return Atan2Fastest(x, One)
}
