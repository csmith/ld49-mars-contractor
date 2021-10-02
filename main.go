package main

import (
	"github.com/csmith/mars-contractor/resources/sounds"
	"github.com/csmith/mars-contractor/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	windowWidth  = 1280
	windowHeight = 768
)

type Game struct {
	scene scenes.Scene
}

func (g *Game) Update() error {
	g.scene = g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Mars contractor")
	ebiten.SetWindowResizable(false)
	sounds.PlayBackground()
	if err := ebiten.RunGame(&Game{
		&scenes.Title{},
	}); err != nil {
		log.Fatal(err)
	}
}
