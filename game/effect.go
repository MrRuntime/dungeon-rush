package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Effect struct {
	duration     int
	currentFrame int
	length       uint
	keys         []color.RGBA
	mode         ebiten.Blend
}

func CreateEffect(duration int, length uint, mode ebiten.Blend) Effect {
	return Effect{duration: duration, length: length, mode: mode}
}

func (e *Effect) AddKey(key color.RGBA) {
	e.keys = append(e.keys, key)
}

func (e *Effect) copy() *Effect {
	newEffect := &Effect{
		duration:     e.duration,
		currentFrame: e.currentFrame,
		length:       e.length,
		keys:         e.keys,
		mode:         e.mode,
	}
	return newEffect
}
