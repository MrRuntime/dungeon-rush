package game

const (
	// 	GAME_NAME     = "Dungeon Rush: Go-Edition v1.0 - by @MrRuntime"
	UNIT          = 32
	SCALE_FACTOR  = 2
	SCREEN_FACTOR = 2
	SCREEN_WIDTH  = 1440 * SCREEN_FACTOR
	SCREEN_HEIGHT = 960 * SCREEN_FACTOR
	N             = SCREEN_WIDTH / UNIT
	M             = SCREEN_HEIGHT / UNIT

	NONE = -1

	MAP_SIZE               = 100
	MAP_HOW_OLD            = 0.05
	MAP_WALL_HOW_DECORATED = 0.1

	TITLE = 202

	BUFF_FROZEN             = 0
	BUFF_SLOWDOWN           = 1
	BUFF_DEFENCE            = 2
	BUFF_ATTACK             = 3
	BUFF_END                = 4
	ANIMATION_LINK_LIST_NUM = 16

	WEAPONS_SIZE         = 128
	WEAPON_SWORD         = 0
	WEAPON_MONSTER_CLAW  = 1
	WEAPON_FIREBALL      = 2
	WEAPON_THUNDER       = 3
	WEAPON_ARROW         = 4
	WEAPON_MONSTER_CLAW2 = 5
	WEAPON_THROW_AXE     = 6
	WEAPON_MANY_AXES     = 7
	WEAPON_SOLID         = 8
	WEAPON_SOLID_GREEN   = 9
	WEAPON_ICEPICK       = 10
	WEAPON_ICE_SWORD     = 12
	WEAPON_HOLY_SWORD    = 13
	WEAPON_PURPLE_BALL   = 14
	WEAPON_PURPLE_STAFF  = 15
	WEAPON_THUNDER_STAFF = 16
	WEAPON_SOLID_CLAW    = 17
	WEAPON_POWERFUL_BOW  = 18

	SPRITE_ANIMATION_DURATION = 30
	SPRITE_KNIGHT             = 0

	UI_MAIN_GAP     = 40
	UI_MAIN_GAP_ALT = 22

	AUDIO_SWORD_HIT      = 4
	AUDIO_ARROW_HIT      = 6
	AUDIO_SHOOT          = 7
	AUDIO_FIREBALL_EXP   = 8
	AUDIO_ICE_SHOOT      = 9
	AUDIO_THUNDER        = 12
	AUDIO_CLAW_HIT_HEAVY = 15
	AUDIO_AXE_FLY        = 19
	AUDIO_BOW_FIRE       = 34
	AUDIO_BOW_HIT        = 35

	// 	SHINE                = 176
	// 	HALO_EXPLOSION1      = 165
	// 	THUNDER              = 177
	// 	THUNDER_YELLOW       = 207
	// 	ARROW                = 179
	// 	CLAWFX               = 175
	// 	PURPLE_EXP           = 204
	// 	PURPLE_BALL          = 203
	// 	PURPLE_FIRE_BALL     = 210
	// 	GOLDEN_CROSS_HIT     = 199

	// 	FIREBALL = 167

	// 	SWORD_FX = 174

	// 	HALO_EXPLOSION2 = 166

	// 	CHORT       = 128

	// 	WALL_TOP_LEFT                 = 0
	// 	WALL_TOP_MID                  = 1
	// 	WALL_TOP_RIGHT                = 2
	// 	WALL_MID                      = 4
	// 	WALL_HOLE_1                   = 11
	// 	WALL_HOLE_2                   = 12
	// 	WALL_SIDE_TOP_LEFT            = 35
	// 	WALL_SIDE_TOP_RIGHT           = 36
	// 	WALL_SIDE_MID_LEFT            = 37
	// 	WALL_SIDE_MID_RIGHT           = 38
	// 	WALL_SIDE_FRONT_LEFT          = 39
	// 	WALL_SIDE_FRONT_RIGHT         = 40
	// 	WALL_CORNER_TOP_LEFT          = 41
	// 	WALL_CORNER_TOP_RIGHT         = 42
	// 	WALL_CORNER_LEFT              = 43
	// 	WALL_CORNER_RIGHT             = 44
	// 	WALL_CORNER_BOTTOM_LEFT       = 45
	// 	WALL_CORNER_BOTTOM_RIGHT      = 46
	// 	WALL_CORNER_FRONT_LEFT        = 47
	// 	WALL_CORNER_FRONT_RIGHT       = 48
	// 	WALL_INNER_CORNER_MID_LEFT    = 51
	// 	WALL_INNER_CORNER_MID_RIGHT   = 52
	// 	WALL_INNER_CORNER_T_TOP_LEFT  = 53
	// 	WALL_INNER_CORNER_T_TOP_RIGHT = 54
	// 	WALL_BANNER_RED               = 13

	/// NOTE: these render list act as layers where MAP_ID is drawn first.
	/// and UI_ID is drawn last (on top).
	LIST_MAP_ID = 0
	// 	LIST_MAP_SPECIAL_ID = 1
	// 	LIST_MAP_ITEMS_ID   = 2
	// 	LIST_DEATH_ID       = 3
	LIST_SPRITE_ID    = 4
	LIST_EFFECT_ID    = 5
	LIST_MAP_FOREWALL = 6
	LIST_UI_ID        = 7
)
