package fixmath

import (
    "fmt"
    "reflect"

    "github.com/camry/fp/fix64"
)

var Identity = QuatFromRaw(fix64.Zero, fix64.Zero, fix64.Zero, fix64.One)

type F64Quat struct {
    RawX int64
    RawY int64
    RawZ int64
    RawW int64
}

func QuatFromRaw(x, y, z, w int64) F64Quat {
    return F64Quat{
        RawX: x,
        RawY: y,
        RawZ: z,
        RawW: w,
    }
}

func FromF64(x, y, z, w F64) F64Quat {
    return F64Quat{
        RawX: x.Raw,
        RawY: y.Raw,
        RawZ: z.Raw,
        RawW: w.Raw,
    }
}

func FromVector(v F64Vec3, w F64) F64Quat {
    return F64Quat{
        RawX: v.RawX,
        RawY: v.RawY,
        RawZ: v.RawZ,
        RawW: w.Raw,
    }
}

func FromAxisAngle(axis F64Vec3, angle F64) F64Quat {
    halfAngle := angle.Div2()
    halfAngleSinFastest := halfAngle.SinFastest()
    return FromVector(axis.Mul(Vec3FromF64(halfAngleSinFastest, halfAngleSinFastest, halfAngleSinFastest)), halfAngle.CosFastest())
}

