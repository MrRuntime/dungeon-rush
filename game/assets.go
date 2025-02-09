package game

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Texture struct {
	name   string
	origin *ebiten.Image
	width  int
	height int
	frames int
	crops  []image.Rectangle
}

type Text = struct {
	text   string
	width  int
	height int
	origin *text.GoTextFace
	color  color.RGBA
}

type Font = *text.GoTextFaceSource

type Audio struct {
	sounds [][]byte
	bgms   [][]byte
}

type Assets struct {
	texts         []Text
	textures      []*Texture
	commonSprites SpriteList
	effects       []Effect
	weapons       [WEAPONS_SIZE]Weapon
	font          Font
	audio         Audio
	animations    [ANIMATION_LINK_LIST_NUM]*AnimationsList
}

var textList = []string{
	"DungeonRush",
	"By Rapiz",
	"PLACEHOLDER",
	"PLACEHOLDER",
	"Player 1",
	"Player 2",
	"Singleplayer",
	"Multiplayers",
	"Ranklist",
	"Exit",
	"Normal",
	"Hard",
	"Insane",
	"Local",
	"Lan",
	"Host a game",
	"Join a game",
	"Go Edition: by @MrRuntime - (c) 2025", // <-- that's me!
}

var tilesetName = []string{
	"0x72_DungeonTilesetII_v1_3", "fireball_explosion1",
	"halo_explosion1", "halo_explosion2", "fireball",
	"floor_spike", "floor_exit", "HpMed", "SwordFx",
	"ClawFx", "Shine", "Thunder", "BloodBound",
	"arrow", "explosion-2", "ClawFx2", "Axe",
	"cross_hit", "blood", "SolidFx", "IcePick",
	"IceShatter", "Ice", "SwordPack", "HolyShield",
	"golden_cross_hit", "ui", "title", "purple_ball",
	"purple_exp", "staff", "Thunder_Yellow", "attack_up",
	"powerful_bow", "purple_fire_ball", "starfield",
}

var bgmsPath = []string{
	"assets/audio/main_title.ogg",
	"assets/audio/bg1.ogg",
	"assets/audio/bg2.ogg",
	"assets/audio/bg3.ogg",
}

// More: https://microstudio.dev/community/resources/essential-retro-video-game-sound-effects-collection-512-sounds/181/
var soundfxList = []string{
	"win.wav",
	"lose_2v.wav",
	"powerloss.wav",
	"hit_0.5v.wav",
	"sword_hit.wav",
	"claw_hit.wav",
	"arrow_hit.wav",
	"shoot.wav",
	"fireball_explosion.wav",
	"ice_shoot_0.5v.wav",
	"interaction1_0.75v.wav",
	"button1.wav",
	"thunder_2v.wav",
	"light_shoot.wav",
	"human_death.wav",
	"claw_hit_heavy.wav",
	"coin.wav",
	"med.wav",
	"holy.wav",
	"axe.wav",
	"hyessir1.wav", // Hero acknowledgements
	"hyessir2.wav",
	"hyessir3.wav",
	"hyessir4.wav",
	"wzyessr1.wav",
	"wzyessr2.wav",
	"wzyessr3.wav",
	"eyessir1.wav",
	"eyessir2.wav",
	"eyessir3.wav",
	"eyessir4.wav",
	"tryessr1.wav",
	"tryessr2.wav",
	"tryessr3.wav",
	"bowfire.wav",
	"bowhit.wav",
}

func InitTileSet(csvPath string, origin *ebiten.Image) ([]*Texture, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	t := make([]*Texture, 0)

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()

	for _, line := range lines {
		if len(strings.TrimSpace(line[0])) > 0 && line[0] != "#" {
			name := line[0]
			x, err := strconv.Atoi(line[1])
			y, err := strconv.Atoi(line[2])
			w, err := strconv.Atoi(line[3])
			h, err := strconv.Atoi(line[4])
			fc, err := strconv.Atoi(line[5])

			if err != nil {
				panic(err)
			}
			texture := &Texture{name, origin, w, h, fc, make([]image.Rectangle, fc)}

			for i := range fc {
				xx := x + i*w
				texture.crops[i] = image.Rect(xx, y, w+xx, h+y)
			}

			t = append(t, texture)
		}
	}
	return t, err
}

