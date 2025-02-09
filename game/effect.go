package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Effect struct {
	duration     int
	currentFrame int
	length       int
	keys         []color.RGBA
	mode         ebiten.Blend
}

func CreateEffect(duration int, length int, mode ebiten.Blend) Effect {
	return Effect{duration: duration, length: length, mode: mode}
}

func (e *Effect) AddKey(key color.RGBA) {
	e.keys = append(e.keys, key)
}

func (e *Effect) copy() *Effect {
	return &Effect{
		duration:     e.duration,
		currentFrame: e.currentFrame,
		length:       e.length,
		keys:         e.keys,
		mode:         e.mode,
	}
}

func SetEffect(texture *Texture, ef *Effect) {
	if ef == nil {
		return
	}
	effect := ef
	// _ = c.SDL_SetTextureBlendMode(texture.origin, effect.mode);

	interval := float64(effect.duration)/float64(effect.length) - 1.0
	progress := float64(effect.currentFrame)
	stage := int(progress / interval)
	progress -= float64(stage) * interval
	progress /= interval

	prev := effect.keys[stage]
	nxt := effect.keys[min(stage+1, effect.length-1)]

	mixed := color.RGBA{}
	mixed.R = uint8(float64(prev.R)*(1-progress) + float64(nxt.R)*progress)
	mixed.G = uint8(float64(prev.G)*(1-progress) + float64(nxt.G)*progress)
	mixed.B = uint8(float64(prev.B)*(1-progress) + float64(nxt.B)*progress)
	mixed.A = uint8(float64(prev.A)*(1-progress) + float64(nxt.A)*progress)

	log.Println("NOT IMPL")
	// op := &ebiten.DrawImageOptions{}
	// op.ColorScale
	// colorm.ColorM{}

	// _ = c.SDL_SetTextureColorMod(texture.origin, mixed.R, mixed.G, mixed.B)
	// _ = c.SDL_SetTextureAlphaMod(texture.origin, mixed.A)
}

func UnsetEffect(texture *Texture) {
	log.Println("NOT IMPL")
	// _ = c.SDL_SetTextureBlendMode(texture.origin, c.SDL_BLENDMODE_BLEND);
	// _ = c.SDL_SetTextureColorMod(texture.origin, 255, 255, 255);
	// _ = c.SDL_SetTextureAlphaMod(texture.origin, 255);
}
