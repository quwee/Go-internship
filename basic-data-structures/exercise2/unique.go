package exercise2

func RemoveDuplicates(strings []string) []string {
	uniqueSet := make(map[string]struct{}, 1)
	result := make([]string, 0, 1)

	for _, s := range strings {
		_, ok := uniqueSet[s]

		if !ok {
			uniqueSet[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}
