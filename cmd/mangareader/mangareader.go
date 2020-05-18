// package mangareader
package main

import (
	"context"
	"log"
	"mangafox/store"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	str, err := store.New(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Panic(err)
	}
	defer str.Client.Disconnect(ctx)
}
