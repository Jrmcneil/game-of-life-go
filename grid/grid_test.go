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

func TestCellCountIncreasesByOneByAddingACell(t *testing.T) {
    grid := Grid{}
    newCell := cell.NewCell()

    grid.AddCell(&newCell)

    if grid.CellCount() != 1 {
        t.Errorf("Grid cell count should increase by one")
    }
}

func TestCellCountIncreasesByTwoByAddingTwoCells(t *testing.T) {
    grid := Grid{}
    newCell := cell.NewCell()
    newCell2 := cell.NewCell()

    grid.AddCell(&newCell)
    grid.AddCell(&newCell2)

    if grid.CellCount() != 2 {
        t.Errorf("Grid cell count should increase by two")
    }
}

func TestSeedingAGridWithCellsAssignsTheirNeighbors(t *testing.T) {}



