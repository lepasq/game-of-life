package game

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winWidth   = 640
	winHeight  = 480
	colAmount  = 80
	rowAmount  = 60
	cellWidth  = winWidth / colAmount
	cellHeight = winHeight / rowAmount
	alive      = 1
	dead       = 0
	frequency  = 12
)

type Game struct {
	world  *World
	pixels []byte
}

type World struct {
	Cells [colAmount][rowAmount]int
	next  [colAmount][rowAmount]int
}

// type Cell struct {
// 	i   int
// 	j   int
// 	val int
// }

// var Updates chan Cell = make(chan Cell)

func (w *World) generateWorld() {
	rand.Seed(time.Now().UnixNano())
	for i, rows := range w.Cells {
		for j := range rows {
			n := rand.Intn(7)
			if n == 1 {
				w.Cells[i][j] = alive
			}
		}
	}
	w.next = w.Cells
}

func (w *World) printCells() {
	for x, v := range w.Cells {
		fmt.Println(x, v)
	}
	fmt.Println()
}

func (w *World) runAllCells() {
	for i, rows := range w.Cells {
		for j := range rows {
			w.updateCell(i, j)
			// newValue := w.next[i][j]
			// if v != newValue {
			// Updates <- Cell{i, j, newValue}
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

func (w *World) update() {
	now := time.Now()
	w.runAllCells()
	time.Sleep(time.Second/frequency - time.Since(now))
}

func (g *Game) Update() error {
	g.world.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	now := time.Now()
	defer time.Sleep(time.Second/frequency - time.Since(now))

	// for cell := range Updates {
	// 	g.UpdateCell(screen, cell.i, cell.j, cell.val)
	// }

	for i, row := range g.world.Cells {
		for j, value := range row {
			g.drawCell(screen, i, j, value)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return winWidth, winHeight
}

func (g *Game) drawCell(screen *ebiten.Image, i, j, value int) {
	if value == alive {
		cellColor := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
		ebitenutil.DrawRect(screen, float64(cellWidth*i), float64(cellHeight*j), cellWidth, cellHeight, cellColor)
	}
}

func Start() {
	g := &Game{
		world: &World{},
	}

	g.world.generateWorld()
	ebiten.SetWindowSize(winWidth*2, winHeight*2)
	ebiten.SetWindowTitle("Game Of Life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
