package resources

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func NewSheet(filename string, width, height int) (*Sheet, error) {
	im, _, err := LoadImageAsset(filename)
	if err != nil {
		return nil, err
	}
	return &Sheet{
		im:     im,
		width:  width,
		height: height,
		cols:   im.Bounds().Dx() / width,
	}, nil
}

type Sheet struct {
	im     *ebiten.Image
	width  int
	height int
	cols   int
}

func (s *Sheet) Sprite(index int) *ebiten.Image {
	x := (index % s.cols) * s.width
	y := (index / s.cols) * s.height
	return s.im.SubImage(image.Rect(x, y, x+s.width, y+s.height)).(*ebiten.Image)
}

