package fix64

import (
    "strconv"

    "github.com/camry/fp/fixutil"
)

// Direct fixed point (signed 32.32) functions.

const (
    Shift        int32 = 32
    FractionMask int64 = (1 << Shift) - 1
    IntegerMask        = ^FractionMask
)

// Constants
const (
    Zero   int64 = 0
    Neg1   int64 = -1 << Shift
    One    int64 = 1 << Shift
    Two    int64 = 2 << Shift
    Three  int64 = 3 << Shift
    Four   int64 = 4 << Shift
    Half         = One >> 1
    Pi     int64 = 13493037705 // (int64)(Math.PI * 65536.0) << 16;
    Pi2    int64 = 26986075409
    PiHalf int64 = 6746518852
    E      int64 = 11674931555
)

const (
    MinValue int64 = -9223372036854775808
    MaxValue int64 = 9223372036854775807
)

// Private constants
const (
    RcpLn2    int64 = 0x171547652 // 1.0 / log(2.0) ~= 1.4426950408889634
    RcpLog2E  int64 = 2977044471  // 1.0 / log2(e) ^= 0.6931471805599453
    RcpHalfPi int32 = 683565276   // 1.0 / (4.0 * 0.5 * Math.PI);  // the 4.0 factor converts directly to s2.30
)

// FromInt32 Converts an integer to a fp-point value.
func FromInt32(v int32) int64 {
    return int64(v) << Shift
}

// FromInt64 Converts an integer to a fp-point value.
func FromInt64(v int64) int64 {
    return v << Shift
}

// FromFloat32 Converts a float32 to a fp-point value.
func FromFloat32(v float32) int64 {
    return int64(v * 4294967296.0)
}

// FromFloat64 Converts a float64 to a fp-point value.
func FromFloat64(v float64) int64 {
    return int64(v * 4294967296.0)
}

// CeilToInt Converts a fp-point value into an integer by rounding it up to nearest integer.
func CeilToInt(v int64) int32 {
    return int32((v + (One - 1)) >> Shift)
}

// FloorToInt Converts a fp-point value into an integer by rounding it down to nearest integer.
func FloorToInt(v int64) int32 {
    return int32(v >> Shift)
}

// RoundToInt Converts a fp-point value into an integer by rounding it to nearest integer.
func RoundToInt(v int64) int32 {
    return int32((v + Half) >> Shift)
}

// ToFloat64 Converts a fp-point value into a float64.
func ToFloat64(v int64) float64 {
    return float64(v) * (1.0 / 4294967296.0)
}

// ToFloat32 Converts a FP value into a float.
func ToFloat32(v int64) float32 {
    return float32(v) * (1.0 / 4294967296.0)
}

// ToString Converts the value to a human readable string.
func ToString(v int64) string {
    return strconv.FormatFloat(ToFloat64(v), 'f', 32, 64)
}

// Abs Returns the absolute (positive) value of v.
func Abs(v int64) int64 {
    // \note fails with LONG_MIN
    mask := v >> 63
    return (v + mask) ^ mask
}

// Nabs Negative absolute value (returns -abs(x)).
func Nabs(v int64) int64 {
    return -Abs(v)
}

// Ceil Round up to nearest integer.
func Ceil(v int64) int64 {
    return (v + FractionMask) & IntegerMask
}

// Floor Round down to nearest integer.
func Floor(v int64) int64 {
    return v & IntegerMask
}

// Round to nearest integer.
func Round(v int64) int64 {
    return (v + Half) & IntegerMask
}

// Fract Returns the fractional part of x. Equal to 'x - floor(x)'.
func Fract(v int64) int64 {
    return v & FractionMask
}

// Min Returns the minimum of the two values.
func Min(a, b int64) int64 {
    if a < b {
        return a
    } else {
        return b
    }
}

// Max Returns the maximum of the two values.
func Max(a, b int64) int64 {
    if a > b {
        return a
    } else {
        return b
    }
}

