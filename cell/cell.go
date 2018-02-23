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

func (cell *Cell) Live() {
    var liveNeighbors int

    for _, c := range cell.neighbors {
        if c.IsAlive {
            liveNeighbors++
        }
    }

    if liveNeighbors <= 1 {
        cell.IsAlive = false
    }
}

func NewCell() Cell {
	return Cell{false, []*Cell {}}
}

