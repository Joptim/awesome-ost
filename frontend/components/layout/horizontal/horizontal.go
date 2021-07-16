package horizontal

import "github.com/joptim/awesome-ost/frontend/components/layout"

type Horizontal struct {
	layout.Layout
}

func (h *Horizontal) Render() error {
	x0, y0, x1, y1 := h.GetDimensions()
	padding := h.GetPadding()
	children := h.GetChildren()
	childrenSizes := h.GetChildrenSizes()

	width := x1 - x0
	allPaddings := (len(h.GetChildrenSizes()) + 1) * padding
	available := width - allPaddings
	position := x0

	for i, h := range children {
		position += padding

		fSize := float64(childrenSizes[i])
		fAvailable := float64(available)
		childWidth := int((fSize / 12.0) * fAvailable)

		_, cy0, _, cy1 := h.GetDimensions()
		cy0, cy1 = layout.Clamp(cy0, cy1, y0, y1)
		if err := h.SetDimensions(position, cy0, position+childWidth, cy1); err != nil {
			panic(err)
		}
		if err := h.Render(); err != nil {
			panic(err)
		}
		position += childWidth
	}
	// Add padding at the end
	position += padding
	return nil
}

func New() *Horizontal {
	return &Horizontal{}
}
