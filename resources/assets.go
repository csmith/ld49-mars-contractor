package resources

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

//go:embed *.png
var assets embed.FS

func LoadImageAsset(path string) (*ebiten.Image, image.Image, error) {
	file, err := assets.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		_ = file.Close()
	}()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, nil, err
	}
	img2 := ebiten.NewImageFromImage(img)
	return img2, img, err
}
