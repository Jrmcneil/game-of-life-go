package cell

type Cell struct {
	IsAlive bool
	neighbors [1]*Cell
}

func (cell *Cell) resurrect() {
    cell.IsAlive = true
}

func (cell *Cell) AddNeighbor(neighbor *Cell) {
    cell.neighbors[0] = neighbor
}

func NewCell() Cell {
	return Cell{false, [1]*Cell {}}
}

