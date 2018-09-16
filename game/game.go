package game

import (
    "game-of-life-go/grid"
    "sync"
    "game-of-life-go/cell"
    "fmt"
)

type Game struct {
    grid *grid.Grid
}

func (game *Game) Build(dimension int) *grid.Grid {
    game.grid = grid.NewGrid(dimension, dimension)
    game.grid.Fill()
    return game.grid
}

func (game *Game) RandomlySeedByPercentage(percentage int) {
    game.grid.RandomSeedByPercentage(percentage)
}

func (game *Game) PlayFor(duration int) {

    for i := 0; i < duration; i++ {
        singleIteration(game)
        game.Print(i + 1)
    }
}

func (game *Game) Print(step int) {
    width := game.grid.GetWidth()

    fmt.Printf("\nStep: %d\n", step)
    var counter int
    var line string
    for index, c := range game.grid.List {
        counter ++
        isAlive := <- c.Pulse
        if isAlive {
           line = line + "x"
        } else {
            line = line + "-"
        }

        if (index + 1) % width != 0 {
            line = line + " "
        } else {
            line = line + "\n"
        }
    }
    fmt.Printf(line)
}

func singleIteration(game *Game) {
    var wg sync.WaitGroup
    wg.Add(game.grid.CellCount())

    for _, c := range game.grid.List {
        go func(c *cell.Cell) {
            defer wg.Done()
            c.Life <- true
        }(c)
    }
    wg.Wait()
}