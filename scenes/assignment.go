package scenes

import (
	"fmt"
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

var assignmentBackground *ebiten.Image

func init() {
	var err error
	assignmentBackground, _, err = resources.LoadImageAsset("assignment.png")
	if err != nil {
		panic(err)
	}
}

type Assignment struct {
	Day          int
	Name         string
	Text         string
	ShouldReject bool
}

func (a Assignment) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(assignmentBackground, nil)
	ebitenutil.DebugPrintAt(screen, "Outline: AI message console", 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("DAY: %2d", a.Day), 10, 30)
	ebitenutil.DebugPrintAt(screen, "CONTRACTOR: Bob McBobFace", 10, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ASSIGNMENT: %s", a.Name), 10, 70)
	ebitenutil.DebugPrintAt(screen, "----- MESSAGE -----", 10, 90)
	ebitenutil.DebugPrintAt(screen, a.Text, 10, 110)
	ebitenutil.DebugPrintAt(screen, "ACCEPT ASSIGNMENT", 10, 750)
	ebitenutil.DebugPrintAt(screen, "IGNORE ASSIGNMENT", 900, 750)
}

func (a Assignment) Update() Scene {
	if x, y := ebiten.CursorPosition(); inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && y > 700 {
		if x < 450 {
			if a.ShouldReject {
				return &GameOver{Days: a.Day}
			} else {
				switch a.Day {
				case 1:
					s := &PipesGame{}
					s.Init()
					return s
				}
			}
		} else {
			if a.ShouldReject {
				// TODO: Win state
			} else {
				return &GameOver{Days: a.Day}
			}
		}

		return &Assignment{
			Day:  1,
			Name: resources.WorkAssignmentOneTitle,
			Text: resources.WorkAssignmentOneBody,
		}
	}

	return a
}
