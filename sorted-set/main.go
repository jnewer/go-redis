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

func TestZAddAndZRange() {
	m1 := &redis.Z{
		Score:  1,
		Member: "java",
	}

	m2 := &redis.Z{
		Score:  2,
		Member: "python",
	}

	m3 := &redis.Z{
		Score:  3,
		Member: "golang",
	}

	client.Del(ctx, "lang")
	client.ZAdd(ctx, "lang", m1, m2, m3)

	ssc := client.ZRange(ctx, "lang", 0, -1)
	fmt.Println(ssc.Result())
}

func main() {
	TestZAddAndZRange()
}
