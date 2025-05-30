
A scalable Golang pipeline for large binary object uploads, supporting chunk-level deduplication, similarity search, and optimized storage with sharding, Bloom filters, and Protocol Buffers.

## Features

- Streaming chunk processing with overlap (default: 512 KB chunks, 50% overlap)
- MinHash, SimHash, feature vector extraction
- Parallel chunk workers
- Sharded chunk signature storage
- Bloom filter for fast existence checks
- Top-N chunk summary index
- Compact Protocol Buffers serialization

## Quick Start

1. Install Go, BadgerDB, and all dependencies (`go get`).
2. Compile protobuf:

```bash
protoc --go_out=. model/chunk_signature.proto
```

3. Run the server:

```bash
go run cmd/server/main.go
```

4. POST to `/upload` with `multipart/form-data` fields:
- `file` (your binary object)
- `description`, `price`

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/path/to/your/big-object.bin" \
  -F "description=My sample object" \
  -F "price=10"
```

## Customization

Edit chunk size, overlap, worker count in code/config as needed for your deployment!
EOF

cd ..
echo "✅ All files and folders created! Next steps:"
echo "1. cd knowledge-dedup-pipeline"
echo "2. go mod init knowledge-dedup-pipeline && go mod tidy"
echo "3. protoc --go_out=. model/chunk_signature.proto"
echo "4. go run cmd/server/main.go"
echo "Done!"
