package async

import (
	"context"
	"os"
	"testing"

	"github.com/PaBah/url-shortener.git/internal/models"
	"github.com/PaBah/url-shortener.git/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	fs := storage.NewInFileStorage("/tmp/.test_store")
	defer fs.Close()
	//"bc2c0be9"
	_ = fs.Store(context.Background(), models.NewShortURL("test", "test"))
	inputCh := make(chan string)
	Delete(&fs, inputCh)
	inputCh <- "bc2c0be9"
	_, err := fs.FindByID(context.Background(), "bc2c0be9")
	assert.NoError(t, err, "no error URL")
	_ = os.Remove("/tmp/.test_store")
}
