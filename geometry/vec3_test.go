package geometry

import (
    "math"
    "testing"
)

func TestVec3Add(t *testing.T) {
    v1 := Vec3{1, 2, 3}
    v2 := Vec3{4, 5, 6}

    expected := Vec3{5, 7, 9}
    v1.Add(v2)

    if v1 != expected {
        t.Errorf("Add(%v, %v) = %v; want %v", v1, v2, v1, expected)
    }
}

func TestVec3Sub(t *testing.T) {
    v1 := Vec3{4, 5, 6}
    v2 := Vec3{1, 2, 3}

    expected := Vec3{3, 3, 3}
    v1.Sub(v2)

    if v1 != expected {
        t.Errorf("Sub(%v, %v) = %v; want %v", v1, v2, v1, expected)
    }
}

func TestVec3Dot(t *testing.T) {
    v1 := Vec3{1, 2, 3}
    v2 := Vec3{4, 5, 6}

    expected := float64(1*4 + 2*5 + 3*6)
    result := v1.Dot(v2)

    if result != expected {
        t.Errorf("Dot(%v, %v) = %f; want %f", v1, v2, result, expected)
    }
}

func TestVec3ZeroVector(t *testing.T) {
    v := Vec3{0, 0, 0}

    v.Add(v)
    if v != (Vec3{0, 0, 0}) {
        t.Errorf("Add zero vector failed: got %v, want %v", v, Vec3{0, 0, 0})
    }

    v.Sub(v)
    if v != (Vec3{0, 0, 0}) {
        t.Errorf("Sub zero vector failed: got %v, want %v", v, Vec3{0, 0, 0})
    }

    dotResult := v.Dot(v)
    if dotResult != 0 {
        t.Errorf("Dot zero vector failed: got %f, want %f", dotResult, 0.0)
    }
}

func TestVec3Magnitude(t *testing.T) {
    v := Vec3{3, 4, 0}

    expected := float64(5) // 3-4-5 triangle
    result := math.Sqrt(v.Dot(v))

    if result != expected {
        t.Errorf("Magnitude of %v failed: got %f, want %f", v, result, expected)
    }
}