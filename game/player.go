package game

import "container/list"

type PlayerType int

const (
	LOCAL PlayerType = iota
	REMOTE
	COMPUTER
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

type Score struct {
	damage int
	stand  int
	killed int
	got    int // Bumped, when a snake has a hero added to the chain
	rank   float64
}

type Snake struct {
	sprites  *list.List
	moveStep int
	team     int

	// num is how many sprites (heroes or baddies) form the snake.
	// num: int,
	buffs      [BUFF_END]int // r.c. - verified these should stay integers
	score      *Score
	playerType PlayerType
}

func CreateSnake(step int, team int, playerType PlayerType) *Snake {
	snake := &Snake{
		moveStep:   step,
		team:       team,
		sprites:    list.New(),
		score:      &Score{},
		playerType: playerType,
		//num: 0,
	}
	return snake
}
