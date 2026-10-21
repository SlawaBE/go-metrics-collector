package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SlawaBE/go-metrics-collector/internal/utils/checksum"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run parses the command-line flags, validates the signing key and prints
// the HMAC SHA-256 signature of the given value.
func run() error {
	value := flag.String("v", "", "string to be signed")
	key := flag.String("k", os.Getenv("KEY"), "secret key for HMAC SHA-256 signing")
	flag.Parse()

	if *key == "" {
		return errors.New("missing secret key, use -k <key> or set KEY env")
	}

	signature := checksum.Sign([]byte(*value), []byte(*key))
	fmt.Println(hex.EncodeToString(signature))
	return nil
}
