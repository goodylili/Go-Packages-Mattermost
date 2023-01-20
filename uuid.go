package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID
	uuidV4 := uuid.New()
	fmt.Printf("UUIDv4: %s\n", uuidV4)

	// Generate a UUID based on a string input
	uuidV5 := uuid.NewSHA1(uuid.NameSpaceDNS, []byte("example.com"))
	fmt.Printf("UUIDv5: %s\n", uuidV5)

	// Parse a string into a UUID
	stringUUID, _ := uuid.Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	fmt.Printf("Parsed UUID: %s\n", stringUUID)
}
