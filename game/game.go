package game

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winWidth   = 640
	winHeight  = 480
	cellWidth  = 32
	cellHeight = 32
	colAmount  = 20
	rowAmount  = 15
	alive      = 1
	dead       = 0
)

type Game struct {
	world  *World
	pixels []byte
}

type World struct {
	Cells [colAmount][rowAmount]int
	next  [colAmount][rowAmount]int
}

type Cell struct {
	i   int
	j   int
	val int
}

func (w *World) generateWorld() {
	w.Cells[2][1], w.Cells[2][2], w.Cells[2][2], w.Cells[2][3], w.Cells[4][4], w.Cells[4][5],
		w.Cells[5][4], w.Cells[5][5], w.Cells[6][6], w.Cells[6][7], w.Cells[7][6], w.Cells[7][7] =
		alive, alive, alive, alive, alive, alive, alive, alive, alive, alive, alive, alive

	w.next = w.Cells
	w.printCells()
}

func (w *World) printCells() {
	for x, v := range w.Cells {
		fmt.Println(x, v)
	}
	fmt.Println()
}

func (w *World) runAllCells() {
	for i, v := range w.Cells {
		for j := range v {
			w.updateCell(i, j)
			newValue := w.next[i][j]
			// if v != newValue {
			Updates <- Cell{i, j, newValue}
		}
	}
	w.Cells = w.next
}

func (w *World) updateCell(row, col int) {
	counter := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if row == 0 || col == 0 || row == colAmount-1 || col == rowAmount-1 {
				break
			}
			if w.Cells[i][j] == 1 && !(i == row && col == j) { // second condition: middle cell is not counted
				counter += 1
			}
		}
	}
	w.updateCellValue(counter, row, col)
}

func (w *World) updateCellValue(counter, row, col int) {
	if counter == 2 {
		return
	} else if counter == 3 {
		w.next[row][col] = alive
	} else {
		w.next[row][col] = dead
	}
}

var Updates chan Cell = make(chan Cell)

func (g *Game) Update() error {
	for {
		g.world.runAllCells()
		g.world.printCells()
		time.Sleep(time.Second)
	}
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

	g := &Game{
		world: &World{},
	}

	g.world.generateWorld()

	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
