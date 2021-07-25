package main

/*func main() {
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
}*/

import (
	"github.com/joptim/awesome-ost/backend/server"
	"log"
)

func main() {
	log.Fatalln(server.New().ListenAndServe(":8000"))
}
