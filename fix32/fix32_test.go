package fix32_test

import (
    "testing"

    "github.com/camry/fp/fix32"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    f1 := fix32.FromInt32(32767)
    f2 := fix32.FromInt32(-32768)
    f3 := fix32.FromInt32(32767)
    f4 := fix32.FromFloat32(134.8765346)
    assert.Equal(t, fix32.RoundToInt(fix32.Add(f1, f2)), int32(-1))
    assert.Equal(t, fix32.RoundToInt(fix32.Add(fix32.Add(f1, f2), f3)), int32(32766))
    assert.Equal(t, fix32.ToFloat32(fix32.Add(fix32.Add(f2, f3), f4)), float32(133.87654))
    assert.Equal(t, fix32.ToFloat64(fix32.Add(fix32.Add(f2, f3), f4)), 133.8765411376953)
}

func TestSub(t *testing.T) {
    f1 := fix32.FromInt32(32767)
    f2 := fix32.FromInt32(32721)
    f3 := fix32.FromInt32(32767)
    f4 := fix32.FromFloat64(32721.4523)
    assert.Equal(t, fix32.RoundToInt(fix32.Sub(f1, f2)), int32(46))
    assert.Equal(t, fix32.ToFloat32(fix32.Sub(f3, f4)), float32(45.547714))
}

func TestMul(t *testing.T) {
    f1 := fix32.Mul(fix32.FromInt32(32767), fix32.FromFloat32(0.30))
    assert.Equal(t, fix32.CeilToInt(f1), int32(9830))
    assert.Equal(t, fix32.FloorToInt(f1), int32(9829))
    assert.Equal(t, fix32.RoundToInt(f1), int32(9830))
}

func TestDiv(t *testing.T) {
    f1 := fix32.FromInt32(32767)
    f2 := fix32.FromFloat32(32157)
    assert.Equal(t, fix32.ToFloat32(fix32.Div(f1, f2)), float32(1.0189667))
}

func TestFix64(t *testing.T) {
    f1 := fix32.FromInt32(300)
    f2 := fix32.FromInt32(10)
    f3 := fix32.FromInt32(5)
    f4 := fix32.FromInt32(1)
    f5 := fix32.FromFloat32(0)
    f6 := fix32.Div(fix32.Mul(f1, f1), fix32.Add(f1, fix32.Mul(f2, fix32.Mul(f3, fix32.Sub(f4, f5)))))
    assert.Equal(t, fix32.ToFloat32(f6), float32(69.89714))
    assert.Equal(t, fix32.ToFloat64(f6), 69.89714050292969)
}
