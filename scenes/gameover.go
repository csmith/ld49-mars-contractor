package scenes

import (
	"fmt"
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

var gameOverBackground *ebiten.Image

func init() {
	var err error
	gameOverBackground, _, err = resources.LoadImageAsset("failure.png")
	if err != nil {
		panic(err)
	}
}

type GameOver struct {
	Days int
	RCA  string
}

func (g GameOver) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(gameOverBackground, nil)
	ebitenutil.DebugPrintAt(screen, g.RCA, 132, 264)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%2d ", g.Days), 1032, 484)

}

func (g GameOver) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return &Title{}
	}
	return g
}
