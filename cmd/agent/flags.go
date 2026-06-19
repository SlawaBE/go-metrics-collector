package main

import (
	"flag"
)

var (
	flagRunAddress     string
	flagPollInterval   int
	flagReportInterval int
)

func parseFlags() {
	flag.StringVar(&flagRunAddress, "a", "localhost:8080", "address and port of metric server")
	flag.IntVar(&flagPollInterval, "p", 2, "metrics polling interval")
	flag.IntVar(&flagReportInterval, "r", 10, "metrics report interval")

	flag.Parse()
}
