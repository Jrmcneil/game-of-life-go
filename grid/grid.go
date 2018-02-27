package grid

type Grid struct {
    cells int
}

func (grid *Grid) CellCount() int {
    return grid.cells
}
