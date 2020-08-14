package examples

import (
	"testing"

	"github.com/zozowind/image-mix/elements"
	"github.com/zozowind/image-mix/poster"
)

func TestPoster(t *testing.T) {
	rect := &elements.Rectangle{
		X:          float64(10),
		Y:          float64(10),
		W:          float64(30),
		H:          float64(40),
		Background: "#ababab",
		Border:     "#00FF00",
	}

	line := &elements.Line{
		SX:    float64(0),
		SY:    float64(0),
		EX:    float64(150),
		EY:    float64(150),
		W:     float64(50),
		Color: "#ff0000",
		Alpha: 1,
	}

	img := &elements.Image{
		X:            float64(0),
		Y:            float64(0),
		W:            float64(80),
		H:            float64(80),
		Path:         "https://png.zhaoxi.net/upload/18/01/1400/63d99185de9844fc44d9e8ab9ebaa408.png",
		BorderRadius: 80,
	}

	text := &elements.Text{
		X:          float64(120),
		Y:          float64(50),
		W:          float64(80),
		Text:       "我是一只小花猫，喵喵喵 abc134",
		Font:       "../font/pingfangsr.ttf",
		FontSize:   12,
		Color:      "#00abff",
		Alpha:      0.7,
		LineHeight: 26,
		TextAlign:  elements.TextAlignCenter,
	}

	poster := poster.Poster{
		W:          300,
		H:          200,
		Background: "#003333",
		Alpha:      0.5,
		Elements:   []elements.Element{line, rect, img, text},
	}
	err := poster.DrawToFile("./test.png")
	if nil != err {
		t.Error(err.Error())
	}
}
