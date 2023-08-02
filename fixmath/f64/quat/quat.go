package quat

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix64"
    "github.com/camry/fp/fixmath/f64"
    "github.com/camry/fp/fixmath/f64/vec3"
)

var Identity = FromRaw(fix64.Zero, fix64.Zero, fix64.Zero, fix64.One)

type Quat struct {
    RawX int64
    RawY int64
    RawZ int64
    RawW int64
}

func FromRaw(x, y, z, w int64) Quat {
    return Quat{
        RawX: x,
        RawY: y,
        RawZ: z,
        RawW: w,
    }
}

func FromF64(x, y, z, w f64.F64) Quat {
    return Quat{
        RawX: x.Raw,
        RawY: y.Raw,
        RawZ: z.Raw,
        RawW: w.Raw,
    }
}

func FromVector(v vec3.Vec3, w f64.F64) Quat {
    return Quat{
        RawX: v.RawX,
        RawY: v.RawY,
        RawZ: v.RawZ,
        RawW: w.Raw,
    }
}

func FromAxisAngle(axis vec3.Vec3, angle f64.F64) Quat {
    halfAngle := angle.Div2()
    halfAngleSinFastest := halfAngle.SinFastest()
    return FromVector(axis.Mul(vec3.FromF64(halfAngleSinFastest, halfAngleSinFastest, halfAngleSinFastest)), halfAngle.CosFastest())
}

func FromYawPitchRoll(yawY, pitchX, rollZ f64.F64) Quat {
    //  Roll first, about axis the object is facing, then
    //  pitch upward, then yaw to face into the new heading
    halfRoll := rollZ.Div2()
    sr := halfRoll.SinFastest()
    cr := halfRoll.CosFastest()

    halfPitch := pitchX.Div2()
    sp := halfPitch.SinFastest()
    cp := halfPitch.CosFastest()

    halfYaw := yawY.Div2()
    sy := halfYaw.SinFastest()
    cy := halfYaw.CosFastest()

    return FromF64(
        cy.Mul(sp).Mul(cr).Add(sy.Mul(cp).Mul(sr)),
        sy.Mul(cp).Mul(cr).Sub(cy.Mul(sp).Mul(sr)),
        cy.Mul(cp).Mul(sr).Sub(sy.Mul(sp).Mul(cr)),
        cy.Mul(cp).Mul(cr).Add(sy.Mul(sp).Mul(sr)),
    )
}

func FromTwoVectors(a, b vec3.Vec3) Quat {
    // From: http://lolengine.net/blog/2014/02/24/quaternion-from-two-vectors-final
    epsilon := f64.Ratio(1, 1000000)

    normANormB := (a.LengthSqr().Mul(b.LengthSqr())).SqrtFastest()
    realPart := normANormB.Add(a.Dot(b))

    var v vec3.Vec3

    if realPart.LT(epsilon.Mul(normANormB)) {
        /* If u and v are exactly opposite, rotate 180 degrees
         * around an arbitrary orthogonal axis. Axis normalization
         * can happen later, when we normalize the quaternion. */
        realPart = f64.Zero
        cond := a.X().Abs().GT(a.Z().Abs())
        if cond {
            v = vec3.FromF64(a.Y().Negate(), a.X(), f64.Zero)
        } else {
            v = vec3.FromF64(f64.Zero, a.Z().Negate(), a.Y())
        }
    } else {
        /* Otherwise, build quaternion the standard way. */
        v = a.Cross(b)
    }

    return FromVector(v, realPart).NormalizeFastest()
}

func LookRotation(dir, up vec3.Vec3) Quat {
    // From: https://answers.unity.com/questions/819699/calculate-quaternionlookrotation-manually.html
    if dir == vec3.Zero {
        return Identity
    }

    if up != dir {
        up = up.NormalizeFastest()
        v := dir.Add(up).MulF64(up.Dot(dir).Negate())
        q := FromTwoVectors(vec3.AxisZ, v)
        return FromTwoVectors(v, dir).Mul(q)
    } else {

        return FromTwoVectors(vec3.AxisZ, dir)
    }
}

