package poster

import (
	"fmt"

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
}

//DrawToFile 保存到文件
func (p *Poster) DrawToFile(path string) (err error) {
	//构建画布
	g := gg.NewContext(p.W, p.H)
	//设置背景色

	g.DrawRectangle(0, 0, float64(p.W), float64(p.H))
	if p.Background != "" {
		err = util.SetColor(g, p.Background, p.Alpha)
		if nil != err {
			return err
		}
	}
	g.Fill()

	for _, ele := range p.Elements {
		err = ele.Draw(g)
		if nil != err {
			fmt.Println(err.Error())
		}
	}
	err = g.SavePNG(path)
	return
}
