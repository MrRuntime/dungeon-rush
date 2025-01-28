package game

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

func (s *Sprite) InitCommonSprite(weaponId int, textureId int, hp int) {
	ani := createAnimation(
		textures[textureId],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		0,
		0,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	s.x = 0
	s.y = 0
	s.hp = hp
	s.totalHp = hp
	s.weapon = &weapons[weaponId]
	s.ani = ani
	s.face = RIGHT
	s.direction = RIGHT
	s.lastAttack = 0
	s.dropRate = 1
}

func InitCommonSprites() {
	// Heroes
	commonSprites[SPRITE_KNIGHT].InitCommonSprite(WEAPON_SWORD, RES_KNIGHT_M, 150)
	// initCommonSprite(&commonSprites[SPRITE_ELF], &wp.weapons[wp.WEAPON_ARROW], RES_ELF_M, 100)
	// initCommonSprite(&commonSprites[SPRITE_WIZZARD], &wp.weapons[wp.WEAPON_FIREBALL], RES_WIZZARD_M, 95)
	// initCommonSprite(&commonSprites[SPRITE_LIZARD], &wp.weapons[wp.WEAPON_MONSTER_CLAW], RES_ZIGGY_M, 120)

	// // Baddies
	// initCommonSprite(&commonSprites[SPRITE_TINY_ZOMBIE], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_TINY_ZOMBIE, 50);
	// initCommonSprite(&commonSprites[SPRITE_GOBLIN], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_GOBLIN, 100);
	// initCommonSprite(&commonSprites[SPRITE_IMP], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_IMP, 100);
	// initCommonSprite(&commonSprites[SPRITE_SKELET], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_SKELET, 100);
	// initCommonSprite(&commonSprites[SPRITE_MUDDY], &wp.weapons[wp.WEAPON_SOLID], RES_MUDDY, 150);
	// initCommonSprite(&commonSprites[SPRITE_SWAMPY], &wp.weapons[wp.WEAPON_SOLID_GREEN], RES_SWAMPY, 150);
	// initCommonSprite(&commonSprites[SPRITE_ZOMBIE], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_ZOMBIE, 120);
	// initCommonSprite(&commonSprites[SPRITE_ICE_ZOMBIE], &wp.weapons[wp.WEAPON_ICEPICK], RES_ICE_ZOMBIE, 120);
	// initCommonSprite(&commonSprites[SPRITE_MASKED_ORC], &wp.weapons[wp.WEAPON_THROW_AXE], RES_MASKED_ORC, 120);
	// initCommonSprite(&commonSprites[SPRITE_ORC_WARRIOR], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_ORC_WARRIOR, 200);
	// initCommonSprite(&commonSprites[SPRITE_ORC_SHAMAN], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_ORC_SHAMAN, 120);
	// initCommonSprite(&commonSprites[SPRITE_NECROMANCER], &wp.weapons[wp.WEAPON_PURPLE_BALL], RES_NECROMANCER, 120);
	// initCommonSprite(&commonSprites[SPRITE_WOGOL], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_WOGOL, 150);
	// initCommonSprite(&commonSprites[SPRITE_CHROT], &wp.weapons[wp.WEAPON_MONSTER_CLAW2], RES_CHORT, 150);
	// initCommonSprite(&commonSprites[SPRITE_GREEN_HOOD_SKEL], &wp.weapons[wp.WEAPON_PURPLE_BALL], RES_GREEN_HOOD_SKEL, 150);

	// var now: *spr.Sprite = undefined;

	// now = &commonSprites[SPRITE_BIG_ZOMBIE];
	// now.dropRate = 100;
	// initCommonSprite(now, &wp.weapons[wp.WEAPON_THUNDER], RES_BIG_ZOMBIE, 3000);

	// now = &commonSprites[SPRITE_ORGRE];
	// now.dropRate = 100;
	// initCommonSprite(now, &wp.weapons[wp.WEAPON_MANY_AXES], RES_ORGRE, 3000);

	// now = &commonSprites[SPRITE_BIG_DEMON];
	// now.dropRate = 100;
	// initCommonSprite(now, &wp.weapons[wp.WEAPON_THUNDER], RES_BIG_DEMON, 2500);
}
