package game

import "log"

type WeaponType int

const (
	WEAPON_SWORD_POINT WeaponType = iota
	WEAPON_SWORD_RANGE
	WEAPON_GUN_RANGE
	WEAPON_GUN_POINT
	WEAPON_GUN_POINT_MULTI
)

type WeaponBuff struct {
	chance   float64
	duration int
}

type Weapon struct {
	wp WeaponType
	// distance for the projectile to fire, too far and it won't fire
	shootRange int
	// not sure
	effectRange int
	// how much hp damage
	damage int
	// fire rate of weapon
	gap int
	// speed of projectile
	bulletSpeed int

	birthAni *Animation
	deathAni *Animation
	flyAni   *Animation

	birthAudio int
	deathAudio int

	effects [BUFF_END]WeaponBuff
}

func (w *Weapon) InitWeapon(birthTextureId, deathTextureId, flyTextureId int) {
	var birthAni *Animation = nil
	var deathAni *Animation = nil
	var flyAni *Animation = nil

	if birthTextureId != NONE {
		birthAni = createAnimation(
			textures[birthTextureId],
			nil,
			LOOP_ONCE,
			30,
			0,
			0,
			FLIP_NONE,
			0,
			AT_CENTER,
		)
	}
	if deathTextureId != NONE {
		deathAni = createAnimation(
			textures[deathTextureId],
			nil,
			LOOP_ONCE,
			30,
			0,
			0,
			FLIP_NONE,
			0,
			AT_BOTTOM_CENTER,
		)
	}
	if flyTextureId != NONE {
		flyAni = createAnimation(
			textures[flyTextureId],
			nil,
			LOOP_INFI,
			30,
			0,
			0,
			FLIP_NONE,
			0,
			AT_CENTER,
		)
	}
	w.wp = WEAPON_SWORD_POINT
	w.shootRange = 32 * 2
	w.effectRange = 40
	w.damage = 10
	w.gap = 60
	w.bulletSpeed = 6
	w.birthAudio = -1
	w.deathAudio = 5
	w.birthAni = birthAni
	w.deathAni = deathAni
	w.flyAni = flyAni
}