func FromYawPitchRoll(yawY, pitchX, rollZ F64) F64Quat {
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

func FromTwoVectors(a, b F64Vec3) F64Quat {
    // From: http://lolengine.net/blog/2014/02/24/quaternion-from-two-vectors-final
    epsilon := F64Ratio(1, 1000000)

    normANormB := (a.LengthSqr().Mul(b.LengthSqr())).SqrtFastest()
    realPart := normANormB.Add(a.Dot(b))

    var v F64Vec3

    if realPart.LT(epsilon.Mul(normANormB)) {
        /* If u and v are exactly opposite, rotate 180 degrees
         * around an arbitrary orthogonal axis. Axis normalization
         * can happen later, when we normalize the quaternion. */
        realPart = F64Zero
        cond := a.X().Abs().GT(a.Z().Abs())
        if cond {
            v = Vec3FromF64(a.Y().Negate(), a.X(), F64Zero)
        } else {
            v = Vec3FromF64(F64Zero, a.Z().Negate(), a.Y())
        }
    } else {
        /* Otherwise, build quaternion the standard way. */
        v = a.Cross(b)
    }

    return FromVector(v, realPart).NormalizeFastest()
}

func LookRotation(dir, up F64Vec3) F64Quat {
    // From: https://answers.unity.com/questions/819699/calculate-quaternionlookrotation-manually.html
    if dir == Vec3Zero {
        return Identity
    }

    if up != dir {
        up = up.NormalizeFastest()
        v := dir.Add(up).MulF64(up.Dot(dir).Negate())
        q := FromTwoVectors(Vec3AxisZ, v)
        return FromTwoVectors(v, dir).Mul(q)
    } else {

        return FromTwoVectors(Vec3AxisZ, dir)
    }
}

func LookAtRotation(from, to, up F64Vec3) F64Quat {
    dir := (to.Sub(from)).NormalizeFastest()
    return LookRotation(dir, up)
}

func (q F64Quat) QuatX() F64 {
    return F64FromRaw(q.RawX)
}

func (q F64Quat) QuatY() F64 {
    return F64FromRaw(q.RawY)
}

func (q F64Quat) QuatZ() F64 {
    return F64FromRaw(q.RawZ)
}

func (q F64Quat) QuatW() F64 {
    return F64FromRaw(q.RawW)
}

func (q F64Quat) Mul(b F64Quat) F64Quat {
    return q.Multiply(b)
}

// EQ q == b
func (q F64Quat) EQ(b F64Quat) bool {
    return q.RawX == b.RawX && q.RawY == b.RawY && q.RawZ == b.RawZ && q.RawW == b.RawW
}

// NE q != b
func (q F64Quat) NE(b F64Quat) bool {
    return q.RawX != b.RawX || q.RawY != b.RawY || q.RawZ != b.RawZ || q.RawW != b.RawW
}

// Negate -q.RawX, -q.RawY, -q.RawZ, -q.RawW
func (q F64Quat) Negate() F64Quat {
    return QuatFromRaw(-q.RawX, -q.RawY, -q.RawZ, -q.RawW)
}

// Conjugate -q.RawX, -q.RawY, -q.RawZ, q.RawW
func (q F64Quat) Conjugate() F64Quat {
    return QuatFromRaw(-q.RawX, -q.RawY, -q.RawZ, q.RawW)
}

func (q F64Quat) Inverse() F64Quat {
    invNorm := q.LengthSqr().Rcp().Raw
    return QuatFromRaw(
        -fix64.Mul(q.RawX, invNorm),
        -fix64.Mul(q.RawY, invNorm),
        -fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q F64Quat) InverseUnit() F64Quat {
    return QuatFromRaw(-q.RawX, -q.RawY, -q.RawZ, q.RawW)
}

func (q F64Quat) Multiply(b F64Quat) F64Quat {
    q1x := q.QuatX()
    q1y := q.QuatY()
    q1z := q.QuatZ()
    q1w := q.QuatW()

    q2x := b.QuatX()
    q2y := b.QuatY()
    q2z := b.QuatZ()
    q2w := b.QuatW()

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

func (q F64Quat) Length() F64 {
    return q.LengthSqr().Sqrt()
}

func (q F64Quat) LengthFast() F64 {
    return q.LengthSqr().SqrtFast()
}

func (q F64Quat) LengthFastest() F64 {
    return q.LengthSqr().SqrtFastest()
}

func (q F64Quat) LengthSqr() F64 {
    return F64FromRaw(fix64.Mul(q.RawX, q.RawX) + fix64.Mul(q.RawY, q.RawY) + fix64.Mul(q.RawZ, q.RawZ) + fix64.Mul(q.RawW, q.RawW))
}

func (q F64Quat) Normalize() F64Quat {
    invNorm := q.Length().Rcp().Raw
    return QuatFromRaw(
        fix64.Mul(q.RawX, invNorm),
        fix64.Mul(q.RawY, invNorm),
        fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q F64Quat) NormalizeFast() F64Quat {
    invNorm := q.LengthFast().RcpFast().Raw
    return QuatFromRaw(
        fix64.Mul(q.RawX, invNorm),
        fix64.Mul(q.RawY, invNorm),
        fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q F64Quat) NormalizeFastest() F64Quat {
    invNorm := q.LengthFastest().RcpFastest().Raw
    return QuatFromRaw(
        fix64.Mul(q.RawX, invNorm),
        fix64.Mul(q.RawY, invNorm),
        fix64.Mul(q.RawZ, invNorm),
        fix64.Mul(q.RawW, invNorm),
    )
}

func (q F64Quat) Slerp(q2 F64Quat, t F64) F64Quat {
    epsilon := F64Ratio(1, 1000000)
    cosOmega := q.QuatX().Mul(q2.QuatX()).Add(q.QuatY().Mul(q2.QuatY())).Add(q.QuatZ().Mul(q2.QuatZ())).Add(q.QuatW().Mul(q2.QuatW()))

    flip := false

    if cosOmega.LT(F64FromInt32(0)) {
        flip = true
        cosOmega = cosOmega.Negate()
    }

    var s1, s2 F64
    if cosOmega.GT(F64One.Sub(epsilon)) {
        // Too close, do straight linear interpolation.
        s1 = F64One.Sub(t)
        if flip {
            s2 = t.Negate()
        } else {
            s2 = t
        }
    } else {
        omega := cosOmega.AcosFastest()
        invSinOmega := omega.SinFastest().RcpFastest()

        s1 = F64One.Sub(t).Mul(omega).SinFastest().Mul(invSinOmega)
        if flip {
            s2 = t.Mul(omega).SinFastest().Negate().Mul(invSinOmega)
        } else {
            s2 = t.Mul(omega).SinFastest().Mul(invSinOmega)
        }
    }

    return FromF64(
        s1.Mul(q.QuatX()).Add(s2.Mul(q2.QuatX())),
        s1.Mul(q.QuatY()).Add(s2.Mul(q2.QuatY())),
        s1.Mul(q.QuatZ()).Add(s2.Mul(q2.QuatZ())),
        s1.Mul(q.QuatW()).Add(s2.Mul(q2.QuatW())),
    )
}

func (q F64Quat) Lerp(q2 F64Quat, t F64) F64Quat {
    t1 := F64One.Sub(t)
    dot := q.QuatX().Mul(q2.QuatX()).Add(q.QuatY().Mul(q2.QuatY())).Add(q.QuatZ().Mul(q2.QuatZ())).Add(q.QuatW().Mul(q2.QuatW()))

    var r F64Quat
    if dot.GE(F64FromInt32(0)) {
        r = FromF64(
            t1.Mul(q.QuatX()).Add(t.Mul(q2.QuatX())),
            t1.Mul(q.QuatY()).Add(t.Mul(q2.QuatY())),
            t1.Mul(q.QuatZ()).Add(t.Mul(q2.QuatZ())),
            t1.Mul(q.QuatW()).Add(t.Mul(q2.QuatW())),
        )
    } else {
        r = FromF64(
            t1.Mul(q.QuatX()).Sub(t.Mul(q2.QuatX())),
            t1.Mul(q.QuatY()).Sub(t.Mul(q2.QuatY())),
            t1.Mul(q.QuatZ()).Sub(t.Mul(q2.QuatZ())),
            t1.Mul(q.QuatW()).Sub(t.Mul(q2.QuatW())),
        )
    }

    return r.NormalizeFastest()
}

// Concatenate two Quaternions; the result represents the value1 rotation followed by the value2 rotation.
func (q F64Quat) Concatenate(q2 F64Quat) F64Quat {
    return q.Mul(q2)
}

// RotateVector Rotates a vector by the unit quaternion.
func (q F64Quat) RotateVector(v F64Vec3) F64Vec3 {
    // From https://gamedev.stackexchange.com/questions/28395/rotating-vector3-by-a-quaternion
    u := Vec3FromF64(q.QuatX(), q.QuatY(), q.QuatZ())
    s := q.QuatW()

    return u.MulF64(F64Two.Mul(u.Dot(v))).Add(v.MulF64(s.Mul(s).Sub(u.Dot(u)))).Add(F64Two.Mul(s).MulVec3(u.Cross(v)))
}

func (q F64Quat) Equals(obj F64Quat) bool {
    return reflect.DeepEqual(q, obj)
}

func (q F64Quat) ToString() string {
    return fmt.Sprintf(`(%s, %s, %s, %s)`, fix64.ToString(q.RawX), fix64.ToString(q.RawY), fix64.ToString(q.RawZ), fix64.ToString(q.RawW))
}
