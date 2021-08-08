package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winWidth   = 640
	winHeight  = 480
	cellWidth  = 32
	cellHeight = 32
)

type Game struct {
}

var Updates chan Cell = make(chan Cell)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for cell := range Updates {
		g.UpdateCell(screen, cell.i, cell.j, cell.val)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) UpdateCell(screen *ebiten.Image, i, j, value int) {
	var cellColor color.Color
	if value == 1 {
		cellColor = color.RGBA{0xFF, 0, 0, 0xFF}
	} else {
		cellColor = color.RGBA{0, 0xFF, 0, 0xFF}
	}
	ebitenutil.DrawRect(screen, cellWidth*1, cellHeight*1, cellWidth, cellHeight, cellColor)
}

func Start() {
	w := World{}

	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle("Hello, World!")

	go w.StartGame()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
