package fix64_test

import (
    "testing"

    "github.com/camry/fp/fix64"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    f1 := fix64.FromInt32(2147483647)
    f2 := fix64.FromInt32(-2147483647)
    f3 := fix64.FromInt64(2147483647)
    f4 := fix64.FromFloat32(134.8765346)
    assert.Equal(t, fix64.RoundToInt(fix64.Add(f1, f2)), int32(0))
    assert.Equal(t, fix64.RoundToInt(fix64.Add(fix64.Add(f1, f2), f3)), int32(2147483647))
    assert.Equal(t, fix64.ToFloat32(fix64.Add(fix64.Add(f2, f3), f4)), float32(134.87654))
    assert.Equal(t, fix64.ToFloat64(fix64.Add(fix64.Add(f2, f3), f4)), 134.8765411376953)
}

func TestSub(t *testing.T) {
    f1 := fix64.FromInt32(2147483647)
    f2 := fix64.FromInt32(2147483123)
    f3 := fix64.FromInt32(2147483647)
    f4 := fix64.FromFloat64(2147483123.4523)
    assert.Equal(t, fix64.RoundToInt(fix64.Sub(f1, f2)), int32(524))
    assert.Equal(t, fix64.ToFloat32(fix64.Sub(f3, f4)), float32(523.5477))
}

func TestMul(t *testing.T) {
    f1 := fix64.Mul(fix64.FromInt32(2147483647), fix64.FromFloat32(0.30))
    assert.Equal(t, fix64.CeilToInt(f1), int32(644245120))
    assert.Equal(t, fix64.FloorToInt(f1), int32(644245119))
    assert.Equal(t, fix64.RoundToInt(f1), int32(644245120))
}

func TestDiv(t *testing.T) {
    f1 := fix64.FromInt32(2147483647)
    f2 := fix64.FromFloat32(4567822)
    assert.Equal(t, fix64.ToFloat32(fix64.Div(f1, f2)), float32(470.13297))
}

func TestFix64(t *testing.T) {
    f1 := fix64.FromInt32(300)
    f2 := fix64.FromInt32(10)
    f3 := fix64.FromInt32(5)
    f4 := fix64.FromInt32(1)
    f5 := fix64.FromFloat32(0)
    f6 := fix64.Div(fix64.Mul(f1, f1), fix64.Add(f1, fix64.Mul(f2, fix64.Mul(f3, fix64.Sub(f4, f5)))))
    assert.Equal(t, fix64.ToFloat32(f6), float32(257.14285))
    assert.Equal(t, fix64.ToFloat64(f6), 257.1428517694585)
}
