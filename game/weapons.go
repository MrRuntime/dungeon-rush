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

func (w *Weapon) InitWeapon(textures []*Texture, birthTextureId, deathTextureId, flyTextureId int) {
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
