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

func (g *GameOver) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(gameOverBackground, nil)
	ebitenutil.DebugPrintAt(screen, g.RCA, 132, 264)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%2d ", g.Days), 1032, 484)

}

func (g *GameOver) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		switch g.Days {
		case 1:
			return &Assignment{
				Day:          g.Days,
				Name:         resources.WorkAssignmentOneTitle,
				Text:         resources.WorkAssignmentOneBody,
				ShouldReject: false,
			}
		case 2:
			return &Assignment{
				Day:          g.Days,
				Name:         resources.WorkAssignmentTwoTitle,
				Text:         resources.WorkAssignmentTwoBody,
				ShouldReject: false,
			}
		case 3:
			return &Assignment{
				Day:          g.Days,
				Name:         resources.WorkAssignmentThreeTitle,
				Text:         resources.WorkAssignmentThreeBody,
				ShouldReject: false,
			}
		case 4:
			return &Assignment{
				Day:          g.Days,
				Name:         resources.WorkAssignmentFourTitle,
				Text:         resources.WorkAssignmentFourBody,
				ShouldReject: false,
			}
		case 5:
			return &Assignment{
				Day:          g.Days,
				Name:         resources.WorkAssignmentSixTitle,
				Text:         resources.WorkAssignmentSixBody,
				ShouldReject: true,
			}
		}
		return &Title{}
	}
	return g
}
