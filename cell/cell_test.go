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

func TestLiveCellWithNoLiveNeighborDies(t *testing.T) {
    cell := NewCell()
    cell.resurrect()
    neighbor := NewCell()
    cell.AddNeighbor(&neighbor)

    cell.Live()

    if cell.IsAlive != false {
        t.Errorf("Live cell has 0 neighbor so should be dead")
    }
}

func TestLiveCellWithOneLiveNeighborDies(t *testing.T) {
    cell := NewCell()
    cell.resurrect()
    neighbor := NewCell()
    cell.AddNeighbor(&neighbor)
    neighbor.resurrect()

    cell.Live()

    if cell.IsAlive != false {
        t.Errorf("Live cell has 1 neighbor so should be dead")
    }
}

func TestLiveCellWithTwoLiveNeighborsLives(t *testing.T) {
    cell := NewCell()
    cell.resurrect()
    neighbor1 := NewCell()
    neighbor2 := NewCell()
    cell.AddNeighbor(&neighbor1)
    cell.AddNeighbor(&neighbor2)
    neighbor1.resurrect()
    neighbor2.resurrect()

    cell.Live()

    if cell.IsAlive != true {
        t.Errorf("Live cell has 2 neighbor so should be alive")
    }
}

func TestLiveCellWithThreeLiveNeighborsLives(t *testing.T) {
    cell := NewCell()
    cell.resurrect()
    neighbor1 := NewCell()
    neighbor2 := NewCell()
    neighbor3 := NewCell()
    cell.AddNeighbor(&neighbor1)
    cell.AddNeighbor(&neighbor2)
    cell.AddNeighbor(&neighbor3)
    neighbor1.resurrect()
    neighbor2.resurrect()
    neighbor3.resurrect()

    cell.Live()

    if cell.IsAlive != true {
        t.Errorf("Live cell has 3 neighbor so should be alive")
    }
}

func TestLiveCellWithFourLiveNeighborsDies(t *testing.T) {
    cell := NewCell()
    cell.resurrect()
    var neighborCount int
    for i := 0; i < 4;  i++ {
        neighborCount = i + 1
        neighbor :=NewCell()
        cell.AddNeighbor(&neighbor)
        neighbor.resurrect()
    }

    cell.Live()

    if cell.IsAlive != false {
        t.Errorf("Live cell has %d neighbors so should be dead", neighborCount)
    }
}






