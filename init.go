package main

import (
	"fmt"
	"log"
)

func init() {
	if err := envSetup(); err != nil {
		log.Fatal(fmt.Errorf("cannot detect env: %w", err))
	}
}
