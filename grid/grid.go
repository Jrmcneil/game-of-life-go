package grid

import "game-of-life-go/cell"

type Grid struct {
    cells int
}

func (grid *Grid) CellCount() int {
    return grid.cells
}

func (grid *Grid) AddCell(cell *cell.Cell) {
    grid.cells++
}


