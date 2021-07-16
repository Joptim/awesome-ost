package manager

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/joptim/awesome-ost/frontend/components"
	"github.com/joptim/awesome-ost/frontend/components/layout/horizontal"
	"github.com/joptim/awesome-ost/frontend/components/layout/vertical"
	"github.com/joptim/awesome-ost/frontend/components/wrappers"
	"github.com/joptim/awesome-ost/frontend/components/wrappers/asciiart"
)

func Layout(gui *gocui.Gui) error {
	navbarLayout := CreateHorizontalNavbar(gui)
	//tracksLayout := CreateControlTrackList(gui)
	navigateLayout := CreateControlsMenu(gui)

	x, y := gui.Size()
	vLayout := vertical.New()
	vLayout.SetDimensions(0, 0, x, y)
	vLayout.SetChildren(
		[]components.IRenderer{
			navbarLayout,
			navigateLayout,
		},
	)
	vLayout.SetChildrenSizes([]int{6, 6})
	return vLayout.Render()
}

func CreateHorizontalNavbar(gui *gocui.Gui) *horizontal.Horizontal {
	x, y := gui.Size()
	hLayout := horizontal.New()
	hLayout.SetDimensions(0, 0, x, y)
	hLayout.SetPadding(2)
	hLayout.SetChildren(
		[]components.IRenderer{
			CreateNavBarItem("Section A", gui),
			CreateNavBarItem("Section B", gui),
			CreateNavBarItem("Section C", gui),
			CreateNavBarItem("Section D", gui),
			CreateNavBarItem("Section E", gui),
			CreateNavBarItem("Finish", gui),
		},
	)
	hLayout.SetChildrenSizes([]int{2, 2, 2, 2, 2, 2})
	return hLayout
}

func CreateControlsMenu(gui *gocui.Gui) *horizontal.Horizontal {
	x, y := gui.Size()
	hLayout := horizontal.New()
	hLayout.SetDimensions(0, 0, x, y)
	hLayout.SetPadding(2)
	hLayout.SetChildren(
		[]components.IRenderer{
			CreateControlNavigateKeysItem(gui),
			CreateControlSelectItem(gui),
			CreateControlPlayItem(gui),
		},
	)
	hLayout.SetChildrenSizes([]int{4, 4, 4})
	return hLayout
}

func CreateNavBarItem(text string, gui *gocui.Gui) *wrappers.View {
	wrapper := wrappers.GetView(fmt.Sprintf("navbar_%s", text), gui)
	view := wrapper.GetView()
	view.Frame = true
	view.FrameRunes = []rune{'═', '║', '╔', '╗', '╚', '╝', '╠', '╣', '╦', '╩', '╬'}
	view.Clear()
	view.Write([]byte(text))
	return wrapper
}

func CreateControlNavigateKeysItem(gui *gocui.Gui) *asciiart.View {
	wrapper := asciiart.GetView(asciiart.KEYS, "arrow_keys", gui)
	view := wrapper.GetView()
	view.Frame = true
	return wrapper
}

func CreateControlSelectItem(gui *gocui.Gui) *asciiart.View {
	wrapper := asciiart.GetView(asciiart.ENTERKEY, "select_keys", gui)
	wrapper.GetView().Frame = true
	return wrapper
}

func CreateControlPlayItem(gui *gocui.Gui) *asciiart.View {
	wrapper := asciiart.GetView(asciiart.SPACEBARKEY, "play_keys", gui)
	wrapper.GetView().Frame = true
	return wrapper
}
