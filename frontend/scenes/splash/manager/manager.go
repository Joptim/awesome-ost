package manager

import (
	"github.com/awesome-gocui/gocui"
	"github.com/joptim/awesome-ost/frontend/components"
	"github.com/joptim/awesome-ost/frontend/components/layout/horizontal"
	"github.com/joptim/awesome-ost/frontend/components/layout/vertical"
	"github.com/joptim/awesome-ost/frontend/components/wrappers"
	"github.com/joptim/awesome-ost/frontend/components/wrappers/asciiart"
)

func Layout(gui *gocui.Gui) error {
	vLayout := CreateVerticalLayout(gui)
	hLayout := CreateHorizontalLayout(vLayout, gui)
	return hLayout.Render()
}

func CreateVerticalLayout(gui *gocui.Gui) *vertical.Vertical {
	x, y := gui.Size()
	vLayout := vertical.New()
	vLayout.SetDimensions(0, 0, x, y)
	vLayout.SetChildren(
		[]components.IRenderer{
			CreateEmptySlot("empty_slot_1", gui),
			CreateTitle(gui),
			CreateMenu(gui),
			CreateEmptySlot("empty_slot_2", gui),
		},
	)
	vLayout.SetChildrenSizes([]int{1, 5, 5, 1})
	vLayout.SetPadding(2)
	return vLayout
}

func CreateHorizontalLayout(renderer components.IRenderer, gui *gocui.Gui) *horizontal.Horizontal {
	x, y := gui.Size()
	hLayout := horizontal.New()
	hLayout.SetDimensions(0, 0, x, y)
	hLayout.SetChildren(
		[]components.IRenderer{
			CreateEmptySlot("empty_slot_3", gui),
			renderer,
			CreateEmptySlot("empty_slot_4", gui),
		},
	)
	hLayout.SetChildrenSizes([]int{2, 8, 2})
	return hLayout
}

func CreateTitle(gui *gocui.Gui) *asciiart.View {
	wrapper := asciiart.GetView(asciiart.TITLE, "title", gui)
	view := wrapper.GetView()
	view.Frame = true
	view.FrameRunes = []rune{'═', '║', '╔', '╗', '╚', '╝'}
	return wrapper
}

func CreateMenu(gui *gocui.Gui) *wrappers.View {
	wrapper := wrappers.GetView("menu", gui)
	view := wrapper.GetView()
	view.Clear()
	view.Wrap = true
	view.Frame = false
	view.Write([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin id tristique urna, mattis posuere ipsum. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin id tristique urna, mattis posuere ipsum."))
	return wrapper
}

func CreateEmptySlot(name string, gui *gocui.Gui) *wrappers.View {
	wrapper := wrappers.GetView(name, gui)
	view := wrapper.GetView()
	view.Frame = false
	return wrapper
}
