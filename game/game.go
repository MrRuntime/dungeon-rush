package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenSize struct {
	width  int
	height int
}

type Game struct {
	screenSize ScreenSize
	assets     *Assets
	opts       []*Text
	player     *Snake
}

func New() (*Game, error) {
	assets, err := LoadAssets()

	// InitUI maybe move to update
	BaseUI(assets, 30, 12)

	if err != nil {
		return nil, err
	}

	w, h := ebiten.Monitor().Size()

	g := &Game{
		screenSize: ScreenSize{w, h},
		assets:     assets,
	}

	const optsNum = 4
	for i := range optsNum {
		g.opts = append(g.opts, &assets.texts[i+6])
	}
	ebiten.SetFullscreen(true)

	// mainUI()
	return g, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.player == nil {
		player := CreateSnake(2, 0, LOCAL)

		AppendSpriteToSnake(
			g.assets,
			player,
			SPRITE_KNIGHT,
			g.screenSize.width/2,
			g.screenSize.height/2,
			UP,
		)
		g.player = player
		fmt.Println("!!!!")
	}

	cursorPos = 0
	lineGap := FONT_SIZE + FONT_SIZE/2
	totalHeight := lineGap * (len(g.opts) - 1)
	startY := (g.screenSize.height - totalHeight) / 2

	sprite := g.player.sprites.Front().Value.(*Sprite)
	sprite.ani.at = AT_CENTER
	sprite.x = g.screenSize.width/2 - g.opts[cursorPos].width/2 - UNIT/2
	sprite.y = startY + cursorPos*lineGap

	UpdateAnimationOfSprite(sprite)
	DrawUI(screen, g.assets)

	for i, opt := range g.opts {
		DrawCenteredText(screen, opt, g.screenSize.width/2, startY+i*lineGap, 1.5)
	}

	author := g.assets.texts[17]
	version := g.assets.texts[len(g.assets.texts)-1]
	DrawText(screen, &author, 10, g.screenSize.height-author.height-10, 1)
	DrawText(screen, &version, g.screenSize.width-version.width-10, g.screenSize.height-version.height-10, 1)
}

func (g *Game) Layout(width, height int) (int, int) {
	// return int(float64(g.screenSize.width) * 1.05), int(float64(g.screenSize.height) * 1.05)
	return g.screenSize.width, g.screenSize.height
}

func (g *Game) Update() error {
	return nil
}
