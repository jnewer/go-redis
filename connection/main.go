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

func TestPing() {
	sc := client.Ping(ctx)
	fmt.Println(sc.Result())
}

func TestEcho() {
	sc := client.Echo(ctx, "go-redis-echo")
	fmt.Println(sc.Result())
}

func TestSelect() {
	sc := client.Do(ctx, "select", 1)
	fmt.Println(sc.Result())
}
func main() {
	//TestPing()
	//TestEcho()
	TestSelect()
}
