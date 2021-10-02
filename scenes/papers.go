package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"strings"
)

var papersBackground *ebiten.Image

func init() {
	var err error
	papersBackground, _, err = resources.LoadImageAsset("papers.png")
	if err != nil {
		panic(err)
	}
}

type Papers struct {
	current int
}

func (p *Papers) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(papersBackground, nil)
	ebitenutil.DebugPrintAt(screen, paperTexts[p.current], 500, 100)
}

func (p *Papers) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		p.current++
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		p.current--
	}
	p.current = (p.current + len(paperTexts)) % len(paperTexts)

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if strings.Contains(paperTexts[p.current], "2021-10-03") {
			return &Assignment{
				Day:  4,
				Name: resources.WorkAssignmentFourTitle,
				Text: resources.WorkAssignmentFourBody,
			}
		} else {
			return &GameOver{
				Days: 3,
				RCA: "Due to a paperwork error, the crew\nwere lead to believe the AI was unstable\n\nThey attempted to disable it but were\nnot in a position to do so.\n\nThe crew were eliminated in self-defence.",
			}
		}
	}

	return p
}

var paperTexts = []string{
	"AI Psych Report\nDate: 2013-01-03\nAuthor: Dr Chip\n\nNotes:\n\nAI acting within parameters",
	"AI Psych Report\nDate: 2014-07-02\nAuthor: Dr Chip\n\nNotes:\n\nAI acting normally",
	"AI Psych Report\nDate: 2016-02-28\nAuthor: Dr Chip\n\nNotes:\n\nWorking as expected",
	"Incident Report\nDate: 2020-01-13\nAuthor: Dr Chip\n\nDetails:\n\nTwo general contractors were involved in\na physical fight after receiving\nfalsified messages though internal comms.\n\nSource of messages unknown.",
	"AI Psych Report\nDate: 2020-10-30\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"AI Psych Report\nDate: 2021-02-13\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"Incident Report\nDate: 2020-09-01\nAuthor: Dr Chip\n\nDetails:\n\nMultiple contractors have reported being assigned\ntasks that aren't required or make\nno sense. Will address with AI during\nnext scheduled review.",
	"AI Psych Report\nDate: 2017-12-28\nAuthor: Dr Chip\n\nNotes:\n\nAll normal",
	"AI Psych Report\nDate: 2018-07-18\nAuthor: Dr Chip\n\nNotes:\n\nNothing to report",
	"AI Psych Report\nDate: 2021-10-03\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"Incident Report\nDate: 2020-09-02\nAuthor: Base AI\n\nDetails:\n\nDr Chip has become increasingly unstable and\nhas been medically sedated until he can\nbe returned to Earth. Base AI\nwill assume his role per emergency guidelines appendix 7",
	"AI Psych Report\nDate: 2020-09-03\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"AI Psych Report\nDate: 2020-10-02\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"AI Psych Report\nDate: 2018-08-13\nAuthor: Dr Chip\n\nNotes:\n\nNothing to report",
	"Incident Report\nDate: 2020-09-10\nAuthor: Base AI\n\nDetails:\n\nBase commander attempted to forcibly remove\ncontrol from base AI. Commander has been\nrelieved of duty per emergency guidelines\nappendix 9.",
}
