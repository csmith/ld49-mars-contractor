package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"image"
	"math"
)

var rateBackground *ebiten.Image

func init() {
	var err error
	rateBackground, _, err = resources.LoadImageAsset("ratecontrol.png")
	if err != nil {
		panic(err)
	}
}

type RateGame struct {
	rawRate  float64
	rate     float64
	ticks    int
	damage   int
	end      bool
	endTicks int
}

func (r *RateGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(rateBackground, nil)
	bar := screen.SubImage(image.Rect(800, 130, 850, 530)).(*ebiten.Image)
	bar.Fill(colornames.Gray)

	color := colornames.Blue
	if r.rate < -1.2 || r.rate > 1.2 {
		color = colornames.Red
	} else if r.rate < -0.8 || r.rate > 0.5 {
		color = colornames.Yellow
	}

	bar.SubImage(image.Rect(800, 330+int(r.rate*100), 850, 530)).(*ebiten.Image).Fill(color)
}

func (r *RateGame) Update() Scene {
	if !r.end {
		r.ticks++
		r.rawRate = math.Sin(float64(r.ticks)/120) + math.Pow(math.Sin(float64(r.ticks)/60), 2) - math.Pow(math.Sin(5+float64(r.ticks)/60), 2)
		r.rate += (r.rawRate - r.rate) / 25

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			r.rate -= 0.34
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			r.rate += 0.34
		}

		if r.rate < -1.2 || r.rate > 1.2 {
			r.damage++
			r.end = r.damage > 60
		}

		if r.ticks > 60*45 {
			r.end = true
		}
	} else {
		r.endTicks++
		if r.endTicks > 120 {
			if r.rate < -1.2 || r.rate > 1.2 {
				return &GameOver{
					Days: 2,
					RCA: "Contractor engaged in manual coolant rate\ncontrol failed to keep coolant within\nacceptable parameters.\n\nA reactor cascade occurred.\n\nAll hands were lost.",
				}
			} else {
				return &Assignment{
					Day:       3,
					Name:      resources.WorkAssignmentThreeTitle,
					Text:      resources.WorkAssignmentThreeBody,
				}
			}
		}
	}

	return r
}
