package util

import (
	"github.com/fogleman/gg"
	"gopkg.in/go-playground/colors.v1"
)

//HexToRGBA 16进制ToRGBA
func HexToRGBA(hex string, a float64) (color *colors.RGBAColor, err error) {
	hexColor, err := colors.ParseHEX(hex)
	if nil != err {
		return
	}
	color = hexColor.ToRGBA()
	color.A = a
	return
}

//SetColor 设置颜色
func SetColor(g *gg.Context, hex string, a float64) (err error) {
	color, err := HexToRGBA(hex, a)
	if nil != err {
		return
	}
	g.SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A*255))
	return
}