func LoadTileSet() ([]*Texture, error) {
	res := []*Texture{}
	for _, it := range tilesetName {
		csvPath := fmt.Sprintf("assets/images/%s", it)
		imgPath := fmt.Sprintf("assets/images/%s.png", it)
		origin, _, _ := ebitenutil.NewImageFromFile(imgPath)
		t, err := InitTileSet(csvPath, origin)
		if err != nil {
			return nil, err
		}
		res = append(res, t...)
	}
	log.Println("| load tileset - [DONE]")
	return res, nil
}

func LoadFont() (Font, error) {
	var err error

	file, err := os.Open("assets/font/m5x7.ttf")
	if err != nil {
		log.Fatalf("Failed to open font file: %v", err)
	}
	defer file.Close()

	fontData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read font file: %v", err)
	}

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fontData))
	if err != nil {
		log.Fatalf("Failed to create font source: %v", err)
	}

	log.Println("| load font - [DONE]")
	return s, nil
}

func LoadTextSet(font Font) []Text {
	texts := []Text{}
	for _, str := range textList {
		f := &text.GoTextFace{Source: font, Size: FONT_SIZE}
		w, h := text.Measure(str, f, 0)
		texts = append(texts, Text{
			origin: f,
			text:   str,
			width:  int(w),
			height: int(h),
			color:  color.RGBA{0xff, 0xff, 0xff, 0xff},
		})
	}

	version := fmt.Sprintf("Version: %.1f", VERSION)
	f := &text.GoTextFace{Source: font, Size: FONT_SIZE}
	w, h := text.Measure(version, f, 0)
	texts = append(texts, Text{
		origin: f,
		text:   version,
		width:  int(w),
		height: int(h),
		color:  color.RGBA{0xff, 0xff, 0xff, 0xff},
	})

	log.Println("| load textset - [DONE]")
	return texts
}

