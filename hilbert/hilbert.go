package hilbert

// XYToD converts a two-dimensional cell position with coordinates
// (x, y) to a one-dimensional distance (d) along a discrete Hilbert
// curve. The curve is constructed by dividing a square into n X n
// cells.
//
// The cell count n must be a power of 2, and the coordinates x and y
// must range from (0, 0) representing the cell at the lower left-hand
// corner of the square, to (n-1, n-1) representing the cell at the
// upper right-hand corner of the square.
//
// The return value d is a number in the range [0, n^2-1].
//
// The complementary function DToXY performs the inverse mapping from
// one-dimensional distance (d) back to two-dimensional position (x, y).
func XYToD(n, x, y int) (d int) {
	for s:=n/2; s>0; s/=2 {
		var rx, ry int
		if x&s > 0 {
			rx = 1
		}
		if y&s > 0 {
			ry = 1
		}
		d += s * s * ((3 * rx) ^ ry)
		rot(s, rx, ry, &x, &y)
	}
	return
}

// DToXY converts a one-dimensional distance (d) along a discrete
// Hilbert curve to a two-dimensional cell position with coordinates
// (x, y).
//
// The cell count n must be a power of 2, and the distance d must range
// from 0, representing the cell in the lower left-hand corner of the
// square, to n^2-1, representing the cell in the upper right-hand
// corner of the square.
//
// The return value (x, y) where x and y are between 0 and n-1, is the
// position of the cell to which d corresponds, where cell (0, 0)
// represents cell at the lower left-hand corner of the square and
// (n-1, n-1) represents the cell at the upper right-hand corner of the
// square.
//
// The complementary function XYToD performs the inverse transformation.
func DToXY(n, d int) (x, y int) {
	t := d
	for s:=1; s<n; s*=2 {
		rx := 1 & (t/2)
		ry := 1 & (t ^ rx)
		rot(s, rx, ry, &x, &y)
		x += s * rx
		y += s * ry
		t /= 4
	}
	return
}

func rot(n, rx, ry int, x, y *int) {
	if ry == 0 {
		if rx == 1 {
			*x = n-1 - *x
			*y = n-1 - *y
		}
		*x, *y = *y, *x // Swap x and y
	}
}
