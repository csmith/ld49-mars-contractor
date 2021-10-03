package scenes

import (
	"fmt"
	"github.com/csmith/mars-contractor/resources"
	"github.com/csmith/mars-contractor/resources/sounds"
	"github.com/hajimehoshi/ebiten/v2"
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
	resources.RenderTextLarge(
		screen,
		fmt.Sprintf(
			"Assignment issued by: Base AI                      Sol: 30%d\n"+
				"          Contractor: Bob Hines\n"+
				"                Task: %s\n\n"+
				"Additional information follows_\n\n%s",
			a.Day,
			a.Name,
			a.Text,
		),
		100,
		130,
	)
}

func (a Assignment) Update() Scene {
	sounds.EnableAiBackground = true
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
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
		case 5:
			return &BackupPower{}
		}
	}

	return a
}
