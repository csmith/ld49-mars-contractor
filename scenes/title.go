package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

var titleBackground *ebiten.Image

func init() {
	var err error
	titleBackground, _, err = resources.LoadImageAsset("splash.png")
	if err != nil {
		panic(err)
	}
}

type Title struct {

}

func (t Title) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(titleBackground, nil)
	ebitenutil.DebugPrintAt(screen, "You are a general contractor at SpaceY's newest base on Mars.", 100, 350)
	ebitenutil.DebugPrintAt(screen, "You are assigned tasks by the base's AI. While your job may", 100, 360)
	ebitenutil.DebugPrintAt(screen, "seem menial, SpaceY assures you that your work is vital to", 100, 370)
	ebitenutil.DebugPrintAt(screen, "the success of the mission.", 100, 380)
}

func (t Title) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return &Assignment{
			Day:  1,
			Name: resources.WorkAssignmentOneTitle,
			Text: resources.WorkAssignmentOneBody,
		}
	}

	return t
}
