package N_Queens_II

func totalNQueens(n int) int {
	if n < 4 && n != 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	mark := NewMark(n)
	result := 0

	backTracking(&result, n, 0, mark)

	return result
}

func backTracking(result *int, n, row int, m *mark) {
	if row == n {
		*result++
		return
	}

	for c := 0; c < n; c++ {
		if m.isUnderAttack(row, c) {
			continue
		}

		m.place(row, c)
		backTracking(result, n, row+1, m)
		m.remove(row, c)
	}

}

type mark struct {
	n       int
	hills   []bool
	dales   []bool
	rows    []bool
	columns []bool
}

func NewMark(n int) *mark {
	return &mark{
		n:       n,
		hills:   make([]bool, 2*n-1),
		dales:   make([]bool, 2*n),
		rows:    make([]bool, n),
		columns: make([]bool, n),
	}
}

func (m *mark) getHillAndDale(row, column int) (int, int) {
	return row + column, row - column + m.n
}

func (m *mark) isUnderAttack(row, column int) bool {
	hill, dale := m.getHillAndDale(row, column)

	return m.rows[row] || m.columns[column] || m.hills[hill] || m.dales[dale]
}

func (m *mark) place(row, column int) {
	hill, dale := m.getHillAndDale(row, column)

	m.rows[row] = true
	m.columns[column] = true
	m.hills[hill] = true
	m.dales[dale] = true
}

func (m *mark) remove(row, column int) {
	hill, dale := m.getHillAndDale(row, column)

	m.rows[row] = false
	m.columns[column] = false
	m.hills[hill] = false
	m.dales[dale] = false

}
