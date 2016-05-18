package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			x1, y1 := float64(px)/width*(xmax-xmin)+xmin, float64(py)/height*(ymax-ymin)+ymin
			x2, y2 := float64(px+1)/width*(xmax-xmin)+xmin, float64(py+1)/height*(ymax-ymin)+ymin
			r1, g1, b1, a1 := mandelbrot(complex(x1, y1)).RGBA()
			r2, g2, b2, a2 := mandelbrot(complex(x1, y2)).RGBA()
			r3, g3, b3, a3 := mandelbrot(complex(x2, y1)).RGBA()
			r4, g4, b4, a4 := mandelbrot(complex(x2, y2)).RGBA()
			c := color.RGBA{
				uint8(r1+r2+r3+r4) / 4,
				uint8(g1+g2+g3+g4) / 4,
				uint8(b1+b2+b3+b4) / 4,
				uint8(a1+a2+a3+a4) / 4,
			}
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
