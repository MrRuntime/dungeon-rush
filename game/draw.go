package game

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Rect struct {
	x, y int
	w, h int
}

type Point struct {
	x, y int
}

func DrawCenteredText(screen *ebiten.Image, t *Text, x int, y int, scale float64) {
	width := int(float64(t.width)*scale + 0.5)
	height := int(float64(t.height)*scale + 0.5)
	op := &text.DrawOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(x-width/2), float64(y-height/2))
	text.Draw(screen, t.text, t.origin, op)
}

func DrawText(screen *ebiten.Image, t *Text, x int, y int, scale float64) {
	op := &text.DrawOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(x), float64(y))
	text.Draw(screen, t.text, t.origin, op)
}

func DrawAnimation(a *Animation) {
	if a == nil {
		return
	}
	ani := a

	UpdateAnimationFromBind(ani)

	width := ani.origin.width
	height := ani.origin.height
	poi := Point{
		x: ani.origin.width,
		y: ani.origin.height / 2,
	}

	if ani.scaled {
		width *= SCALE_FACTOR
		height *= SCALE_FACTOR
	}

	dst := Rect{
		x: ani.x - width/2,
		y: ani.y - height,
		w: width,
		h: height,
	}

	if ani.at == AT_TOP_LEFT {
		dst.x = ani.x
		dst.y = ani.y
	} else if ani.at == AT_CENTER {
		dst.x = ani.x - width/2
		dst.y = ani.y - height/2
		poi.x = ani.origin.width / 2
	} else if ani.at == AT_BOTTOM_LEFT {
		dst.x = ani.x
		dst.y = ani.y + UNIT - height - 3
	}
	if ani.effect != nil {
		SetEffect(ani.origin, ani.effect)
		ani.effect.currentFrame = ani.effect.currentFrame % ani.effect.duration
	}

	//     std.debug.assert(ani.duration >= ani.origin.frames);

	// IMPL:
	// rc: stage just means which animation frame.
	// stage := 0
	// if ani.origin.frames > 1 {
	// 	interval := float64(ani.duration) / float64(ani.origin.frames)
	// 	stage = int(math.Floor(float64(ani.currentFrame) / interval))
	// }

	//     _ = c.SDL_RenderCopyEx(
	//         renderer,
	//         ani.origin.origin,
	//         &(ani.origin.crops[stage]),
	//         &dst,
	//         ani.angle,
	//         &poi,
	//         ani.flip,
	//     );

	if ani.effect != nil {
		UnsetEffect(ani.origin)
	}

	//     // When left-shift key is held down (eXtreme Developer Mode)
	//     // Show the various debug bounding boxes of the sprites.
	//     const state = c.SDL_GetKeyboardState(null);
	//     if (state[c.SDL_SCANCODE_SPACE] > 0) {
	//         if (ani.at == .AT_BOTTOM_CENTER) {
	//             var tmp: c.SDL_Rect = undefined;
	//             var fake: spr.Sprite = undefined;
	//             fake.ani = ani;

	//             // Debug draw bounded box
	//             tmp = hlp.getSpriteBoundBox(&fake);
	//             _ = c.SDL_SetRenderDrawColor(renderer, 0, 255, 0, 200);
	//             _ = c.SDL_RenderDrawRect(renderer, &tmp);

	//             // Debug draw feet box.
	//             tmp = hlp.getSpriteFeetBox(&fake);
	//             _ = c.SDL_SetRenderDrawColor(renderer, 255, 0, 0, 200);
	//             _ = c.SDL_RenderDrawRect(renderer, &tmp);

	//	        // Debug draw dst box.
	//	        _ = c.SDL_SetRenderDrawColor(renderer, 0, 0, 255, 200);
	//	        _ = c.SDL_RenderDrawRect(renderer, &dst);
	//	    }
	//	}
}

type SortByY []*Animation

func (a SortByY) Len() int           { return len(a) }
func (a SortByY) Less(i, j int) bool { return a[i].y < a[j].y }
func (a SortByY) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func DrawAnimationLinkListWithSort(al *AnimationsList) {
	buffer := make([]*Animation, 0, al.Len())

	count := 0
	for e := al.Front(); e != nil; e = e.Next() {
		buffer = append(buffer, e.Value.(*Animation))
		count++
	}

	sort.Sort(SortByY(buffer))

	// 4. Render that shit in reverse order.
	for count > 0 {
		count -= 1
		DrawAnimation(buffer[count])
	}
}

func DrawAnimationLinkList(al *AnimationsList) {
	for p := al.Front(); p != nil; p = p.Next() {
		DrawAnimation(p.Value.(*Animation))
	}
}

func DrawUI(screen *ebiten.Image, assets *Assets) {
	// screen.Fill(color.RGBA{0x19, 0x11, 0x17, 0xff})
	screen.Fill(color.RGBA{0x19, 0x24, 0x17, 0xff})
	for i := range ANIMATION_LINK_LIST_NUM {
		UpdateAnimationLinkList(assets.animations[i])
		if i == LIST_SPRITE_ID {
			DrawAnimationLinkListWithSort(assets.animations[i])
		} else {
			DrawAnimationLinkList(assets.animations[i])
		}
	}
}
