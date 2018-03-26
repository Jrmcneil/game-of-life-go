package grid

import (
    "game-of-life-go/cell"
)

type Grid struct {
    Surface [][] *cell.Cell
    cells int
    width int
    height int
}

func NewGrid(width, height int) Grid {
    surface := make([][]*cell.Cell, width)
    for i := range surface {
        surface[i] = make([]*cell.Cell, height)
    }

    return Grid{Surface: surface, width: width, height: width}
}

func (grid *Grid) CellCount() int {
    return grid.cells
}

func (grid *Grid) AddCell(cell *cell.Cell) {
    width := grid.cells % grid.width
    height := grid.cells / grid.width

    if width > 0 {
        grid.Surface[width - 1][height].AddNeighbor(cell)
    }

    if height > 0 {
        grid.Surface[width][height - 1].AddNeighbor(cell)
    }

    grid.Surface[width][height] = cell
    grid.cells++
}



