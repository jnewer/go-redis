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

func TestHSetAndHGet() {
	client.HSet(ctx, "site", "name", "lanhouzi.top")
	sc := client.HGet(ctx, "site", "name")
	fmt.Println(sc.Result())
}

func TestHMSetAndHMGet() {
	client.HMSet(ctx, "site", "name", "go-redis-hmset", "url", "lanhouzi.top", "author", "kong")
	sc := client.HMGet(ctx, "site", "name", "url", "author")
	fmt.Println(sc.Result())
}

func TestHKeysAndHVals() {
	client.HMSet(ctx, "site", "name", "go-redis-hmset", "url", "lanhouzi.top", "author", "kong")
	ssc := client.HKeys(ctx, "site")
	fmt.Println(ssc.Result())

	ssc2 := client.HVals(ctx, "site")
	fmt.Println(ssc2.Result())

}
func main() {
	//TestHSetAndHGet()
	//TestHMSetAndHMGet()
	TestHKeysAndHVals()
}
