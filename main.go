package main

import (
    "flag"
    "game-of-life-go/game"
)

var (
    gridSize = flag.Int("s", 10, "Size of the side of the square game grid (defaults to 10)")
    randomSeedPercentage = flag.Int("p", 30, "Percentage of the grid (as an integer) to seed with live cells (defaults to 30)")
    duration = flag.Int("d", 10, "Iterations of the game to play (defaults to 10)")
)

func main() {
    flag.Parse()

    g := new(game.Game)
    g.Build(*gridSize)
    g.RandomlySeedByPercentage(*randomSeedPercentage)
    g.PlayFor(*duration)
}