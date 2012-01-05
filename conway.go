package conway

import (
	"rand"
   "bytes"
   "fmt"
)

type ConwaySim struct {
	Cells [][]bool
	Rows  int
	Cols  int
}

func MakeConwaySim(rows int, cols int) (b *ConwaySim) {
	b = new(ConwaySim)
	b.Rows = rows
	b.Cols = cols

	b.Cells = make([][]bool, rows)

	for i := 0; i < rows; i++ {
		b.Cells[i] = make([]bool, cols)
	}
	return
}

func (b *ConwaySim) Randomize(seed int64) {
   rand.Seed(seed)
	for i := 0; i < b.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			b.Cells[i][j] = (rand.Intn(2) == 0)
		}
	}
	return
}

func (b *ConwaySim) String() string {
   buffer := bytes.NewBufferString("")
   for i := 0; i < b.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			if b.Cells[i][j] {
            fmt.Fprintf(buffer, "X")
			} else {
            fmt.Fprintf(buffer, " ")
			}
		}
      fmt.Fprintf(buffer, "\n")
	}
   return string(buffer.Bytes())
}

func (b *ConwaySim) Step() {
	var newCells [][]bool

	newCells = make([][]bool, b.Rows)
	for i := 0; i < b.Rows; i++ {
		newCells[i] = make([]bool, b.Cols)
		for j := 0; j < b.Cols; j++ {
			newCells[i][j] = b.nextState(i, j)
		}
	}

	b.Cells = newCells
}

func (b *ConwaySim) nextState(row int, col int) (next bool) {
	var neighbors = b.cntneighbors(row, col)

	if b.Cells[row][col] { // Cell was alive
		if neighbors < 2 {
			next = false // Starvation
		} else if neighbors > 3 {
			next = false // Overpopulation
		}
	} else { // Cell was dead
		if neighbors == 3 {
			next = true // Reproduction
		}
	}
	return
}

func (b *ConwaySim) cntneighbors(row int, col int) (neighbors int) {
	neighbors = 0
	// Top Row
	if row > 0 {
		if col > 0 {
			if b.Cells[row-1][col-1] {
				neighbors++
			}
		}
		if b.Cells[row-1][col] {
			neighbors++
		}
		if col < b.Cols-1 {
			if b.Cells[row-1][col+1] {
				neighbors++
			}
		}
	}

	// Middle Row
	if col > 0 {
		if b.Cells[row][col-1] {
			neighbors++
		}
	}
	if col < b.Cols-1 {
		if b.Cells[row][col+1] {
			neighbors++
		}
	}

	// Bottom Row
	if row < b.Rows-1 {
		if col > 0 {
			if b.Cells[row+1][col-1] {
				neighbors++
			}
		}
		if b.Cells[row+1][col] {
			neighbors++
		}
		if col < b.Cols-1 {
			if b.Cells[row+1][col+1] {
				neighbors++
			}
		}
	}

	return
}
