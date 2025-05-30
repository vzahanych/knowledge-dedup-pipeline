package storage

import (
	"encoding/binary"
	"fmt"
	"github.com/dgraph-io/badger/v4"
)

func (s *Store) IncrementChunkCount(chunkSig []byte) error {
	key := []byte(fmt.Sprintf("chunkcount:%x", chunkSig))
	return s.DB.Update(func(txn *badger.Txn) error {
		var cnt uint64 = 1
		item, err := txn.Get(key)
		if err == nil {
			val, _ := item.ValueCopy(nil)
			cnt = binary.BigEndian.Uint64(val) + 1
		}
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, cnt)
		return txn.Set(key, buf)
	})
}
