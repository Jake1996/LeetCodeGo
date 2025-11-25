package solutions

// https://leetcode.com/problems/rotating-the-box/
func RotateTheBox(boxGrid [][]byte) [][]byte {
	n := len(boxGrid)
	if n == 0 {
		return [][]byte{}
	}
	m := len(boxGrid[0])
	out := make([][]byte, m)
	for i := range m {
		out[i] = make([]byte, n)
	}
	for i := range n {
		bi := m - 1
		oi := m - 1
		for bi >= 0 {
			switch boxGrid[n-i-1][bi] {
			case '#': // stone
				out[oi][i] = '#'
				oi--
				bi--
			case '*': // obstacle
				oi = bi
				out[oi][i] = '*'
				oi--
				bi--
			case '.': // empty
				bi--
			}
		}
	}
	for i := range m {
		for j := range n {
			if out[i][j] == 0 {
				out[i][j] = '.'
			}
		}
	}
	return out
}
