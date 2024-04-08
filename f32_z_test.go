package fp_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/fp"
)

func TestF32_Pow(t *testing.T) {
    f1 := fp.F32FromInt32(1).Add(fp.F32FromFloat32(0.08)).Pow(fp.F32FromInt32(3))
    assert.Equal(t, f1.Float32(), float32(1.2595978))
}
