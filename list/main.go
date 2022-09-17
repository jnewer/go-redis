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

func TestLPushAndLRange() {
	client.LPush(ctx, "lang", "java", "python", "golang")
	ssc := client.LRange(ctx, "lang", 0, 1)
	fmt.Println(ssc.Result())
}

func TestLPopAndRPop() {
	client.LPush(ctx, "lang", "java", "python", "golang")
	sc := client.LPop(ctx, "lang")
	fmt.Println(sc.Result())
	sc2 := client.RPop(ctx, "lang")
	fmt.Println(sc2.Result())
}

func main() {
	//TestLPushAndLRange()
	TestLPopAndRPop()
}
