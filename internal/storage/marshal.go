package storage

import (
	"bytes"
	"encoding/gob"
)

func MarshalSimpleSignature(minHash []uint64, simHash uint64) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(minHash)
	_ = enc.Encode(simHash)
	return buf.Bytes()
}
