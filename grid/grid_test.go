package grid

import (
    "testing"
)

func TestGridIsEmptyByDefault(t *testing.T) {
    grid := Grid{}

    if len(grid.Cells) != 0 {
        t.Errorf("Grid's list should be empty")
    }
}
