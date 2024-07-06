package async

import "github.com/PaBah/url-shortener.git/internal/dto"

// BulkDeletionDataGenerator - initiate shortened URLs batch deletion process via generation through channels
func BulkDeletionDataGenerator(input dto.DeleteURLsRequest) chan string {
	inputCh := make(chan string)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			inputCh <- data
		}
	}()

	return inputCh
}