func InitWeapons() {
	weapons[WEAPON_SWORD].InitWeapon(NONE, RES_SWORDFX, NONE)
	weapons[WEAPON_SWORD].damage = 30
	weapons[WEAPON_SWORD].shootRange = 32 * 3
	weapons[WEAPON_SWORD].deathAni.scaled = false
	weapons[WEAPON_SWORD].deathAni.angle = -1.0
	weapons[WEAPON_SWORD].deathAudio = AUDIO_SWORD_HIT

	weapons[WEAPON_MONSTER_CLAW].InitWeapon(NONE, RES_CLAWFX2, NONE)
	weapons[WEAPON_MONSTER_CLAW].wp = WEAPON_SWORD_RANGE
	weapons[WEAPON_MONSTER_CLAW].shootRange = 32*3 + 16
	weapons[WEAPON_MONSTER_CLAW].damage = 24
	weapons[WEAPON_MONSTER_CLAW].deathAni.angle = -1.0
	weapons[WEAPON_MONSTER_CLAW].deathAni.at = AT_CENTER
	weapons[WEAPON_MONSTER_CLAW].deathAudio = AUDIO_CLAW_HIT_HEAVY

	weapons[WEAPON_FIREBALL].InitWeapon(RES_SHINE, RES_HALO_EXPLOSION1, RES_FIREBALL)
	weapons[WEAPON_FIREBALL].wp = WEAPON_GUN_RANGE
	weapons[WEAPON_FIREBALL].damage = 45
	weapons[WEAPON_FIREBALL].effectRange = 50
	weapons[WEAPON_FIREBALL].shootRange = 256
	weapons[WEAPON_FIREBALL].gap = 180
	weapons[WEAPON_FIREBALL].deathAni.angle = -1.0
	weapons[WEAPON_FIREBALL].deathAni.at = AT_CENTER
	weapons[WEAPON_FIREBALL].birthAni.duration = 24
	weapons[WEAPON_FIREBALL].birthAudio = AUDIO_SHOOT
	weapons[WEAPON_FIREBALL].deathAudio = AUDIO_FIREBALL_EXP

	weapons[WEAPON_THUNDER].InitWeapon(RES_BLOODBOUND, RES_THUNDER, NONE)
	weapons[WEAPON_THUNDER].wp = WEAPON_SWORD_RANGE
	weapons[WEAPON_THUNDER].damage = 80
	weapons[WEAPON_THUNDER].shootRange = 128
	weapons[WEAPON_THUNDER].gap = 120
	weapons[WEAPON_THUNDER].deathAni.angle = -1
	weapons[WEAPON_THUNDER].deathAni.scaled = false
	weapons[WEAPON_THUNDER].deathAudio = AUDIO_THUNDER

	weapons[WEAPON_THUNDER_STAFF].InitWeapon(NONE, RES_THUNDER_YELLOW, NONE)
	weapons[WEAPON_THUNDER_STAFF].wp = WEAPON_SWORD_RANGE
	weapons[WEAPON_THUNDER_STAFF].damage = 50
	weapons[WEAPON_THUNDER_STAFF].shootRange = 128
	weapons[WEAPON_THUNDER_STAFF].gap = 120
	weapons[WEAPON_THUNDER_STAFF].deathAni.angle = -1
	weapons[WEAPON_THUNDER_STAFF].deathAni.scaled = false
	weapons[WEAPON_THUNDER_STAFF].deathAudio = AUDIO_THUNDER

	weapons[WEAPON_ARROW].InitWeapon(NONE, RES_HALO_EXPLOSION2, RES_ARROW)
	weapons[WEAPON_ARROW].wp = WEAPON_GUN_POINT
	weapons[WEAPON_ARROW].gap = 40
	weapons[WEAPON_ARROW].damage = 10
	weapons[WEAPON_ARROW].shootRange = 200
	weapons[WEAPON_ARROW].bulletSpeed = 10
	weapons[WEAPON_ARROW].deathAni.angle = -1
	weapons[WEAPON_ARROW].deathAni.at = AT_CENTER
	weapons[WEAPON_ARROW].flyAni.scaled = false
	weapons[WEAPON_ARROW].birthAudio = AUDIO_BOW_FIRE
	weapons[WEAPON_ARROW].deathAudio = AUDIO_BOW_HIT

	weapons[WEAPON_POWERFUL_BOW].InitWeapon(NONE, RES_HALO_EXPLOSION2, RES_ARROW)
	weapons[WEAPON_POWERFUL_BOW].wp = WEAPON_GUN_POINT
	weapons[WEAPON_POWERFUL_BOW].gap = 60
	weapons[WEAPON_POWERFUL_BOW].damage = 25
	weapons[WEAPON_POWERFUL_BOW].shootRange = 320
	weapons[WEAPON_POWERFUL_BOW].bulletSpeed = 7
	weapons[WEAPON_POWERFUL_BOW].deathAni.angle = -1
	weapons[WEAPON_POWERFUL_BOW].deathAni.at = AT_CENTER
	weapons[WEAPON_POWERFUL_BOW].birthAudio = AUDIO_BOW_FIRE
	weapons[WEAPON_POWERFUL_BOW].deathAudio = AUDIO_BOW_HIT
	weapons[WEAPON_POWERFUL_BOW].effects[BUFF_ATTACK] = WeaponBuff{chance: 0.5, duration: 240}

	weapons[WEAPON_MONSTER_CLAW2].InitWeapon(NONE, RES_CLAWFX, NONE)

	weapons[WEAPON_THROW_AXE].InitWeapon(NONE, RES_CROSS_HIT, RES_AXE)
	weapons[WEAPON_THROW_AXE].wp = WEAPON_GUN_POINT
	weapons[WEAPON_THROW_AXE].damage = 12
	weapons[WEAPON_THROW_AXE].shootRange = 160
	weapons[WEAPON_THROW_AXE].bulletSpeed = 10
	weapons[WEAPON_THROW_AXE].flyAni.duration = 24
	weapons[WEAPON_THROW_AXE].flyAni.angle = -1
	weapons[WEAPON_THROW_AXE].flyAni.scaled = false
	weapons[WEAPON_THROW_AXE].deathAni.scaled = false
	weapons[WEAPON_THROW_AXE].deathAni.at = AT_CENTER
	weapons[WEAPON_THROW_AXE].birthAudio = AUDIO_AXE_FLY //res.AUDIO_LIGHT_SHOOT
	weapons[WEAPON_THROW_AXE].deathAudio = AUDIO_ARROW_HIT

	weapons[WEAPON_MANY_AXES].InitWeapon(NONE, RES_CROSS_HIT, RES_AXE)
	weapons[WEAPON_MANY_AXES].wp = WEAPON_GUN_POINT_MULTI
	weapons[WEAPON_MANY_AXES].shootRange = 180
	weapons[WEAPON_MANY_AXES].gap = 70
	weapons[WEAPON_MANY_AXES].effectRange = 50
	weapons[WEAPON_MANY_AXES].damage = 50
	weapons[WEAPON_MANY_AXES].bulletSpeed = 4
	weapons[WEAPON_MANY_AXES].flyAni.duration = 24
	weapons[WEAPON_MANY_AXES].flyAni.angle = -1
	weapons[WEAPON_MANY_AXES].deathAni.at = AT_CENTER
	weapons[WEAPON_MANY_AXES].birthAudio = AUDIO_AXE_FLY //res.AUDIO_LIGHT_SHOOT;res.AUDIO_LIGHT_SHOOT
	weapons[WEAPON_MANY_AXES].deathAudio = AUDIO_ARROW_HIT

	weapons[WEAPON_SOLID].InitWeapon(NONE, RES_SOLIDFX, NONE)
	weapons[WEAPON_SOLID].deathAni.scaled = false
	weapons[WEAPON_SOLID].deathAni.angle = -1
	weapons[WEAPON_SOLID].effects[BUFF_SLOWDOWN] = WeaponBuff{chance: 0.3, duration: 180}

	weapons[WEAPON_SOLID_GREEN].InitWeapon(NONE, RES_SOLID_GREENFX, NONE)
	weapons[WEAPON_SOLID_GREEN].shootRange = 96
	weapons[WEAPON_SOLID_GREEN].deathAni.scaled = false
	weapons[WEAPON_SOLID_GREEN].deathAni.angle = -1
	weapons[WEAPON_SOLID_GREEN].effects[BUFF_SLOWDOWN] = WeaponBuff{chance: 0.3, duration: 180}

	weapons[WEAPON_SOLID_CLAW].InitWeapon(NONE, RES_SOLID_GREENFX, NONE)
	weapons[WEAPON_SOLID_CLAW].wp = WEAPON_SWORD_RANGE
	weapons[WEAPON_SOLID_CLAW].shootRange = 32*3 + 16
	weapons[WEAPON_SOLID_CLAW].damage = 35
	weapons[WEAPON_SOLID_CLAW].deathAni.scaled = false
	weapons[WEAPON_SOLID_CLAW].deathAni.angle = -1
	weapons[WEAPON_SOLID_CLAW].deathAudio = AUDIO_CLAW_HIT_HEAVY
	weapons[WEAPON_SOLID_CLAW].effects[BUFF_SLOWDOWN] = WeaponBuff{chance: 0.7, duration: 60}

	weapons[WEAPON_ICEPICK].InitWeapon(NONE, RES_ICESHATTER, RES_ICEPICK)
	weapons[WEAPON_ICEPICK].wp = WEAPON_GUN_RANGE
	weapons[WEAPON_ICEPICK].damage = 30
	weapons[WEAPON_ICEPICK].effectRange = 50
	weapons[WEAPON_ICEPICK].shootRange = 256
	weapons[WEAPON_ICEPICK].gap = 180
	weapons[WEAPON_ICEPICK].bulletSpeed = 8
	weapons[WEAPON_ICEPICK].deathAni.angle = -1
	weapons[WEAPON_ICEPICK].flyAni.scaled = false
	weapons[WEAPON_ICEPICK].deathAni.at = AT_CENTER
	weapons[WEAPON_ICEPICK].effects[BUFF_FROZEN] = WeaponBuff{chance: 0.2, duration: 60}
	weapons[WEAPON_ICEPICK].birthAudio = AUDIO_ICE_SHOOT

	weapons[WEAPON_PURPLE_BALL].InitWeapon(NONE, RES_PURPLE_EXP, RES_PURPLE_BALL)
	weapons[WEAPON_PURPLE_BALL].wp = WEAPON_GUN_RANGE
	weapons[WEAPON_PURPLE_BALL].damage = 20
	weapons[WEAPON_PURPLE_BALL].effectRange = 50
	weapons[WEAPON_PURPLE_BALL].shootRange = 256
	weapons[WEAPON_PURPLE_BALL].gap = 100
	weapons[WEAPON_PURPLE_BALL].bulletSpeed = 6
	weapons[WEAPON_PURPLE_BALL].deathAni.angle = -1
	weapons[WEAPON_PURPLE_BALL].deathAni.scaled = false
	weapons[WEAPON_PURPLE_BALL].flyAni.scaled = false
	weapons[WEAPON_PURPLE_BALL].deathAni.at = AT_CENTER
	weapons[WEAPON_PURPLE_BALL].birthAudio = AUDIO_ICE_SHOOT
	weapons[WEAPON_PURPLE_BALL].deathAudio = AUDIO_ARROW_HIT

	weapons[WEAPON_PURPLE_STAFF].InitWeapon(NONE, RES_PURPLE_EXP, RES_PURPLE_FIRE_BALL)
	weapons[WEAPON_PURPLE_STAFF].wp = WEAPON_GUN_POINT_MULTI
	weapons[WEAPON_PURPLE_STAFF].damage = 45
	weapons[WEAPON_PURPLE_STAFF].effectRange = 50
	weapons[WEAPON_PURPLE_STAFF].shootRange = 256
	weapons[WEAPON_PURPLE_STAFF].gap = 100
	weapons[WEAPON_PURPLE_STAFF].bulletSpeed = 7
	weapons[WEAPON_PURPLE_STAFF].deathAni.angle = -1
	weapons[WEAPON_PURPLE_STAFF].deathAni.scaled = false
	weapons[WEAPON_PURPLE_STAFF].flyAni.scaled = true
	weapons[WEAPON_PURPLE_STAFF].deathAni.at = AT_CENTER
	weapons[WEAPON_PURPLE_STAFF].birthAudio = AUDIO_ICE_SHOOT
	weapons[WEAPON_PURPLE_STAFF].deathAudio = AUDIO_ARROW_HIT

	weapons[WEAPON_HOLY_SWORD].InitWeapon(NONE, RES_GOLDEN_CROSS_HIT, NONE)
	weapons[WEAPON_HOLY_SWORD].wp = WEAPON_SWORD_RANGE
	weapons[WEAPON_HOLY_SWORD].damage = 30
	weapons[WEAPON_HOLY_SWORD].shootRange = 32 * 4
	weapons[WEAPON_HOLY_SWORD].effects[BUFF_DEFENCE] = WeaponBuff{chance: 0.6, duration: 180}

	weapons[WEAPON_ICE_SWORD].InitWeapon(NONE, RES_ICESHATTER, NONE)
	weapons[WEAPON_ICE_SWORD].wp = WEAPON_SWORD_RANGE
	weapons[WEAPON_ICE_SWORD].shootRange = 32*3 + 16
	weapons[WEAPON_ICE_SWORD].damage = 80
	weapons[WEAPON_ICE_SWORD].gap = 30
	weapons[WEAPON_ICE_SWORD].deathAni.angle = -1
	weapons[WEAPON_ICE_SWORD].deathAni.at = AT_CENTER
	weapons[WEAPON_ICE_SWORD].effects[BUFF_FROZEN] = WeaponBuff{chance: 0.6, duration: 80}
	weapons[WEAPON_ICE_SWORD].deathAudio = AUDIO_SWORD_HIT
	log.Println("| init weapons - [DONE]")
}
