package cell

import "testing"

func TestCellCanBeAlive(t *testing.T) {
	cell := Cell{true}

	if cell.IsAlive != true {
		t.Errorf("Cell is not alive")
	}
}