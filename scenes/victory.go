package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var victoryBackground *ebiten.Image

func init() {
	var err error
	victoryBackground, _, err = resources.LoadImageAsset("victory.png")
	if err != nil {
		panic(err)
	}
}

type Victory struct {

}

func (v *Victory) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(victoryBackground, nil)
}

func (v *Victory) Update() Scene {
	return v
}

