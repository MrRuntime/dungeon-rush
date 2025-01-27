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

func InitCommonSprites() {

}
