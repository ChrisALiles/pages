package main

import (
	"fmt"

	"github.com/ChrisALiles/pages"
)

func main() {
	runAck := make(chan int64)
	go pages.Run(runAck)
	// As a sanity check, Run returns the number of pages counted
	// by the workers, the processes fetching pages.
	fmt.Println(<-runAck, "pages processed.")
}
