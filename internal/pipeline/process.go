package pipeline

import (
	"hash/fnv"
)

func HashToIndex(data []byte, vecLen int) int {
	h := fnv.New32a()
	h.Write(data)
	return int(h.Sum32()) % vecLen
}

func CreateFeatureVec(shingles [][]byte, vecLen int) []float64 {
	vec := make([]float64, vecLen)
	for _, shingle := range shingles {
		idx := HashToIndex(shingle, vecLen)
		vec[idx]++
	}
	return vec
}
