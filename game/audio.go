package game

var nowBgmId = -1

func playBgm(id int) {
	if nowBgmId == id {
		return
	}

	if nowBgmId == -1 {
		// play(bgms[id], -1)
	} else {
		// fadeIn(bgms[id], -1, BGM_FADE_DURATION)
	}
	nowBgmId = id
}
