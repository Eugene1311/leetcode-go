package repeated_dna_sequences

func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	subStringToCount := make(map[string]int)

	for i := 0; i <= len(s)-10; i++ {
		subString := s[i : i+10]
		count, ok := subStringToCount[subString]
		if ok {
			subStringToCount[subString] = count + 1
		} else {
			subStringToCount[subString] = 1
		}
	}

	for k, v := range subStringToCount {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}
