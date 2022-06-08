func numEquivDominoPairs(dominoes [][]int) int {
	dominoOccurrences := make(map[[2]int]int)

	// sort each domino
	for i := 0; i < len(dominoes); i++ { 
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}
	}

	// count how many times sorted dominos are found
	for i := 0; i < len(dominoes); i++ { 
		var domino [2]int
		copy(domino[:], dominoes[i][:])
		dominoOccurrences[domino] += 1
	}

	// count the number of duplicate dominoes
	numPairs := 0
	for _, v := range dominoOccurrences {
		// only check occurrences > 1 because each domino exists at least once
		if v > 1 {
			// the magic is that for any n things there can be only
			// (n * (n - 1)) / 2 pairings for them
			numPairs += (v * (v-1)) / 2
		}
	}
	return numPairs
}