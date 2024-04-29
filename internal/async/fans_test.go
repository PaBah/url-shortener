package async

import (
	"context"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestDeletionFanIn_and_FanOut(t *testing.T) {
	inputCh := BulkDeletionDataGenerator([]string{"test"})

	userID := "test"
	fs := storage.NewInFileStorage("/tmp/.test_store")
	//"bc2c0be9"
	_ = fs.Store(context.Background(), models.NewShortURL("test", userID))

	channels := DeletionFanOut(userID, &fs, inputCh)
	addResultCh := DeletionFanIn(channels...)
	data := <-addResultCh
	assert.Equal(t, "", data, "fanIn and fanOut works correctly")
}