func LookAtRotation(from, to, up vec3.Vec3) Quat {
    dir := (to.Sub(from)).NormalizeFastest()
    return LookRotation(dir, up)
}

func (q Quat) X() f64.F64 {
    return f64.FromRaw(q.RawX)
}

func (q Quat) Y() f64.F64 {
    return f64.FromRaw(q.RawY)
}

func (q Quat) Z() f64.F64 {
    return f64.FromRaw(q.RawZ)
}

func (q Quat) W() f64.F64 {
    return f64.FromRaw(q.RawW)
}

func (q Quat) Mul(b Quat) Quat {
    return q.Multiply(b)
}

// EQ q == b
func (q Quat) EQ(b Quat) bool {
    return q.RawX == b.RawX && q.RawY == b.RawY && q.RawZ == b.RawZ && q.RawW == b.RawW
}

// NE q != b
func (q Quat) NE(b Quat) bool {
    return q.RawX != b.RawX || q.RawY != b.RawY || q.RawZ != b.RawZ || q.RawW != b.RawW
}

// Negate -q.RawX, -q.RawY, -q.RawZ, -q.RawW
func (q Quat) Negate() Quat {
    return FromRaw(-q.RawX, -q.RawY, -q.RawZ, -q.RawW)
}

// Conjugate -q.RawX, -q.RawY, -q.RawZ, q.RawW
func (q Quat) Conjugate() Quat {
    return FromRaw(-q.RawX, -q.RawY, -q.RawZ, q.RawW)
}

