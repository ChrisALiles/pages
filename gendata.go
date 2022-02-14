package pages

import (
	"math/rand"
)

// Gendata generates the transaction data and sends the txns
// to the scheduler.
func Gendata(txnChan chan request) {
	var req request
	for i := 0; i < numTxns; i++ {
		req.txnNum = i
		for j := 0; j < numReqPages; j++ {
			req.pageNums[j] = rand.Int63n(numDBPages)
		}
		txnChan <- req
	}
}
