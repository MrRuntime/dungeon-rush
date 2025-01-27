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

func createSnake(step int, team int, playerType PlayerType) *Snake {
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

func appendSpriteToSnake(
	snake *Snake,
	spriteId int,
	x int, // x ,y, dir only matter when empty snake
	y int,
	direction Direction,
) {
	//snake.num += 1;
	snake.score.got += 1
	newX := x
	newY := y

	// at head
	// const node = gAllocator.create(ll.GenericNode) catch unreachable;
	// tps.initLinkNode(node);

	// create a sprite
	var snakeHead *Sprite = nil
	if head := snake.sprites.Front(); head != nil {
		snakeHead = head.Value.(*Sprite)
		newX = snakeHead.x
		newY = snakeHead.y
		delta := (snakeHead.ani.origin.width*SCALE_FACTOR +
			commonSprites[spriteId].ani.origin.width*SCALE_FACTOR) / 2
		if snakeHead.direction == LEFT {
			newX -= delta
		} else if snakeHead.direction == RIGHT {
			newX += delta
		} else if snakeHead.direction == UP {
			newY -= delta
		} else {
			newY += delta
		}
	}
	// const sprite = spr.Sprite.create(&res.commonSprites[@intCast(spriteId)], newX, newY);
	// sprite.direction = direction;
	// if (direction == .LEFT) {
	//     sprite.face = .LEFT;
	// }
	// if (snakeHead) |sh| {
	//     sprite.direction = sh.direction;
	//     sprite.face = sh.face;
	//     sprite.ani.currentFrame = sh.ani.currentFrame;
	// }
	// // insert the sprite
	// node.data = sprite;
	// tps.pushLinkNodeAtHead(snake.sprites, node);

	// // push ani
	// ren.pushAnimationToRender(ren.RENDER_LIST_SPRITE_ID, sprite.ani);

	// // r.c. - I think the buffs array should be booleans (possibly, confirm later)
	// // Confirmed they should not be booleans, because they are counted down.
	// if (snake.buffs[tps.BUFF_DEFENCE] > 0) {
	//     shieldSprite(sprite, snake.buffs[tps.BUFF_DEFENCE]);
	// }
}
