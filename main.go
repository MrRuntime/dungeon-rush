package main

import (
	"dungeon-rush/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.New()
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(game.SCREEN_WIDTH, game.SCREEN_HEIGHT)
	ebiten.SetWindowTitle(game.NAME)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
