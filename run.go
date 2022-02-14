// Pages is an exercise in concurrency.
// It is intended to mimic part of what a DBMS has to do when
// storing a row in a table. This almost always involves fetching
// the page which is the primary location and also other pages
// where indexes are stored.
// For the purposes of the exercise, I assume that all pages
// have to be fetched before processing can continue - this
// allows deadlock to occur.
// The processing is controlled by the values in the "constants"
// file. Note that these are compile-time constants, not run-time.
package pages

var workChans [numWorkers]chan workreq

// Run starts the data generation and transaction run
// processes, then passes page counts back to the
// driver process.
func Run(runAck chan int64) {
	txnChan := make(chan request)
	runChan := make(chan int64)
	var pageCount int64
	for i := 0; i < numWorkers; i++ {
		workChans[i] = make(chan workreq)
	}
	// Start data generation.
	go func() {
		Gendata(txnChan)
		close(txnChan)
	}()
	// Start the transaction processes.
	for i := 0; i < numRunTxns; i++ {
		go RunTxn(txnChan, runChan)
	}
	// Start the workers.
	for i := 0; i < numWorkers; i++ {
		go Worker(workChans[i])
	}
	// Wait for transactions to finish and accumulate the database
	// page counts.
	for i := 0; i < numRunTxns; i++ {
		pageCount += <-runChan
	}
	// Close the worker channels.
	for i := 0; i < numWorkers; i++ {
		close(workChans[i])
	}
	// Respond to the driver program.
	runAck <- pageCount
}
