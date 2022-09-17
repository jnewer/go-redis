package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func TestGetAndSet() {
	err := client.Set(ctx, "name", "go-redis", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "name").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("name", val)
}

func TestMSetAndMGet() {
	client.MSet(ctx, "name", "go redis mset", "url", "go-redis-mset.com", "author", "kong")
	sc := client.MGet(ctx, "name", "url", "author")

	fmt.Printf("sc: %v\n", sc)
}

func TestIncrAndDecr() {
	const KEY = "score"
	client.Set(ctx, KEY, "100", 0)
	client.Incr(ctx, KEY)
	client.Incr(ctx, KEY)
	sc := client.Get(ctx, KEY)
	fmt.Printf("sc: %v\n", sc)

	client.Decr(ctx, KEY)
	sc = client.Get(ctx, KEY)
	fmt.Printf("sc: %v\n", sc)
}
func main() {
	//TestGetAndSet()

	//TestMSetAndMGet()

	TestIncrAndDecr()
}
