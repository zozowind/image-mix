package elements

import (
	"github.com/fogleman/gg"
)

//Element 元素
type Element interface {
	Draw(*gg.Context) error
}
