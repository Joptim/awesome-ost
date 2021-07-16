package vertical

import (
	"github.com/joptim/awesome-ost/frontend/components/layout"
)

type Vertical struct {
	layout.Layout
}

func (v *Vertical) Render() error {
	x0, y0, x1, y1 := v.GetDimensions()
	padding := v.GetPadding()
	children := v.GetChildren()
	childrenSizes := v.GetChildrenSizes()

	height := y1 - y0
	allPaddings := (len(v.GetChildrenSizes()) + 1) * padding
	available := height - allPaddings
	position := y0

	for i, h := range children {
		position += padding

		fSize := float64(childrenSizes[i])
		fAvailable := float64(available)
		childHeight := int((fSize / 12.0) * fAvailable)

		cx0, _, cx1, _ := h.GetDimensions()
		cx0, cx1 = layout.Clamp(cx0, cx1, x0, x1)
		if err := h.SetDimensions(cx0, position, cx1, position+childHeight); err != nil {
			panic(err)
		}

		if err := h.Render(); err != nil {
			panic(err)
		}
		position += childHeight
	}
	// Add padding at the end
	position += padding
	return nil
}

func New() *Vertical {
	return &Vertical{}
}
