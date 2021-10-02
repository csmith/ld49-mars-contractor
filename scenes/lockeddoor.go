package scenes

import (
	"github.com/csmith/mars-contractor/resources"
	"github.com/csmith/mars-contractor/resources/sounds"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	"image"
)

var doorBackground *ebiten.Image

func init() {
	var err error
	doorBackground, _, err = resources.LoadImageAsset("doorlock.png")
	if err != nil {
		panic(err)
	}
}

type LockedDoor struct {
	entered int
	correct int
}

func (l *LockedDoor) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	screen.DrawImage(doorBackground, nil)

	x, y := ebiten.CursorPosition()
	horizontal := (x - 664) / 104
	vertical := (y - 248) / 120
	if (x-664)%104 <= 92 && (y-248)%116 <= 96 && horizontal >= 0 && vertical >= 0 && horizontal <= 4 && vertical <= 1 {
		// Valid button
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(horizontal*104+664), float64(vertical*120+248))
		op.ColorM.Scale(1, 0.3, 0.3, 1)
		screen.DrawImage(doorBackground.SubImage(image.Rect(horizontal*104+664, vertical*120+248, horizontal*104+664+92, vertical*120+248+96)).(*ebiten.Image), op)
	}
}

func (l *LockedDoor) Update() Scene {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		horizontal := (x - 664) / 104
		vertical := (y - 248) / 120
		if (x-664)%104 <= 92 && (y-248)%116 <= 96 && horizontal >= 0 && vertical >= 0 && horizontal <= 4 && vertical <= 1 {
			sounds.PlayBeep()
			correct := []int{3, 6, 9, 2}
			number := (vertical*5 + horizontal + 1) % 10
			if correct[l.entered] == number {
				l.correct++
			}
			l.entered++
			if l.entered == 4 {
				if l.correct == 4 {
					return &Papers{}
				} else {
					sounds.PlayError()
					l.entered = 0
					l.correct = 0
				}
			}
		}
	}
	return l
}
