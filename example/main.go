package main

import apm "github.com/inhochoi/go-simple-apm/collector"

func main() {
	apm.Start()

	forever := make(chan bool)
	<-forever
}
