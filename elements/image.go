package elements

import (
	"fmt"
	"image"
	"net/http"
	"net/url"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/zozowind/image-mix/util"
)

//Image 图片
type Image struct {
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Z            float64 `json:"z"` //顺序
	W            float64 `json:"w"`
	H            float64 `json:"h"`
	Path         string  `json:"path"`
	BorderRadius int     `json:"borderRadius"`
}

// func (img Image) Ax() float64 {
// 	return float64(img.X) - float64(img.Width)*0.5
// }

// func (img Image) Ay() float64 {
// 	return float64(img.Y) - float64(img.Height)*0.5
// }

//Index 顺序
func (e *Image) Index() float64 {
	return e.Z
}

//Draw 绘制
func (e *Image) Draw(g *gg.Context) (err error) {
	u, err := url.Parse(e.Path)
	if nil != err {
		err = fmt.Errorf("img path %s err %s", e.Path, err.Error())
		return
	}
	var img image.Image
	if u.Host == "" {
		img, err = gg.LoadImage(e.Path)
		if nil != err {
			err = fmt.Errorf("load img %s err %s", e.Path, err.Error())
			return
		}
	} else {
		var rsp *http.Response
		rsp, err = http.DefaultClient.Get(e.Path)
		if nil != err {
			err = fmt.Errorf("img %s request err %s", e.Path, err.Error())
			return
		}

		img, _, err = image.Decode(rsp.Body)
		if nil != err {
			err = fmt.Errorf("img %s decode err %s", e.Path, err.Error())
			return
		}
	}

	img = resize.Resize(uint(e.W), uint(e.H), img, resize.Lanczos3)

	if e.BorderRadius > 0 {
		img = util.Circle(img)
	}
	g.DrawImage(img, int(e.X), int(e.Y))
	return
}
