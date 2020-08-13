package util

import (
	"image"
	"image/color"
	"math"
)

//Circle 图片取圆
func Circle(img image.Image) image.Image {
	w := img.Bounds().Max.X - img.Bounds().Min.X
	h := img.Bounds().Max.Y - img.Bounds().Min.Y

	d := w
	if w > h {
		d = h
	}

	return &CircleMask{img, image.Point{0, 0}, d}
}

//CircleMask 圆形遮罩
type CircleMask struct {
	image    image.Image
	point    image.Point
	diameter int
}

//ColorModel 颜色模型
func (ci *CircleMask) ColorModel() color.Model {
	return ci.image.ColorModel()
}

//Bounds 边界信息
func (ci *CircleMask) Bounds() image.Rectangle {
	return image.Rect(0, 0, ci.diameter, ci.diameter)
}

//At 位置
func (ci *CircleMask) At(x, y int) color.Color {
	d := ci.diameter
	dis := math.Sqrt(math.Pow(float64(x-d/2), 2) + math.Pow(float64(y-d/2), 2))
	if dis > float64(d)/2 {
		return ci.image.ColorModel().Convert(color.RGBA{255, 255, 255, 0})
	}
	return ci.image.At(ci.point.X+x, ci.point.Y+y)
}
