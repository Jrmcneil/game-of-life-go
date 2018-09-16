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

    newgrid.AddCell(newcell)

    if newgrid.Surface[0][0] != newcell {
        t.Errorf("Grid surface should contain the cell")
    }
}


func TestCellCountIncreasesByOneByAddingACell(t *testing.T) {
    newgrid := NewGrid(1,1)
    newCell := cell.NewCell()

    newgrid.AddCell(newCell)

    if newgrid.CellCount() != 1 {
        t.Errorf("Grid cell count should increase by one")
    }
}

func TestCellCountIncreasesByTwoByAddingTwoCells(t *testing.T) {
    newgrid := NewGrid(2,1)
    newCell := cell.NewCell()
    newCell2 := cell.NewCell()

    newgrid.AddCell(newCell)
    newgrid.AddCell(newCell2)

    if newgrid.CellCount() != 2 {
        t.Errorf("Grid cell count should increase by two")
    }
}

func TestCellsAreAddedAcrossThenDownTheSurface(t *testing.T) {
    newcell1 := cell.NewCell()
    newcell2 := cell.NewCell()
    newcell3 := cell.NewCell()
    newcell4 := cell.NewCell()
    newgrid := NewGrid(2,2)

    newgrid.AddCell(newcell1)
    newgrid.AddCell(newcell2)
    newgrid.AddCell(newcell3)
    newgrid.AddCell(newcell4)

    pos1 := newgrid.Surface[0][0] != newcell1
    pos2 := newgrid.Surface[1][0] != newcell2
    pos3 := newgrid.Surface[0][1] != newcell3
    pos4 := newgrid.Surface[1][1] != newcell4

    if pos1 || pos2 || pos3 || pos4 {
        t.Errorf("Grid surface should contain the cells in correct order")
    }
}

func TestAddingCellsNextToEachOtherAssignsThemAsNeighbors(t *testing.T) {
    newgrid := NewGrid(3,1)
    newCell := cell.NewCell()
    newCell2 := cell.NewCell()
    newCell3 := cell.NewCell()

    newgrid.AddCell(newCell)
    newgrid.AddCell(newCell2)
    newgrid.AddCell(newCell3)

    if !newCell.HasNeighbor(newCell2) || !newCell2.HasNeighbor(newCell) {
        t.Errorf("First and second cells should be neighbors")
    }

    if !newCell2.HasNeighbor(newCell3) || !newCell2.HasNeighbor(newCell3) {
        t.Errorf("Second and third cells should be neighbors")
    }

    if newCell.HasNeighbor(newCell3) || newCell3.HasNeighbor(newCell) {
        t.Errorf("First and third cells should be not neighbors")
    }
}

func TestAddingCellsAboveEachOtherAssignsThemAsNeighbors(t *testing.T) {
    newgrid := NewGrid(1,3)
    newCell := cell.NewCell()
    newCell2 := cell.NewCell()
    newCell3 := cell.NewCell()

    newgrid.AddCell(newCell)
    newgrid.AddCell(newCell2)
    newgrid.AddCell(newCell3)

    if !newCell.HasNeighbor(newCell2) || !newCell2.HasNeighbor(newCell) {
        t.Errorf("First and second cells should be neighbors")
    }

    if !newCell2.HasNeighbor(newCell3) || !newCell2.HasNeighbor(newCell3) {
        t.Errorf("Second and third cells should be neighbors")
    }

    if newCell.HasNeighbor(newCell3) || newCell3.HasNeighbor(newCell) {
        t.Errorf("First and third cells should be not neighbors")
    }
}

func TestCellsAddedDiagonalToEachOtherAreNotNeighbors(t *testing.T) {
    newgrid := NewGrid(2,2)
    newCell := cell.NewCell()
    newCell2 := cell.NewCell()
    newCell3 := cell.NewCell()
    newCell4 := cell.NewCell()

    newgrid.AddCell(newCell)
    newgrid.AddCell(newCell2)
    newgrid.AddCell(newCell3)
    newgrid.AddCell(newCell4)

    if newCell.HasNeighbor(newCell4) || newCell4.HasNeighbor(newCell) || newCell2.HasNeighbor(newCell3) || newCell3.HasNeighbor(newCell2) {
        t.Errorf("Diagonal cells should not be neighbors")
    }
}

func TestFillingGridWithCells(t *testing.T) {
    newgrid := NewGrid(10, 10)
    newgrid.Fill()

    if newgrid.GetCells() != 100 {
        t.Errorf("Grid should have 100 cells")
    }
}

