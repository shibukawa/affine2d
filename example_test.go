package affine2d_test

import (
	"fmt"
	"github.com/shibukawa/affine2d"
	"math"
)

func ExampleTranslateMatrix() {
	// move 10, 20
	translate := affine2d.TranslateMatrix(10, 20)
	x, y := translate.TransformPoint(5, 5)
	fmt.Println(x, y)
	// Output: 15 25
}

func ExampleRotateMatrix() {
	// rotate 90 degree
	rotate := affine2d.RotateMatrix(affine2d.Pi * 0.5)
	x, y := rotate.TransformPoint(10, 10)
	fmt.Println(x, y)
	// Output: -10 10
}

func ExampleScaleMatrix() {
	// scale x1.5 for x-axis, x2.5 for y-axis
	scale := affine2d.ScaleMatrix(1.5, 2.5)
	x, y := scale.TransformPoint(10, 10)
	fmt.Println(x, y)
	// Output: 15 25
}

func ExampleMatrix() {
	translate := affine2d.TranslateMatrix(10, 20)
	rotate := affine2d.RotateMatrix(affine2d.Pi * 0.5)
	scale := affine2d.ScaleMatrix(1.5, 2.5)
	x, y := translate.Multiply(rotate).Multiply(scale).TransformPoint(10, 10)
	fmt.Println(math.Floor(float64(x) + 0.5), math.Floor(float64(y) + .5))
	// Output: -45 50
}