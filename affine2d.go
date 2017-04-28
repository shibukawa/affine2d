// +build !js

// affine2d provides affine transformation for 2D graphics.
// This code is comes from nanovgo and added gopher.js optimization.
// https://github.com/shibukawa/nanovgo
package affine2d

import (
	"github.com/rkusa/gm/math32"
)

// Scala is a type of element of vector and matrix.
// Scala is a float32 on regular environment, and float64 for Gopher.js
// (https://github.com/gopherjs/gopherjs#performance-tips).
type Scala float32

const (
	// Pi is a constant of Pi value in Scala type.
	Pi Scala = Scala(math32.Pi)
)

// The following functions can be used to make calculations on 2x3 transformation matrices.

// TransformMatrix is a 2x3 matrix is represented as float[6].
type Matrix [6]Scala

// IdentityMatrix makes the transform to identity matrix.
func IdentityMatrix() Matrix {
	return Matrix{1.0, 0.0, 0.0, 1.0, 0.0, 0.0}
}

// TranslateMatrix makes the transform to translation matrix matrix.
func TranslateMatrix(tx, ty Scala) Matrix {
	return Matrix{1.0, 0.0, 0.0, 1.0, tx, ty}
}

// ScaleMatrix makes the transform to scale matrix.
func ScaleMatrix(sx, sy Scala) Matrix {
	return Matrix{sx, 0.0, 0.0, sy, 0.0, 0.0}
}

// RotateMatrix makes the transform to rotate matrix. Angle is specified in radians.
func RotateMatrix(a Scala) Matrix {
	sin, cos := math32.Sincos(float32(a))
	return Matrix{Scala(cos), Scala(sin), Scala(-sin), Scala(cos), 0.0, 0.0}
}

// SkewXMatrix makes the transform to skew-x matrix. Angle is specified in radians.
func SkewXMatrix(a Scala) Matrix {
	return Matrix{1.0, 0.0, Scala(math32.Tan(float32(a))), 1.0, 0.0, 0.0}
}

// SkewYMatrix makes the transform to skew-y matrix. Angle is specified in radians.
func SkewYMatrix(a Scala) Matrix {
	return Matrix{1.0, Scala(math32.Tan(float32(a))), 0.0, 1.0, 0.0, 0.0}
}

// Multiply makes the transform to the result of multiplication of two transforms, of A = A*B.
func (t Matrix) Multiply(s Matrix) Matrix {
	t0 := t[0]*s[0] + t[1]*s[2]
	t2 := t[2]*s[0] + t[3]*s[2]
	t4 := t[4]*s[0] + t[5]*s[2] + s[4]
	t[1] = t[0]*s[1] + t[1]*s[3]
	t[3] = t[2]*s[1] + t[3]*s[3]
	t[5] = t[4]*s[1] + t[5]*s[3] + s[5]
	t[0] = t0
	t[2] = t2
	t[4] = t4
	return t
}

// PreMultiply makes the transform to the result of multiplication of two transforms, of A = B*A.
func (t Matrix) PreMultiply(s Matrix) Matrix {
	return s.Multiply(t)
}

// Inverse makes the destination to inverse of specified transform.
// Returns 1 if the inverse could be calculated, else 0.
func (t Matrix) Inverse() Matrix {
	det := t[0]*t[3] - t[2]*t[1]
	if det > -1e-6 && det < 1e-6 {
		return IdentityMatrix()
	}
	invdet := 1.0 / det
	return Matrix{
		t[3] * invdet,
		-t[1] * invdet,
		-t[2] * invdet,
		t[0] * invdet,
		(t[2]*t[5] - t[3]*t[4]) * invdet,
		(t[1]*t[4] - t[0]*t[5]) * invdet,
	}
}

// TransformPoint transforms a point by given TransformMatrix.
func (t Matrix) TransformPoint(sx, sy Scala) (dx, dy Scala) {
	dx = sx*t[0] + sy*t[2] + t[4]
	dy = sx*t[1] + sy*t[3] + t[5]
	return
}

// ToMat3x4 makes 3x4 matrix.
func (t Matrix) ToMat3x4() []Scala {
	return []Scala{
		t[0], t[1], 0.0, 0.0,
		t[2], t[3], 0.0, 0.0,
		t[4], t[5], 1.0, 0.0,
	}
}

func (t Matrix) getAverageScale() Scala {
	sx := math32.Sqrt(float32(t[0]*t[0] + t[2]*t[2]))
	sy := math32.Sqrt(float32(t[1]*t[1] + t[3]*t[3]))
	return Scala((sx + sy) * 0.5)
}