// Clamp Returns the value clamped between min and max.
func Clamp(a, min, max int64) int64 {
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
func Sign(v int64) int32 {
    return int32((v >> 63) | int64((uint64(-v))>>63))
}

// Add Adds the two FP numbers together.
func Add(a, b int64) int64 {
    return a + b
}

// Sub Subtracts the two FP numbers from each other.
func Sub(a, b int64) int64 {
    return a - b
}

// Mul Multiplies two FP values together.
func Mul(a, b int64) int64 {
    ai := a >> Shift
    af := a & FractionMask
    bi := b >> Shift
    bf := b & FractionMask
    return fixutil.LogicalShiftRight(af*bf, Shift) + ai*b + af*bi
}

func MulIntLongLow(a int32, b int64) int32 {
    bi := b >> Shift
    bf := b & FractionMask
    return int32(fixutil.LogicalShiftRight(int64(a)*bf, Shift) + int64(a)*bi)
}

func MulIntLongLong(a int32, b int64) int64 {
    bi := b >> Shift
    bf := b & FractionMask
    return fixutil.LogicalShiftRight(int64(a)*bf, Shift) + int64(a)*bi
}

// Lerp Linearly interpolate from a to b by t.
func Lerp(a, b, t int64) int64 {
    return Mul(a, t) + Mul(b, One-t)
}

func nlz(v uint64) int32 {
    var n int32 = 0
    if v <= 0x00000000FFFFFFFF {
        n = n + 32
        v = v << 32
    }
    if v <= 0x0000FFFFFFFFFFFF {
        n = n + 16
        v = v << 16
    }
    if v <= 0x00FFFFFFFFFFFFFF {
        n = n + 8
        v = v << 8
    }
    if v <= 0x0FFFFFFFFFFFFFFF {
        n = n + 4
        v = v << 4
    }
    if v <= 0x3FFFFFFFFFFFFFFF {
        n = n + 2
        v = v << 2
    }
    if v <= 0x7FFFFFFFFFFFFFFF {
        n = n + 1
    }
    if v == 0 {
        return 64
    }
    return n
}

// DivPrecise Divides two FP values.
func DivPrecise(argA, argB int64) int64 {
    signDif := argA ^ argB

    const b uint64 = 0x100000000 // Number base (32 bits)
    var absArgA uint64
    if argA < 0 {
        absArgA = uint64(-argA)
    } else {
        absArgA = uint64(argA)
    }
    u1 := absArgA >> 32
    u0 := absArgA << 32
    var v uint64
    if argB < 0 {
        v = uint64(-argB)
    } else {
        v = uint64(argB)
    }

    // Overflow?
    if u1 >= v {
        // rem = 0;
        return 0x7fffffffffffffff
    }

    // Shift amount for norm
    s := nlz(v)    // 0 <= s <= 63
    v = v << s     // Normalize the divisor
    vn1 := v >> 32 // Break the divisor into two 32-bit digits
    vn0 := v & 0xffffffff

    un32 := (u1 << s) | (u0>>(64-s))&uint64(int64(-s)>>63)
    un10 := u0 << s // Shift dividend left

    un1 := un10 >> 32 // Break the right half of dividend into two digits
    un0 := un10 & 0xffffffff

    // Compute the first quotient digit, q1
    q1 := un32 / vn1
    rHat := un32 - q1*vn1
    for rHat < b {
        if (q1 >= b) || ((q1 * vn0) > (b*rHat + un1)) {
            q1 = q1 - 1
            rHat = rHat + vn1
        } else {
            break
        }
    }

    un21 := un32*b + un1 - q1*v // Multiply and subtract

    // Compute the second quotient digit, q0
    q0 := un21 / vn1
    rHat = un21 - q0*vn1
    for rHat < b {
        if (q0 >= b) || ((q0 * vn0) > (b*rHat + un0)) {
            q0 = q0 - 1
            rHat = rHat + vn1
        } else {
            break
        }
    }

    // Calculate the remainder
    // uint64 r = (un21 * b + un0 - q0 * v) >> s;
    // rem = (long)r;
    ret := q1*b + q0
    if signDif < 0 {
        return -int64(ret)
    } else {
        return int64(ret)
    }
}

// Div Calculates division approximation.
func Div(a, b int64) int64 {
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
    b *= int64(sign)

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 31 - nlz(uint64(b))
    n := int32(fixutil.Int64ShiftRight(b, offset+2))
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4Lut8(n - ONE)

    // Apply exponent, convert back to s32.32.
    y := MulIntLongLong(res, a) << 2
    return fixutil.Int64ShiftRight(int64(sign)*y, offset)
}

// DivFast Calculates division approximation.
func DivFast(a, b int64) int64 {
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
    b *= int64(sign)

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 31 - nlz(uint64(b))
    n := int32(fixutil.Int64ShiftRight(b, offset+2))
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly6(n - ONE)

    // Apply exponent, convert back to s32.32.
    y := MulIntLongLong(res, a) << 2
    return fixutil.Int64ShiftRight(int64(sign)*y, offset)
}

// DivFastest Calculates division approximation.
func DivFastest(a, b int64) int64 {
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
    b *= int64(sign)

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 31 - nlz(uint64(b))
    n := int32(fixutil.Int64ShiftRight(b, offset+2))
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4(n - ONE)

    // Apply exponent, convert back to s32.32.
    y := MulIntLongLong(res, a) << 2
    return fixutil.Int64ShiftRight(int64(sign)*y, offset)
}

// Mod Divides two FP values and returns the modulus.
func Mod(a, b int64) int64 {
    if b == 0 {
        return 0
    }
    return a % b
}

// SqrtPrecise Calculates the square root of the given number.
func SqrtPrecise(a int64) int64 {
    if a <= 0 {
        return 0
    }
    r := uint64(a)
    b := uint64(0x4000000000000000)
    q := uint64(0)
    for b > 0x40 {
        t := q + b
        if r >= t {
            r -= t
            q = t + b
        }
        r <<= 1
        b >>= 1
    }
    q >>= 16
    return int64(q)
}

func Sqrt(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const SQRT2 int32 = 1518500249 // sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := fixutil.SqrtPoly3Lut8(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = SQRT2
    } else {

        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s32.32.
    yr := int64(fixutil.Qmul30(adjust, y) << 2)
    if offset >= 0 {
        return yr << offset
    } else {
        return yr >> -offset
    }
}

func SqrtFast(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const SQRT2 int32 = 1518500249 // sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := fixutil.SqrtPoly4(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = SQRT2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s32.32.
    yr := int64(fixutil.Qmul30(adjust, y) << 2)
    if offset >= 0 {
        return yr << offset
    } else {
        return yr >> -offset
    }
}

func SqrtFastest(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const SQRT2 int32 = 1518500249 // sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := fixutil.SqrtPoly3(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = SQRT2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s32.32.
    yr := int64(fixutil.Qmul30(adjust, y) << 2)
    if offset >= 0 {
        return yr << offset
    } else {
        return yr >> -offset
    }
}

// RSqrt Calculates the reciprocal square root.
func RSqrt(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const HalfSqrt2 int32 = 759250125 // 0.5 * sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := fixutil.RSqrtPoly3Lut16(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = HalfSqrt2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s32.32.
    yr := int64(fixutil.Qmul30(adjust, y) << 2)
    if offset >= 0 {
        return yr >> offset
    } else {
        return yr << -offset
    }
}

// RSqrtFast Calculates the reciprocal square root.
func RSqrtFast(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const HalfSqrt2 int32 = 759250125 // 0.5 * sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := fixutil.RSqrtPoly5(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = HalfSqrt2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s32.32.
    yr := int64(fixutil.Qmul30(adjust, y) << 2)
    if offset >= 0 {
        return yr >> offset
    } else {
        return yr << -offset
    }
}

// RSqrtFastest Calculates the reciprocal square root.
func RSqrtFastest(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Constants (s2.30).
    const ONE int32 = 1 << 30
    const HalfSqrt2 int32 = 759250125 // 0.5 * sqrt(2.0)

    // Normalize input into [1.0, 2.0( range (as s2.30).
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := fixutil.RSqrtPoly3(n - ONE)

    // Divide offset by 2 (to get sqrt), compute adjust value for odd exponents.
    var adjust int32
    if (offset & 1) != 0 {
        adjust = HalfSqrt2
    } else {
        adjust = ONE
    }
    offset = offset >> 1

    // Apply exponent, convert back to s32.32.
    yr := int64(fixutil.Qmul30(adjust, y) << 2)
    if offset >= 0 {
        return yr >> offset
    } else {
        return yr << -offset
    }
}

// Rcp Calculates reciprocal approximation.
func Rcp(x int64) int64 {
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
    x *= int64(sign)

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 31 - nlz(uint64(x))
    n := int32(fixutil.Int64ShiftRight(x, offset+2))
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4Lut8(n - ONE)
    y := int64(sign*res) << 2

    // Apply exponent, convert back to s32.32.
    return fixutil.Int64ShiftRight(y, offset)
}

// RcpFast Calculates reciprocal approximation.
func RcpFast(x int64) int64 {
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
    x *= int64(sign)

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 31 - nlz(uint64(x))
    n := int32(fixutil.Int64ShiftRight(x, offset+2))
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly6(n - ONE)
    y := int64(sign*res) << 2

    // Apply exponent, convert back to s32.32.
    return fixutil.Int64ShiftRight(y, offset)
}

// RcpFastest Calculates reciprocal approximation.
func RcpFastest(x int64) int64 {
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
    x *= int64(sign)

    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    offset := 31 - nlz(uint64(x))
    n := int32(fixutil.Int64ShiftRight(x, offset+2))
    const ONE int32 = 1 << 30

    // Polynomial approximation.
    res := fixutil.RcpPoly4(n - ONE)
    y := int64(sign*res) << 2

    // Apply exponent, convert back to s32.32.
    return fixutil.Int64ShiftRight(y, offset)
}

// Exp2 Calculates the base 2 exponent.
func Exp2(x int64) int64 {
    // Handle values that would under or overflow.
    if x >= 32*One {
        return MaxValue
    }
    if x <= -32*One {
        return 0
    }

    // Compute exp2 for fractional part.
    k := int32((x & FractionMask) >> 2)
    y := int64(fixutil.Exp2Poly5(k)) << 2

    // Combine integer and fractional result, and convert back to s32.32.
    intPart := int32(x >> Shift)
    if intPart >= 0 {
        return y << intPart
    } else {
        return y >> -intPart
    }
}

// Exp2Fast Calculates the base 2 exponent.
func Exp2Fast(x int64) int64 {
    // Handle values that would under or overflow.
    if x >= 32*One {
        return MaxValue
    }
    if x <= -32*One {
        return 0
    }

    // Compute exp2 for fractional part.
    k := int32((x & FractionMask) >> 2)
    y := int64(fixutil.Exp2Poly4(k)) << 2

    // Combine integer and fractional result, and convert back to s32.32.
    intPart := int32(x >> Shift)
    if intPart >= 0 {
        return y << intPart
    } else {
        return y >> -intPart
    }
}

// Exp2Fastest Calculates the base 2 exponent.
func Exp2Fastest(x int64) int64 {
    // Handle values that would under or overflow.
    if x >= 32*One {
        return MaxValue
    }
    if x <= -32*One {
        return 0
    }

    // Compute exp2 for fractional part.
    k := int32((x & FractionMask) >> 2)
    y := int64(fixutil.Exp2Poly3(k)) << 2

    // Combine integer and fractional result, and convert back to s32.32.
    intPart := int32(x >> Shift)
    if intPart >= 0 {
        return y << intPart
    } else {
        return y >> -intPart
    }
}

func Exp(x int64) int64 {
    // e^x == 2^(x / ln(2))
    return Exp2(Mul(x, RcpLn2))
}

func ExpFast(x int64) int64 {
    // e^x == 2^(x / ln(2))
    return Exp2Fast(Mul(x, RcpLn2))
}

func ExpFastest(x int64) int64 {
    // e^x == 2^(x / ln(2))
    return Exp2Fastest(Mul(x, RcpLn2))
}

// Log Natural logarithm (base e).
func Log(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    const ONE int32 = 1 << 30
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := int64(fixutil.LogPoly5Lut8(n-ONE) << 2)

    // Combine integer and fractional parts (into s32.32).
    return int64(offset)*RcpLog2E + y
}

// LogFast Natural logarithm (base e).
func LogFast(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    const ONE int32 = 1 << 30
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := int64(fixutil.LogPoly3Lut8(n-ONE) << 2)

    // Combine integer and fractional parts (into s32.32).
    return int64(offset)*RcpLog2E + y
}

// LogFastest Natural logarithm (base e).
func LogFastest(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    const ONE int32 = 1 << 30
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    y := int64(fixutil.LogPoly5(n-ONE) << 2)

    // Combine integer and fractional parts (into s32.32).
    return int64(offset)*RcpLog2E + y
}

func Log2(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2

    // Polynomial approximation of mantissa.
    const ONE int32 = 1 << 30
    y := int64(fixutil.Log2Poly4Lut16(n-ONE) << 2)

    // Combine integer and fractional parts (into s32.32).
    return (int64(offset) << Shift) + y
}

func Log2Fast(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2

    // Polynomial approximation of mantissa.
    const ONE int32 = 1 << 30
    y := int64(fixutil.Log2Poly3Lut16(n-ONE) << 2)

    // Combine integer and fractional parts (into s32.32).
    return (int64(offset) << Shift) + y
}

func Log2Fastest(x int64) int64 {
    if x <= 0 {
        return 0
    }

    // Normalize value to range [1.0, 2.0( as s2.30 and extract exponent.
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2

    // Polynomial approximation of mantissa.
    const ONE int32 = 1 << 30
    y := int64(fixutil.Log2Poly5(n-ONE) << 2)

    // Combine integer and fractional parts (into s32.32).
    return (int64(offset) << Shift) + y
}

// Pow Calculates x to the power of the exponent.
func Pow(x, exponent int64) int64 {
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
func PowFast(x, exponent int64) int64 {
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
func PowFastest(x, exponent int64) int64 {
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

func unitSin(z int32) int32 {
    // See: http://www.coranac.com/2009/07/sines/

    // Handle quadrants 1 and 2 by mirroring the [1, 3] range to [-1, 1] (by calculating 2 - z).
    // The if condition uses the fact that for the quadrants of interest are 0b01 and 0b10 (top two bits are different).
    if (z ^ (z << 1)) < 0 {
        z = int32((1 << 31) - int64(z))
    }

    // Polynomial approximation.
    zz := fixutil.Qmul30(z, z)
    res := fixutil.Qmul30(fixutil.SinPoly4(zz), z)

    // Return s2.30 value.
    return res
}

func unitSinFast(z int32) int32 {
    // See: http://www.coranac.com/2009/07/sines/

    // Handle quadrants 1 and 2 by mirroring the [1, 3] range to [-1, 1] (by calculating 2 - z).
    // The if condition uses the fact that for the quadrants of interest are 0b01 and 0b10 (top two bits are different).
    if (z ^ (z << 1)) < 0 {
        z = int32((1 << 31) - int64(z))
    }

    // Polynomial approximation.
    zz := fixutil.Qmul30(z, z)
    res := fixutil.Qmul30(fixutil.SinPoly3(zz), z)

    // Return s2.30 value.
    return res
}

func unitSinFastest(z int32) int32 {
    // See: http://www.coranac.com/2009/07/sines/

    // Handle quadrants 1 and 2 by mirroring the [1, 3] range to [-1, 1] (by calculating 2 - z).
    // The if condition uses the fact that for the quadrants of interest are 0b01 and 0b10 (top two bits are different).
    if (z ^ (z << 1)) < 0 {
        z = int32((1 << 31) - int64(z))
    }

    // Polynomial approximation.
    zz := fixutil.Qmul30(z, z)
    res := fixutil.Qmul30(fixutil.SinPoly2(zz), z)

    // Return s2.30 value.
    return res
}

func Sin(x int64) int64 {
    // Map [0, 2pi] to [0, 4] (as s2.30).
    // This also wraps the values into one period.
    z := MulIntLongLow(RcpHalfPi, x)

    // Compute sine and convert to s32.32.
    return int64(unitSin(z)) << 2
}

func SinFast(x int64) int64 {
    // Map [0, 2pi] to [0, 4] (as s2.30).
    // This also wraps the values into one period.
    z := MulIntLongLow(RcpHalfPi, x)

    // Compute sine and convert to s32.32.
    return int64(unitSinFast(z)) << 2
}

func SinFastest(x int64) int64 {
    // Map [0, 2pi] to [0, 4] (as s2.30).
    // This also wraps the values into one period.
    z := MulIntLongLow(RcpHalfPi, x)

    // Compute sine and convert to s32.32.
    return int64(unitSinFastest(z)) << 2
}

func Cos(x int64) int64 {
    return Sin(x + PiHalf)
}

func CosFast(x int64) int64 {
    return SinFast(x + PiHalf)
}

func CosFastest(x int64) int64 {
    return SinFastest(x + PiHalf)
}

func Tan(x int64) int64 {
    z := MulIntLongLow(RcpHalfPi, x)
    sinX := int64(unitSin(z)) << 32
    cosX := int64(unitSin(z+(1<<30))) << 32
    return Div(sinX, cosX)
}

func TanFast(x int64) int64 {
    z := MulIntLongLow(RcpHalfPi, x)
    sinX := int64(unitSinFast(z)) << 32
    cosX := int64(unitSinFast(z+(1<<30))) << 32
    return DivFast(sinX, cosX)
}

func TanFastest(x int64) int64 {
    z := MulIntLongLow(RcpHalfPi, x)
    sinX := int64(unitSinFastest(z)) << 32
    cosX := int64(unitSinFastest(z+(1<<30))) << 32
    return DivFastest(sinX, cosX)
}

func atan2Div(y, x int64) int32 {
    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    const ONE int32 = 1 << 30
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    k := n - ONE

    // Polynomial approximation of reciprocal.
    oox := fixutil.RcpPoly4Lut8(k)

    // Apply exponent and multiply.
    var yr int64
    if offset >= 0 {
        yr = y >> offset
    } else {
        yr = y << -offset
    }
    return fixutil.Qmul30(int32(yr>>2), oox)
}

func atan2DivFast(y, x int64) int32 {
    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    const ONE int32 = 1 << 30
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    k := n - ONE

    // Polynomial approximation of reciprocal.
    oox := fixutil.RcpPoly6(k)

    // Apply exponent and multiply.
    var yr int64
    if offset >= 0 {
        yr = y >> offset
    } else {
        yr = y << -offset
    }
    return fixutil.Qmul30(int32(yr>>2), oox)
}

func atan2DivFastest(y, x int64) int32 {
    // Normalize input into [1.0, 2.0( range (convert to s2.30).
    const ONE int32 = 1 << 30
    offset := 31 - nlz(uint64(x))
    var n int32
    if offset >= 0 {
        n = int32(x >> offset)
    } else {
        n = int32(x << -offset)
    }
    n >>= 2
    k := n - ONE

    // Polynomial approximation of reciprocal.
    oox := fixutil.RcpPoly4(k)

    // Apply exponent and multiply.
    var yr int64
    if offset >= 0 {
        yr = y >> offset
    } else {
        yr = y << -offset
    }
    return fixutil.Qmul30(int32(yr>>2), oox)
}

func Atan2(y, x int64) int64 {
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

    // \note these round negative numbers slightly
    nx := x ^ (x >> 63)
    ny := y ^ (y >> 63)
    negMask := (x ^ y) >> 63

    if nx >= ny {
        k := atan2Div(ny, nx)
        z := fixutil.AtanPoly5Lut8(k)
        angle := negMask ^ (int64(z) << 2)
        if x > 0 {
            return angle
        }
        if y >= 0 {
            return angle + Pi
        }
        return angle - Pi
    } else {
        k := atan2Div(nx, ny)
        z := fixutil.AtanPoly5Lut8(k)
        angle := negMask ^ (int64(z) << 2)
        if y > 0 {
            return PiHalf - angle
        } else {
            return -PiHalf - angle
        }
    }
}

func Atan2Fast(y, x int64) int64 {
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

    // \note these round negative numbers slightly
    nx := x ^ (x >> 63)
    ny := y ^ (y >> 63)
    negMask := (x ^ y) >> 63

    if nx >= ny {
        k := atan2DivFast(ny, nx)
        z := fixutil.AtanPoly3Lut8(k)
        angle := negMask ^ (int64(z) << 2)
        if x > 0 {
            return angle
        }
        if y >= 0 {
            return angle + Pi
        }
        return angle - Pi
    } else {
        k := atan2DivFast(nx, ny)
        z := fixutil.AtanPoly3Lut8(k)
        angle := negMask ^ (int64(z) << 2)
        if y > 0 {
            return PiHalf - angle
        } else {
            return -PiHalf - angle
        }
    }
}

func Atan2Fastest(y, x int64) int64 {
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

    // \note these round negative numbers slightly
    nx := x ^ (x >> 63)
    ny := y ^ (y >> 63)
    negMask := (x ^ y) >> 63

    if nx >= ny {
        k := atan2DivFastest(ny, nx)
        z := fixutil.AtanPoly4(k)
        angle := negMask ^ (int64(z) << 2)
        if x > 0 {
            return angle
        }
        if y >= 0 {
            return angle + Pi
        }
        return angle - Pi
    } else {
        k := atan2DivFastest(nx, ny)
        z := fixutil.AtanPoly4(k)
        angle := negMask ^ (int64(z) << 2)
        if y > 0 {
            return PiHalf - angle
        } else {
            return -PiHalf - angle
        }
    }
}

func Asin(x int64) int64 {
    if x < -One || x > One {
        return 0
    }
    return Atan2(x, Sqrt(Mul(One+x, One-x)))
}

func AsinFast(x int64) int64 {
    if x < -One || x > One {
        return 0
    }
    return Atan2Fast(x, SqrtFast(Mul(One+x, One-x)))
}

func AsinFastest(x int64) int64 {
    if x < -One || x > One {
        return 0
    }
    return Atan2Fastest(x, SqrtFastest(Mul(One+x, One-x)))
}

func Acos(x int64) int64 {
    if x < -One || x > One {
        return 0
    }
    return Atan2(Sqrt(Mul(One+x, One-x)), x)
}

func AcosFast(x int64) int64 {
    if x < -One || x > One {
        return 0
    }
    return Atan2Fast(SqrtFast(Mul(One+x, One-x)), x)
}

func AcosFastest(x int64) int64 {
    if x < -One || x > One {
        return 0
    }
    return Atan2Fastest(SqrtFastest(Mul(One+x, One-x)), x)
}

func Atan(x int64) int64 {
    return Atan2(x, One)
}

func AtanFast(x int64) int64 {
    return Atan2Fast(x, One)
}

func AtanFastest(x int64) int64 {
    return Atan2Fastest(x, One)
}
