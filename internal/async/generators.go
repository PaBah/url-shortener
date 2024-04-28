package async

import "github.com/PaBah/url-shortener.git/internal/dto"

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
