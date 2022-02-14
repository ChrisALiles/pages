package pages

// RunTxn processes transactions.
func RunTxn(txnChan chan request, runAck chan int64) {
	// The workAck channel is buffered to prevent deadlock
	// caused by worker processes not being able to send
	// on the channel.
	workAck := make(chan any, numReqPages)
	var pageCount int64
	for req := range txnChan {
		// Convert the array into a slice to avoid
		// having to call the sort using pointers.
		// The page numbers are sorted before they are
		// requested to avoid deadlock between concurrent
		// transactions.
		pn := req.pageNums[:]
		quicksort(pn, 0, numReqPages-1)
		// Send the page requests to workers based on the
		// page number to prevent concurrent access.
		for i := 0; i < numReqPages; i++ {
			var wr workreq
			wr.pageNum = pn[i]
			wr.ackChan = workAck
			workChans[pn[i]%numWorkers] <- wr
		}
		// Wait for the page fetches to complete.
		for i := 0; i < numReqPages; i++ {
			<-workAck
		}
		pageCount += numReqPages
	}
	runAck <- pageCount
}
