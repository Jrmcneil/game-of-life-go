package game

import "testing"

func TestGameTakesParametersAndCreatesAGridOfCells(t *testing.T) {
    game := new(Game)
    game.Build(5)
    grid := game.grid

    if  grid.CellCount() != 25 || grid.GetHeight() != 5 || grid.GetWidth() != 5 {
        t.Errorf("Initialised game should create a Grid of the correct size and number of Cells")
    }
}

func TestRandomSeedLiveCells(t *testing.T) {
}

func TestRunForDuration(t *testing.T) {
    game := new(Game)
    game.Build(5)

}
