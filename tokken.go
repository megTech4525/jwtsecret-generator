package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
}
