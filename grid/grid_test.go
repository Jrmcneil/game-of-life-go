package grid

import (
    "testing"
    "game-of-life-go/cell"
)

func TestGridIsEmptyByDefault(t *testing.T) {
    grid := Grid{}

    if grid.CellCount() != 0 {
        t.Errorf("Grid's list should be empty")
    }
}

func TestCellCountIncreasesByAddingACell(t *testing.T) {
    grid := Grid{}
    cell := cell.NewCell()

    grid.AddCell(&cell)

    if grid.CellCount() != 1 {
        t.Errorf("Grid cell count should increase by one")
    }
}



