package cell

type Cell struct {
    IsAlive   bool
    neighbors NeigbhorSet
    Life      chan bool
    Pulse     chan bool
}

type NeigbhorSet struct {
    set map[*Cell]bool
}

func (neighbors *NeigbhorSet) Add(cell *Cell) bool {
    defer func() {neighbors.set[cell] = true}()
    return !neighbors.Has(cell)
}

func (neighbors *NeigbhorSet) Has(cell *Cell) bool {
    _, found := neighbors.set[cell]
    return found
}

func (cell *Cell) Resurrect() {
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
        cell.Pulse <- false
    case !cell.IsAlive && liveNeighbors == 3:
        cell.Resurrect()
        cell.Pulse <- true
    default:
        cell.Pulse <- cell.IsAlive
    }
}

func waitForLife(cell *Cell) {
    for {
        select {
        case <-cell.Life:
            cell.live()
        }
    }
}

func NewCell() *Cell {
    cell := new(Cell)
    cell.Life = make(chan bool)
    cell.Pulse = make(chan bool, 1)
    cell.neighbors.set = make(map[*Cell]bool)
    go waitForLife(cell)
    return cell
}
