package affine2d

import (
	"testing"
	"math"
)

func nearlyEqual(a, b Scala) bool {
	return math.Abs(float64(a) - float64(b)) < 0.01
}

func TestNoTransform(t *testing.T) {
	noTransform := IdentityMatrix()
	a, b := noTransform.TransformPoint(10, 20)
	if !nearlyEqual(a, 10.0) {
		t.Errorf("a should be 10.0, but %f\n", a)
	}
	if !nearlyEqual(b, 20.0) {
		t.Errorf("a should be 20.0, but %f\n", b)
	}
}

func TestTranslate(t *testing.T) {
	translate := TranslateMatrix(40, 60)
	a, b := translate.TransformPoint(10, 20)
	if !nearlyEqual(a, 50.0) {
		t.Errorf("a should be 50.0, but %f\n", a)
	}
	if !nearlyEqual(b, 80.0) {
		t.Errorf("a should be 80.0, but %f\n", b)
	}
}

func TestRotate(t *testing.T) {
	rotate := RotateMatrix(Pi * 0.5)
	a, b := rotate.TransformPoint(10, 20)
	if !nearlyEqual(a, -20.0) {
		t.Errorf("a should be -20.0, but %f\n", a)
	}
	if !nearlyEqual(b, 10.0) {
		t.Errorf("a should be 10.0, but %f\n", b)
	}
}

func TestScale(t *testing.T) {
	scale := ScaleMatrix(2, 3)
	a, b := scale.TransformPoint(10, 20)
	if !nearlyEqual(a, 20.0) {
		t.Errorf("a should be 20.0, but %f\n", a)
	}
	if !nearlyEqual(b, 60.0) {
		t.Errorf("a should be 60.0, but %f\n", b)
	}
}

func TestSkewX(t *testing.T) {
	skewX := SkewXMatrix(Pi * 0.25)
	a, b := skewX.TransformPoint(10, 20)
	if !nearlyEqual(a, 30.0) {
		t.Errorf("a should be 30.0, but %f\n", a)
	}
	if !nearlyEqual(b, 20.0) {
		t.Errorf("a should be 20.0, but %f\n", b)
	}
}

func TestSkewY(t *testing.T) {
	skewY := SkewYMatrix(Pi * 0.25)
	a, b := skewY.TransformPoint(10, 20)
	if !nearlyEqual(a, 10.0) {
		t.Errorf("a should be 10.0, but %f\n", a)
	}
	if !nearlyEqual(b, 30.0) {
		t.Errorf("a should be 30.0, but %f\n", b)
	}
}

func TestMultiply(t *testing.T) {
	t1 := TranslateMatrix(10, 20)
	a, b := t1.Multiply(RotateMatrix(0.5 * Pi)).TransformPoint(10, 20)
	if !nearlyEqual(a, -40.0) {
		t.Errorf("a should be -20.0, but %f\n", a)
	}
	if !nearlyEqual(b, 20.0) {
		t.Errorf("a should be 10.0, but %f\n", b)
	}
}