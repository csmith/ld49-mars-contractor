package scenes

import (
	"fmt"
	"github.com/csmith/mars-contractor/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"

	_ "image/png"
)

var pumpBackground *ebiten.Image

func init() {
	var err error
	pumpBackground, _, err = resources.LoadImageAsset("pumpcontrol.png")
	if err != nil {
		panic(err)
	}
}

const (
	pipeScale = 3
	pipeSize  = 32
)

var pipeSheet *resources.Sheet

func init() {
	var err error
	pipeSheet, err = resources.NewSheet("pipes.png", pipeSize, pipeSize)
	if err != nil {
		panic(err)
	}
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type StraightPipe struct {
	Horizontal bool
}

func (s *StraightPipe) Connections() []Direction {
	if s.Horizontal {
		return []Direction{Left, Right}
	} else {
		return []Direction{Up, Down}
	}
}

func (s *StraightPipe) Rotate() {
	s.Horizontal = !s.Horizontal
}

func (s *StraightPipe) Sprite() int {
	if s.Horizontal {
		return 0
	} else {
		return 1
	}
}

type TeePipe struct {
	Direction Direction
}

func (t *TeePipe) Connections() []Direction {
	return []Direction{
		(t.Direction + 3) % 4,
		t.Direction,
		(t.Direction + 1) % 4,
	}
}

func (t *TeePipe) Rotate() {
	t.Direction = (t.Direction + 1) % 4
}

func (t *TeePipe) Sprite() int {
	return int(t.Direction) + 2
}

type CornerPipe struct {
	Direction Direction
}

func (c *CornerPipe) Connections() []Direction {
	return []Direction{
		c.Direction,
		(c.Direction + 1) % 4,
	}
}

func (c *CornerPipe) Rotate() {
	c.Direction = (c.Direction + 1) % 4
}

func (c *CornerPipe) Sprite() int {
	return int(c.Direction) + 6
}

type Pipe interface {
	Connections() []Direction
	Rotate()
	Sprite() int
}

type PipesGame struct {
	pipes         [4][6]Pipe
	filled        [4][6]bool
	complete      bool
	completeTicks int
}

func (p *PipesGame) Init() {
	p.pipes = [4][6]Pipe{
		{
			&CornerPipe{Direction: Left},
			&StraightPipe{Horizontal: false},
			&StraightPipe{Horizontal: true},
			&StraightPipe{Horizontal: true},
			&TeePipe{Direction: Down},
			&StraightPipe{Horizontal: false},
		},
		{
			&CornerPipe{Direction: Left},
			&CornerPipe{Direction: Left},
			&CornerPipe{Direction: Left},
			&StraightPipe{Horizontal: true},
			&TeePipe{Direction: Down},
			&StraightPipe{Horizontal: false},
		},
		{
			&CornerPipe{Direction: Left},
			&StraightPipe{Horizontal: false},
			&TeePipe{Direction: Left},
			&StraightPipe{Horizontal: true},
			&StraightPipe{Horizontal: false},
			&StraightPipe{Horizontal: false},
		},
		{
			&CornerPipe{Direction: Left},
			&CornerPipe{Direction: Left},
			&CornerPipe{Direction: Left},
			&StraightPipe{Horizontal: false},
			&StraightPipe{Horizontal: true},
			&StraightPipe{Horizontal: false},
		},
	}
}

func (p *PipesGame) Draw(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	selectedX := (mouseX - 400) / (pipeSize * pipeScale)
	selectedY := (mouseY - 150) / (pipeSize * pipeScale)

	screen.Fill(colornames.Black)
	screen.DrawImage(pumpBackground, nil)

	for y := range p.pipes {
		for x := range p.pipes[y] {
			pipe := p.pipes[y][x]
			op := &ebiten.DrawImageOptions{}
			if selectedX == x && selectedY == y && !p.complete {
				op.ColorM.Scale(1, 0, 0, 1)
			} else if p.filled[y][x] {
				op.ColorM.Scale(0, 0, 1, 1)
			}
			op.GeoM.Scale(pipeScale, pipeScale)
			op.GeoM.Translate(float64(400+x*(pipeSize*pipeScale)), float64(150+y*(pipeSize*pipeScale)))
			screen.DrawImage(pipeSheet.Sprite(pipe.Sprite()), op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(0.3, 0.3, 0.3, 1)
	op.GeoM.Scale(pipeScale, pipeScale)
	op.GeoM.Translate(float64(400-(pipeSize*pipeScale)), float64(150+2*(pipeSize*pipeScale)))
	screen.DrawImage(pipeSheet.Sprite(0), op)

	for i := 0; i < 4; i++ {
		op.GeoM.Reset()
		op.GeoM.Scale(pipeScale, pipeScale)
		op.GeoM.Translate(float64(400+6*(pipeSize*pipeScale)), float64(150+i*(pipeSize*pipeScale)))
		screen.DrawImage(pipeSheet.Sprite(0), op)
	}
}

func (p *PipesGame) Update() Scene {
	_, dy := ebiten.Wheel()
	if !p.complete && (dy != 0 || inpututil.IsKeyJustPressed(ebiten.KeyR)) {
		mouseX, mouseY := ebiten.CursorPosition()
		selectedX := (mouseX - 400) / (pipeSize * pipeScale)
		selectedY := (mouseY - 150) / (pipeSize * pipeScale)
		if selectedX >= 0 && selectedX < 6 && selectedY >= 0 && selectedY < 4 {
			// TOOD: Allow backwards rotation
			p.pipes[selectedY][selectedX].Rotate()
			p.calculateFlow()
			p.checkVictory()
		}
	}

	if p.complete {
		p.completeTicks++
		if p.completeTicks > 120 {
			return &Assignment{
				Day:          2,
				Name:         "",
				Text:         "",
				ShouldReject: false,
			}
		}
	}
	return p
}

func (p *PipesGame) calculateFlow() {
	for y := range p.pipes {
		for x := range p.pipes[y] {
			p.filled[y][x] = false
		}
	}

	p.checkFlow(2, 0, Left)
}

func (p *PipesGame) checkFlow(y, x int, from Direction) {
	if !p.connected(p.pipes[y][x], from) {
		return
	}

	p.filled[y][x] = true

	connections := p.pipes[y][x].Connections()
	for i := range connections {
		if connections[i] != from {
			nextY, nextX, ok := p.pipeFrom(y, x, connections[i])
			if ok && !p.filled[nextY][nextX] {
				p.checkFlow(nextY, nextX, (connections[i]+2)%4)
			}
		}
	}
}

func (p *PipesGame) checkVictory() {
	for i := 0; i < 4; i++ {
		if !p.filled[i][5] || !p.connected(p.pipes[i][5], Right) {
			return
		}
	}
	p.complete = true
}

func (p *PipesGame) pipeFrom(y int, x int, direction Direction) (int, int, bool) {
	switch direction {
	case Up:
		return y - 1, x, y > 0
	case Right:
		return y, x + 1, x < 5
	case Down:
		return y + 1, x, y < 3
	case Left:
		return y, x - 1, x > 0
	}

	panic(fmt.Sprintf("Unknown direction: %d", direction))
}

func (p *PipesGame) connected(pipe Pipe, d Direction) bool {
	connections := pipe.Connections()
	for i := range connections {
		if connections[i] == d {
			return true
		}
	}
	return false
}
