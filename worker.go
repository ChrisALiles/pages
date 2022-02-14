package pages

import (
	"math/rand"
	"time"
)

// Pretend to fetch the database page.
func Worker(input chan workreq) {
	for wr := range input {
		time.Sleep(time.Duration(rand.Intn(50) * int(time.Microsecond)))
		wr.ackChan <- wr.pageNum
	}
}
