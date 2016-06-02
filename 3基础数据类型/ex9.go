package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)

	height, err := toInt(r.URL.Query().Get("height"), 1024)
	if err != nil {
		log.Println(err)
		http.Error(w, "height should be integer", http.StatusBadRequest)
		return
	}
	width, err := toInt(r.URL.Query().Get("width"), 1024)
	if err != nil {
		log.Println(err)
		http.Error(w, "width should be integer", http.StatusBadRequest)
		return
	}
	zoom, err := toInt(r.URL.Query().Get("zoom"), 1)
	if err != nil {
		log.Println(err)
		http.Error(w, "zoom should be integer", http.StatusBadRequest)
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x/float64(zoom), y/float64(zoom))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
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

func toInt(param string, d int) (int, error) {
	if len(param) == 0 {
		return d, nil
	}
	h, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return h, nil
}
