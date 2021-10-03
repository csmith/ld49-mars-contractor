package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

var backupPowerBackground *ebiten.Image

func init() {
	var err error
	backupPowerBackground, _, err = resources.LoadImageAsset("backuppower.png")
	if err != nil {
		panic(err)
	}
}

type BackupPower struct {
}

func (p *BackupPower) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(backupPowerBackground, nil)
}

func (p *BackupPower) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x > 251*4 && y > 171*4 {
			return &Victory{}
		} else {
			return &GameOver{
				Days: 5,
				RCA:  "A SpaceY contractor deliberately repowered the\nunstable base AI.\n\nIt immediately vented atmosphere in retaliation for\nbeing cut off.\n\nAll hands were lost.",
			}
		}
	}

	return p
}
