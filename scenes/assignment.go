package scenes

import (
	"fmt"
	"github.com/csmith/mars-contractor/resources"
	"github.com/csmith/mars-contractor/resources/sounds"
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
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%02d", a.Day), 1170, 60)
	ebitenutil.DebugPrintAt(screen, "Bob McBobFace", 370, 130)
	ebitenutil.DebugPrintAt(screen, a.Name, 370, 180)
	ebitenutil.DebugPrintAt(screen, a.Text, 100, 400)
}

func (a Assignment) Update() Scene {
	sounds.EnableAiBackground = true
	if x, y := ebiten.CursorPosition(); inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && y > 600 {
		if x < 450 {
			if a.ShouldReject {
				return &GameOver{
					Days: a.Day,
					RCA: "A SpaceY contractor deliberately repowered\nthe known-unstable base AI.\n\nIt immediately vented atmosphere\nin retaliation for being cut off.",
				}
			} else {
				sounds.EnableAiBackground = false
				switch a.Day {
				case 1:
					s := &PipesGame{Day: 1}
					s.Init()
					return s
				case 2:
					return &RateGame{}
				case 3:
					return &LockedDoor{}
				case 4:
					return &Power{}
				}
			}
		} else {
			if a.ShouldReject {
				sounds.EnableAiBackground = false
				return &Victory{}
			} else {
				sounds.EnableAiBackground = false
				return &GameOver{
					Days: a.Day,
					RCA: rejectReasons[a.Day-1],
				}
			}
		}

		sounds.EnableAiBackground = false
		return &Assignment{
			Day:  1,
			Name: resources.WorkAssignmentOneTitle,
			Text: resources.WorkAssignmentOneBody,
		}
	}

	return a
}

var rejectReasons = []string{
	"Coolant flow to the reactor was compromised.\n\nThe general contractor assigned ignored the work order.\n\nThe reactor went super-critical; all hands were lost.",
	"Coolant flow to the reactor was compromised.\n\nThe general contractor assigned ignored the work order.\n\nThe reactor went super-critical; all hands were lost.",
	"Due to a lack of paperwork, the crew\nwere lead to believe the AI was unstable\n\nThey attempted to disable it but were\nnot in a position to do so.\n\nThe crew were eliminated in self-defence.",
	"Excess power draw caused rolling blackouts.\n\nThe general contractor assigned ignored the work order.\n\nLife support, meteor defence and hydroponics all suffered.\n\nThe consequences were dire.",
}
