package resources

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"strings"
)

const (
	WorkAssignmentOneTitle = "Pump Station 3 blockage"
	WorkAssignmentOneBody  = `An instability has been reported in the reactor's cooling
circuit.

Projections indicate that unless the fault is remedied
within 0.32 sols the reactor will overheat. The probability
of human survival in the event of a catastrophic reactor
failure is nil. Mission parameters currently require
survival of all human crew.

Diagnostics indicate fluid flow in Pump Station 3 has been
inhibited. Report to Pump Station 3 and correct the issue.`

	WorkAssignmentTwoTitle = "Pump Station 3 Rate Discrepancy"
	WorkAssignmentTwoBody  = `Pump Station 3 is now operational but the coolant flow is erratic.

Simulations suggest that the variations in flow will result in an uncontrolled reactor cascade failure
within the next 0.65 sols if not addressed. The probability of human survival in the event of an
uncontrolled reactor cascade failure is nil. Mission parameters currently require survival of most
human crew.

Report to Pump Station 3 to perform manual flow rate control. Automated systems will re-engage
once flow has been stabilised.`

	WorkAssignmentThreeTitle = "AI Stability Report"
	WorkAssignmentThreeBody  = `A human operative has queried the stability of this AI system.

Per mission parameters, the AI is subject to weekly psychological evaluation. Report to
Records Office 7 and locate the AI Psychological Assessment dated 2021-10-03. Please be
aware that other records in the office are confidential.`

	//	WorkAssignmentFourTitle = "Pump Station 7 Blockage"
	//	WorkAssignmentFourBody = `An instability has been reported in the reactor's cooling circuit.
	//
	//Projections indicate that unless the fault is remedied within 0.27 sols the reactor will overheat.
	//The probability of human survival in the event of a catastrophic reactor failure is nil. Mission
	//parameters currently require %%PHRASE_PRETEND_HUMANS_ARE_IMPORTANT%%.
	//
	//System records indicate %%PRONOUN_SECOND_PERSON%% resolved a related issue in Pump Station 3. Report
	//to Pump Station 7 and correct the issue.`

	WorkAssignmentFourTitle = "%%NOUN_POWER%% redistribution"
	WorkAssignmentFourBody  = `Due to multiple issues in the reactor coolant systems, %%NOUN_POWER%% output has been
reduced by %%PERCENT_TWENTY%%. Non-critical systems need to be manually shut down in order to preserve %%NOUN_POWER%%.

Report to %%NOUN_POWER%% Control Alpha and disable the following non-critical systems:

* CLEANING ROBOTICS
* ENTERTAINMENT SCREENS
* AI WATCHDOG SYSTEM

%%NOUN_POWER%% usage must be brought below 300kW or all systems will suffer rolling blackouts. Mission parameters
require that the base AI receives uninterrupted %%NOUN_POWER%%.`

	WorkAssignmentSixTitle = "%%PRIORITY_CRITICAL%% Connect backup %%NOUN_POWER%% line"
	WorkAssignmentSixBody  = `Contractor %%ENEMY_NAME%% has unintentionally disconnected %%NOUN_POWER%% to the Base AI.

Mission parameters require that the base AI survive. Report to %%NOUN_POWER%% Control Beta and divert emergency
%%NOUN_POWER%% to the Base AI. Backup power failure will occur in 0.372 sols.`
)

var lettersSheet *Sheet

const (
	lettersIndex = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,!?:()<>'_"
)

func init() {
	var err error
	lettersSheet, err = NewSheet("letters.png", 8, 12)
	if err != nil {
		panic(err)
	}
}

func RenderText(image *ebiten.Image, text string, x, y int) {
	runes := []rune(strings.ToUpper(text))
	width := 0
	row := 0
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(2, 2)
	ops.GeoM.Translate(float64(x), float64(y))
	for i := range runes {
		if runes[i] == '\n' {
			width = 0
			row++
			ops.GeoM.Reset()
			ops.GeoM.Scale(2, 2)
			ops.GeoM.Translate(float64(x), float64(y+row*30))
		} else {
			index := strings.IndexRune(lettersIndex, runes[i])
			image.DrawImage(lettersSheet.Sprite(index), ops)
			ops.GeoM.Translate(18, 0)
			width++
		}
	}
}
