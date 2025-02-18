package game

import (
	"log"
)

var cursorPos int = 0

func BaseUI(assets *Assets, w, h int) {
	InitAnimList(assets)
	InitBlankMap(assets, w, h)
	PushMapToRender(assets)
}

func mainUI(assets *Assets) {
	log.Println("| MAIN UI:")
	BaseUI(assets, 30, 12)
	// playBgm(0)
	startY := (SCREEN_HEIGHT / 2) - 70
	startX := (SCREEN_WIDTH / 3) + 125

	cpa := CreateAndPushAnimation

	// Title - Logo
	_ = cpa(
		assets.animations[LIST_UI_ID],
		assets.textures[TITLE],
		nil,
		LOOP_INFI,
		80,
		SCREEN_WIDTH/2,
		(SCREEN_HEIGHT/2)-275,
		FLIP_NONE,
		0,
		AT_CENTER,
	)

	// Knight
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_KNIGHT_M],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX,
		startY,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)

	// Sword effect
	ani := cpa(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_SWORDFX],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP_ALT*2,
		startY-32,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	ani.scaled = false

	// Red bad-guy (Knight enemy)
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_CHORT],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP_ALT*2,
		startY-32,
		FLIP_H,
		0,
		AT_BOTTOM_CENTER,
	)

	startX += UI_MAIN_GAP_ALT * (6 + 2*int(randDouble()))
	startY += UI_MAIN_GAP * (1 + int(randDouble()))

	// Green elf
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_ELF_M],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX,
		startY,
		FLIP_H,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_HALO_EXPLOSION2],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX-int(UI_MAIN_GAP*1.5),
		startY,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_ZOMBIE],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX-int(UI_MAIN_GAP*1.5),
		startY,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)

	startX -= UI_MAIN_GAP_ALT * (1 + 2*int(randDouble()))
	startY += UI_MAIN_GAP * (2 + int(randDouble()))

	// Blue wizard and fireball.
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_WIZZARD_M],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX,
		startY,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_FIREBALL],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP,
		startY,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)

	startX += int(UI_MAIN_GAP_ALT * (18.0 + 4.0*randDouble()))
	startY -= int(UI_MAIN_GAP * (1.0 + 3.0*randDouble()))

	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_ZIGGY_M],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX,
		startY,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_CLAWFX2],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX,
		startY-UI_MAIN_GAP+16,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_MUDDY],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX,
		startY-UI_MAIN_GAP,
		FLIP_H,
		0,
		AT_BOTTOM_CENTER,
	)

	_ = cpa(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_CLAWFX2],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP,
		startY-UI_MAIN_GAP+16,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_SWAMPY],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP,
		startY-UI_MAIN_GAP,
		FLIP_H,
		0,
		AT_BOTTOM_CENTER,
	)

	_ = cpa(
		assets.animations[LIST_EFFECT_ID],
		assets.textures[RES_CLAWFX2],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP,
		startY+16,
		FLIP_NONE,
		0,
		AT_BOTTOM_CENTER,
	)
	_ = cpa(
		assets.animations[LIST_SPRITE_ID],
		assets.textures[RES_SWAMPY],
		nil,
		LOOP_INFI,
		SPRITE_ANIMATION_DURATION,
		startX+UI_MAIN_GAP,
		startY,
		FLIP_H,
		0,
		AT_BOTTOM_CENTER,
	)

	const optsNum = 4
	var opts []*Text
	for i := range optsNum {
		opts = append(opts, &assets.texts[i+6])
	}
	// opt := chooseOptions(assets, optsNum, opts)

	// HERE:
	// ren.blackout();
	// ren.clearRenderer();

	// switch (opt) {
	//     0 => {
	//         if (try chooseLevelUi()) {
	//             try launchLocalGame(1);
	//         }
	//         std.debug.print("option 0 - local game!!\n", .{});
	//     },
	//     1 => {
	//         std.debug.print("option 1 - LAN game - (NOT BUILT OUT)!!\n", .{});
	//     },
	//     2 => {
	//         std.debug.print("option 2 - show ranks (NOT BUILT OUT)!!\n", .{});
	//     },
	//     else => {},
	// }
	// if (opt == optsNum) return;
	// if (opt != 3) {
	//     try mainUi();
	// }
}

func ChooseOptions(assets *Assets, optionsNum int, options []*Text) int {
	cursorPos = 0
	player := CreateSnake(2, 0, LOCAL)

	AppendSpriteToSnake(
		assets,
		player,
		SPRITE_KNIGHT,
		SCREEN_WIDTH/2,
		SCREEN_HEIGHT/2,
		UP,
	)

	// lineGap := FONT_SIZE + FONT_SIZE/2
	// totalHeight := lineGap * (optionsNum - 1)
	// startY := (SCREEN_HEIGHT - totalHeight) / 2

	// HERE:
	// var throttler = th.Throttler.init();
	// while (!moveCursor(optionsNum)) {
	//     if (throttler.shouldWait()) {
	//         continue;
	//     }

	//     const sprite: *spr.Sprite = @alignCast(@ptrCast(player.sprites.first.?.data));
	//     sprite.ani.at = .AT_CENTER;
	//     sprite.x = (res.SCREEN_WIDTH / 2) - @divTrunc(options[@intCast(cursorPos)].width, 2) - (res.UNIT / 2);
	//     sprite.y = startY + cursorPos * lineGap;
	//     ren.updateAnimationOfSprite(sprite);
	//     try ren.renderUi();

	//     const optsNum: usize = @intCast(optionsNum);
	//     for (0..optsNum) |i| {
	//         const ii: c_int = @intCast(i);
	//         _ = ren.renderCenteredText(options[i], res.SCREEN_WIDTH / 2, startY + ii * lineGap, 1);
	//     }

	//     // Wedge in Zig-Edition
	//     // by @deckarep text.
	//     _ = ren.renderCenteredText(&res.texts[17], res.SCREEN_WIDTH / 2, 920 * res.SCREEN_FACTOR, 1);

	//     // Update Screen
	//     c.SDL_RenderPresent(ren.renderer);
	//     ren.renderFrames += 1;

	//     throttler.tick();
	// }

	// aud.playAudio(res.AUDIO_BUTTON1);
	// gm.destroySnake(player);
	// tps.destroyAnimationsByLinkList(&ren.animationsList[ren.RENDER_LIST_SPRITE_ID]);
	return cursorPos
}
