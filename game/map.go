package game

import (
	"log"
)

//go:generate stringer -type=BlockType -linecomment
type BlockType int

const (
	BLOCK_TRAP BlockType = iota
	BLOCK_WALL
	BLOCK_FLOOR
	BLOCK_EXIT
)

type Block struct {
	bp     BlockType
	x      int
	y      int
	bid    int  // block id
	enable bool // Used for trap block
	ani    *Animation
}

type GameMap = [MAP_SIZE][MAP_SIZE]Block

var Game_Map GameMap

var Has_Map [MAP_SIZE][MAP_SIZE]bool
var exitX int = -1
var exitY int = -1

func initBlock(block *Block, bp BlockType, x int, y int, bid int, enable bool) {
	block.x = x
	block.y = y
	block.bid = bid
	block.enable = enable

	if bp == BLOCK_TRAP {
		var floor_spike int
		if enable {
			floor_spike = FLOOR_SPIKE_ENABLED
		} else {
			floor_spike = FLOOR_SPIKE_DISABLED
		}
		block.ani = createAnimation(Textures[floor_spike], nil, LOOP_INFI, 1, x, y, FLIP_NONE, 0, AT_TOP_LEFT)
	} else if bp == BLOCK_EXIT {
		var floor_spike int
		if enable {
			floor_spike = FLOOR_EXIT
		} else {
			floor_spike = FLOOR_2
		}
		block.ani = createAnimation(Textures[floor_spike], nil, LOOP_INFI, 1, x, y, FLIP_NONE, 0, AT_TOP_LEFT)
	} else {
		block.ani = createAnimation(Textures[bid], nil, LOOP_INFI, 1, x, y, FLIP_NONE, 0, AT_TOP_LEFT)
	}
}

func InitBlankMap(w, h int) {
	// clearMapGenerator();
	si := N/2 - w/2
	sj := M/2 - h/2
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			ii := si + i
			jj := sj + j
			Has_Map[ii][jj] = true
			initBlock(&Game_Map[ii][jj], BLOCK_FLOOR, ii*UNIT, jj*UNIT, FLOOR_1, false)
		}
	}
	// decorateMap();
	log.Println("| init blank map - [DONE]")
	// fmt.Println(Game_Map)
}

func pushMapToRender() {
	cpa := createAndPushAnimation
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !Has_Map[i][j] {
				if inr(j+1, 0, M-1) && Has_Map[i][j+1] {
					if inr(i+1, 0, N-1) && Has_Map[i+1][j] {
						_ = cpa(LIST_MAP_ID, WALL_CORNER_FRONT_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, j*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)

						_ = cpa(LIST_MAP_ID, WALL_CORNER_BOTTOM_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)
					} else if inr(i-1, 0, N-1) && Has_Map[i-1][j] {
						_ = cpa(LIST_MAP_ID, WALL_CORNER_FRONT_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, j*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)

						_ = cpa(LIST_MAP_ID, WALL_CORNER_BOTTOM_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)
					} else {
						var bid int
						if randDouble() < MAP_HOW_OLD*5 {
							bid = WALL_HOLE_1 + randInt(0, 1)
						} else {
							bid = WALL_MID
						}
						if randDouble() < MAP_WALL_HOW_DECORATED {
							bid = WALL_BANNER_RED + randInt(0, 3)
						}
						_ = cpa(LIST_MAP_ID, int(bid), nil, LOOP_INFI, 1,
							i*UNIT, j*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)

						_ = cpa(LIST_MAP_ID, WALL_TOP_MID, nil, LOOP_INFI,
							1, i*UNIT, (j-1)*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)
					}
				}
				if inr(j-1, 0, M-1) && Has_Map[i][j-1] {
					var bid int

					if randDouble() < MAP_HOW_OLD*2 {
						bid = WALL_HOLE_1 + randInt(0, 1)
					} else {
						bid = WALL_MID
					}
					_ = cpa(LIST_MAP_ID, int(bid), nil, LOOP_INFI, 1,
						i*UNIT, j*UNIT, FLIP_NONE, 0, AT_TOP_LEFT)

					if Has_Map[i-1][j] {
						_ = cpa(LIST_MAP_FOREWALL, WALL_CORNER_TOP_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					} else if Has_Map[i+1][j] {
						_ = cpa(LIST_MAP_FOREWALL, WALL_CORNER_TOP_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					} else {
						_ = cpa(LIST_MAP_FOREWALL, WALL_TOP_MID,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
				}
				if inr(i+1, 0, N-1) && Has_Map[i+1][j] {
					if inr(j+1, 0, M-1) && Has_Map[i][j+1] {
						// just do not render
					} else {
						_ = cpa(LIST_MAP_ID, WALL_SIDE_MID_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, j*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
					if !Has_Map[i+1][j+1] {
						_ = cpa(LIST_MAP_ID, WALL_SIDE_FRONT_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, (j+1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
					if !Has_Map[i+1][j-1] {
						_ = cpa(LIST_MAP_ID, WALL_SIDE_MID_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)

						_ = cpa(LIST_MAP_ID, WALL_SIDE_TOP_LEFT,
							nil, LOOP_INFI, 1, i*UNIT, (j-2)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
				}
				if inr(i-1, 0, N-1) && Has_Map[i-1][j] {
					if inr(j+1, 0, M-1) && Has_Map[i][j+1] {
						// just do not render
					} else {
						_ = cpa(LIST_MAP_ID, WALL_SIDE_MID_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, j*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
					if !Has_Map[i-1][j+1] {
						_ = cpa(LIST_MAP_ID, WALL_SIDE_FRONT_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, (j+1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
					if !Has_Map[i-1][j-1] {
						_ = cpa(LIST_MAP_ID, WALL_SIDE_MID_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, (j-1)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)

						_ = cpa(LIST_MAP_ID, WALL_SIDE_TOP_RIGHT,
							nil, LOOP_INFI, 1, i*UNIT, (j-2)*UNIT,
							FLIP_NONE, 0, AT_TOP_LEFT)
					}
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !Has_Map[i][j] {
				continue
			}
			animationsList[LIST_MAP_ID].PushBack(Game_Map[i][j].ani)
		}
	}
	log.Println("| push map to render - [DONE]")
}
