package utils

func MoveLookAhead(pointer *int, currentLine *string, lineCounter int, programSize int, lookAhead *rune) {
	if (*pointer + 1) > len(*currentLine) {
		lineCounter++
		*pointer = 0
		if lineCounter > programSize {
			*lookAhead = -10
		} else {
			*currentLine += "\n"
			*lookAhead = rune((*currentLine)[*pointer])
		}
	} else {
		*lookAhead = rune((*currentLine)[*pointer])
	}

	if (*lookAhead >= 'a') &&
		(*lookAhead <= 'z') {
		*lookAhead = (*lookAhead - 'a' + 'A')
	}

	*pointer++
}
