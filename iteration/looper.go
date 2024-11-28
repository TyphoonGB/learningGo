package iteration

func Loop(character string, reps int) string {
	var looped string
	for i := 0; i < reps; i++ {
		looped += character
	}

	return looped
}
