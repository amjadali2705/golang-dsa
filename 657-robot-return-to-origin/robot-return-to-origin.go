func judgeCircle(moves string) bool {
    X, Y := 0, 0

	for _, i := range moves {
		switch i {
		case 'U':
			Y++
		case 'D':
			Y--
		case 'L':
			X++
		case 'R':
			X--
		}
	}

	return X == 0 && Y == 0
}