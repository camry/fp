package f64_test

import (
    "testing"

    "github.com/camry/fp/fixmath/f64"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    f1 := f64.FromInt32(2147483647)
    f2 := f64.FromInt32(-2147483647)
    f3 := f64.FromInt64(2147483647)
    f4 := f64.FromFloat32(134.8765346)
    assert.Equal(t, f1.Add(f2).RoundToInt(), int32(0))
    assert.Equal(t, f1.Add(f2).Add(f3).RoundToInt(), int32(2147483647))
    assert.Equal(t, f2.Add(f3).Add(f4).Float32(), float32(134.87654))
    assert.Equal(t, f2.Add(f3).Add(f4).Float64(), 134.8765411376953)
}

func TestSub(t *testing.T) {
    f1 := f64.FromInt32(2147483647)
    f2 := f64.FromInt32(2147483123)
    f3 := f64.FromInt32(2147483647)
    f4 := f64.FromFloat64(2147483123.4523)
    assert.Equal(t, f1.Sub(f2).RoundToInt(), int32(524))
    assert.Equal(t, f3.Sub(f4).Float32(), float32(523.5477))
}

func TestMul(t *testing.T) {
    f1 := f64.FromInt32(2147483647).Mul(f64.FromFloat32(0.30))
    assert.Equal(t, f1.CeilToInt(), int32(644245120))
    assert.Equal(t, f1.FloorToInt(), int32(644245119))
    assert.Equal(t, f1.RoundToInt(), int32(644245120))
}

func TestDiv(t *testing.T) {
    f1 := f64.FromInt32(2147483647).Div(f64.FromFloat32(4567822))
    assert.Equal(t, f1.Float32(), float32(470.13297))
}

func TestFix64(t *testing.T) {
    f1 := f64.FromInt32(300).Mul(f64.FromInt32(300)).DivPrecise(f64.FromInt32(300).Add(f64.FromInt32(10).Mul(f64.FromInt32(5).Mul(f64.FromInt32(1).Sub(f64.FromFloat32(0))))))
    assert.Equal(t, f1.Float32(), float32(257.14285))
    assert.Equal(t, f1.Float64(), 257.1428571427241)
}
