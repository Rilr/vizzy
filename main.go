package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	showSine bool
	phase    float64 // phase offset for rolling effect
}

const (
	width  = 640
	height = 480
)

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.showSine = true
		g.phase += 0.05 // advance phase for rolling effect
	} else {
		g.showSine = false
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.showSine {
		// Draw rolling sine wave
		for x := 0; x < width; x++ {
			y := int(height/2 + 100*math.Sin(float64(x)*2*math.Pi/float64(width)+g.phase))
			if y >= 0 && y < height {
				screen.Set(x, y, color.White)
			}
		}
	} else {
		// Draw straight line
		for x := 0; x < width; x++ {
			y := height / 2
			screen.Set(x, y, color.White)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Line to Sine Wave (press SPACE)")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
