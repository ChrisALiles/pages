package pages

const (
	// Number of pages in the database.
	// Making this smaller increases the chance of contention and
	// deadlock.
	numDBPages = 10000
	// Number of transactions to run.
	numTxns = 10000000
	// Number of concurrent processes fetching database pages.
	numWorkers = 20
	// Number of transactions running concurrently.
	numRunTxns = 20
	// Number of pages per transaction.
	numReqPages = 10
)
