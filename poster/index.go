package poster

import (
	"fmt"
	"image"
	"sort"

	"github.com/fogleman/gg"
	"github.com/zozowind/image-mix/elements"
	"github.com/zozowind/image-mix/util"
)

//ElementSlice 元素组
type ElementSlice []elements.Element

func (s ElementSlice) Len() int {
	return len(s)
}
func (s ElementSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ElementSlice) Less(i, j int) bool {
	return s[i].Index() < s[j].Index()
}

//Poster 海报或混合图片
type Poster struct {
	W          int     `json:"w"`
	H          int     `json:"h"`
	Background string  `json:"background"` //背景颜色
	Alpha      float64 `json:"alpha"`
	Elements   ElementSlice
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

	//排序
	sort.Sort(p.Elements)
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
