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

func TestSAddAndSMembers() {
	client.Del(ctx, "lang")
	client.SAdd(ctx, "lang", "java", "python", "golang", "golang")
	ssc := client.SMembers(ctx, "lang")
	fmt.Println(ssc.Result())
}

//TestSUnion 并集
func TestSUnion() {
	client.Del(ctx, "lang1")
	client.Del(ctx, "lang2")
	client.SAdd(ctx, "lang1", "java", "python", "php", "object-c")
	client.SAdd(ctx, "lang2", "cpp", "c#", "java", "golang")

	ssc := client.SUnion(ctx, "lang1", "lang2")
	fmt.Println(ssc.Result())
}

//TestSInter 交集
func TestSInter() {
	client.Del(ctx, "lang1")
	client.Del(ctx, "lang2")

	client.SAdd(ctx, "lang1", "java", "python", "php", "golang")
	client.SAdd(ctx, "lang2", "cpp", "c#", "java", "golang")

	ssc := client.SInter(ctx, "lang1", "lang2")
	fmt.Println(ssc.Result())

}
func main() {
	//TestSAddAndSMembers()
	//TestSUnion()
	TestSInter()
}
