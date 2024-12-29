package geometry

import (
    "math"
    "testing"
)

func TestVec4Add(t *testing.T) {
    v1 := Vec4{1, 2, 3, 4}
    v2 := Vec4{5, 6, 7, 8}

    expected := Vec4{6, 8, 10, 12}
    v1.Add(v2)

    if v1 != expected {
        t.Errorf("Add(%v, %v) = %v; want %v", v1, v2, v1, expected)
    }
}

func TestVec4Subtract(t *testing.T) {
    v1 := Vec4{5, 6, 7, 8}
    v2 := Vec4{1, 2, 3, 4}

    expected := Vec4{4, 4, 4, 4}
    v1.Sub(v2)

    if v1 != expected {
        t.Errorf("Sub(%v, %v) = %v; want %v", v1, v2, v1, expected)
    }
}

func TestVec4Dot(t *testing.T) {
    v1 := Vec4{1, 2, 3, 4}
    v2 := Vec4{5, 6, 7, 8}

    expected := float64(1*5 + 2*6 + 3*7 + 4*8)
    result := v1.Dot(v2)

    if result != expected {
        t.Errorf("Dot(%v, %v) = %f; want %f", v1, v2, result, expected)
    }
}

func TestVec4Magnitude(t *testing.T) {
    v := Vec4{1, 2, 3, 4}

    expected := math.Sqrt(1*1 + 2*2 + 3*3 + 4*4)
    result := v.Length()

    if result != expected {
        t.Errorf("Magnitude of %v failed: got %f, want %f", v, result, expected)
    }
}

func TestVec4Normalize(t *testing.T) {
    v := Vec4{1, 2, 3, 4}
    v.Normalize()

    length := v.Length()
    expected := float64(1)

    if length != expected {
        t.Errorf("Normalize(%v) failed: got length %f, want %f", v, length, expected)
    }
}