package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

func New() (*Game, error) {
	g := &Game{}

	err := LoadAssets(g)
	if err != nil {
		return g, err
	}

	// mainUI()
	return g, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(width, height int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) Update() error {
	return nil
}
