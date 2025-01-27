package main

import (
	"dungeon-rush/game"
)

func main() {
	_, err := game.New()
	if err != nil {
		panic(err)
	}

	// ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	// ebiten.SetWindowTitle(game.NAME)

	// if err := ebiten.RunGame(g); err != nil {
	// 	panic(err)
	// }
}
