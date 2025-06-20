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
		// Draw a rolling spiked (starburst) circle with overshoot and recoil effect
		radius := 200.0
		cx, cy := width/2, height/2
		spikeCount := 60.0  // number of spikes
		spikeHeight := 100.0 // how tall the spikes are
		// Use a damped sine for overshoot/recoil
		for i := range width {
			angle := float64(i) * 2 * math.Pi / float64(width)
			// Damped oscillation: overshoot and recoil
			osc := math.Sin(spikeCount*angle + g.phase)
			decay := math.Exp(-2 * math.Abs(osc)) // fast decay for recoil
			overshoot := spikeHeight * osc * decay
			spikedRadius := radius + overshoot
			xCircle := int(float64(cx) + spikedRadius*math.Cos(angle+g.phase))
			yCircle := int(float64(cy) + spikedRadius*math.Sin(angle+g.phase))
			if xCircle >= 0 && xCircle < width && yCircle >= 0 && yCircle < height {
				screen.Set(xCircle, yCircle, color.White)
			}
		}
	} else {
		// Draw a Circle
		radius := 100.0
		cx, cy := width/2, height/2
		for i := range 360 {
			angle := float64(i) * 2 * math.Pi / 360
			xCircle := int(float64(cx) + radius*math.Cos(angle))
			yCircle := int(float64(cy) + radius*math.Sin(angle))
			if xCircle >= 0 && xCircle < width && yCircle >= 0 && yCircle < height {
				screen.Set(xCircle, yCircle, color.White)
			}
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
