package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
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
	resources.RenderTextLarge(screen,
		"You are a general contractor at SpaceY's\n"+
			"newest base on Mars.\n\n"+
			"You are assigned tasks by the base's AI.\n\n"+
			"While your job may seem menial, SpaceY\n"+
			"assures you that your work is vital to\n"+
			"the success of the mission.", 60, 300)
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
