package poster

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
	"github.com/zozowind/image-mix/elements"
	"github.com/zozowind/image-mix/util"
)

//Poster 海报或混合图片
type Poster struct {
	W          int     `json:"w"`
	H          int     `json:"h"`
	Background string  `json:"background"` //背景颜色
	Alpha      float64 `json:"alpha"`
	Elements   []elements.Element
	Context    *gg.Context
}

//Draw 绘制海报
func (p *Poster) Draw() (err error) {
	//构建画布
	p.Context = gg.NewContext(p.W, p.H)
	//设置背景色

	p.Context.DrawRectangle(0, 0, float64(p.W), float64(p.H))
	if p.Background != "" {
		err = util.SetColor(p.Context, p.Background, p.Alpha)
		if nil != err {
			return
		}
	}

	p.Context.Fill()

	for _, ele := range p.Elements {
		err = ele.Draw(p.Context)
		if nil != err {
			fmt.Println(err.Error())
		}
	}
	return
}

//DrawToImage 返回Image
func (p *Poster) DrawToImage() (img image.Image, err error) {
	err = p.Draw()
	if nil != err {
		return
	}
	img = p.Context.Image()
	return
}

//DrawToFile 保存到文件
func (p *Poster) DrawToFile(path string) (err error) {
	err = p.Draw()
	if nil != err {
		return
	}
	err = p.Context.SavePNG(path)
	return
}
