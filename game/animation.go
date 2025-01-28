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
	at           AtType // align
	bindTo       *any
	dieWithBind  bool // force kill anim
	lifeSpan     int  //anim play in secs
}

var renderFrames uint = 0
var animationsList [ANIMATION_LINK_LIST_NUM]*list.List

func createAnimation(
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

func createAndPushAnimation(
	listId int,
	textureId int,
	effect *Effect,
	lp LoopType,
	duration int,
	x, y int,
	flip FlipType,
	angle float64,
	at AtType,
) *Animation {
	ani := createAnimation(
		textures[textureId],
		effect,
		lp,
		duration,
		x, y,
		flip,
		angle,
		at,
	)
	animationsList[listId].PushBack(ani)
	return ani
}

func initAnimList() {
	renderFrames = 0
	for idx := 0; idx < ANIMATION_LINK_LIST_NUM; idx++ {
		animationsList[idx] = list.New()
	}
	log.Println("| init animation list - [DONE]")
}
