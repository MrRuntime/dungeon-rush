package game

import (
	"log"
)

var cursorPos int = 0

func baseUI(w, h int) {
	initAnimList() //init array of link list
	InitBlankMap(w, h)
	pushMapToRender()
}

func mainUI() {
	log.Println("| MAIN UI:")
	baseUI(30, 12)
	// playBgm(0)
	startY := (SCREEN_HEIGHT / 2) - 70
	startX := (SCREEN_WIDTH / 3) + 125

	cpa := createAndPushAnimation

	// Title - Logo
	_ = cpa(
		LIST_UI_ID,
		TITLE,
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
		LIST_SPRITE_ID,
		RES_KNIGHT_M,
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
		LIST_EFFECT_ID,
		RES_SWORDFX,
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
		LIST_SPRITE_ID,
		RES_CHORT,
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
		LIST_SPRITE_ID,
		RES_ELF_M,
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
		LIST_EFFECT_ID,
		RES_HALO_EXPLOSION2,
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
		LIST_SPRITE_ID,
		RES_ZOMBIE,
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
		LIST_SPRITE_ID,
		RES_WIZZARD_M,
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
		LIST_EFFECT_ID,
		RES_FIREBALL,
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
		LIST_SPRITE_ID,
		RES_ZIGGY_M,
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
		LIST_EFFECT_ID,
		RES_CLAWFX2,
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
		LIST_SPRITE_ID,
		RES_MUDDY,
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
		LIST_EFFECT_ID,
		RES_CLAWFX2,
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
		LIST_SPRITE_ID,
		RES_SWAMPY,
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
		LIST_EFFECT_ID,
		RES_CLAWFX2,
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
		LIST_SPRITE_ID,
		RES_SWAMPY,
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
	var opts [optsNum]*Text
	for i := 0; i < optsNum; i++ {
		opts[i] = &Texts[i+6]
	}
	// const opt = try chooseOptions(optsNum, &opts);

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

func chooseOptions(optionsNum int, options []*Text) int {
	cursorPos = 0
	// player := createSnake(2, 0, LOCAL)

	// gm.appendSpriteToSnake(
	//     player,
	//     res.SPRITE_KNIGHT,
	//     res.SCREEN_WIDTH / 2,
	//     res.SCREEN_HEIGHT / 2,
	//     .UP,
	// );
	// const lineGap: c_int = res.FONT_SIZE + res.FONT_SIZE / 2;
	// const totalHeight: c_int = lineGap * (optionsNum - 1);
	// const startY: c_int = @divTrunc((res.SCREEN_HEIGHT - totalHeight), 2);

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
