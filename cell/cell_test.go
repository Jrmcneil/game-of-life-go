package cell

import (
    "testing"
)

func TestCellCanBeAlive(t *testing.T) {
	cell := NewCell()

    cell.resurrect()

	if cell.IsAlive != true {
		t.Errorf("Cell is not alive")
	}
}

func TestCellIsNotAliveByDefault(t *testing.T) {
    cell := NewCell()

    if cell.IsAlive != false {
        t.Errorf("Cell is not dead")
    }
}

func TestCellHasNoNeighborsByDefault(t *testing.T) {
    cell := NewCell()

    if len(cell.neighbors) != 0 {
        t.Errorf("Cell should have no neighbors by default")
    }
}

func TestCellCanHaveOneNeighbor(t *testing.T) {
    cell := NewCell()
    neighbor := NewCell()

    cell.AddNeighbor(&neighbor)

    neighbors := len(cell.neighbors)

    if neighbors != 1 {
        t.Errorf("Cell should have 1 neighbor but has %d", neighbors)
    }
}

func TestCellCanHaveTwoNeighbors(t *testing.T) {
    cell := NewCell()
    neighbor := NewCell()
    otherNeighbor := NewCell()

    cell.AddNeighbor(&neighbor)
    cell.AddNeighbor(&otherNeighbor)

    neighbors := len(cell.neighbors)

    if neighbors != 2 {
        t.Errorf("Cell should have 2 neighbor but has %d", neighbors)
    }
}