func (q Quat) Inverse() Quat {
    invNorm := q.LengthSqr().Rcp().Raw
    return FromRaw(
        -fix64.Mul(q.RawX, invNorm),
        -fix64.Mul(q.RawY, invNorm),
        -fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q Quat) InverseUnit() Quat {
    return FromRaw(-q.RawX, -q.RawY, -q.RawZ, q.RawW)
}

func (q Quat) Multiply(b Quat) Quat {
    q1x := q.X()
    q1y := q.Y()
    q1z := q.Z()
    q1w := q.W()

    q2x := b.X()
    q2y := b.Y()
    q2z := b.Z()
    q2w := b.W()

    // cross(av, bv)
    cx := q1y.Mul(q2z).Sub(q1z.Mul(q2y))
    cy := q1z.Mul(q2x).Sub(q1x.Mul(q2z))
    cz := q1x.Mul(q2y).Sub(q1y.Mul(q2x))

    dot := q1x.Mul(q2x).Add(q1y.Mul(q2y)).Add(q1z.Mul(q2z))

    return FromF64(
        q1x.Mul(q2w).Add(q2x.Mul(q1w)).Add(cx),
        q1y.Mul(q2w).Add(q2y.Mul(q1w)).Add(cy),
        q1z.Mul(q2w).Add(q2z.Mul(q1w)).Add(cz),
        q1w.Mul(q2w).Sub(dot),
    )
}

func (q Quat) Length() f64.F64 {
    return q.LengthSqr().Sqrt()
}

func (q Quat) LengthFast() f64.F64 {
    return q.LengthSqr().SqrtFast()
}

func (q Quat) LengthFastest() f64.F64 {
    return q.LengthSqr().SqrtFastest()
}

func (q Quat) LengthSqr() f64.F64 {
    return f64.FromRaw(fix64.Mul(q.RawX, q.RawX) + fix64.Mul(q.RawY, q.RawY) + fix64.Mul(q.RawZ, q.RawZ) + fix64.Mul(q.RawW, q.RawW))
}

func (q Quat) Normalize() Quat {
    invNorm := q.Length().Rcp().Raw
    return FromRaw(
        fix64.Mul(q.RawX, invNorm),
        fix64.Mul(q.RawY, invNorm),
        fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q Quat) NormalizeFast() Quat {
    invNorm := q.LengthFast().RcpFast().Raw
    return FromRaw(
        fix64.Mul(q.RawX, invNorm),
        fix64.Mul(q.RawY, invNorm),
        fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q Quat) NormalizeFastest() Quat {
    invNorm := q.LengthFastest().RcpFastest().Raw
    return FromRaw(
        fix64.Mul(q.RawX, invNorm),
        fix64.Mul(q.RawY, invNorm),
        fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q Quat) Slerp(q2 Quat, t f64.F64) Quat {
    epsilon := f64.Ratio(1, 1000000)
    cosOmega := q.X().Mul(q2.X()).Add(q.Y().Mul(q2.Y())).Add(q.Z().Mul(q2.Z())).Add(q.W().Mul(q2.W()))

    flip := false

    if cosOmega.LT(f64.FromInt32(0)) {
        flip = true
        cosOmega = cosOmega.Negate()
    }

    var s1, s2 f64.F64
    if cosOmega.GT(f64.One.Sub(epsilon)) {
        // Too close, do straight linear interpolation.
        s1 = f64.One.Sub(t)
        if flip {
            s2 = t.Negate()
        } else {
            s2 = t
        }
    } else {
        omega := cosOmega.AcosFastest()
        invSinOmega := omega.SinFastest().RcpFastest()

        s1 = f64.One.Sub(t).Mul(omega).SinFastest().Mul(invSinOmega)
        if flip {
            s2 = t.Mul(omega).SinFastest().Negate().Mul(invSinOmega)
        } else {
            s2 = t.Mul(omega).SinFastest().Mul(invSinOmega)
        }
    }

    return FromF64(
        s1.Mul(q.X()).Add(s2.Mul(q2.X())),
        s1.Mul(q.Y()).Add(s2.Mul(q2.Y())),
        s1.Mul(q.Z()).Add(s2.Mul(q2.Z())),
        s1.Mul(q.W()).Add(s2.Mul(q2.W())),
    )
}

func (q Quat) Lerp(q2 Quat, t f64.F64) Quat {
    t1 := f64.One.Sub(t)
    dot := q.X().Mul(q2.X()).Add(q.Y().Mul(q2.Y())).Add(q.Z().Mul(q2.Z())).Add(q.W().Mul(q2.W()))

    var r Quat
    if dot.GE(f64.FromInt32(0)) {
        r = FromF64(
            t1.Mul(q.X()).Add(t.Mul(q2.X())),
            t1.Mul(q.Y()).Add(t.Mul(q2.Y())),
            t1.Mul(q.Z()).Add(t.Mul(q2.Z())),
            t1.Mul(q.W()).Add(t.Mul(q2.W())),
        )
    } else {
        r = FromF64(
            t1.Mul(q.X()).Sub(t.Mul(q2.X())),
            t1.Mul(q.Y()).Sub(t.Mul(q2.Y())),
            t1.Mul(q.Z()).Sub(t.Mul(q2.Z())),
            t1.Mul(q.W()).Sub(t.Mul(q2.W())),
        )
    }

    return r.NormalizeFastest()
}

// Concatenate two Quaternions; the result represents the value1 rotation followed by the value2 rotation.
func (q Quat) Concatenate(q2 Quat) Quat {
    return q.Mul(q2)
}

// RotateVector Rotates a vector by the unit quaternion.
func (q Quat) RotateVector(v vec3.Vec3) vec3.Vec3 {
    // From https://gamedev.stackexchange.com/questions/28395/rotating-vector3-by-a-quaternion
    u := vec3.FromF64(q.X(), q.Y(), q.Z())
    s := q.W()

    return u.MulF64(f64.Two.Mul(u.Dot(v))).Add(v.MulF64(s.Mul(s).Sub(u.Dot(u)))).Add(f64.Two.Mul(s).MulVec3(u.Cross(v)))
}

func (q Quat) Equals(obj Quat) bool {
    return reflect.DeepEqual(q, obj)
}

func (q Quat) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s, %s)`, fix64.ToString(q.RawX), fix64.ToString(q.RawY), fix64.ToString(q.RawZ), fix64.ToString(q.RawW))
}
