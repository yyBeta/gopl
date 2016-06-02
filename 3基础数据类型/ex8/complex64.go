package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	// "math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		// for instance, convert complex64 -> complex128
		// if cmplx.Abs(complex128(v)) > 2 {
		if abs32(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// abs for float32 ( not care about inf and NaN)
func abs32(x complex64) float32 {
	// borrow from math.hypot
	p, q := real(x), imag(x)
	if p < 0 {
		p = -p
	}
	if q < 0 {
		q = -q
	}
	if p < q {
		p, q = q, p
	}
	if p == 0 {
		return 0
	}
	q = q / p
	return p * float32(math.Sqrt(float64(1+q*q)))
}
