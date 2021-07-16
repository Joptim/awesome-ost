package asciiart

import (
	"github.com/awesome-gocui/gocui"
	"github.com/joptim/awesome-ost/frontend/components/wrappers"
	"strings"
)

type View struct {
	wrappers.View
	library Library
}

func (t *View) Render() error {
	x0, y0, x1, y1 := t.GetDimensions()
	width := x1 - x0
	height := y1 - y0

	title_, _ := New(t.library, width, height)
	v := t.GetView()
	v.Clear()

	text := center(title_, width, height)
	if _, err := v.Write([]byte(text)); err != nil {
		return err
	}
	return t.View.Render()
}

func center(title_ AsciiArt, width, height int) string {
	center := strings.Repeat("\n", (height-title_.Height())/2)
	lines := strings.Split(title_.String(), "\n")
	for _, line := range lines {
		center += strings.Repeat(" ", (width-title_.Width())/2) + line + "\n"
	}
	return center
}

// GetView returns a view with the given name, or creates a new one if it does not exist.
func GetView(library Library, name string, gui *gocui.Gui) *View {
	return &View{
		View:    *wrappers.GetView(name, gui),
		library: library,
	}
}
