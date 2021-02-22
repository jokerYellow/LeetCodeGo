package _66_toeplitz_matrix

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	if m == 0 {
		return true
	}
	n := len(matrix[0])
	for i := 0; i < m-1; i++ {
		t := matrix[i][0 : n-1]
		t1 := matrix[i+1][1:n]
		if !equal(t, t1) {
			return false
		}
	}
	return true
}

func equal(n1, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}
	for i, e := range n1 {
		if e != n2[i] {
			return false
		}
	}
	return true
}
