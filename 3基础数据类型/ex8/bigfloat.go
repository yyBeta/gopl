package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	limit := new(big.Float).SetFloat64(2.0)
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if hypot(big.NewFloat(real(v)), big.NewFloat(imag(v))).Cmp(limit) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// hypot for big.Float
func hypot(p, q *big.Float) *big.Float {
	// special cases
	switch {
	case p.IsInf() || q.IsInf():
		return big.NewFloat(math.Inf(1))
	}
	p = p.Abs(p)
	q = q.Abs(q)
	if p.Cmp(p) < 0 {
		p, q = q, p
	}
	if p.Cmp(big.NewFloat(0)) == 0 {
		return big.NewFloat(0)
	}
	q = q.Quo(q, p)
	return sqrt(q.Mul(q, q).Add(q, big.NewFloat(1))).Mul(q, p)
}

// sqrt for big.Float
func sqrt(given *big.Float) *big.Float {
	const prec = 200
	steps := int(math.Log2(prec))
	given.SetPrec(prec)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	x := new(big.Float).SetPrec(prec).SetInt64(1)
	t := new(big.Float)
	for i := 0; i <= steps; i++ {
		t.Quo(given, x)
		t.Add(x, t)
		t.Mul(half, t)
	}
	return x
}
