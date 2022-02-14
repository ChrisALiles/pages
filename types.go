package pages

// Request defines a transaction request.
// Created by Gendata and processed by RunTxn.
type request struct {
	txnNum   int
	pageNums [numReqPages]int64
}

// Workreq defines a page request to a worker process.
type workreq struct {
	ackChan chan any
	pageNum int64
}
