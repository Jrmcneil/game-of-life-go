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
    neighbor.neighbors = append(neighbor.neighbors, cell)
}

func (cell *Cell) Live() {
    var liveNeighbors int

    for _, c := range cell.neighbors {
        if c.IsAlive {
            liveNeighbors++
        }
    }

    if cell.IsAlive && (liveNeighbors <= 1 || liveNeighbors > 3) {
        cell.IsAlive = false
    }

    if cell.IsAlive == false && liveNeighbors == 3 {
        cell.resurrect()
    }
}

func NewCell() Cell {
	return Cell{false, []*Cell {}}
}

