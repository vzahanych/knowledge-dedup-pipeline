package pipeline

import (
	"knowledge-dedup-pipeline/internal/storage"
	"knowledge-dedup-pipeline/internal/similarity"
	"sync"
	"runtime"
)

type ChunkTask struct {
	Data []byte
	Idx  int
}

type ChunkResult struct {
	Idx         int
	MinHashSig  []uint64
	SimHashSig  uint64
	FeatureIdx  int
	FeatureVal  float64
	ChunkDup    bool
}

func ProcessChunksParallel(
	chunks [][]byte, vecLen, numPerm int, store *storage.Store, bloom *storage.BloomManager,
) []ChunkResult {
	numWorkers := runtime.NumCPU()
	tasks := make(chan ChunkTask, len(chunks))
	results := make(chan ChunkResult, len(chunks))

	worker := func() {
		for task := range tasks {
			mh := similarity.CreateMinHash([][]byte{task.Data}, numPerm)
			simSig := similarity.CreateSimHash([][]byte{task.Data})
			featureIdx := HashToIndex(task.Data, vecLen)
			sigBytes := storage.MarshalSimpleSignature(mh, simSig)

			// Bloom filter check
			chunkDup := bloom.Test(sigBytes)
			results <- ChunkResult{
				Idx:        task.Idx,
				MinHashSig: mh,
				SimHashSig: simSig,
				FeatureIdx: featureIdx,
				FeatureVal: 1.0,
				ChunkDup:   chunkDup,
			}
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); worker() }()
	}
	for i, chunk := range chunks {
		tasks <- ChunkTask{Data: chunk, Idx: i}
	}
	close(tasks)
	go func() { wg.Wait(); close(results) }()

	out := make([]ChunkResult, len(chunks))
	for res := range results {
		out[res.Idx] = res
	}
	return out
}
