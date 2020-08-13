package elements

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/zozowind/image-mix/util"
)

//TextAlignCenter 居中对齐
const TextAlignCenter = "center"

//Text 文本
type Text struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Text       string  `json:"text"`
	W          float64 `json:"w"`
	Font       string  `json:"font"`
	FontSize   int     `json:"fontSize"`
	Color      string  `json:"color"`
	Alpha      float64 `json:"alpha"`
	LineHeight int     `json:"lineHeight"`
	TextAlign  string  `json:"textAlign"`
}

//ModifyX 修改x
func (e *Text) ModifyX(w float64) float64 {
	if e.TextAlign == TextAlignCenter {
		return float64(e.X) - w/2
	}
	return float64(e.X)
}

//Draw 绘制
func (e *Text) Draw(g *gg.Context) (err error) {
	err = util.SetColor(g, e.Color, e.Alpha)
	if nil != err {
		err = fmt.Errorf("Text color %s %.2f err %s", e.Color, e.Alpha, err.Error())
		return
	}
	err = g.LoadFontFace(e.Font, float64(e.FontSize))
	if nil != err {
		err = fmt.Errorf("Text load font %s %d err %s", e.Font, e.FontSize, err.Error())
		return
	}
	w, _ := g.MeasureString(e.Text)
	words := g.WordWrap(e.Text, e.W)
	for index, word := range words {
		g.DrawString(word, e.ModifyX(w), e.Y+float64(e.LineHeight*index))
	}
	return
}
