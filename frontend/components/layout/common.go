package layout

func Clamp(x0, x1, min, max int) (int, int) {
	size := x1 - x0
	totalSize := max - min
	if size <= totalSize {
		if x0 < min {
			return min, min + size
		} else if x1 < max {
			return x0, x1
		} else {
			return max - size, max
		}
	} else {
		return min, max
	}
}
