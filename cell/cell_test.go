package cell

import (
    "testing"
    "math/rand"
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

    if len(cell.neighbors) != 0 {
        t.Errorf("Cell should have no neighbors by default")
    }
}

func TestCellCanHaveNeighbors(t *testing.T) {
    cell := NewCell()
    var neighborCount int
    for i := 0; i < rand.Intn(10);  i++ {
        neighborCount = i + 1
        neighbor :=NewCell()
        cell.AddNeighbor(&neighbor)
    }

    if len(cell.neighbors) != neighborCount  {
        t.Errorf("Cell should have %d neighbors but has %d", neighborCount, len(cell.neighbors))
    }
}

func TestCellWithNoLiveNeighborDies(t *testing.T) {
    cell := NewCell()
    cell.resurrect()
    neighbor := NewCell()
    cell.AddNeighbor(&neighbor)

    cell.Live()

    if cell.IsAlive != false {
        t.Errorf("Cell has 0 neighbor so should be dead")
    }
}






