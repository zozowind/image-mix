package elements

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/zozowind/image-mix/util"
)

//Line 线条
type Line struct {
	SX    float64 `json:"sx"`
	SY    float64 `json:"sy"`
	EX    float64 `json:"ex"`
	EY    float64 `json:"ey"`
	W     float64 `json:"w"`
	Color string  `json:"color"`
	Alpha float64 `json:"alpha"`
}

//Draw 绘制
func (e *Line) Draw(g *gg.Context) (err error) {
	err = util.SetColor(g, e.Color, e.Alpha)
	if nil != err {
		err = fmt.Errorf("Line color %s %.2f err %s", e.Color, e.Alpha, err.Error())
		return
	}
	g.SetLineWidth(e.W)
	g.DrawLine(e.SX, e.SY, e.EX, e.EY)
	g.Stroke()
	return
}
