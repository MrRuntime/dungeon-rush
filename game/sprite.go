package game

import "log"

type PositionBufferSlot struct {
	x         int
	y         int
	direction Direction
}

type PositionBufferQueue Queue[PositionBufferSlot]

type Sprite struct {
	x       int
	y       int
	hp      int
	totalHp int

	weapon    *Weapon
	ani       *Animation
	face      Direction
	direction Direction

	posQueue PositionBufferQueue

	lastAttack int // Timestamp of the last attack
	dropRate   float64
}

type SpriteList struct {
	list map[int]*Sprite
}

func (self *SpriteList) AddSprite(id int, weapon *Weapon, texture *Texture, hp int) {
	ani := createAnimation(
		texture,
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		0,
		0,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	self.list[id] = &Sprite{
		x:          0,
		y:          0,
		hp:         hp,
		totalHp:    hp,
		weapon:     weapon,
		ani:        ani,
		face:       RIGHT,
		direction:  RIGHT,
		lastAttack: 0,
		dropRate:   1,
	}
}

func (self *SpriteList) GetSprite(id int) *Sprite {
	return self.list[id]
}
func (self *SpriteList) GetCopy(id int) *Sprite {
	s := self.list[id]
	return &Sprite{
		x:          s.x,
		y:          s.y,
		hp:         s.hp,
		totalHp:    s.totalHp,
		weapon:     s.weapon,
		ani:        s.ani,
		face:       s.face,
		direction:  s.direction,
		lastAttack: s.lastAttack,
		dropRate:   s.dropRate,
	}
}

func CreateSpriteList() SpriteList {
	return SpriteList{
		list: make(map[int]*Sprite),
	}
}
