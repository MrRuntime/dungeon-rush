package game

import (
	"container/list"
	"log"
)

//go:generate stringer -type=LoopType -linecomment
type LoopType int

const (
	LOOP_ONCE LoopType = iota
	LOOP_INFI
	LOOP_LIFESPAN
)

//go:generate stringer -type=FlipType -linecomment
type FlipType int

const (
	FLIP_NONE FlipType = iota
	FLIP_H
	FLIP_V
)

//go:generate stringer -type=AtType -linecomment
type AtType int

const (
	AT_TOP_LEFT AtType = iota
	AT_BOTTOM_LEFT
	AT_BOTTOM_CENTER
	AT_CENTER
)

type AnimationsList = list.List

type Animation struct {
	lp           LoopType
	origin       *Texture
	effect       *Effect
	duration     int
	currentFrame int
	x            int
	y            int
	angle        float64
	flip         FlipType
	scaled       bool
	at           AtType  // align
	bindTo       *Sprite // any or interface{}
	dieWithBind  bool    // force kill anim
	lifeSpan     int     //anim play in secs
}

var renderFrames uint = 0

func CreateAnimation(
	origin *Texture,
	effect *Effect,
	lp LoopType,
	duration int,
	x, y int,
	flip FlipType,
	angle float64,
	at AtType,
) *Animation {
	var ef *Effect
	if effect != nil {
		ef = effect.copy()
	} else {
		ef = nil
	}

	return &Animation{
		origin:       origin,
		effect:       ef,
		lp:           lp,
		duration:     duration,
		currentFrame: 0,
		x:            x,
		y:            y,
		flip:         flip,
		angle:        angle,
		at:           at,
		bindTo:       nil,
		dieWithBind:  false,
		scaled:       true,
		lifeSpan:     duration,
	}
}

func CreateAndPushAnimation(
	animationsList *AnimationsList,
	texture *Texture,
	effect *Effect,
	lp LoopType,
	duration int,
	x, y int,
	flip FlipType,
	angle float64,
	at AtType,
) *Animation {
	ani := CreateAnimation(
		texture,
		effect,
		lp,
		duration,
		x, y,
		flip,
		angle,
		at,
	)
	animationsList.PushBack(ani)
	return ani
}

func InitAnimList(assets *Assets) {
	renderFrames = 0
	for idx := range ANIMATION_LINK_LIST_NUM {
		assets.animations[idx] = list.New()
	}
	log.Println("| init animation list - [DONE]")
}

func UpdateAnimationFromBind(ani *Animation) {
	if bnd := ani.bindTo; bnd != nil {
		sprite := bnd
		ani.x = sprite.x
		ani.y = sprite.y
		ani.flip = sprite.ani.flip
	}
}

func UpdateAnimationOfSprite(sprite *Sprite) {
	ani := sprite.ani
	ani.x = sprite.x
	ani.y = sprite.y
	if sprite.face == RIGHT {
		ani.flip = FLIP_NONE
	} else {
		ani.flip = FLIP_H
	}
}

func UpdateAnimationLinkList(al *AnimationsList) {
	for p := al.Front(); p != nil; p = p.Next() {
		ani := p.Value.(*Animation)
		ani.currentFrame += 1
		ani.lifeSpan -= 1

		if ani.effect != nil {
			ani.effect.currentFrame += 1
			ani.effect.currentFrame = ani.effect.currentFrame % ani.effect.duration
		}

		if ani.lp == LOOP_ONCE {
			if ani.currentFrame == ani.duration {
				ani = nil
				al.Remove(p)
			}
		} else {
			if ani.lp == LOOP_LIFESPAN && ani.lifeSpan <= 0 {
				ani = nil
				al.Remove(p)
			} else {
				ani.currentFrame = ani.currentFrame % ani.duration
			}
		}
	}
}
