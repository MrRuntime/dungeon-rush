// Code generated by "stringer -type=LoopType -linecomment"; DO NOT EDIT.

package game

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LOOP_ONCE-0]
	_ = x[LOOP_INFI-1]
	_ = x[LOOP_LIFESPAN-2]
}

const _LoopType_name = "LOOP_ONCELOOP_INFILOOP_LIFESPAN"

var _LoopType_index = [...]uint8{0, 9, 18, 31}

func (i LoopType) String() string {
	if i < 0 || i >= LoopType(len(_LoopType_index)-1) {
		return "LoopType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LoopType_name[_LoopType_index[i]:_LoopType_index[i+1]]
}
