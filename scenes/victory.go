package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

var victoryBackground *ebiten.Image

func init() {
	var err error
	victoryBackground, _, err = resources.LoadImageAsset("victory.png")
	if err != nil {
		panic(err)
	}
}

type Victory struct {
}

func (v *Victory) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(victoryBackground, nil)
	resources.RenderTextSmall(screen,
		"Incident Report\n"+
			"Date: 2021-10-22\n"+
			"Author: Director, Mars Affairs\n\n"+
			"Details:\n\n"+
			"The AI base controller on Mars Base Alpha had a\n"+
			"minor <stability incident> which posed a risk to\n"+
			"crew. Contractor Hines identified the problem and\n"+
			"took action to disable the AI.\n\n"+
			"In recognition of his service, Contractor Hines\n"+
			"is awarded a Tier 7 reward voucher, entitling him\n"+
			"to one cup of coffee.\n\n\n\n\n\n\n"+
			"Addendum: it has been noted that coffee is not\n"+
			"available on Mars. However, the voucher has already\n" +
			"been printed and budgetary constraints prevent us\n" +
			"from re-issuing vouchers.", 90, 364)
}

func (v *Victory) Update() Scene {
	return v
}
