package common

func FixSize(size uint) uint {
	size |= size >> 1
	size |= size >> 2
	size |= size >> 4
	size |= size >> 8
	size |= size >> 16

	return size+1
}

func IsPowOf2(num uint)  bool {
	return (num&(num-1)) == 0
}

func MaxI(x, y int) (maxI int) {
	maxI = x
	if x < y {
		maxI = y
	}

	return
}

func MinI(x, y int) (minI int) {
	minI = x
	if x > y {
		minI = y
	}

	return
}