package storage

import (
	"fmt"
	"hash/fnv"
)

const ShardCount = 32

func chunkShardID(chunkSig []byte) int {
	h := fnv.New32a()
	h.Write(chunkSig)
	return int(h.Sum32() % ShardCount)
}

func chunkShardKey(chunkSig []byte) []byte {
	shardID := chunkShardID(chunkSig)
	return []byte(fmt.Sprintf("chunkshard:%02d:%x", shardID, chunkSig))
}
