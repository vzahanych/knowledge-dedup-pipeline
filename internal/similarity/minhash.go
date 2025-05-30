package similarity

import "github.com/ekzhu/minhash"

// MinHashJaccard returns the Jaccard similarity of two MinHash signatures.
func MinHashJaccard(sig1, sig2 []uint64) float64 {
	if len(sig1) != len(sig2) {
		return 0
	}
	matches := 0
	for i := range sig1 {
		if sig1[i] == sig2[i] {
			matches++
		}
	}
	return float64(matches) / float64(len(sig1))
}

// CreateMinHash processes shingles into a MinHash signature.
func CreateMinHash(shingles [][]byte, numPerm int) []uint64 {
	mh := minhash.NewMinHash(numPerm)
	for _, shingle := range shingles {
		mh.Push(shingle)
	}
	return mh.Signature()
}
