package similarity

import (
	"github.com/mfonda/simhash"
)

// CreateSimHash computes SimHash signature for shingles.
func CreateSimHash(shingles [][]byte) uint64 {
	features := []simhash.Feature{}
	for _, shingle := range shingles {
		features = append(features, simhash.NewBaseFeature(string(shingle)))
	}
	return simhash.NewSimhash().Compute(features)
}
