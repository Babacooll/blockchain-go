package main

import (
	"fmt"
	"crypto/sha256"
)

func ComputeHashSha256(bytes []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(bytes))
}
