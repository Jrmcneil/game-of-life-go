package grid

import (
    "game-of-life-go/cell"
)

type Grid struct {
    Surface [][]*cell.Cell
    List    []*cell.Cell
    width   int
    height  int
}

func (grid *Grid) CellCount() int {
    return len(grid.List)
}

func (grid *Grid) GetHeight() int {
    return grid.height
}

func (grid *Grid) GetWidth() int {
    return grid.width
}

func NewGrid(width, height int) *Grid {
    surface := make([][]*cell.Cell, width)
    for i := range surface {
        surface[i] = make([]*cell.Cell, height)
    }

    list := make([]*cell.Cell, 0, width * height)
    grid := new(Grid)
    grid.Surface = surface
    grid.List = list
    grid.height = height
    grid.width = width

    return grid
}

func (grid *Grid) AddCell(cell *cell.Cell) {
    width := grid.CellCount() % grid.width
    height := grid.CellCount() / grid.width

    if width > 0 {
        grid.Surface[width - 1][height].AddNeighbor(cell)
    }

    if height > 0 {
        grid.Surface[width][height - 1].AddNeighbor(cell)
    }

    grid.Surface[width][height] = cell
    grid.List = append(grid.List, cell)
}

func (grid *Grid) Fill() {
    cells := grid.width * grid.height

    for i := 0; i < cells; i++ {
        grid.AddCell(cell.NewCell())
    }
}



