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

func CreateSpriteList() SpriteList {
	return SpriteList{
		list: make(map[int]*Sprite),
	}
}

func InitCommonSprites() {
	// Heroes
	commonSprites.AddSprite(SPRITE_KNIGHT, &weapons[WEAPON_SWORD], textures[RES_KNIGHT_M], 150)
	commonSprites.AddSprite(SPRITE_ELF, &weapons[WEAPON_ARROW], textures[RES_ELF_M], 100)
	commonSprites.AddSprite(SPRITE_WIZZARD, &weapons[WEAPON_FIREBALL], textures[RES_WIZZARD_M], 95)
	commonSprites.AddSprite(SPRITE_LIZARD, &weapons[WEAPON_MONSTER_CLAW], textures[RES_ZIGGY_M], 120)

	// Baddies
	commonSprites.AddSprite(SPRITE_TINY_ZOMBIE, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_TINY_ZOMBIE], 50)
	commonSprites.AddSprite(SPRITE_GOBLIN, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_GOBLIN], 100)
	commonSprites.AddSprite(SPRITE_IMP, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_IMP], 100)
	commonSprites.AddSprite(SPRITE_SKELET, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_SKELET], 100)
	commonSprites.AddSprite(SPRITE_MUDDY, &weapons[WEAPON_SOLID], textures[RES_MUDDY], 150)
	commonSprites.AddSprite(SPRITE_SWAMPY, &weapons[WEAPON_SOLID_GREEN], textures[RES_SWAMPY], 150)
	commonSprites.AddSprite(SPRITE_ZOMBIE, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_ZOMBIE], 120)
	commonSprites.AddSprite(SPRITE_ICE_ZOMBIE, &weapons[WEAPON_ICEPICK], textures[RES_ICE_ZOMBIE], 120)
	commonSprites.AddSprite(SPRITE_MASKED_ORC, &weapons[WEAPON_THROW_AXE], textures[RES_MASKED_ORC], 120)
	commonSprites.AddSprite(SPRITE_ORC_WARRIOR, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_ORC_WARRIOR], 200)
	commonSprites.AddSprite(SPRITE_ORC_SHAMAN, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_ORC_SHAMAN], 120)
	commonSprites.AddSprite(SPRITE_NECROMANCER, &weapons[WEAPON_PURPLE_BALL], textures[RES_NECROMANCER], 120)
	commonSprites.AddSprite(SPRITE_WOGOL, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_WOGOL], 150)
	commonSprites.AddSprite(SPRITE_CHROT, &weapons[WEAPON_MONSTER_CLAW2], textures[RES_CHORT], 150)
	commonSprites.AddSprite(SPRITE_GREEN_HOOD_SKEL, &weapons[WEAPON_PURPLE_BALL], textures[RES_GREEN_HOOD_SKEL], 150)

	commonSprites.AddSprite(SPRITE_BIG_ZOMBIE, &weapons[WEAPON_THUNDER], textures[RES_BIG_ZOMBIE], 3000)
	commonSprites.GetSprite(SPRITE_BIG_ZOMBIE).dropRate = 100

	commonSprites.AddSprite(SPRITE_ORGRE, &weapons[WEAPON_MANY_AXES], textures[RES_ORGRE], 3000)
	commonSprites.GetSprite(SPRITE_ORGRE).dropRate = 100

	commonSprites.AddSprite(SPRITE_BIG_DEMON, &weapons[WEAPON_THUNDER], textures[RES_BIG_DEMON], 2500)
	commonSprites.GetSprite(SPRITE_BIG_DEMON).dropRate = 100
	log.Println("| init common sprites - [DONE]")
}
