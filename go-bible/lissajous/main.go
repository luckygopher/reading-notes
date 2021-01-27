// @Description: 1.4 GIF动画
// @Author: Arvin
// @date: 2021/1/27 2:41 下午
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var green = color.RGBA{R: uint8(0), G: uint8('x'), B: uint8('G'), A: uint8('G')}

var palette = []color.Color{color.White, color.Black, green}

const (
	whiteIndex = 0 // 调色板中的白色
	blackIndex = 1 // 调色板中的黑色
	greenIndex = 2 // 调色板中的绿色
)

func main() {
	// 图像序列是确定的，除非我们seed
	// 使用当前时间的伪随机数生成器
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整x振荡器转数
		res     = 0.001 // 角分辨率
		size    = 100   // 图像画布封面尺寸 [-size..+size]
		nframes = 64    // 动画帧数
		delay   = 8     // 以8ms为单位的帧间延迟
	)

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 相位差
	for i := 0; i < nframes; i++ {
		// 生成一个单独的动画帧，包含两种颜色的201*201大小的图片，白色和黑色
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 新建调色板
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// 向标准输出流打印信息
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
