package elements

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/zozowind/image-mix/util"
)

//Rectangle 块
type Rectangle struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	W          float64 `json:"w"`
	H          float64 `json:"h"`
	Background string  `json:"background"`
	Alpha      float64 `json:"alpha"`
	Border     string  `json:"border"`
}

//Draw 画
func (e *Rectangle) Draw(g *gg.Context) (err error) {
	g.DrawRectangle(e.X, e.Y, e.W, e.H)
	if e.Background != "" {
		err = util.SetColor(g, e.Background, e.Alpha)
		if nil != err {
			err = fmt.Errorf("Rectangle Background color %s %.2f err %s", e.Background, e.Alpha, err.Error())
			return
		}
		g.Fill()
	}
	if e.Border != "" {
		err = util.SetColor(g, e.Border, e.Alpha)
		if nil != err {
			err = fmt.Errorf("Rectangle Border color %s %.2f err %s", e.Border, e.Alpha, err.Error())
			return
		}
		g.Stroke()
	}
	return
}
