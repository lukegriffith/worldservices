package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"github.com/lukegriffith/worldservices/internal/world"
)

var (
	frameRate time.Duration
	simLength = 500
	cycle     = 0
)

func renderCreature(x float64, y float64, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Clear()
	imd.Color = colornames.Navy
	imd.Push(pixel.V(x, y))
	imd.Ellipse(pixel.V(5, 5), 0)
	imd.Draw(win)
}

func run() {

	world.NewWorldService()
	world.RegisterWorld("sim")
	sim := world.NewWorld(769, 100)
	sim.Run(500)
	world.SetWorld("sim", sim)

	frameRate = 33 * time.Millisecond

	cfg := pixelgl.WindowConfig{
		Title:  "Creature Simulation",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Aliceblue)

	tick := time.Tick(frameRate)

	imd := imdraw.New(nil)

	for !win.Closed() {

		select {
		case <-tick:
			win.Clear(colornames.Aliceblue)
			w, err := world.GetWorldBoard("sim", cycle)
			c := w.Objects()
			if err != nil {
				panic(err)
			}

			for _, creature := range c {
				fmt.Println(creature)
				renderCreature(float64(creature.X), float64(creature.Y), imd, win)
			}
		}

		if cycle < 498 {
			cycle++
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
