package cell

type Cell struct {
	IsAlive bool
	neighbors []*Cell
}

func (cell *Cell) resurrect() {
    cell.IsAlive = true
}

func NewCell() Cell {
	return Cell{false, make([]*Cell, 0)}
}

