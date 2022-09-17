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

func Test() {
	ic := client.Exists(ctx, "lang")
	fmt.Println(ic.Result())
	client.Set(ctx, "name", "go-key", 0)
	sc := client.Type(ctx, "name")
	fmt.Println(sc.Result())
	client.Del(ctx, "name")
}
func main() {
	Test()
}
