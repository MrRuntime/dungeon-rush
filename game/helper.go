package game

func inr(val, lower, upper int) bool {
	return lower <= val && val <= upper
}

func randDouble() float64 {
	return float64(prngRand()) / float64(PRNG_MAX)
}

// Generates a random int from l to r inclusive on both ends.
func randInt(l, r int) int {
	rdm := int(prngRand())
	return (rdm % (r - l + 1)) + l // Generate a number between l and r
}
