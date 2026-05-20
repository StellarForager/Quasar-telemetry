package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	privKey, err := base64.StdEncoding.DecodeString(os.Getenv("TELEMETRY_PRI_KEY"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	configBytes, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = os.WriteFile("config.json.sig", []byte(
		base64.StdEncoding.EncodeToString(
			ed25519.Sign(ed25519.PrivateKey(privKey), configBytes))), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
