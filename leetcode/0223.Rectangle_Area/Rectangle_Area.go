package Rectangle_Area

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if C < A {
		C, A = A, C
	}

	if D < B {
		D, B = B, D
	}

	if G < E {
		G, E = E, G
	}

	if H < F {
		H, F = F, H
	}

	rectA, rectB := (C-A)*(D-B), (G-E)*(H-F)
	if E > C || A > G || B > H || F > D {
		return rectA + rectB
	}

	length := min(C, G) - max(A, E)
	width := min(D, H) - max(B, F)

	return rectA + rectB - length*width

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
