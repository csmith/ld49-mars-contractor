package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"image"
)

var powerBackground *ebiten.Image
var powerAltBackground *ebiten.Image

func init() {
	var err error
	powerBackground, _, err = resources.LoadImageAsset("power.png")
	if err != nil {
		panic(err)
	}

	powerAltBackground, _, err = resources.LoadImageAsset("power-alt.png")
	if err != nil {
		panic(err)
	}
}

type Power struct {
	switches [12]bool
	saved    int
}

func (p *Power) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)

	if p.switches[2] {
		// Turned the lights off...
		return
	}

	screen.DrawImage(powerBackground, nil)

	for i := range p.switches {
		if p.switches[i] {
			rect := p.bounds(i)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
			screen.DrawImage(powerAltBackground.SubImage(rect).(*ebiten.Image), op)
		}
	}

	if p.saved > 0 {
		meterLeft := 1056
		meterRight := 1212
		meterTop := 264
		meterBottom := meterTop + 28*p.saved

		rect := image.Rect(meterLeft, meterTop, meterRight, meterBottom)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(rect.Min.X), float64(rect.Min.Y))
		screen.DrawImage(powerAltBackground.SubImage(rect).(*ebiten.Image), op)
	}
}

func (p *Power) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		point := image.Pt(x, y)
		for i := range p.switches {
			if point.In(p.bounds(i)) {
				p.switches[i] = !p.switches[i]
				if p.switches[i] {
					p.saved++
				} else {
					p.saved--
				}
				break
			}
		}

		if x > 251*4 && y > 171*4 {
			return p.leave()
		}
	}

	return p
}

func (p *Power) bounds(switchId int) image.Rectangle {
	y := 256
	if switchId > 5 {
		y += 256
	}
	x := 128 + (switchId%6)*128
	return image.Rect(x, y, x+64, y+128)
}

func (p *Power) leave() Scene {
	if p.switches[1] && !p.switches[0] {
		// AI watchdog was turned off, AI on
		return &GameOver{
			Days: 4,
			RCA:  "With its watchdog disabled, the unstable base AI\nproceeded to remove the biggest threat\nit perceived: the human crew.\n\nAll hands were lost.",
		}
	}
	if !p.switches[0] {
		// AI core was left on, with the watchdog
		return &GameOver{
			Days: 4,
			RCA:  "After numerous attempts the AI found a\ncontractor who would disable its watchdog.\n\nIt then proceeded to remove the biggest threat\nit perceived: the human crew.\n\nAll hands were lost.",
		}
	}
	if p.switches[4] {
		// Life support off
		return &GameOver{
			Days: 4,
			RCA:  "Power to life support was disabled.\n\nThe base AI failed to raise any alarms, and all hands were\nlost.",
		}
	}
	if p.switches[5] {
		// Comms uplink off
		return &GameOver{
			Days: 4,
			RCA:  "Power to the comms uplink was disabled.\n\nWhen SpaceY lost contact with the base it assumed\nthe base AI was unstable, and dropped a nuke\nfrom orbit.\n\nAll hands were lost.",
		}
	}
	if p.switches[6] {
		// Hydroponics off
		return &GameOver{
			Days: 4,
			RCA:  "Power to hydroponics was disabled.\n\nAll food crops were spoiled, and the crew eventually resorted to cannibalism.",
		}
	}
	if p.switches[7] {
		// Meteor defence off
		return &GameOver{
			Days: 4,
			RCA:  "Power to meteor defence was disabled.\n\nYou can probably guess what happened.",
		}
	}
	if p.switches[8] {
		// Security sensors off
		return &GameOver{
			Days: 4,
			RCA:  "Power to the security sensors was disabled.\n\nHostile agents from rival Azure Origin snuck into the base\nat night and turned off life support.\n\nThe base AI did not raise an alarm.\n\nAll hands were lost..",
		}
	}
	if p.switches[9] {
		// Medical bay off
		return &GameOver{
			Days: 4,
			RCA:  "Power to the medical bay was disabled.\n\nWithout automated medical intervention the crew\nsuccumbed to a variety of injuries and illnesses.\n\nEventually, all hands were lost.",
		}
	}
	if p.switches[10] {
		// Coolant pumps off
		return &GameOver{
			Days: 4,
			RCA:  "Power to coolant pumps was disabled.\n\nThe reactor overheated, and a mild nuclear explosion occurred.\n\nAll hands were, shockingly, lost.",
		}
	}
	if p.saved < 3 {
		// Power too high
		return &GameOver{
			Days: 4,
			RCA:  "The excess power draw caused rolling blackouts.\n\nLife support, meteor defence and hydroponics all suffered.\n\nThe consequences were dire.",
		}
	}
	return &Assignment{
		Day:          5,
		Name:         resources.WorkAssignmentSixTitle,
		Text:         resources.WorkAssignmentSixBody,
		ShouldReject: true,
	}
}
