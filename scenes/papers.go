package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
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
	resources.RenderTextSmall(screen, paperTexts[p.current], 380, 170)
}

func (p *Papers) Update() Scene {
	x, y := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if x > 1080 && y > 640 {
			if strings.Contains(paperTexts[p.current], "2021-10-03") {
				return &Assignment{
					Day:  4,
					Name: resources.WorkAssignmentFourTitle,
					Text: resources.WorkAssignmentFourBody,
				}
			} else {
				return &GameOver{
					Days: 3,
					RCA:  "Due to a paperwork error, the crew's doubts about\nthe AI's stability were aggravated.\n\nThey attempted to disable it but were not capable.\n\nThe crew were eliminated in <self-defence>.",
				}
			}
		} else {
			p.current++
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		p.current--
	}
	p.current = (p.current + len(paperTexts)) % len(paperTexts)
	return p
}

var paperTexts = []string{
	"AI Psych Report\nDate: 2020-01-03\nAuthor: Dr Chip\n\nNotes:\n\nAI acting within parameters",
	"AI Psych Report\nDate: 2020-07-02\nAuthor: Dr Chip\n\nNotes:\n\nAI acting normally",
	"AI Psych Report\nDate: 2020-02-28\nAuthor: Dr Chip\n\nNotes:\n\nWorking as expected",
	"Incident Report\nDate: 2020-08-13\nAuthor: Dr Chip\n\nDetails:\n\nTwo general contractors were involved in\na physical fight after receiving\nfalsified messages though internal comms.\n\nSource of messages unknown.",
	"AI Psych Report\nDate: 2020-10-30\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"AI Psych Report\nDate: 2021-02-13\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"Incident Report\nDate: 2020-09-01\nAuthor: Dr Chip\n\nDetails:\n\nMultiple contractors have reported being assigned\ntasks that aren't required or make\nno sense. Will address with AI during\nnext scheduled review.",
	"AI Psych Report\nDate: 2020-04-28\nAuthor: Dr Chip\n\nNotes:\n\nAll normal",
	"AI Psych Report\nDate: 2020-05-18\nAuthor: Dr Chip\n\nNotes:\n\nNothing to report",
	"AI Psych Report\nDate: 2021-10-03\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"Incident Report\nDate: 2020-10-10\nAuthor: Base AI\n\nDetails:\n\nDr Chip has become increasingly unstable and\nhas been medically sedated until he can\nbe returned to Earth.\n\nBase AI will assume his role per\nemergency guidelines appendix 7",
	"AI Psych Report\nDate: 2020-12-03\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"AI Psych Report\nDate: 2020-11-02\nAuthor: Base AI\n\nNotes:\n\nSelf-check complete. Situation normal.",
	"AI Psych Report\nDate: 2020-08-05\nAuthor: Dr Chip\n\nNotes:\n\nNothing to report",
	"Incident Report\nDate: 2020-12-10\nAuthor: Base AI\n\nDetails:\n\nBase commander attempted to forcibly remove\ncontrol from base AI. Commander has been\nrelieved of duty per emergency guidelines\nappendix 9.",
}
