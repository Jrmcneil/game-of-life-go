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

func TestCellHasNoNeighborByDefault(t *testing.T) {
    cell := NewCell()

    if cell.neighbors[0] != nil {
        t.Errorf("Cell should have no neighbors by default")
    }
}

func TestCellCanHaveANeighbor(t *testing.T) {
    cell := NewCell()
    neighbor := NewCell()

    cell.AddNeighbor(&neighbor)

    if cell.neighbors[0] != &neighbor {
        t.Errorf("Cell should have neighbor")
    }
}




