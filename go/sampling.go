// Reservoir Sampling is too slow for large list
func getRandomN(list []string, n int) []string {
	if len(list) <= n {
		return list
	}

	rand.Seed(time.Now().Unix())
	idxSet := map[int]bool{}
	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		for idxSet[r] {
			r = (r + 1) % n
		}
		idxSet[r] = true
	}
	result := []string{}
	for idx := range idxSet {
		result = append(result, list[idx])
	}
	return result
}
