package cell

type Cell struct {
    IsAlive   bool
    neighbors NeigbhorSet
    life      chan bool
    pulse     chan bool
}

type NeigbhorSet struct {
    set map[*Cell]bool
}

func (neighbors *NeigbhorSet) Add(cell *Cell) bool {
    _, found := neighbors.set[cell]
    neighbors.set[cell] = true
    return !found
}

func (neighbors *NeigbhorSet) Has(cell *Cell) bool {
    _, found := neighbors.set[cell]
    return found
}

func (cell *Cell) resurrect() {
    cell.IsAlive = true
}

func (cell *Cell) kill() {
    cell.IsAlive = false
}

func (cell *Cell) AddNeighbor(neighbor *Cell) {
    cell.neighbors.Add(neighbor)
    neighbor.neighbors.Add(cell)
}

func (cell *Cell) HasNeighbor(neighbor *Cell) bool {
    return cell.neighbors.Has(neighbor)
}

func (cell *Cell) live() {
    var liveNeighbors int

    for c := range cell.neighbors.set {
        if c.IsAlive {
            liveNeighbors++
        }
    }

    switch {
    case cell.IsAlive && (liveNeighbors <= 1 || liveNeighbors > 3):
        cell.kill()
        cell.pulse <- false
    case !cell.IsAlive && liveNeighbors == 3:
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
    cell.neighbors.set = make(map[*Cell]bool)
    go waitForLife(cell)
    return cell
}