func loadOgg(audioPath string) ([]byte, error) {
	file, err := os.Open(audioPath)
	if err != nil {
		log.Fatalf("Failed to open audio file [%v]: %v", file.Name(), err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to read audio file [%v] stat: %v", file.Name(), err)
	}
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil {
		log.Fatalf("Failed to read audio file [%v] bytes: %v", file.Name(), err)
	}

	var o io.Reader
	o, err = vorbis.DecodeWithoutResampling(bytes.NewReader(bs))
	if err != nil {
		log.Fatalf("Failed to decode audio file [%v]: %v", file.Name(), err)
	}
	b, err := io.ReadAll(o)
	if err != nil {
		log.Fatalf("Failed to read bytes from file [%v]: %v", file.Name(), err)
	}

	return b, nil
}

func loadWav(audioPath string) ([]byte, error) {
	file, err := os.Open(audioPath)
	if err != nil {
		log.Fatalf("Failed to open wav file [%v]: %v", file.Name(), err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to read wav file [%v]: %v", file.Name(), err)
	}
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil {
		log.Fatalf("Failed to read audio file [%v] bytes: %v", file.Name(), err)
	}

	var o io.Reader
	o, err = wav.DecodeWithoutResampling(bytes.NewReader(bs))
	if err != nil {
		log.Fatalf("Failed to decode wav file [%v]: %v", file.Name(), err)
	}
	b, err := io.ReadAll(o)
	if err != nil {
		log.Fatalf("Failed to read bytes from file [%v]: %v", file.Name(), err)
	}

	return b, nil
}

func LoadAudio() (Audio, error) {
	a := Audio{}
	for _, it := range bgmsPath {
		tmp, err := loadOgg(it)
		if err != nil {
			return a, err
		}
		a.bgms = append(a.bgms, tmp)
	}
	for _, it := range soundfxList {
		tmp, err := loadWav(fmt.Sprintf("assets/audio/%s", it))
		if err != nil {
			return a, err
		}
		a.sounds = append(a.sounds, tmp)
	}
	log.Println("| load audio - [DONE]")
	return a, nil
}

func (a *Assets) InitCommonEffects() {
	deathEffect := CreateEffect(30, 4, ebiten.BlendDestinationOver)
	death := color.RGBA{0xff, 0xff, 0xff, 0xff}
	deathEffect.AddKey(death)

	death.R = 0xa8
	death.G = 0x0
	death.B = 0x0
	deathEffect.AddKey(death)

	death.R = 0x50
	deathEffect.AddKey(death)

	death.R = 0x0
	death.A = 0x0
	deathEffect.AddKey(death)
	a.effects = append(a.effects, deathEffect)
	// log.Println("| effect #0: Death (30frames) loaded")

	blinkEffect := CreateEffect(30, 3, ebiten.BlendLighter)
	blink := color.RGBA{0x0, 0x0, 0x0, 0xff}
	blinkEffect.AddKey(blink)
	blink.R = 0xc8
	blink.G = 0xc8
	blink.B = 0xc8
	blinkEffect.AddKey(blink)
	blink.R = 0x0
	blink.G = 0x0
	blink.B = 0x0
	blinkEffect.AddKey(blink)
	a.effects = append(a.effects, blinkEffect)
	// log.Println("| effect #1: Blink (white) (30frames) loaded")

	vanishEffect := CreateEffect(30, 2, ebiten.BlendDestinationOver)
	vanish := color.RGBA{0xff, 0xff, 0xff, 0xff}
	vanishEffect.AddKey(vanish)
	vanish.A = 0x0
	vanishEffect.AddKey(vanish)
	a.effects = append(a.effects, vanishEffect)
	// log.Println("| effect #2: Vanish (30frames) loaded")
	log.Println("| load effects - [DONE]")
}

func (a *Assets) InitWeapons() {
	log.Println(len(a.textures))
	a.weapons[WEAPON_SWORD].InitWeapon(a.textures, NONE, RES_SWORDFX, NONE)
	a.weapons[WEAPON_SWORD].damage = 30
	a.weapons[WEAPON_SWORD].shootRange = 32 * 3
	a.weapons[WEAPON_SWORD].deathAni.scaled = false
	a.weapons[WEAPON_SWORD].deathAni.angle = -1.0
	a.weapons[WEAPON_SWORD].deathAudio = AUDIO_SWORD_HIT
	log.Println("here")

	a.weapons[WEAPON_MONSTER_CLAW].InitWeapon(a.textures, NONE, RES_CLAWFX2, NONE)
	a.weapons[WEAPON_MONSTER_CLAW].wp = WEAPON_SWORD_RANGE
	a.weapons[WEAPON_MONSTER_CLAW].shootRange = 32*3 + 16
	a.weapons[WEAPON_MONSTER_CLAW].damage = 24
	a.weapons[WEAPON_MONSTER_CLAW].deathAni.angle = -1.0
	a.weapons[WEAPON_MONSTER_CLAW].deathAni.at = AT_CENTER
	a.weapons[WEAPON_MONSTER_CLAW].deathAudio = AUDIO_CLAW_HIT_HEAVY

	a.weapons[WEAPON_FIREBALL].InitWeapon(a.textures, RES_SHINE, RES_HALO_EXPLOSION1, RES_FIREBALL)
	a.weapons[WEAPON_FIREBALL].wp = WEAPON_GUN_RANGE
	a.weapons[WEAPON_FIREBALL].damage = 45
	a.weapons[WEAPON_FIREBALL].effectRange = 50
	a.weapons[WEAPON_FIREBALL].shootRange = 256
	a.weapons[WEAPON_FIREBALL].gap = 180
	a.weapons[WEAPON_FIREBALL].deathAni.angle = -1.0
	a.weapons[WEAPON_FIREBALL].deathAni.at = AT_CENTER
	a.weapons[WEAPON_FIREBALL].birthAni.duration = 24
	a.weapons[WEAPON_FIREBALL].birthAudio = AUDIO_SHOOT
	a.weapons[WEAPON_FIREBALL].deathAudio = AUDIO_FIREBALL_EXP

	a.weapons[WEAPON_THUNDER].InitWeapon(a.textures, RES_BLOODBOUND, RES_THUNDER, NONE)
	a.weapons[WEAPON_THUNDER].wp = WEAPON_SWORD_RANGE
	a.weapons[WEAPON_THUNDER].damage = 80
	a.weapons[WEAPON_THUNDER].shootRange = 128
	a.weapons[WEAPON_THUNDER].gap = 120
	a.weapons[WEAPON_THUNDER].deathAni.angle = -1
	a.weapons[WEAPON_THUNDER].deathAni.scaled = false
	a.weapons[WEAPON_THUNDER].deathAudio = AUDIO_THUNDER

	a.weapons[WEAPON_THUNDER_STAFF].InitWeapon(a.textures, NONE, RES_THUNDER_YELLOW, NONE)
	a.weapons[WEAPON_THUNDER_STAFF].wp = WEAPON_SWORD_RANGE
	a.weapons[WEAPON_THUNDER_STAFF].damage = 50
	a.weapons[WEAPON_THUNDER_STAFF].shootRange = 128
	a.weapons[WEAPON_THUNDER_STAFF].gap = 120
	a.weapons[WEAPON_THUNDER_STAFF].deathAni.angle = -1
	a.weapons[WEAPON_THUNDER_STAFF].deathAni.scaled = false
	a.weapons[WEAPON_THUNDER_STAFF].deathAudio = AUDIO_THUNDER

	a.weapons[WEAPON_ARROW].InitWeapon(a.textures, NONE, RES_HALO_EXPLOSION2, RES_ARROW)
	a.weapons[WEAPON_ARROW].wp = WEAPON_GUN_POINT
	a.weapons[WEAPON_ARROW].gap = 40
	a.weapons[WEAPON_ARROW].damage = 10
	a.weapons[WEAPON_ARROW].shootRange = 200
	a.weapons[WEAPON_ARROW].bulletSpeed = 10
	a.weapons[WEAPON_ARROW].deathAni.angle = -1
	a.weapons[WEAPON_ARROW].deathAni.at = AT_CENTER
	a.weapons[WEAPON_ARROW].flyAni.scaled = false
	a.weapons[WEAPON_ARROW].birthAudio = AUDIO_BOW_FIRE
	a.weapons[WEAPON_ARROW].deathAudio = AUDIO_BOW_HIT

	a.weapons[WEAPON_POWERFUL_BOW].InitWeapon(a.textures, NONE, RES_HALO_EXPLOSION2, RES_ARROW)
	a.weapons[WEAPON_POWERFUL_BOW].wp = WEAPON_GUN_POINT
	a.weapons[WEAPON_POWERFUL_BOW].gap = 60
	a.weapons[WEAPON_POWERFUL_BOW].damage = 25
	a.weapons[WEAPON_POWERFUL_BOW].shootRange = 320
	a.weapons[WEAPON_POWERFUL_BOW].bulletSpeed = 7
	a.weapons[WEAPON_POWERFUL_BOW].deathAni.angle = -1
	a.weapons[WEAPON_POWERFUL_BOW].deathAni.at = AT_CENTER
	a.weapons[WEAPON_POWERFUL_BOW].birthAudio = AUDIO_BOW_FIRE
	a.weapons[WEAPON_POWERFUL_BOW].deathAudio = AUDIO_BOW_HIT
	a.weapons[WEAPON_POWERFUL_BOW].effects[BUFF_ATTACK] = WeaponBuff{chance: 0.5, duration: 240}

	a.weapons[WEAPON_MONSTER_CLAW2].InitWeapon(a.textures, NONE, RES_CLAWFX, NONE)

	a.weapons[WEAPON_THROW_AXE].InitWeapon(a.textures, NONE, RES_CROSS_HIT, RES_AXE)
	a.weapons[WEAPON_THROW_AXE].wp = WEAPON_GUN_POINT
	a.weapons[WEAPON_THROW_AXE].damage = 12
	a.weapons[WEAPON_THROW_AXE].shootRange = 160
	a.weapons[WEAPON_THROW_AXE].bulletSpeed = 10
	a.weapons[WEAPON_THROW_AXE].flyAni.duration = 24
	a.weapons[WEAPON_THROW_AXE].flyAni.angle = -1
	a.weapons[WEAPON_THROW_AXE].flyAni.scaled = false
	a.weapons[WEAPON_THROW_AXE].deathAni.scaled = false
	a.weapons[WEAPON_THROW_AXE].deathAni.at = AT_CENTER
	a.weapons[WEAPON_THROW_AXE].birthAudio = AUDIO_AXE_FLY //res.AUDIO_LIGHT_SHOOT
	a.weapons[WEAPON_THROW_AXE].deathAudio = AUDIO_ARROW_HIT

	a.weapons[WEAPON_MANY_AXES].InitWeapon(a.textures, NONE, RES_CROSS_HIT, RES_AXE)
	a.weapons[WEAPON_MANY_AXES].wp = WEAPON_GUN_POINT_MULTI
	a.weapons[WEAPON_MANY_AXES].shootRange = 180
	a.weapons[WEAPON_MANY_AXES].gap = 70
	a.weapons[WEAPON_MANY_AXES].effectRange = 50
	a.weapons[WEAPON_MANY_AXES].damage = 50
	a.weapons[WEAPON_MANY_AXES].bulletSpeed = 4
	a.weapons[WEAPON_MANY_AXES].flyAni.duration = 24
	a.weapons[WEAPON_MANY_AXES].flyAni.angle = -1
	a.weapons[WEAPON_MANY_AXES].deathAni.at = AT_CENTER
	a.weapons[WEAPON_MANY_AXES].birthAudio = AUDIO_AXE_FLY //res.AUDIO_LIGHT_SHOOT;res.AUDIO_LIGHT_SHOOT
	a.weapons[WEAPON_MANY_AXES].deathAudio = AUDIO_ARROW_HIT

	a.weapons[WEAPON_SOLID].InitWeapon(a.textures, NONE, RES_SOLIDFX, NONE)
	a.weapons[WEAPON_SOLID].deathAni.scaled = false
	a.weapons[WEAPON_SOLID].deathAni.angle = -1
	a.weapons[WEAPON_SOLID].effects[BUFF_SLOWDOWN] = WeaponBuff{chance: 0.3, duration: 180}

	a.weapons[WEAPON_SOLID_GREEN].InitWeapon(a.textures, NONE, RES_SOLID_GREENFX, NONE)
	a.weapons[WEAPON_SOLID_GREEN].shootRange = 96
	a.weapons[WEAPON_SOLID_GREEN].deathAni.scaled = false
	a.weapons[WEAPON_SOLID_GREEN].deathAni.angle = -1
	a.weapons[WEAPON_SOLID_GREEN].effects[BUFF_SLOWDOWN] = WeaponBuff{chance: 0.3, duration: 180}

	a.weapons[WEAPON_SOLID_CLAW].InitWeapon(a.textures, NONE, RES_SOLID_GREENFX, NONE)
	a.weapons[WEAPON_SOLID_CLAW].wp = WEAPON_SWORD_RANGE
	a.weapons[WEAPON_SOLID_CLAW].shootRange = 32*3 + 16
	a.weapons[WEAPON_SOLID_CLAW].damage = 35
	a.weapons[WEAPON_SOLID_CLAW].deathAni.scaled = false
	a.weapons[WEAPON_SOLID_CLAW].deathAni.angle = -1
	a.weapons[WEAPON_SOLID_CLAW].deathAudio = AUDIO_CLAW_HIT_HEAVY
	a.weapons[WEAPON_SOLID_CLAW].effects[BUFF_SLOWDOWN] = WeaponBuff{chance: 0.7, duration: 60}

	a.weapons[WEAPON_ICEPICK].InitWeapon(a.textures, NONE, RES_ICESHATTER, RES_ICEPICK)
	a.weapons[WEAPON_ICEPICK].wp = WEAPON_GUN_RANGE
	a.weapons[WEAPON_ICEPICK].damage = 30
	a.weapons[WEAPON_ICEPICK].effectRange = 50
	a.weapons[WEAPON_ICEPICK].shootRange = 256
	a.weapons[WEAPON_ICEPICK].gap = 180
	a.weapons[WEAPON_ICEPICK].bulletSpeed = 8
	a.weapons[WEAPON_ICEPICK].deathAni.angle = -1
	a.weapons[WEAPON_ICEPICK].flyAni.scaled = false
	a.weapons[WEAPON_ICEPICK].deathAni.at = AT_CENTER
	a.weapons[WEAPON_ICEPICK].effects[BUFF_FROZEN] = WeaponBuff{chance: 0.2, duration: 60}
	a.weapons[WEAPON_ICEPICK].birthAudio = AUDIO_ICE_SHOOT

	a.weapons[WEAPON_PURPLE_BALL].InitWeapon(a.textures, NONE, RES_PURPLE_EXP, RES_PURPLE_BALL)
	a.weapons[WEAPON_PURPLE_BALL].wp = WEAPON_GUN_RANGE
	a.weapons[WEAPON_PURPLE_BALL].damage = 20
	a.weapons[WEAPON_PURPLE_BALL].effectRange = 50
	a.weapons[WEAPON_PURPLE_BALL].shootRange = 256
	a.weapons[WEAPON_PURPLE_BALL].gap = 100
	a.weapons[WEAPON_PURPLE_BALL].bulletSpeed = 6
	a.weapons[WEAPON_PURPLE_BALL].deathAni.angle = -1
	a.weapons[WEAPON_PURPLE_BALL].deathAni.scaled = false
	a.weapons[WEAPON_PURPLE_BALL].flyAni.scaled = false
	a.weapons[WEAPON_PURPLE_BALL].deathAni.at = AT_CENTER
	a.weapons[WEAPON_PURPLE_BALL].birthAudio = AUDIO_ICE_SHOOT
	a.weapons[WEAPON_PURPLE_BALL].deathAudio = AUDIO_ARROW_HIT

	a.weapons[WEAPON_PURPLE_STAFF].InitWeapon(a.textures, NONE, RES_PURPLE_EXP, RES_PURPLE_FIRE_BALL)
	a.weapons[WEAPON_PURPLE_STAFF].wp = WEAPON_GUN_POINT_MULTI
	a.weapons[WEAPON_PURPLE_STAFF].damage = 45
	a.weapons[WEAPON_PURPLE_STAFF].effectRange = 50
	a.weapons[WEAPON_PURPLE_STAFF].shootRange = 256
	a.weapons[WEAPON_PURPLE_STAFF].gap = 100
	a.weapons[WEAPON_PURPLE_STAFF].bulletSpeed = 7
	a.weapons[WEAPON_PURPLE_STAFF].deathAni.angle = -1
	a.weapons[WEAPON_PURPLE_STAFF].deathAni.scaled = false
	a.weapons[WEAPON_PURPLE_STAFF].flyAni.scaled = true
	a.weapons[WEAPON_PURPLE_STAFF].deathAni.at = AT_CENTER
	a.weapons[WEAPON_PURPLE_STAFF].birthAudio = AUDIO_ICE_SHOOT
	a.weapons[WEAPON_PURPLE_STAFF].deathAudio = AUDIO_ARROW_HIT

	a.weapons[WEAPON_HOLY_SWORD].InitWeapon(a.textures, NONE, RES_GOLDEN_CROSS_HIT, NONE)
	a.weapons[WEAPON_HOLY_SWORD].wp = WEAPON_SWORD_RANGE
	a.weapons[WEAPON_HOLY_SWORD].damage = 30
	a.weapons[WEAPON_HOLY_SWORD].shootRange = 32 * 4
	a.weapons[WEAPON_HOLY_SWORD].effects[BUFF_DEFENCE] = WeaponBuff{chance: 0.6, duration: 180}

	a.weapons[WEAPON_ICE_SWORD].InitWeapon(a.textures, NONE, RES_ICESHATTER, NONE)
	a.weapons[WEAPON_ICE_SWORD].wp = WEAPON_SWORD_RANGE
	a.weapons[WEAPON_ICE_SWORD].shootRange = 32*3 + 16
	a.weapons[WEAPON_ICE_SWORD].damage = 80
	a.weapons[WEAPON_ICE_SWORD].gap = 30
	a.weapons[WEAPON_ICE_SWORD].deathAni.angle = -1
	a.weapons[WEAPON_ICE_SWORD].deathAni.at = AT_CENTER
	a.weapons[WEAPON_ICE_SWORD].effects[BUFF_FROZEN] = WeaponBuff{chance: 0.6, duration: 80}
	a.weapons[WEAPON_ICE_SWORD].deathAudio = AUDIO_SWORD_HIT
	log.Println("| init weapons - [DONE]")
}

func (a *Assets) InitCommonSprites() {
	// Heroes
	a.commonSprites.AddSprite(SPRITE_KNIGHT, &a.weapons[WEAPON_SWORD], a.textures[RES_KNIGHT_M], 150)
	a.commonSprites.AddSprite(SPRITE_ELF, &a.weapons[WEAPON_ARROW], a.textures[RES_ELF_M], 100)
	a.commonSprites.AddSprite(SPRITE_WIZZARD, &a.weapons[WEAPON_FIREBALL], a.textures[RES_WIZZARD_M], 95)
	a.commonSprites.AddSprite(SPRITE_LIZARD, &a.weapons[WEAPON_MONSTER_CLAW], a.textures[RES_ZIGGY_M], 120)

	// Baddies
	a.commonSprites.AddSprite(SPRITE_TINY_ZOMBIE, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_TINY_ZOMBIE], 50)
	a.commonSprites.AddSprite(SPRITE_GOBLIN, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_GOBLIN], 100)
	a.commonSprites.AddSprite(SPRITE_IMP, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_IMP], 100)
	a.commonSprites.AddSprite(SPRITE_SKELET, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_SKELET], 100)
	a.commonSprites.AddSprite(SPRITE_MUDDY, &a.weapons[WEAPON_SOLID], a.textures[RES_MUDDY], 150)
	a.commonSprites.AddSprite(SPRITE_SWAMPY, &a.weapons[WEAPON_SOLID_GREEN], a.textures[RES_SWAMPY], 150)
	a.commonSprites.AddSprite(SPRITE_ZOMBIE, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_ZOMBIE], 120)
	a.commonSprites.AddSprite(SPRITE_ICE_ZOMBIE, &a.weapons[WEAPON_ICEPICK], a.textures[RES_ICE_ZOMBIE], 120)
	a.commonSprites.AddSprite(SPRITE_MASKED_ORC, &a.weapons[WEAPON_THROW_AXE], a.textures[RES_MASKED_ORC], 120)
	a.commonSprites.AddSprite(SPRITE_ORC_WARRIOR, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_ORC_WARRIOR], 200)
	a.commonSprites.AddSprite(SPRITE_ORC_SHAMAN, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_ORC_SHAMAN], 120)
	a.commonSprites.AddSprite(SPRITE_NECROMANCER, &a.weapons[WEAPON_PURPLE_BALL], a.textures[RES_NECROMANCER], 120)
	a.commonSprites.AddSprite(SPRITE_WOGOL, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_WOGOL], 150)
	a.commonSprites.AddSprite(SPRITE_CHROT, &a.weapons[WEAPON_MONSTER_CLAW2], a.textures[RES_CHORT], 150)
	a.commonSprites.AddSprite(SPRITE_GREEN_HOOD_SKEL, &a.weapons[WEAPON_PURPLE_BALL], a.textures[RES_GREEN_HOOD_SKEL], 150)

	a.commonSprites.AddSprite(SPRITE_BIG_ZOMBIE, &a.weapons[WEAPON_THUNDER], a.textures[RES_BIG_ZOMBIE], 3000)
	a.commonSprites.GetSprite(SPRITE_BIG_ZOMBIE).dropRate = 100

	a.commonSprites.AddSprite(SPRITE_ORGRE, &a.weapons[WEAPON_MANY_AXES], a.textures[RES_ORGRE], 3000)
	a.commonSprites.GetSprite(SPRITE_ORGRE).dropRate = 100

	a.commonSprites.AddSprite(SPRITE_BIG_DEMON, &a.weapons[WEAPON_THUNDER], a.textures[RES_BIG_DEMON], 2500)
	a.commonSprites.GetSprite(SPRITE_BIG_DEMON).dropRate = 100
	log.Println("| init common sprites - [DONE]")
}

func LoadAssets() (*Assets, error) {
	log.Println("# LOAD MEDIA:")

	var err error
	a := &Assets{
		commonSprites: CreateSpriteList(),
	}

	a.textures, err = LoadTileSet()
	if err != nil {
		return nil, err
	}

	a.font, err = LoadFont()
	if err != nil {
		return nil, err
	}

	a.texts = LoadTextSet(a.font)

	a.audio, err = LoadAudio()
	if err != nil {
		return nil, err
	}

	a.InitCommonEffects()
	a.InitWeapons()
	a.InitCommonSprites()
	return a, nil
}
