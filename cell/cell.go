package cell

type Cell struct {
	IsAlive bool
	neighbors []*Cell
}

func (cell *Cell) resurrect() {
    cell.IsAlive = true
}

func (cell *Cell) AddNeighbor(neighbor *Cell) {
    cell.neighbors = append(cell.neighbors, neighbor)
}

func NewCell() Cell {
	return Cell{false, make([]*Cell, 0)}
}

