package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	number, dataSize int

	ctx = context.Background()
)

func init() {
	flag.IntVar(&number, "n", 10000, "number of keys")
	flag.IntVar(&dataSize, "d", 100, "data size in bytes")
	flag.Parse()
}

func main() {
	fmt.Printf("Writing %d keys of %d bytes...\n", number, dataSize)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	value := make([]byte, dataSize)
	if _, err := rand.Read(value); err != nil {
		panic(err)
	}

	for i := 0; i < number; i++ {
		key := fmt.Sprint("key:", i)
		err := rdb.Set(ctx, key, value, time.Minute*20).Err()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Done!")
}
