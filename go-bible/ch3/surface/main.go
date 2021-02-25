// @Description: 浮点数画图
// @Author: Arvin
// @date: 2021/2/25 11:51 上午
package main

import (
	"fmt"
	"math"
	"net/http"
	"strings"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30.0
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°),cos(30°)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		svg := make([]string, 0)
		// 如果服务器使用标准的PNG图像格式，可以根据前面的512个字节自动输出对应的头部
		res.Header().Set("Content-Type", "image/svg+xml")

		svg = append(svg, fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height))
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				svg = append(svg, fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy))
			}
		}
		svg = append(svg, "</svg>")

		fmt.Fprintf(res, strings.Join(svg, "\n"))
	})

	http.ListenAndServe("localhost:8000", nil)
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 返回所有参数的 平方和 的 平方根
	return math.Sin(r) / r
}

// math包学习
func MathPackage() {
	var r float64
	x, y := 3.0, 4.0

	// 常用计算函数
	// 计算直角三角形的斜边长 r = √(x^2 + y^2)
	r = math.Hypot(x, y)
	// 平方根函数 √y = 2
	r = math.Sqrt(y)
	// 立方根函数
	r = math.Cbrt(8.0)
	// x 的幂函数
	r = math.Pow(x, y)
	// 10根的幂函数
	r = math.Pow10(2)
	// 绝对值函数
	r = math.Abs(x)
	// 向上取整
	r = math.Ceil(x)
	// 向下取整
	r = math.Floor(x)
	// 取大值
	r = math.Max(x, y)
	// 取小值
	r = math.Min(x, y)
	// 取模
	r = math.Mod(x, y)
	// 取余运算(余数可能为负数)
	r = math.Remainder(x, y)
	// 只保留整数部分的函数
	r = math.Trunc(x)
	// 获取 x 是否为负 (-0 也会返回false)
	b := math.Signbit(-1)

	fmt.Println(r, b)
}
