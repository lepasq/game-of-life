package game

import (
	"fmt"
	"time"
)

const (
	width  = 20
	height = 15
	alive  = 1
	dead   = 0
)

type World struct {
	Cells [width][height]int
	next  [width][height]int
	Game  Game
}

type Cell struct {
	i   int
	j   int
	val int
}

func Height() int {
	return height
}

func Width() int {
	return width
}

func (w *World) StartGame() {
	w.generateWorld()

	// go Start()
	w.gameLoop()
}

func (w *World) generateWorld() {
	w.Cells[2][1] = alive
	w.Cells[2][2] = alive
	w.Cells[2][3] = alive

	w.Cells[4][4] = alive
	w.Cells[4][5] = alive
	w.Cells[5][4] = alive
	w.Cells[5][5] = alive

	w.Cells[6][6] = alive
	w.Cells[6][7] = alive
	w.Cells[7][6] = alive
	w.Cells[7][7] = alive

	w.next = w.Cells
	w.printCells()
}

func (w *World) printCells() {
	for x, v := range w.Cells {
		fmt.Println(x, v)
	}
	fmt.Println()
}

func (w *World) gameLoop() {
	for {
		w.runAllCells()
		w.printCells()
		time.Sleep(time.Second)
	}
}

func (w *World) runAllCells() {
	for i, v := range w.Cells {
		for j, v := range v {
			w.updateCell(i, j)
			newValue := w.next[i][j]
			if v != newValue {
				Updates <- Cell{i, j, newValue}
			}
		}
	}
	w.Cells = w.next
}

func (w *World) updateCell(row, col int) {
	counter := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if row == 0 || col == 0 || row == width-1 || col == height-1 {
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
