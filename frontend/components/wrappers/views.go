package wrappers

import (
	"errors"
	"github.com/awesome-gocui/gocui"
	"github.com/joptim/awesome-ost/frontend/components"
)

// View is a *gocui.Views adapter
type View struct {
	components.Base
	view *gocui.View
	gui  *gocui.Gui
}

func (v *View) Render() error {
	var err error
	x0, y0, x1, y1 := v.GetDimensions()
	v.view, err = v.gui.SetView(v.view.Name(), x0, y0, x1, y1, v.view.Overlaps)
	if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
		panic(err)
	}
	return nil
}

func (v *View) GetView() *gocui.View {
	return v.view
}

// GetView returns a view with the given name, or creates a new one if it does not exist.
func GetView(name string, gui *gocui.Gui) *View {
	x, y := gui.Size()
	view, err := gui.View(name)
	if err != nil && errors.Is(err, gocui.ErrUnknownView) {
		// view does not exist. Create one
		view, err = gui.SetView(name, 0, 0, x, y, 0)
		if err != nil && !errors.Is(err, gocui.ErrUnknownView) {
			// unexpected error
			panic(err)
		}
	}
	wrapper := &View{
		view: view,
		gui:  gui,
	}

	if err := wrapper.SetDimensions(0, 0, x, y); err != nil {
		panic(err)
	}
	return wrapper

}
