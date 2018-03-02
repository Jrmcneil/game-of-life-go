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

func TestAddingACellPutsItOnTheSurface(t *testing.T) {
    newcell := cell.NewCell()
    newgrid := NewGrid(1,1)

    newgrid.AddCell(&newcell)

    if newgrid.Surface[0][0] != &newcell {
        t.Errorf("Grid surface should contain the cell")
    }
}

func TestCellCountIncreasesByOneByAddingACell(t *testing.T) {
    newgrid := NewGrid(1,1)
    newCell := cell.NewCell()

    newgrid.AddCell(&newCell)

    if newgrid.CellCount() != 1 {
        t.Errorf("Grid cell count should increase by one")
    }
}

func TestCellCountIncreasesByTwoByAddingTwoCells(t *testing.T) {
    newgrid := NewGrid(1,1)
    newCell := cell.NewCell()
    newCell2 := cell.NewCell()

    newgrid.AddCell(&newCell)
    newgrid.AddCell(&newCell2)

    if newgrid.CellCount() != 2 {
        t.Errorf("Grid cell count should increase by two")
    }
}

func TestSeedingAGridWithCellsAssignsTheirNeighbors(t *testing.T) {}



