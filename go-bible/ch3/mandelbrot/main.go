// @Description:
// @Author: Arvin
// @date: 2021/3/1 10:30 上午
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		creatImg(resp)
	})
	http.ListenAndServe("localhost:8000", nil)
}

func creatImg(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// image.Rect 矩形  image.NewRGBA 返回一个有边界的RGBA图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px,py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// gray灰色
			//return color.Gray{Y: 255 - contrast*n}
			return color.RGBA{
				R: contrast * n,
				G: 255 - contrast*n,
				B: 255 - contrast*n,
				A: 255,
			}
		}
	}
	return color.Black
}
