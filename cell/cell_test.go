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
        cell.AddNeighbor(neighbor)
    }
    if len(cell.neighbors) != neighborCount  {
        t.Errorf("Cell should have %d neighbors but has %d", neighborCount, len(cell.neighbors))
    }
}

func TestCellNeighborHasCellAsNeighbor(t *testing.T) {
    cell := NewCell()
    neighbor :=NewCell()
    cell.AddNeighbor(neighbor)

    if len(neighbor.neighbors) != 1  {
        t.Errorf("Neighbor should have one neighbor if cell adds it as a neighbor")
    }
}

func TestCellsAreNotNeighborsByDefault(t *testing.T) {
    cell := NewCell()
    neighbor :=NewCell()

    if cell.HasNeighbor(neighbor) != false && neighbor.HasNeighbor(cell) != false {
        t.Errorf("Cells should not be neighbors unless added")
    }
}

func TestCellCanIdentifyANeighbor(t *testing.T) {
    cell := NewCell()
    neighbor :=NewCell()
    cell.AddNeighbor(neighbor)

    if cell.HasNeighbor(neighbor) != true && neighbor.HasNeighbor(cell) != true {
        t.Errorf("Cells should be neighbors if added")
    }
}

func TestLiveCellWithNoLiveNeighborDies(t *testing.T) {
    helper(t, 0, true, false)
}

func TestLiveCellWithOneLiveNeighborDies(t *testing.T) {
    helper(t, 1, true, false)
}

func TestLiveCellWithTwoLiveNeighborsLives(t *testing.T) {
    helper(t, 2, true, true)
}

func TestLiveCellWithThreeLiveNeighborsLives(t *testing.T) {
    helper(t, 3, true, true)
}

func TestLiveCellWithFourLiveNeighborsDies(t *testing.T) {
    helper(t, 4, true, false)
}

func TestLiveCellWithMoreThanThreeLiveNeighborsDies(t *testing.T) {}

func TestDeadCellWithNoLiveNeighborsStaysDead(t *testing.T) {
    helper(t, 0, false, false)
}

func TestDeadCellWithOneLiveNeighborStaysDead(t *testing.T) {
    helper(t, 1, false, false)
}

func TestDeadCellWithTwoLiveNeighborsStaysDead(t *testing.T) {
    helper(t, 2, false, false)
}

func TestDeadCellWithThreeLiveNeighborsResurrects(t *testing.T) {
    helper(t, 3, false, true)
}

func TestDeadCellWithMoreThanThreeLiveNeighborsStaysDead(t *testing.T) {
    helper(t, 4, false, false)
}

func helper(t *testing.T, neighbors int, isAlive bool, shouldLive bool) {
    cell := NewCell()
    if isAlive {cell.resurrect()}
    var neighborCount int
    for i := 0; i < neighbors;  i++ {
        neighborCount = i + 1
        neighbor :=NewCell()
        cell.AddNeighbor(neighbor)
        neighbor.resurrect()
    }

    cell.life <- true
    <- cell.pulse

    var originalStatus string

    if isAlive {
        originalStatus = "Live"
    } else {
        originalStatus =  "Dead"
    }

    var finalStatus string

    if shouldLive {
        finalStatus = "alive"
    } else {
        finalStatus = "dead"
    }

    if cell.IsAlive != shouldLive {
        t.Errorf("%s cell has %d neighbors so should be %s", originalStatus, neighborCount, finalStatus)
    }
}






