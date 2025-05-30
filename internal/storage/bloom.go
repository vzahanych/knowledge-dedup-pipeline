package storage

import (
	"github.com/willf/bloom"
	"sync"
)

type BloomManager struct {
	Filters [ShardCount]*bloom.BloomFilter
	mu      [ShardCount]sync.RWMutex
}

func NewBloomManager() *BloomManager {
	var bm BloomManager
	for i := 0; i < ShardCount; i++ {
		bm.Filters[i] = bloom.NewWithEstimates(1e6, 0.001)
	}
	return &bm
}

func (bm *BloomManager) Add(chunkSig []byte) {
	shardID := chunkShardID(chunkSig)
	bm.mu[shardID].Lock()
	defer bm.mu[shardID].Unlock()
	bm.Filters[shardID].Add(chunkSig)
}

func (bm *BloomManager) Test(chunkSig []byte) bool {
	shardID := chunkShardID(chunkSig)
	bm.mu[shardID].RLock()
	defer bm.mu[shardID].RUnlock()
	return bm.Filters[shardID].Test(chunkSig)
}
