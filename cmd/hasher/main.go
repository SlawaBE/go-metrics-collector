package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
)

func main() {
	var (
		value = flag.String("v", "", "string to be signed")
		key   = flag.String("k", os.Getenv("KEY"), "secret key for HMAC SHA-256 signing")
	)
	flag.Parse()

	if *key == "" {
		fmt.Fprintln(os.Stderr, "missing secret key, use -k <key> or set KEY env")
		os.Exit(1)
	}

	signature := checksum.Sign([]byte(*value), []byte(*key))
	fmt.Println(hex.EncodeToString(signature))
}
