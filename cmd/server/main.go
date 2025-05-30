package main

import (
	"knowledge-dedup-pipeline/internal/pipeline"
	"knowledge-dedup-pipeline/internal/similarity"
	"knowledge-dedup-pipeline/internal/storage"
	"knowledge-dedup-pipeline/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

func main() {
	store, _ := storage.NewStore("metadatadb")
	bloom := storage.NewBloomManager()
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		description := c.PostForm("description")
		price := c.PostForm("price")
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File required"})
			return
		}
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open file"})
			return
		}
		defer src.Close()

		// Shingling/Chunking
		shingles := utils.ShingleStream(src, 512*1024, 256*1024) // 512 KB chunks, 50% overlap

		// Parallel chunk processing
		results := pipeline.ProcessChunksParallel(shingles, 1024, 128, store, bloom)

		// Aggregate for whole-object signatures
		mh := similarity.CreateMinHash(shingles, 128)
		sh := similarity.CreateSimHash(shingles)
		vec := pipeline.CreateFeatureVec(shingles, 1024)

		// Check for global object duplicates if desired here...

		// Store object meta
		objID := strconv.FormatUint(sh, 10)
		store.PutMeta(storage.ObjectMeta{
			ID:          objID,
			Description: description,
			Price:       price,
			MinHashSig:  mh,
			SimHashSig:  sh,
			FeatureVec:  vec,
		})

		c.JSON(http.StatusOK, gin.H{
			"message":  "Upload accepted",
			"objectID": objID,
		})
	})

	r.Run(":8080")
}
