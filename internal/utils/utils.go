package utils

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BatchSplit(src []int, batchSize int) (result [][]int) {
	if batchSize <= 0 {
		panic("Incorrect batchSize")
	}

	remainder := len(src) % batchSize
	fullBatchesCount := len(src) / batchSize

	batchesCount := fullBatchesCount
	if remainder > 0 {
		batchesCount++
	}

	result = make([][]int, 0, batchesCount)

	for batchStart := 0; batchStart < len(src); batchStart += batchSize {
		batchEnd := min(batchStart+batchSize, len(src))
		result = append(result, src[batchStart:batchEnd])
	}

	return
}

func ReverseMap(src map[string]int) (result map[int]string) {
	if src == nil {
		return
	}

	result = make(map[int]string)

	for key, value := range src {
		if _, alreadyExists := result[value]; alreadyExists {
			panic("Two records with same value!")
		}

		result[value] = key
	}
	return
}

func Filter(src []string, wordsToRemove []string) (result []string) {
	if src == nil {
		return nil
	}

	if len(wordsToRemove) == 0 {
		return src
	}

	type void struct{}

	wordsToRemove_set := make(map[string]void)
	for _, word := range wordsToRemove {
		wordsToRemove_set[word] = void{}
	}

	result = make([]string, 0, len(src))

	for _, src_str := range src {
		if _, shouldIgnore := wordsToRemove_set[src_str]; shouldIgnore {
			continue
		}

		result = append(result, src_str)
	}

	return
}
