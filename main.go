package main

import (
	"errors"
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/joptim/awesome-ost/frontend/scenes/mixer/manager"
	"log"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	// Enable mouse support
	//g.Mouse = true

	g.SetManagerFunc(manager.Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		panic(fmt.Errorf("unexpected error: %v", err))
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
