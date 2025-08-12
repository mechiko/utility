package utility

func SplitStringSliceToChunks(slice []string, chunkSize int) [][]string {
	if chunkSize <= 0 {
		// avoid infinite loop; treat as no-op
		return nil
	}
	// pre-size capacity for fewer allocations
	capacity := (len(slice) + chunkSize - 1) / chunkSize
	chunks := make([][]string, 0, capacity)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// avoid slicing beyond slice length
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
