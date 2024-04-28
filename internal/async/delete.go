package async

import (
	"context"

	"github.com/PaBah/url-shortener.git/internal/storage"
)

func Delete(repository storage.Repository, inputCh chan string) {
	deletionBatchSize := 10
	go func() {
		ctx := context.Background()
		var deletionBuffer []string
		for data := range inputCh {
			deletionBuffer = append(deletionBuffer, data)

			if len(deletionBuffer) == deletionBatchSize {
				_ = repository.DeleteShortURLs(ctx, deletionBuffer)
				deletionBuffer = []string{}
			}
		}
		_ = repository.DeleteShortURLs(ctx, deletionBuffer)
	}()
}
