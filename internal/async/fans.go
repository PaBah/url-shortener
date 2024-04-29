package async

import (
	"sync"

	"github.com/PaBah/url-shortener.git/internal/storage"
)

func DeletionFanOut(userID string, repository storage.Repository, inputCh chan string) []chan string {
	numWorkers := 3
	channels := make([]chan string, numWorkers)

	for i := 0; i < numWorkers; i++ {
		addResultCh := repository.AsyncCheckURLsUserID(userID, inputCh)
		channels[i] = addResultCh
	}

	return channels
}

func DeletionFanIn(resultChs ...chan string) chan string {
	finalCh := make(chan string)

	var wg sync.WaitGroup

	for _, ch := range resultChs {
		chClosure := ch

		wg.Add(1)

		go func() {
			defer wg.Done()

			for data := range chClosure {
				finalCh <- data
			}
		}()
	}

	go func() {
		wg.Wait()
		close(finalCh)
	}()

	return finalCh
}
