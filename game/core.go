package game

func AppendSpriteToSnake(
	assets *Assets,
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
			assets.commonSprites.GetSprite(spriteId).ani.origin.width*SCALE_FACTOR) / 2
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

	sprite := assets.commonSprites.GetCopy(spriteId)
	sprite.x = newX
	sprite.y = newY
	sprite.direction = direction

	if direction == LEFT {
		sprite.face = LEFT
	}

	if snakeHead != nil {
		sprite.direction = snakeHead.direction
		sprite.face = snakeHead.face
		sprite.ani.currentFrame = snakeHead.ani.currentFrame
	}

	// insert the sprite
	// node.data = sprite;
	// tps.pushLinkNodeAtHead(snake.sprites, node);
	snake.sprites.PushFront(sprite)

	// push ani
	// ren.pushAnimationToRender(ren.RENDER_LIST_SPRITE_ID, sprite.ani);
	assets.animations[LIST_SPRITE_ID].PushBack(sprite.ani)

	// r.c. - I think the buffs array should be booleans (possibly, confirm later)
	// Confirmed they should not be booleans, because they are counted down.
	// if (snake.buffs[tps.BUFF_DEFENCE] > 0) {
	//     shieldSprite(sprite, snake.buffs[tps.BUFF_DEFENCE]);
	// }
	if snake.buffs[BUFF_DEFENCE] > 0 {
		ShieldSprite(assets, sprite, snake.buffs[BUFF_DEFENCE])
	}
}

func ShieldSprite(assets *Assets, sprite *Sprite, duration int) {
	effect := assets.effects[EFFECT_BLINK].copy()

	ani := CreateAndPushAnimation(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_HOLY_SHIELD],
		effect,
		LOOP_LIFESPAN,
		40,
		sprite.x,
		sprite.y,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)

	bindAnimationToSprite(ani, sprite, true)
	ani.lifeSpan = duration
}

func updateAnimationFromBind(ani *Animation) {
	if bnd := ani.bindTo; bnd != nil {
		ani.x = bnd.x
		ani.y = bnd.y
		ani.flip = bnd.ani.flip
	}
}

// / associates an animation to track and render following the sprites x,y position. If isStrong
// / is true, then the animation should be removed should the sprite need to be destroyed, ie: died.
func bindAnimationToSprite(ani *Animation, sprite *Sprite, isStrong bool) {
	ani.bindTo = sprite
	ani.dieWithBind = isStrong
	updateAnimationFromBind(ani)
}
