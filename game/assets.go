package game

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
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

type Crop struct {
	x, y int
	w, h int
}

type Texture struct {
	origin *ebiten.Image
	width  int
	height int
	frames int
	crops  []Crop
}

type Text = struct {
	text   string
	width  float64
	heigh  float64
	origin *text.GoTextFace
	color  color.RGBA
}

var (
	Texts         []Text
	Textures      []*Texture
	commonSprites []Sprite
	effects       []Effect
	weapons       [WEAPONS_SIZE]Weapon
	Font          *text.GoTextFaceSource
)

var TextList = []string{
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

var bgms [][]byte

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

var sounds [][]byte

func initTileSet(csvPath string, origin *ebiten.Image) {
	// debugLogger := log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	file, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()

	for _, line := range lines {
		if len(strings.TrimSpace(line[0])) > 0 && line[0] != "#" {
			// name := line[0]
			x, err := strconv.Atoi(line[1])
			y, err := strconv.Atoi(line[2])
			w, err := strconv.Atoi(line[3])
			h, err := strconv.Atoi(line[4])
			fc, err := strconv.Atoi(line[5])
			if err != nil {
				panic(err)
			}
			texture := &Texture{
				origin, w, h, fc, make([]Crop, fc),
			}
			for j := 0; j <= fc-1; j++ {
				texture.crops[j].x = x + j*w
				texture.crops[j].y = y
				texture.crops[j].w = w
				texture.crops[j].h = h
			}
			Textures = append(Textures, texture)
			// debugLogger.Println("name:", name, " x:", x, " y:", y, " w:", w, " h:", h, " fc:", fc)
		}
	}
}

func loadTileSet() {
	for _, it := range tilesetName {
		csvPath := fmt.Sprintf("assets/images/%s", it)
		imgPath := fmt.Sprintf("assets/images/%s.png", it)
		origin, _, _ := ebitenutil.NewImageFromFile(imgPath)
		initTileSet(csvPath, origin)
	}
	log.Println("| load tileset - [DONE]")
}

func loadFont() {
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

	Font = s
	log.Println("| load font - [DONE]")
}

func loadTextSet() {
	// TODO: handle error
	for _, str := range TextList {
		f := &text.GoTextFace{Source: Font, Size: 24}
		w, h := text.Measure(str, f, 0)
		Texts = append(Texts, Text{
			origin: f,
			text:   str,
			width:  w,
			heigh:  h,
			color:  color.RGBA{0xff, 0xff, 0xff, 0xff},
		})
	}
	log.Println("| load textset - [DONE]")
}

func loadOgg(audioPath string) {
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

	// fmt.Printf("[%v]: %v\n", file.Name(), b)
	bgms = append(bgms, b)
}

func loadWav(audioPath string) {
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

	// fmt.Printf("[%v]: %v\n", file.Name(), b)
	sounds = append(sounds, b)
}

func LoadAudio() {
	for _, it := range bgmsPath {
		loadOgg(it)
	}
	for _, it := range soundfxList {
		loadWav(fmt.Sprintf("assets/audio/%s", it))
	}
	log.Println("| load audio - [DONE]")
}

func initCommonEffects() {
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
	effects = append(effects, deathEffect)
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
	effects = append(effects, blinkEffect)
	// log.Println("| effect #1: Blink (white) (30frames) loaded")

	vanishEffect := CreateEffect(30, 2, ebiten.BlendDestinationOver)
	vanish := color.RGBA{0xff, 0xff, 0xff, 0xff}
	vanishEffect.AddKey(vanish)
	vanish.A = 0x0
	vanishEffect.AddKey(vanish)
	effects = append(effects, vanishEffect)
	// log.Println("| effect #2: Vanish (30frames) loaded")
	log.Println("| load effects - [DONE]")
}

func LoadMedia() error {
	log.Println("# LOAD MEDIA:")
	initCommonEffects()
	loadTileSet()
	loadFont()
	loadTextSet()
	InitWeapons()
	InitCommonSprites()
	LoadAudio()
	fmt.Println()
	return nil
}
