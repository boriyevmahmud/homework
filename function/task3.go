func vote(x int, y int, z int) int {
	var a, b int
	if x == 0 {
		a += 1
	} else {
		b += 1
	}
	if y == 0 {
		a += 1
	} else {
		b += 1
	}
	if z == 0 {
		a += 1
	} else {
		b += 1
	}
	if a > b {
		return 0
	} else {
		return 1
	}
}