package cell

type Cell struct {
    id        int
    IsAlive   bool
    neighbors []*Cell
    life      chan bool
    pulse     chan bool
}

func (cell *Cell) resurrect() {
    cell.IsAlive = true
}

func (cell *Cell) kill() {
    cell.IsAlive = false
}

func (cell *Cell) AddNeighbor(neighbor *Cell) {
    cell.neighbors = append(cell.neighbors, neighbor)
    neighbor.neighbors = append(neighbor.neighbors, cell)
}

func (cell *Cell) HasNeighbor(neighbor *Cell) bool {
    for i := range cell.neighbors {
        if cell.neighbors[i] == neighbor {
            return true
        }
    }

    return false
}

func (cell *Cell) live() {
    var liveNeighbors int

    for _, c := range cell.neighbors {
        if c.IsAlive {
            liveNeighbors++
        }
    }

    switch {
    case cell.IsAlive && (liveNeighbors <= 1 || liveNeighbors > 3):
        cell.kill()
        cell.pulse <- false
    case cell.IsAlive == false && liveNeighbors == 3:
        cell.resurrect()
        cell.pulse <- true
    default:
        cell.pulse <- cell.IsAlive
    }
}

func waitForLife(cell *Cell) {
    for {
        select {
        case <-cell.life:
            cell.live()
        }
    }
}

func NewCell() *Cell {
    cell := new(Cell)
    cell.life = make(chan bool)
    cell.pulse = make(chan bool, 1)
    go waitForLife(cell)
    return cell
}
